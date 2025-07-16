package create

import (
	_ "embed"
	"fmt"
	"net/http"
	"oj/api"
	"oj/handlers/layout"
	"oj/handlers/render"

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
	r.Get("/", page)
	r.Post("/", s.post)
}

func page(w http.ResponseWriter, r *http.Request) {
	l := layout.FromContext(r.Context())

	layout.Layout(l,
		"Create New Quiz",
		h.Div(
			h.H1(
				g.Text("create new quiz"),
			),
			h.Div(
				h.Class("nes-container is-dark"),
				h.Form(
					h.Method("post"),
					h.Action(fmt.Sprintf("/u/%d/quizzes/create", l.User.ID)),
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
						),
					),
					// h.Div(
					// 	h.Class("nes-field"),
					// 	h.Label(
					// 		h.For("description_field"),
					// 		g.Text("Description"),
					// 	),
					// 	h.Input(
					// 		h.AutoFocus(),
					// 		h.Type("text"),
					// 		h.Name("description"),
					// 		h.ID("description_field"),
					// 		h.Class("nes-input"),
					// 	),
					// ),
					h.Button(
						h.Class("nes-btn is-primary"),
						g.Text("create"),
					),
					h.A(
						h.Href("/me"),
						h.Class("nes-btn"),
						g.Text("cancel"),
					),
				),
			),
		)).Render(w)
}

func (s *service) post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(ctx)

	quiz, err := s.Queries.CreateQuiz(ctx, api.CreateQuizParams{
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		UserID:      l.User.ID,
	})
	if err != nil {
		render.Error(w, fmt.Errorf("CreateQuiz: %w", err), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/u/%d/quizzes/%d", l.User.ID, quiz.ID), http.StatusSeeOther)
}
