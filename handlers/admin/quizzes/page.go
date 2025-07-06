package quizzes

import (
	_ "embed"
	"fmt"
	"net/http"
	"oj/api"
	"oj/handlers/admin/quizzes/create"
	"oj/handlers/admin/quizzes/show"
	"oj/handlers/layout"
	"oj/handlers/render"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

type service struct {
	Queries *api.Queries
}

func NewService(q *api.Queries) *service {
	return &service{Queries: q}
}

func (s *service) Router(r chi.Router) {
	r.Get("/", s.page)
	r.Route("/create", create.NewService(s.Queries).Router)
	r.Route("/{quizID}", show.NewService(s.Queries).Router)
}

var (
	//go:embed page.gohtml
	pageContent  string
	pageTemplate = layout.MustParse(pageContent, pageContent)
)

func (s *service) page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(ctx)

	allQuizzes, err := s.Queries.AllQuizzes(ctx)
	if err != nil && err != pgx.ErrNoRows {
		render.Error(w, fmt.Errorf("AllQuizzes: %w", err), http.StatusInternalServerError)
		return
	}

	render.Execute(w, pageTemplate, struct {
		Layout  layout.Data
		Quizzes []api.Quiz
	}{
		Layout:  l,
		Quizzes: allQuizzes,
	})
}
