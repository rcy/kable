package compose

import (
	_ "embed"
	"log"
	"net/http"
	"oj/api"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type service struct {
	Queries *api.Queries
}

func NewService(q *api.Queries) *service {
	return &service{Queries: q}
}

func (s *service) Router(r chi.Router) {
	r.Get("/", s.page)
	r.Post("/", s.post)
}

var (
	//go:embed page.gohtml
	pageContent string

	pageTemplate = layout.MustParse(pageContent)
)

func (s *service) page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(r.Context())

	connections, err := s.Queries.GetConnections(ctx, l.User.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	render.Execute(w, pageTemplate, struct {
		Layout      layout.Data
		Connections []api.User
	}{
		Layout:      l,
		Connections: connections,
	})
}

func (s *service) post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sender := auth.FromContext(ctx)

	recipient, _ := strconv.Atoi(r.FormValue("recipient"))
	params := api.CreatePostcardParams{
		Sender:    sender.ID,
		Recipient: int64(recipient),
		Subject:   r.FormValue("subject"),
		Body:      r.FormValue("body"),
		State:     "queued",
	}

	log.Print("postcard", params)

	_, err := s.Queries.CreatePostcard(ctx, params)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/postoffice", http.StatusSeeOther)
}
