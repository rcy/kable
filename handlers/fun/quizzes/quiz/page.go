package quiz

import (
	_ "embed"
	"fmt"
	"net/http"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/models/attempts"
	"oj/models/question"
	"oj/models/quizzes"
	"oj/models/users"

	"github.com/go-chi/chi/v5"
)

var (
	//go:embed page.gohtml
	pageContent  string
	pageTemplate = layout.MustParse(pageContent)
)

func Router(r chi.Router) {
	r.Use(quizzes.Provider)
	r.Get("/", page)
}

func page(w http.ResponseWriter, r *http.Request) {
	l, err := layout.FromRequest(r)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	quiz := quizzes.FromContext(r.Context())

	questions, err := quiz.FindQuestions()
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(questions) == 0 {
		render.Error(w, "quiz has no questions", http.StatusInternalServerError)
		return
	}

	render.Execute(w, pageTemplate, struct {
		Layout           layout.Data
		Quiz             quizzes.Quiz
		Questions        []question.Question
		CreateAttemptURL string
	}{
		Layout:           l,
		Quiz:             quiz,
		Questions:        questions,
		CreateAttemptURL: r.URL.Path + "/attempt",
	})
}

func CreateAttempt(w http.ResponseWriter, r *http.Request) {
	user := users.FromContext(r.Context())
	quiz := quizzes.FromContext(r.Context())

	attempt, err := attempts.Create(quiz.ID, user.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("HX-Redirect", fmt.Sprintf("/fun/quizzes/attempts/%d", attempt.ID))
	w.WriteHeader(201)
}
