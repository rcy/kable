package completed

import (
	_ "embed"
	"fmt"
	"net/http"
	"oj/api"
	"oj/handlers/layout"
	"oj/handlers/render"
	"strconv"

	"github.com/go-chi/chi/v5"
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

	attemptID, _ := strconv.Atoi(chi.URLParam(r, "attemptID"))
	attempt, err := s.Queries.GetAttemptByID(ctx, int64(attemptID))
	if err != nil {
		if err == pgx.ErrNoRows {
			render.Error(w, fmt.Errorf("GetAttemptByID: %w", err), http.StatusNotFound)
			return
		}
		render.Error(w, fmt.Errorf("GetAttemptByID: %w", err), http.StatusNotFound)
		return
	}

	quiz, err := s.Queries.Quiz(ctx, attempt.QuizID)
	if err != nil {
		if err == pgx.ErrNoRows {
			render.Error(w, fmt.Errorf("Quiz: %w", err), http.StatusNotFound)
			return
		}
		render.Error(w, fmt.Errorf("Quiz: %w", err), http.StatusInternalServerError)
		return
	}

	questionCount, err := s.Queries.QuestionCount(ctx, quiz.ID)
	if err != nil {
		render.Error(w, fmt.Errorf("QuestionCount: %w", err), http.StatusInternalServerError)
		return
	}

	responses, err := s.Queries.Responses(ctx, attempt.ID)
	if err != nil {
		render.Error(w, fmt.Errorf("Responses: %w", err), http.StatusInternalServerError)
		return
	}

	render.Execute(w, pageTemplate, struct {
		Layout        layout.Data
		Quiz          api.Quiz
		Attempt       api.Attempt
		QuestionCount int64
		Responses     []api.ResponsesRow
	}{
		Layout:        l,
		Quiz:          quiz,
		Attempt:       attempt,
		QuestionCount: questionCount,
		Responses:     responses,
	})
}
