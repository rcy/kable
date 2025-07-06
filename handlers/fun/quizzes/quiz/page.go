package quiz

import (
	_ "embed"
	"errors"
	"fmt"
	"net/http"
	"oj/api"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
	"oj/internal/middleware/quizctx"

	"github.com/go-chi/chi/v5"
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

func (s *service) Router(r chi.Router) {
	r.Use(quizctx.NewService(s.Queries).Provider)
	r.Get("/", s.page)
	r.Post("/attempt", s.createAttempt)
}

func (s *service) page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(ctx)
	quiz := quizctx.Value(ctx)

	questions, err := s.Queries.QuizQuestions(ctx, quiz.ID)
	if err != nil {
		render.Error(w, fmt.Errorf("QuizQuestions: %w", err), http.StatusInternalServerError)
		return
	}
	if len(questions) == 0 {
		render.Error(w, errors.New("quiz has no questions"), http.StatusInternalServerError)
		return
	}

	render.Execute(w, pageTemplate, struct {
		Layout           layout.Data
		Quiz             api.Quiz
		Questions        []api.Question
		CreateAttemptURL string
	}{
		Layout:           l,
		Quiz:             quiz,
		Questions:        questions,
		CreateAttemptURL: r.URL.Path + "/attempt",
	})
}

func (s *service) createAttempt(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := auth.FromContext(ctx)
	quiz := quizctx.Value(ctx)

	attempt, err := s.Queries.CreateAttempt(ctx, api.CreateAttemptParams{
		QuizID: quiz.ID,
		UserID: user.ID,
	})
	if err != nil {
		render.Error(w, fmt.Errorf("CreateAttempt: %w", err), http.StatusInternalServerError)
		return
	}

	w.Header().Add("HX-Redirect", fmt.Sprintf("/fun/quizzes/attempts/%d", attempt.ID))
	w.WriteHeader(201)
}
