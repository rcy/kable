package stickers

import (
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
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
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

	layout.Layout(l, "Sticker Book", stickersPage(images)).Render(w)
}

func (rs Resource) search(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("query")
	keyword := url.QueryEscape("cartoon " + query)
	result := goduckgo.Search(goduckgo.Query{Keyword: keyword})
	resultEl(result.Results[0].Image).Render(w)
}

func (rs Resource) save(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user := auth.FromContext(r.Context())

	imgURL := r.FormValue("url")

	_, err := rs.Conn.Exec(ctx, `insert into images(url, user_id) values($1,$2)`, imgURL, user.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	saveStickerEl(Image{URL: imgURL}).Render(w)
}

func stickersPage(images []Image) g.Node {
	return h.Div(
		h.H1(g.Text("Sticker Book")),
		h.P(g.Text("Type a word to search for stickers")),
		h.Form(g.Attr("hx-post", ""), g.Attr("hx-target", "#result"), g.Attr("hx-swap", "outerHTML"),
			h.Input(h.Type("text"), h.Name("query"), g.Attr("autofocus", ""), h.Style("width: 100%"), g.Attr("placeholder", "dog, cat, etc...")),
		),
		h.Div(h.Style("display:flex; flex-direction: column; gap:1em; align-items: center"),
			h.Div(h.ID("result")),
			h.Div(h.Class("nes-container is-dark"),
				stickerBookEl(images),
			),
		),
	)
}

func resultEl(imgURL string) g.Node {
	return h.Div(h.ID("result"), h.Style("display:flex; flex-direction: column; gap:1em"), h.Class("nes-container is-dark"),
		h.Div(h.Style("height: 400px"),
			h.Img(h.Height("100%"), h.Src(imgURL)),
		),
		h.Form(g.Attr("hx-post", "stickers/save"), g.Attr("hx-target", "#stickerBook"), g.Attr("hx-swap", "afterbegin"),
			h.Input(h.Type("hidden"), h.Name("url"), h.Value(imgURL)),
			h.Button(h.ID("addbutton"), h.Class("nes-btn"), h.Style("width:100%"), g.Text("Add to sticker book")),
		),
	)
}

func stickerBookEl(images []Image) g.Node {
	return h.Div(h.ID("stickerBook"), h.Style("display:flex; flex-wrap: wrap; gap:1em"),
		g.Map(images, func(img Image) g.Node { return stickerEl(img) }),
	)
}

func stickerEl(img Image) g.Node {
	return h.Div(h.Img(h.Src(img.URL), h.Height("100px")))
}

func saveStickerEl(img Image) g.Node {
	return g.Group{
		h.Div(h.ID("addButton"), g.Attr("hx-swap-oob", "true")),
		stickerEl(img),
	}
}
