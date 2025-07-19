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
	r.Delete("/", s.deleteQuiz)
	r.Get("/edit", editQuiz)
	r.Post("/publish", s.publish)
	r.Post("/unpublish", s.unpublish)
	r.Get("/add-question", newQuestion)
	r.Post("/add-question", s.postNewQuestion)
	r.Get("/question/{questionID}/edit", s.editQuestion)
	r.Patch("/question/{questionID}", s.patchQuestion)
}

func (s *service) page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(ctx)
	user := auth.FromContext(ctx)
	quiz := quizctx.Value(ctx)

	questions, err := s.Queries.QuizQuestions(ctx, quiz.ID)
	if err != nil && err != pgx.ErrNoRows {
		render.Error(w, fmt.Errorf("QuizQuestions: %w", err), http.StatusInternalServerError)
		return
	}

	layout.Layout(l, quiz.Name,
		h.Div(
			h.Style("display:flex; flex-direction: column; gap: 2em; margin-bottom: 25vh"),
			h.Div(
				h.Class("nes-container"),
				h.Style("background: #f7d51d; display:flex; flex-direction:column"),
				h.P(
					h.B(g.Text("This quiz is not shared, only you can see it.")),
				),
				h.P(
					g.Text("Share it when you are ready to let your friends take the quiz!"),
				),
				h.Div(h.Style("display:flex; justify-content:space-between"),
					h.Button(
						g.Attr("hx-post", fmt.Sprintf("/u/%d/quizzes/%d/publish", quiz.UserID, quiz.ID)),
						h.Class("nes-btn is-success"),
						g.Text("Share with friends"),
					),
					h.Button(
						g.Attr("hx-delete", fmt.Sprintf("/u/%d/quizzes/%d", user.ID, quiz.ID)),
						g.Attr("hx-confirm", "Are you sure you want to delete this quiz?"),
						h.Href(fmt.Sprintf("/u/%d/quizzes/%d", user.ID, quiz.ID)),
						h.Class("nes-btn is-error"),
						g.Text("Delete quiz"),
					),
				),
			),

			quizHeader(quiz),
			h.Div(
				h.ID("questions"),
				h.Style("display:flex; flex-direction: column; gap:1em"),
				g.Map(questions, func(quest api.Question) g.Node {
					return question(quiz.Published, quiz.UserID, quest)
				}),
			),
			g.If(!quiz.Published,
				h.Div(
					h.Button(
						g.Attr("hx-get", fmt.Sprintf("/u/%d/quizzes/%d/add-question", quiz.UserID, quiz.ID)),
						g.Attr("hx-target", "#questions"),
						g.Attr("hx-swap", "beforeend"),
						h.Class("nes-btn"),
						g.Text("add a question"),
					)),
			),
		)).Render(w)
}

func quizHeader(quiz api.Quiz) g.Node {
	return h.Div(h.Class("hx-target"),
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
				g.If(!quiz.Published,
					h.Button(
						g.Attr("hx-get", fmt.Sprintf("/u/%d/quizzes/%d/edit", quiz.UserID, quiz.ID)),
						g.Attr("hx-swap", "outerHTML"),
						g.Attr("hx-target", "closest .hx-target"),
						h.Class("nes-btn"),
						g.Text("edit name"),
					),
				),
			),
		),
	)
}

func editQuiz(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	quiz := quizctx.Value(ctx)

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
			h.Div(h.Style("display:flex;justify-content:space-between"),
				h.Button(
					h.Class("nes-btn is-primary"),
					g.Text("save"),
				),
				// h.A(
				// 	h.Href(fmt.Sprintf("/u/%d/quizzes/%d", quiz.UserID, quiz.ID)),
				// 	h.Class("nes-btn"),
				// 	g.Text("cancel"),
				// ),
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

func (s *service) deleteQuiz(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	quiz := quizctx.Value(ctx)
	user := auth.FromContext(ctx)

	err := s.Queries.DeleteQuiz(ctx, api.DeleteQuizParams{
		QuizID: quiz.ID,
		UserID: user.ID,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("DeleteQuiz: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Add("HX-Redirect", fmt.Sprintf("/u/%d#quizzes", user.ID))
}

func (s *service) publish(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	quiz := quizctx.Value(r.Context())

	count, err := s.Queries.QuestionCount(ctx, quiz.ID)
	if err != nil {
		http.Error(w, "QuestionCount:"+err.Error(), http.StatusInternalServerError)
		return
	}

	if count < 1 {
		http.Error(w, "Quiz has no questions!", http.StatusBadRequest)
		return
	}

	quiz, err = s.Queries.SetQuizPublished(ctx, api.SetQuizPublishedParams{
		ID:        quiz.ID,
		Published: true,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("SetQuizPublished: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Add("HX-Redirect", fmt.Sprintf("/u/%d/quizzes/%d/view", quiz.UserID, quiz.ID))
}

func (s *service) unpublish(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	quiz := quizctx.Value(r.Context())

	quiz, err := s.Queries.SetQuizPublished(ctx, api.SetQuizPublishedParams{
		ID:        quiz.ID,
		Published: false,
	})
	if err != nil {
		render.Error(w, fmt.Errorf("SetQuizPublished: %w", err), http.StatusInternalServerError)
		return
	}
	w.Header().Add("HX-Redirect", fmt.Sprintf("/u/%d/quizzes/%d", quiz.UserID, quiz.ID))
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
	question(quiz.Published, quiz.UserID, quest).Render(w)
}

func question(published bool, userID int64, quest api.Question) g.Node {
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
		g.If(!published,
			h.Button(
				g.Attr("hx-get", fmt.Sprintf("/u/%d/quizzes/%d/question/%d/edit", userID, quest.QuizID, quest.ID)),
				g.Attr("hx-swap", "outerHTML"),
				g.Attr("hx-target", "closest div"),
				h.Class("nes-btn"),
				g.Text("edit question"),
			),
		),
	)
}

func (s *service) patchQuestion(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := auth.FromContext(ctx)
	quiz := quizctx.Value(ctx)

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

	question(quiz.Published, user.ID, quest).Render(w)
}
