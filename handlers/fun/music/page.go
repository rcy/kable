package music

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"oj/handlers/layout"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/lrstanley/go-ytdlp"
)

type Job struct {
	ID       string
	URL      string
	Status   string
	Title    string
	Uploader string
	Filename string
	Error    string
}

var (
	jobs     = map[string]*Job{}
	jobOrder []string
	jobsMu   sync.Mutex
	initOnce sync.Once
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

func Page(w http.ResponseWriter, r *http.Request) {
	initYtdlp()
	loadExistingFiles()

	l := layout.FromContext(r.Context())
	layout.Layout(l, "Music", musicPage()).Render(w)
}

func Download(w http.ResponseWriter, r *http.Request) {
	initYtdlp()

	url := r.FormValue("url")
	if url == "" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_ = downloadsList().Render(w)
		return
	}

	id := fmt.Sprintf("%d", time.Now().UnixNano())
	job := &Job{ID: id, URL: url, Status: "downloading"}

	jobsMu.Lock()
	jobs[id] = job
	jobOrder = append([]string{id}, jobOrder...)
	jobsMu.Unlock()

	go runDownload(job)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_ = downloadsList().Render(w)
}

func Status(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	jobsMu.Lock()
	job, ok := jobs[id]
	jobsMu.Unlock()

	if !ok {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_ = trackRow(job).Render(w)
}

func File(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "missing name", http.StatusBadRequest)
		return
	}

	path := filepath.Join("downloads", filepath.Base(name))
	if _, err := os.Stat(path); os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filepath.Base(name)))
	http.ServeFile(w, r, path)
}

func runDownload(job *Job) {
	dl := ytdlp.New().
		FormatSort("res,ext:mp4:m4a").
		ExtractAudio().
		AudioFormat("mp3").
		Output("downloads/" + job.ID + ".%(ext)s").
		NoProgress().
		PrintJSON()

	result, err := dl.Run(context.Background(), job.URL)

	jobsMu.Lock()
	defer jobsMu.Unlock()

	if err != nil {
		job.Status = "error"
		job.Error = err.Error()
		return
	}

	if info, e := result.GetExtractedInfo(); e == nil && len(info) > 0 {
		if info[0].Title != nil {
			job.Title = *info[0].Title
		}
		if info[0].Uploader != nil {
			job.Uploader = *info[0].Uploader
		} else if info[0].Channel != nil {
			job.Uploader = *info[0].Channel
		}
	}

	job.Status = "done"
	job.Filename = job.ID + ".mp3"

	meta, _ := json.Marshal(struct {
		Title    string `json:"title"`
		Uploader string `json:"uploader"`
	}{Title: job.Title, Uploader: job.Uploader})
	_ = os.WriteFile(filepath.Join("downloads", job.ID+".json"), meta, 0644)
}

func loadExistingFiles() {
	if err := os.MkdirAll("downloads", 0755); err != nil {
		return
	}
	entries, err := os.ReadDir("downloads")
	if err != nil {
		return
	}
	for _, e := range entries {
		if e.IsDir() || filepath.Ext(e.Name()) != ".json" {
			continue
		}
		data, err := os.ReadFile(filepath.Join("downloads", e.Name()))
		if err != nil {
			continue
		}
		var meta struct {
			Title    string `json:"title"`
			Uploader string `json:"uploader"`
		}
		if err := json.Unmarshal(data, &meta); err != nil {
			continue
		}
		id := e.Name()[:len(e.Name())-len(".json")]
		mp3Name := id + ".mp3"
		if _, err := os.Stat(filepath.Join("downloads", mp3Name)); os.IsNotExist(err) {
			continue
		}
		j := &Job{
			ID:       id,
			Status:   "done",
			Title:    meta.Title,
			Uploader: meta.Uploader,
			Filename: mp3Name,
		}
		if _, exists := jobs[id]; exists {
			continue
		}
		jobs[id] = j
		jobOrder = append(jobOrder, id)
	}
}

func musicPage() g.Node {
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
		h.Div(h.ID("downloads"), downloadsList()),
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
 p.src='/fun/music/file?name='+id+'.mp3';
 p.play();
 btn.textContent='\u23F8';
 btn.classList.remove('is-success');
 btn.classList.add('is-primary');
 p.onended=()=>{btn.textContent='\u25B6';btn.classList.remove('is-primary');btn.classList.add('is-success');currentTrack=null;};
}
</script>`),
	}
}

func trackRow(job *Job) g.Node {
	switch job.Status {
	case "downloading":
		return h.Div(
			h.Class("nes-container"),
			h.Style("margin-bottom:0.5rem"),
			g.Attr("hx-get", "/fun/music/status?id="+job.ID),
			g.Attr("hx-trigger", "every 2s"),
			g.Attr("hx-swap", "outerHTML"),
			h.Div(
				h.Style("display:flex; align-items:center; gap:0.75rem"),
				h.I(h.Class("nes-octocat animate"), h.Style("width:32px;height:32px")),
				h.Span(g.Text("Downloading...")),
				h.Span(
					h.Class("nes-text is-disabled"),
					h.Style("font-size:0.75rem; overflow:hidden; text-overflow:ellipsis; white-space:nowrap"),
					g.Text(job.URL),
				),
			),
		)
	case "done":
		title := job.Title
		if title == "" {
			title = job.Filename
		}
		uploader := job.Uploader
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
						g.Attr("data-track", job.ID),
						g.Attr("onclick", "playTrack('"+job.ID+"',this)"),
						g.Text("▶"),
					),
					h.A(
						h.Class("nes-btn"),
						h.Style("padding:0.35rem 0.5rem; font-size:0.8rem"),
						h.Href("/fun/music/file?name="+job.Filename),
						g.Text("MP3"),
					),
				),
			),
		)
	default:
		return h.Div(
			h.Class("nes-container is-dark"),
			h.Style("margin-bottom:0.5rem"),
			h.Div(
				h.Style("display:flex; align-items:center; gap:0.75rem"),
				h.Div(h.Style("flex:1; min-width:0"),
					h.Strong(
						h.Style("display:block; overflow:hidden; text-overflow:ellipsis; white-space:nowrap"),
						g.Text(job.URL),
					),
					h.Small(
						h.Class("nes-text is-error"),
						g.Text(job.Error),
					),
				),
			),
		)
	}
}

func downloadsList() g.Node {
	jobsMu.Lock()
	ids := make([]string, len(jobOrder))
	copy(ids, jobOrder)
	jobsMu.Unlock()

	if len(ids) == 0 {
		return h.P(
			h.Class("nes-text is-disabled"),
			g.Text("Nothing downloaded yet."),
		)
	}

	var tracks []g.Node
	for _, id := range ids {
		jobsMu.Lock()
		job := jobs[id]
		jobsMu.Unlock()
		if job == nil {
			continue
		}
		tracks = append(tracks, trackRow(job))
	}
	return g.Group(tracks)
}
