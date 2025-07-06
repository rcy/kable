package show

import (
	_ "embed"
	"fmt"
	"net/http"
	"oj/api"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/middleware/quizctx"
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

func (s *service) Router(r chi.Router) {
	r.Use(quizctx.NewService(s.Queries).Provider)
	r.Get("/", s.page)
	r.Patch("/", s.patchQuiz)
	r.Get("/edit", editQuiz)
	r.Post("/toggle-published", s.togglePublished)
	r.Get("/add-question", newQuestion)
	r.Post("/add-question", s.postNewQuestion)
	r.Get("/question/{questionID}/edit", s.editQuestion)
	r.Patch("/question/{questionID}", s.patchQuestion)
}

var (
	//go:embed page.gohtml
	pageContent  string
	pageTemplate = layout.MustParse(pageContent, pageContent)
)

func (s *service) page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(ctx)
	quiz := quizctx.Value(ctx)

	questions, err := s.Queries.QuizQuestions(ctx, quiz.ID)
	if err != nil && err != pgx.ErrNoRows {
		render.Error(w, fmt.Errorf("QuizQuestions: %w", err), http.StatusInternalServerError)
		return
	}

	render.Execute(w, pageTemplate, struct {
		Layout    layout.Data
		Quiz      api.Quiz
		Questions []api.Question
	}{
		Layout:    l,
		Quiz:      quiz,
		Questions: questions,
	})
}

func editQuiz(w http.ResponseWriter, r *http.Request) {
	quiz := quizctx.Value(r.Context())

	render.ExecuteNamed(w, pageTemplate, "quiz-header-edit", quiz)
}

func (s *service) patchQuiz(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	quiz := quizctx.Value(ctx)

	result, err := s.Queries.UpdateQuiz(ctx, api.UpdateQuizParams{
		ID:          quiz.ID,
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
	})
	if err != nil {
		render.Error(w, fmt.Errorf("UpdateQuiz: %w", err), http.StatusInternalServerError)
		return
	}

	render.ExecuteNamed(w, pageTemplate, "quiz-header", result)
}

func (s *service) togglePublished(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	quiz := quizctx.Value(r.Context())

	quiz, err := s.Queries.SetQuizPublished(ctx, api.SetQuizPublishedParams{
		ID:        quiz.ID,
		Published: !quiz.Published,
	})
	if err != nil {
		render.Error(w, fmt.Errorf("SetQuizPublished: %w", err), http.StatusInternalServerError)
		return
	}
	render.ExecuteNamed(w, pageTemplate, "quiz-header", quiz)
}

func newQuestion(w http.ResponseWriter, r *http.Request) {
	quiz := quizctx.Value(r.Context())

	render.ExecuteNamed(w, pageTemplate, "new-question-form", struct{ QuizID int64 }{quiz.ID})
}

func (s *service) editQuestion(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	questionID, _ := strconv.Atoi(chi.URLParam(r, "questionID"))

	quest, err := s.Queries.Question(ctx, int64(questionID))
	if err != nil {
		render.Error(w, fmt.Errorf("Question: %w", err), http.StatusInternalServerError)
		return
	}

	render.ExecuteNamed(w, pageTemplate, "edit-question-form", quest)
}

func (s *service) postNewQuestion(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	quiz := quizctx.Value(r.Context())

	var err error
	var quest api.Question

	if r.FormValue("id") != "" {
		questionID, _ := strconv.Atoi(r.FormValue("id"))
		_, err = s.Queries.Question(ctx, int64(questionID))
		if err != nil && err != pgx.ErrNoRows {
			render.Error(w, fmt.Errorf("Question: %w", err), http.StatusNotFound)
			return
		}
		quest, err = s.Queries.UpdateQuestion(r.Context(), api.UpdateQuestionParams{
			ID:     int64(questionID),
			Text:   r.FormValue("text"),
			Answer: r.FormValue("answer"),
		})
		if err != nil {
			render.Error(w, fmt.Errorf("UpdateQuestion: %w", err), http.StatusInternalServerError)
			return
		}
	} else {
		quest, err = s.Queries.CreateQuestion(r.Context(), api.CreateQuestionParams{
			QuizID: quiz.ID,
			Text:   r.FormValue("text"),
			Answer: r.FormValue("answer"),
		})
		if err != nil {
			render.Error(w, fmt.Errorf("CreateQuestion: %w", err), http.StatusInternalServerError)
			return
		}
	}

	render.ExecuteNamed(w, pageTemplate, "question", quest)
}

func (s *service) patchQuestion(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	questionID, _ := strconv.Atoi(chi.URLParam(r, "questionID"))
	quest, err := s.Queries.Question(ctx, int64(questionID))
	if err != nil {
		render.Error(w, fmt.Errorf("Question: %w", err), http.StatusNotFound)
		return
	}

	quest, err = s.Queries.UpdateQuestion(r.Context(), api.UpdateQuestionParams{
		ID:     quest.ID,
		Text:   r.FormValue("text"),
		Answer: r.FormValue("answer"),
	})
	if err != nil {
		render.Error(w, fmt.Errorf("UpdateQuestion: %w", err), http.StatusInternalServerError)
		return
	}

	render.ExecuteNamed(w, pageTemplate, "question", quest)
}
