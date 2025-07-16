package show

import (
	_ "embed"
	"fmt"
	"net/http"
	"oj/api"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
	"oj/internal/middleware/quizctx"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
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

func (s *service) page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(ctx)
	quiz := quizctx.Value(ctx)

	questions, err := s.Queries.QuizQuestions(ctx, quiz.ID)
	if err != nil && err != pgx.ErrNoRows {
		render.Error(w, fmt.Errorf("QuizQuestions: %w", err), http.StatusInternalServerError)
		return
	}

	layout.Layout(l, quiz.Name,
		h.Div(
			h.Style("display:flex; flex-direction: column; gap: 2em; margin-bottom: 25vh"),
			quizHeader(quiz),
			h.Div(
				h.ID("questions"),
				h.Style("display:flex; flex-direction: column; gap:1em"),
				g.Map(questions, func(quest api.Question) g.Node {
					return question(quiz.UserID, quest)
				}),
			),
			h.Button(
				g.Attr("hx-get", fmt.Sprintf("/u/%d/quizzes/%d/add-question", quiz.UserID, quiz.ID)),
				g.Attr("hx-target", "#questions"),
				g.Attr("hx-swap", "beforeend"),
				h.Class("nes-btn is-primary"),
				g.Text("add question"),
			),
		)).Render(w)
}

func quizHeader(quiz api.Quiz) g.Node {
	return h.Div(
		h.Class("hx-target"),
		g.If(!quiz.Published,
			h.Div(
				h.Class("nes-container"),
				h.Style("background: #f7d51d; display:flex; justify-content:space-between"),
				h.P(
					g.Text("This quiz is not published — so only you can see it."),
				),
				h.Div(
					h.Button(
						g.Attr("hx-post", fmt.Sprintf("/u/%d/quizzes/%d/toggle-published", quiz.UserID, quiz.ID)),
						g.Attr("hx-target", "closest .hx-target"),
						g.Attr("hx-swap", "outerHTML"),
						h.Class("nes-btn is-warning"),
						g.Text("Publish"),
					),
				),
			)),
		g.If(quiz.Published,
			h.Div(
				h.Class("nes-container"),
				h.Style("background: #92cc41; display:flex; justify-content:space-between"),
				h.P(
					g.Text("This quiz is published!"),
				),
				h.Div(
					h.Button(
						g.Attr("hx-post", fmt.Sprintf("/u/%d/quizzes/%d/toggle-published", quiz.UserID, quiz.ID)),
						g.Attr("hx-target", "closest .hx-target"),
						g.Attr("hx-swap", "outerHTML"),
						h.Class("nes-btn is-success"),
						g.Text("Unpublish"),
					),
				),
			)),
		h.Div(
			h.Class("nes-container ghost"),
			h.Style("margin-top: 1em; display:flex;justify-content:space-between"),
			h.Div(
				h.Div(
					h.Style("display:flex; justify-content:space-between"),
					h.H1(g.Text(quiz.Name)),
				),
			),
			h.Div(
				h.Button(
					g.Attr("hx-get", fmt.Sprintf("/u/%d/quizzes/%d/edit", quiz.UserID, quiz.ID)),
					g.Attr("hx-swap", "outerHTML"),
					g.Attr("hx-target", "closest .hx-target"),
					h.Class("nes-btn"),
					g.Text("edit"),
				),
			),
		),
	)
}

func editQuiz(w http.ResponseWriter, r *http.Request) {
	quiz := quizctx.Value(r.Context())

	h.Div(
		h.Class("nes-container is-dark"),
		h.Form(
			g.Attr("hx-patch", fmt.Sprintf("/u/%d/quizzes/%d", quiz.UserID, quiz.ID)),
			g.Attr("hx-swap", "outerHTML"),
			g.Attr("hx-target", "closest div"),
			h.Div(
				h.Class("nes-field"),
				h.Label(
					h.For("name_field"),
					g.Text("Quiz Name"),
				),
				h.Input(
					h.AutoFocus(),
					h.Type("text"),
					h.Name("name"),
					h.ID("name_field"),
					h.Class("nes-input"),
					h.Value(quiz.Name),
				),
			),
			h.Button(
				h.Class("nes-btn is-primary"),
				g.Text("save"),
			),
			h.A(
				h.Href(fmt.Sprintf("/admin/quizzes/%d", quiz.ID)),
				h.Class("nes-btn"),
				g.Text("cancel"),
			),
		),
	).Render(w)
}

