package quizzes

import (
	_ "embed"
	"fmt"
	"net/http"
	"oj/api"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/link"

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

	layout.Layout(l, "quizzes",
		h.Div(
			h.Div(
				h.Style("display:flex; justify-content:space-between; align-items:center"),
				h.H1(
					g.Text("quizzes"),
				),
			),
			g.Map(allQuizzes, func(q api.Quiz) g.Node {
				return h.Div(
					h.Class("nes-container ghost"),
					h.A(
						h.Href(link.Quiz(q)),
						g.Text(q.Name),
					),
				)
			}),
		),
	).Render(w)
}
