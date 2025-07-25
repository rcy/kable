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

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
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

var (
	//go:embed page.gohtml
	pageContent  string
	pageTemplate = layout.MustParse(pageContent, pageContent)
)

func (s *service) page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	l := layout.FromContext(r.Context())

	allUsers, err := s.Queries.AllUsers(ctx)
	if err != nil {
		render.Error(w, fmt.Errorf("AllUsers: %w", err), http.StatusInternalServerError)
		return
	}

	render.Execute(w, pageTemplate, struct {
		Layout layout.Data
		Users  []api.User
	}{
		Layout: l,
		Users:  allUsers,
	})
}
