package view

import (
	_ "embed"
	"errors"
	"fmt"
	"net/http"
	"oj/api"
	"oj/avatar"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
	"oj/internal/middleware/quizctx"
	"oj/templatehelpers"

	"github.com/go-chi/chi/v5"

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
	r.Post("/attempt", s.createAttempt)
}

func (s *service) page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(ctx)
	user := auth.FromContext(ctx)
	quiz := quizctx.Value(ctx)

	author, err := s.Queries.UserByID(ctx, quiz.UserID)
	if err != nil {
		render.Error(w, fmt.Errorf("UserByID: %w", err), http.StatusInternalServerError)
		return
	}

	questions, err := s.Queries.QuizQuestions(ctx, quiz.ID)
	if err != nil {
		render.Error(w, fmt.Errorf("QuizQuestions: %w", err), http.StatusInternalServerError)
		return
	}

	if len(questions) == 0 {
		render.Error(w, errors.New("quiz has no questions"), http.StatusInternalServerError)
		return
	}

	layout.Layout(l,
		l.User.Username,
		h.Div(
			h.Style("display:flex; flex-direction: column; gap: 2em; margin-bottom: 25vh"),
			g.If(author.ID == user.ID,
				h.Div(
					h.Class("nes-container"),
					h.Style("background: #92cc41"),
					h.Div(h.Style("display:flex; justify-content:space-between"),
						h.P(
							g.Text("This quiz is shared with your friends!"),
						),
						h.Button(
							g.Attr("hx-post", fmt.Sprintf("/u/%d/quizzes/%d/unpublish", quiz.UserID, quiz.ID)),
							h.Class("nes-btn xis-warning"),
							g.Text("Edit"),
						),
					),
				),
			),
			h.Div(h.Class("nes-container ghost"),
				QuizTitleEl(author, quiz),
				h.P(),
				h.Div(h.Style("display:flex; gap:1em"),
					h.Button(
						g.Attr("hx-post", fmt.Sprintf("/u/%d/quizzes/%d/view/attempt", quiz.UserID, quiz.ID)),
						h.Class("nes-btn is-primary"),
						g.Text("Take the Quiz!"),
					)),
			))).Render(w)
}

func QuizTitleEl(author api.User, quiz api.Quiz) g.Node {
	avi := avatar.New(fmt.Sprint(quiz.ID), avatar.IconsStyle)
	return h.Div(h.Style("display:flex; gap:1em; align-items:start"),
		h.Img(h.Width("88px"), h.Src(avi.URL())),
		h.Div(
			h.H1(
				g.Text(quiz.Name),
			),
			g.Text("by "),
			h.A(
				h.Href(fmt.Sprintf("/u/%d", author.ID)),
				h.Img(h.Width("32px"), h.Src(author.Avatar.URL())),
				g.Text(author.Username),
			),
			g.Text(" "+templatehelpers.FromNow(quiz.CreatedAt.Time)+" ago"),
		),
	)
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

	w.Header().Add("HX-Redirect", fmt.Sprintf("/u/%d/quizzes/%d/attempts/%d", quiz.UserID, quiz.ID, attempt.ID))
	w.WriteHeader(201)
}
