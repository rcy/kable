package youtubedownload

import (
	"context"
	"fmt"
	"log"
	"os"

	"oj/api"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lrstanley/go-ytdlp"
	"github.com/riverqueue/river"
)

type YoutubeDownloadArgs struct {
	MusicTrackID int64
}

func (YoutubeDownloadArgs) Kind() string { return "youtube_download" }

type Worker struct {
	river.WorkerDefaults[YoutubeDownloadArgs]
	Queries *api.Queries
}

func NewWorker(q *api.Queries) *Worker {
	return &Worker{Queries: q}
}

func dlPath() string {
	p := os.Getenv("MUSIC_DOWNLOAD_PATH")
	if p == "" {
		return "./music-downloads"
	}
	return p
}

func (w *Worker) Work(ctx context.Context, job *river.Job[YoutubeDownloadArgs]) error {
	trackID := job.Args.MusicTrackID

	track, err := w.Queries.MusicTrackByID(ctx, trackID)
	if err != nil {
		return fmt.Errorf("fetching track %d: %w", trackID, err)
	}

	cmd := os.Getenv("YTDLP_EXECUTABLE")
	if cmd == "" {
		return fmt.Errorf("YTDLP_EXECUTABLE not set")
	}

	dl := ytdlp.New().
		SetExecutable(cmd).
		FormatSort("res,ext:mp4:m4a").
		ExtractAudio().
		AudioFormat("mp3").
		Output(dlPath() + "/" + track.Filename).
		NoProgress().
		PrintJSON()

	cookiesFile := os.Getenv("YTDLP_COOKIES_FILE")
	if cookiesFile != "" {
		dl.Cookies(cookiesFile)
	}

	result, err := dl.Run(ctx, track.Url)
	if err != nil {
		_ = w.Queries.UpdateMusicTrackByID(ctx, api.UpdateMusicTrackByIDParams{
			ID:     trackID,
			Status: "error",
			Error:  toPgText(err.Error()),
		})
		return fmt.Errorf("downloading: %w", err)
	}

	title := ""
	uploader := ""
	if info, e := result.GetExtractedInfo(); e == nil && len(info) > 0 {
		if info[0].Title != nil {
			title = *info[0].Title
		}
		if info[0].Uploader != nil {
			uploader = *info[0].Uploader
		} else if info[0].Channel != nil {
			uploader = *info[0].Channel
		}
	}

	err = w.Queries.UpdateMusicTrackByID(ctx, api.UpdateMusicTrackByIDParams{
		ID:       trackID,
		Title:    toPgText(title),
		Uploader: toPgText(uploader),
		Status:   "done",
	})
	if err != nil {
		return fmt.Errorf("updating track %d: %w", trackID, err)
	}

	log.Printf("downloaded track %d: %s - %s", trackID, title, uploader)
	return nil
}

func toPgText(s string) pgtype.Text {
	if s == "" {
		return pgtype.Text{Valid: false}
	}
	return pgtype.Text{String: s, Valid: true}
}
