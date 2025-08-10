package admin

import (
	_ "embed"
	"fmt"
	"net/http"
	"oj/api"
	"oj/gradient"
	"oj/handlers/admin/messages"
	"oj/handlers/admin/middleware/auth"
	"oj/handlers/admin/middleware/background"
	"oj/handlers/admin/quizzes"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/link"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type service struct {
	Conn    *pgxpool.Pool
	Queries *api.Queries
}

func NewService(q *api.Queries, conn *pgxpool.Pool) *service {
	return &service{Queries: q, Conn: conn}
}

func (s *service) Routes() chi.Router {
	r := chi.NewRouter()

	r.Use(auth.EnsureAdmin)
	r.Use(background.Set(gradient.Admin))
	r.Get("/", s.page)
	r.Route("/quizzes", quizzes.NewService(s.Queries).Router)
	r.Route("/messages", messages.NewService(s.Queries).Router)
	return r
}

func (s *service) page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	l := layout.FromContext(r.Context())

	allUsers, err := s.Queries.AllUsers(ctx)
	if err != nil {
		render.Error(w, fmt.Errorf("AllUsers: %w", err), http.StatusInternalServerError)
		return
	}

	layout.Layout(l, "admin", g.Group{
		h.Div(h.Style("display:flex; gap:1em"),
			h.A(
				h.Href("/admin/quizzes"),
				g.Text("quizzes"),
			),
			h.A(
				h.Href("/admin/messages"),
				g.Text("messages"),
			),
		),
		h.Hr(),
		g.Map(allUsers, func(user api.User) g.Node {
			return h.Div(
				h.A(
					h.Href(link.User(user.ID)),
					g.Text(user.Username),
				),
			)
		}),
	}).Render(w)
}
