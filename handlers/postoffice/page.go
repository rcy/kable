package postoffice

import (
	_ "embed"
	"net/http"
	"oj/api"
	"oj/handlers/layout"
	"oj/handlers/postoffice/compose"
	"oj/handlers/render"

	"github.com/go-chi/chi/v5"
)

type service struct {
	Queries *api.Queries
}

func NewService(q *api.Queries) *service {
	return &service{Queries: q}
}

func (s *service) Router(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/postoffice/inbox", http.StatusSeeOther)
	})

	r.Get("/inbox", s.page)
	r.Route("/compose", compose.NewService(s.Queries).Router)
}

var (
	//go:embed page.gohtml
	pageContent string

	pageTemplate = layout.MustParse(pageContent)
)

func (s *service) page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(r.Context())

	received, err := s.Queries.UserPostcardsReceived(ctx, l.User.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sent, err := s.Queries.UserPostcardsSent(ctx, l.User.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Execute(w, pageTemplate, struct {
		Layout   layout.Data
		Received []api.UserPostcardsReceivedRow
		Sent     []api.UserPostcardsSentRow
	}{
		Layout:   l,
		Received: received,
		Sent:     sent,
	})
}