func (s *service) patchQuiz(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	quiz := quizctx.Value(ctx)

	quiz, err := s.Queries.UpdateQuiz(ctx, api.UpdateQuizParams{
		ID:          quiz.ID,
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
	})
	if err != nil {
		render.Error(w, fmt.Errorf("UpdateQuiz: %w", err), http.StatusInternalServerError)
		return
	}

	quizHeader(quiz).Render(w)
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
	quizHeader(quiz).Render(w)
}

func newQuestion(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := auth.FromContext(ctx)
	quiz := quizctx.Value(r.Context())

	h.Div(
		h.Class("nes-container is-dark"),
		h.Form(
			g.Attr("hx-post", fmt.Sprintf("/u/%d/quizzes/%d/add-question", user.ID, quiz.ID)),
			g.Attr("hx-swap", "outerHTML"),
			g.Attr("hx-target", "closest div"),
			h.Div(
				h.Class("nes-field"),
				h.Label(
					h.For("text_field"),
					g.Text("Question"),
				),
				h.Input(
					h.AutoFocus(),
					h.Type("text"),
					h.Name("text"),
					h.ID("text_field"),
					h.Class("nes-input"),
				),
			),
			h.Div(
				h.Class("nes-field"),
				h.Label(
					h.For("answer_field"),
					g.Text("Answer"),
				),
				h.Input(
					h.Type("text"),
					h.Name("answer"),
					h.ID("answer_field"),
					h.Class("nes-input"),
				),
			),
			h.Button(
				h.Class("nes-btn is-primary"),
				g.Text("save"),
			),
		),
	).Render(w)
}

func (s *service) editQuestion(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := auth.FromContext(ctx)
	questionID, _ := strconv.Atoi(chi.URLParam(r, "questionID"))

	quest, err := s.Queries.Question(ctx, int64(questionID))
	if err != nil {
		render.Error(w, fmt.Errorf("Question: %w", err), http.StatusInternalServerError)
		return
	}

	h.Div(
		h.Class("nes-container is-dark"),
		h.Form(
			g.Attr("hx-patch", fmt.Sprintf("/u/%d/quizzes/%d/question/%d", user.ID, quest.QuizID, quest.ID)),
			g.Attr("hx-swap", "outerHTML"),
			g.Attr("hx-target", "closest div"),
			h.Div(
				h.Class("nes-field"),
				h.Label(
					h.For("text_field"),
					g.Text("Question"),
				),
				h.Input(
					h.AutoFocus(),
					h.Type("text"),
					h.Name("text"),
					h.ID("text_field"),
					h.Class("nes-input"),
					h.Value(quest.Text),
				),
			),
			h.Div(
				h.Class("nes-field"),
				h.Label(
					h.For("answer_field"),
					g.Text("Answer"),
				),
				h.Input(
					h.Type("text"),
					h.Name("answer"),
					h.ID("answer_field"),
					h.Class("nes-input"),
					h.Value(quest.Answer),
				),
			),
			h.Button(
				h.Class("nes-btn"),
				g.Text("save"),
			),
		)).Render(w)
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
	question(quiz.UserID, quest).Render(w)
}

func question(userID int64, quest api.Question) g.Node {
	return h.Div(
		h.Class("nes-container ghost"),
		h.Style("display:flex; justify-content:space-between"),
		h.Div(
			h.Div(
				g.Text("Q: "+quest.Text),
			),
			h.Div(
				g.Text("A: "+quest.Answer),
			),
		),
		h.Button(
			g.Attr("hx-get", fmt.Sprintf("/u/%d/quizzes/%d/question/%d/edit", userID, quest.QuizID, quest.ID)),
			g.Attr("hx-swap", "outerHTML"),
			g.Attr("hx-target", "closest div"),
			h.Class("nes-btn"),
			g.Text("edit"),
		),
	)
}

func (s *service) patchQuestion(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := auth.FromContext(ctx)

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

	question(user.ID, quest).Render(w)
}
