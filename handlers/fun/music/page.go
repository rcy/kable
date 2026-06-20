package music

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"oj/api"
	"oj/handlers/layout"
	"oj/internal/middleware/auth"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lrstanley/go-ytdlp"
)

var (
	activeDownloads = map[string]string{}
	activeMu        sync.Mutex
	initOnce        sync.Once
)

func initYtdlp() {
	initOnce.Do(func() {
		log.Println("installing yt-dlp and dependencies...")
		_, err := ytdlp.InstallAll(context.Background())
		if err != nil {
			log.Printf("yt-dlp install warning: %v", err)
		}
		log.Println("yt-dlp ready")
	})
}

type service struct {
	Queries *api.Queries
}

func NewService(q *api.Queries) *service {
	return &service{Queries: q}
}

func dlPath() string {
	p := os.Getenv("MUSIC_DOWNLOAD_PATH")
	if p == "" {
		return "downloads"
	}
	return p
}

func (s *service) Page(w http.ResponseWriter, r *http.Request) {
	initYtdlp()
	os.MkdirAll(dlPath(), 0755)

	ctx := r.Context()
	user := auth.FromContext(ctx)

	_, err := s.Queries.UserMusicTracks(ctx, user.ID)
	if err != nil {
		log.Printf("UserMusicTracks: %v", err)
	}

	l := layout.FromContext(r.Context())
	layout.Layout(l, "Music", s.musicPage(user.ID)).Render(w)
}

func (s *service) Download(w http.ResponseWriter, r *http.Request) {
	initYtdlp()
	os.MkdirAll(dlPath(), 0755)

	ctx := r.Context()
	user := auth.FromContext(ctx)

	url := r.FormValue("url")
	if url == "" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_ = s.downloadsList(user.ID).Render(w)
		return
	}

	id := fmt.Sprintf("%d", time.Now().UnixNano())

	_, err := s.Queries.InsertMusicTrack(ctx, api.InsertMusicTrackParams{
		UserID:   user.ID,
		Url:      url,
		Filename: id + ".mp3",
		Status:   "downloading",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	activeMu.Lock()
	activeDownloads[id] = url
	activeMu.Unlock()

	go s.runDownload(id, url, user.ID)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_ = s.downloadsList(user.ID).Render(w)
}

func (s *service) Status(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	activeMu.Lock()
	_, active := activeDownloads[id]
	activeMu.Unlock()

	if active {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_ = downloadingRow(id, activeDownloads[id]).Render(w)
		return
	}

	user := auth.FromContext(r.Context())
	tracks, err := s.Queries.UserMusicTracks(r.Context(), user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, t := range tracks {
		if t.Filename == id+".mp3" {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			_ = trackRow(t).Render(w)
			return
		}
	}

	http.NotFound(w, r)
}

func (s *service) File(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "missing name", http.StatusBadRequest)
		return
	}

	path := filepath.Join(dlPath(), filepath.Base(name))
	if _, err := os.Stat(path); os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filepath.Base(name)))
	http.ServeFile(w, r, path)
}

