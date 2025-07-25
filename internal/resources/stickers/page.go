package stickers

import (
	_ "embed"
	"fmt"
	"net/http"
	"net/url"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	goduckgo "github.com/minoplhy/duckduckgo-images-api"
)

type Resource struct {
	Conn *pgxpool.Pool
}

func (rs Resource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.page)
	r.Post("/", rs.search)
	r.Post("/save", rs.save)

	return r
}

var (
	//go:embed page.gohtml
	pageContent  string
	pageTemplate = layout.MustParse(pageContent)
)

type Image struct {
	ID        int64
	CreatedAt time.Time `db:"created_at"`
	UserID    int64     `db:"user_id"`
	URL       string    `db:"url"`
}

func (rs Resource) page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	l := layout.FromContext(r.Context())

	var images []Image
	err := pgxscan.Select(ctx, rs.Conn, &images, `select * from images where user_id = $1 order by created_at desc`, l.User.ID)
	if err != nil {
		render.Error(w, fmt.Errorf("select * from images: %w", err), http.StatusInternalServerError)
		return
	}

	d := struct {
		Layout layout.Data
		Images []Image
	}{
		Layout: l,
		Images: images,
	}

	render.Execute(w, pageTemplate, d)
}

func (rs Resource) search(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("query")

	keyword := url.QueryEscape("cartoon " + query)

	result := goduckgo.Search(goduckgo.Query{Keyword: keyword})

	render.ExecuteNamed(w, pageTemplate, "result", struct {
		URL       string
		Thumbnail string
	}{
		URL:       result.Results[0].Image,
		Thumbnail: result.Results[0].Thumbnail,
	})
}

func (rs Resource) save(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user := auth.FromContext(r.Context())

	url := r.FormValue("url")

	img := Image{URL: url}

	_, err := rs.Conn.Exec(ctx, `insert into images(url, user_id) values($1,$2)`, url, user.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.ExecuteNamed(w, pageTemplate, "saveSticker", img)
}
