package quizzes

import (
	"fmt"
	"net/http"
	"oj/api"
	"oj/handlers/layout"
	"oj/handlers/render"

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
	r.Get("/", s.page)
}

func (s *service) page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(ctx)

	allQuizzes, err := s.Queries.AllQuizzes(ctx)
	if err != nil && err != pgx.ErrNoRows {
		render.Error(w, fmt.Errorf("AllQuizzes: %w", err), http.StatusInternalServerError)
		return
	}

	layout.Layout(l, "Quizzes", quizzesPage(allQuizzes)).Render(w)
}

func quizzesPage(quizzes []api.Quiz) g.Node {
	return h.Div(
		h.Div(h.Style("display:flex; justify-content:space-between; align-items:center"),
			h.H1(g.Text("quizzes")),
			h.A(h.Href("/admin/quizzes/create"), h.Class("nes-btn"), g.Text("create quiz")),
		),
		g.Map(quizzes, func(q api.Quiz) g.Node {
			return h.Div(h.Class("nes-container ghost"),
				h.A(h.Href(fmt.Sprintf("/admin/quizzes/%d", q.ID)), g.Text(q.Name)),
			)
		}),
	)
}