func (s *service) runDownload(id, url string, userID int64) {
	dl := ytdlp.New().
		FormatSort("res,ext:mp4:m4a").
		ExtractAudio().
		AudioFormat("mp3").
		Output(dlPath() + "/" + id + ".%(ext)s").
		NoProgress().
		PrintJSON()

	result, err := dl.Run(context.Background(), url)

	activeMu.Lock()
	delete(activeDownloads, id)
	activeMu.Unlock()

	ctx := context.Background()

	if err != nil {
		_ = s.Queries.UpdateMusicTrack(ctx, api.UpdateMusicTrackParams{
			UserID:   userID,
			Filename: id + ".mp3",
			Status:   "error",
			Error:    toPgText(err.Error()),
		})
		return
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

	_ = s.Queries.UpdateMusicTrack(ctx, api.UpdateMusicTrackParams{
		UserID:   userID,
		Filename: id + ".mp3",
		Title:    toPgText(title),
		Uploader: toPgText(uploader),
		Status:   "done",
	})
}

func (s *service) musicPage(userID int64) g.Node {
	return g.Group{
		h.H1(
			h.Class("nes-text is-primary"),
			g.Text("Music Ripper"),
		),
		h.P(
			h.Class("nes-text"),
			g.Text("Download audio from YouTube, SoundCloud, and more."),
		),
		h.Form(
			h.Class("nes-field"),
			h.Style("margin-bottom:2rem"),
			g.Attr("hx-post", "/fun/music/download"),
			g.Attr("hx-target", "#downloads"),
			g.Attr("hx-swap", "innerHTML"),
			g.Attr("hx-on::after-request", "this.reset()"),
			h.Div(
				h.Style("display:flex; gap:0.5rem"),
				h.Input(
					h.Class("nes-input"),
					h.Type("text"),
					h.Name("url"),
					h.Placeholder("Paste a video URL..."),
					h.Required(),
					h.AutoFocus(),
				),
				h.Button(
					h.Class("nes-btn is-primary"),
					h.Type("submit"),
					g.Text("Download"),
				),
			),
		),
		h.H3(h.Class("nes-text"), g.Text("Downloads")),
		h.Div(h.ID("downloads"), s.downloadsList(userID)),
		h.Audio(h.ID("player"), g.Attr("preload", "none"), h.Style("display:none")),
		g.Raw(`<script>
let currentTrack=null;
function playTrack(id,btn){
 const p=document.getElementById('player');
 if(currentTrack===id){
  if(p.paused){p.play();btn.textContent='\u23F8';btn.classList.add('is-primary');}
  else{p.pause();btn.textContent='\u25B6';btn.classList.remove('is-primary');btn.classList.add('is-success');}
  return;
 }
 if(currentTrack){
  const prev=document.querySelector('[data-track="'+currentTrack+'"]');
  if(prev){prev.textContent='\u25B6';prev.classList.remove('is-primary');prev.classList.add('is-success');}
 }
 currentTrack=id;
 p.src='/fun/music/file?name='+id;
 p.play();
 btn.textContent='\u23F8';
 btn.classList.remove('is-success');
 btn.classList.add('is-primary');
 p.onended=()=>{btn.textContent='\u25B6';btn.classList.remove('is-primary');btn.classList.add('is-success');currentTrack=null;};
}
</script>`),
	}
}

func (s *service) downloadsList(userID int64) g.Node {
	tracks, err := s.Queries.UserMusicTracks(context.Background(), userID)
	if err != nil || len(tracks) == 0 {
		activeMu.Lock()
		hasActive := len(activeDownloads) > 0
		activeMu.Unlock()

		if !hasActive {
			return h.P(
				h.Class("nes-text is-disabled"),
				g.Text("Nothing downloaded yet."),
			)
		}

		var nodes []g.Node
		activeMu.Lock()
		for id, url := range activeDownloads {
			nodes = append(nodes, downloadingRow(id, url))
		}
		activeMu.Unlock()
		return g.Group(nodes)
	}

	var nodes []g.Node
	for _, t := range tracks {
		id := t.Filename[:len(t.Filename)-len(".mp3")]
		activeMu.Lock()
		_, active := activeDownloads[id]
		activeMu.Unlock()

		if active && t.Status == "downloading" {
			nodes = append(nodes, downloadingRow(id, t.Url))
		} else {
			nodes = append(nodes, trackRow(t))
		}
	}
	return g.Group(nodes)
}

func downloadingRow(id, url string) g.Node {
	return h.Div(
		h.Class("nes-container"),
		h.Style("margin-bottom:0.5rem"),
		g.Attr("hx-get", "/fun/music/status?id="+id),
		g.Attr("hx-trigger", "every 2s"),
		g.Attr("hx-swap", "outerHTML"),
		h.Div(
			h.Style("display:flex; align-items:center; gap:0.75rem"),
			h.I(h.Class("nes-octocat animate"), h.Style("width:32px;height:32px")),
			h.Span(g.Text("Downloading...")),
			h.Span(
				h.Class("nes-text is-disabled"),
				h.Style("font-size:0.75rem; overflow:hidden; text-overflow:ellipsis; white-space:nowrap"),
				g.Text(url),
			),
		),
	)
}

func trackRow(t api.MusicTrack) g.Node {
	switch t.Status {
	case "downloading":
		return downloadingRow(t.Filename[:len(t.Filename)-len(".mp3")], t.Url)
	case "done":
		title := t.Title.String
		if title == "" {
			title = t.Filename
		}
		uploader := t.Uploader.String
		if uploader == "" {
			uploader = "—"
		}
		return h.Div(
			h.Class("nes-container"),
			h.Style("margin-bottom:0.5rem"),
			h.Div(
				h.Style("display:flex; align-items:center; gap:0.75rem"),
				h.Div(
					h.Style("flex:1; min-width:0"),
					h.Strong(
						h.Style("display:block; overflow:hidden; text-overflow:ellipsis; white-space:nowrap"),
						g.Text(title),
					),
					h.Small(
						h.Class("nes-text is-disabled"),
						g.Text(uploader),
					),
				),
				h.Div(
					h.Style("display:flex; gap:0.35rem; flex-shrink:0"),
					h.Button(
						h.Class("nes-btn is-success"),
						h.Style("padding:0.35rem 0.5rem; font-size:0.8rem"),
						g.Attr("data-track", t.Filename),
						g.Attr("onclick", "playTrack('"+t.Filename+"',this)"),
						g.Text("▶"),
					),
					h.A(
						h.Class("nes-btn"),
						h.Style("padding:0.35rem 0.5rem; font-size:0.8rem"),
						h.Href("/fun/music/file?name="+t.Filename),
						g.Text("MP3"),
					),
				),
			),
		)
	default:
		errMsg := t.Error.String
		if errMsg == "" {
			errMsg = "Unknown error"
		}
		return h.Div(
			h.Class("nes-container is-dark"),
			h.Style("margin-bottom:0.5rem"),
			h.Div(
				h.Style("display:flex; align-items:center; gap:0.75rem"),
				h.Div(h.Style("flex:1; min-width:0"),
					h.Strong(
						h.Style("display:block; overflow:hidden; text-overflow:ellipsis; white-space:nowrap"),
						g.Text(t.Url),
					),
					h.Small(
						h.Class("nes-text is-error"),
						g.Text(errMsg),
					),
				),
			),
		)
	}
}

func toPgText(s string) pgtype.Text {
	if s == "" {
		return pgtype.Text{Valid: false}
	}
	return pgtype.Text{String: s, Valid: true}
}
