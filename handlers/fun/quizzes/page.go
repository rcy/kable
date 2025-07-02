package quizzes

import (
	_ "embed"
	"fmt"
	"net/http"
	"oj/api"
	"oj/handlers/layout"
	"oj/handlers/render"

	"github.com/jackc/pgx/v5"
)

type service struct {
	Queries *api.Queries
}

func NewService(q *api.Queries) *service {
	return &service{Queries: q}
}

var (
	//go:embed page.gohtml
	pageContent  string
	pageTemplate = layout.MustParse(pageContent)
)

func (s *service) Page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(ctx)

	allQuizzes, err := s.Queries.PublishedQuizzes(ctx)
	if err != nil && err != pgx.ErrNoRows {
		render.Error(w, fmt.Errorf("PublishedQuizzes: %w", err), http.StatusInternalServerError)
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
