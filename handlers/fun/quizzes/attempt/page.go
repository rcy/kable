package attempt

import (
	_ "embed"
	"fmt"
	"net/http"
	"oj/api"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
	"strconv"
	"strings"

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
			render.Error(w, fmt.Errorf("GetAttemptID: %w", err), http.StatusNotFound)
			return
		}
		render.Error(w, fmt.Errorf("GetAttemptID: %w", err), http.StatusInternalServerError)
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

	responseCount, err := s.Queries.ResponseCount(ctx, attempt.ID)
	if err != nil {
		render.Error(w, fmt.Errorf("ResponseCount: %w", err), http.StatusInternalServerError)
		return
	}

	nextQuestion, err := s.Queries.AttemptNextQuestion(ctx, api.AttemptNextQuestionParams{
		QuizID:    attempt.QuizID,
		AttemptID: attempt.ID,
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			url := fmt.Sprintf("/fun/quizzes/attempts/%d/done", attempt.ID)
			if r.Header.Get("HX-Request") == "true" {
				w.Header().Add("HX-Redirect", url)
				w.WriteHeader(http.StatusSeeOther)
				return
			}
			http.Redirect(w, r, url, http.StatusSeeOther)
			return
		}
		render.Error(w, fmt.Errorf("AttemptNextQuestion: %w", err), http.StatusInternalServerError)
		return
	}

	render.Execute(w, pageTemplate, struct {
		Layout        layout.Data
		Quiz          api.Quiz
		Attempt       api.Attempt
		Question      api.Question
		QuestionCount int64
		ResponseCount int64
	}{
		Layout:        l,
		Quiz:          quiz,
		Attempt:       attempt,
		Question:      nextQuestion,
		QuestionCount: questionCount,
		ResponseCount: responseCount,
	})
}

func (s *service) PostResponse(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user := auth.FromContext(ctx)

	attemptID, err := strconv.Atoi(chi.URLParam(r, "attemptID"))
	if err != nil {
		render.Error(w, fmt.Errorf("Atoi: %w", err), http.StatusInternalServerError)
		return
	}
	attempt, err := s.Queries.GetAttemptByID(ctx, int64(attemptID))
	if err != nil {
		render.Error(w, fmt.Errorf("GetAttemptByID: %w", err), http.StatusInternalServerError)
		return
	}

	questionID, err := strconv.Atoi(chi.URLParam(r, "questionID"))
	if err != nil {
		render.Error(w, fmt.Errorf("Atoi: %w", err), http.StatusInternalServerError)
		return
	}

	text := strings.TrimSpace(r.FormValue("response"))

	if text != "" {
		_, err := s.Queries.CreateResponse(ctx, api.CreateResponseParams{
			QuizID:     attempt.QuizID,
			UserID:     user.ID,
			AttemptID:  int64(attemptID),
			QuestionID: int64(questionID),
			Text:       text,
		})

		if err != nil {
			render.Error(w, fmt.Errorf("CreateResponse: %w", err), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Add("HX-Location", fmt.Sprintf("/fun/quizzes/attempts/%d", attemptID))
	w.WriteHeader(201)
}
