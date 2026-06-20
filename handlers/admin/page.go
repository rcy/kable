package admin

import (
	"fmt"
	"log/slog"
	"net/http"
	"oj/api"
	"oj/gradient"
	"oj/handlers/admin/messages"
	"oj/handlers/admin/middleware/auth"
	"oj/handlers/admin/middleware/background"
	"oj/handlers/admin/quizzes"
	"oj/handlers/layout"
	"oj/handlers/render"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/riverqueue/river"
	"riverqueue.com/riverui"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

var RiverUIHandler *riverui.Handler

type service struct {
	Conn        *pgxpool.Pool
	Queries     *api.Queries
	RiverClient *river.Client[pgx.Tx]
}

func NewService(q *api.Queries, conn *pgxpool.Pool, riverClient *river.Client[pgx.Tx]) *service {
	return &service{Queries: q, Conn: conn, RiverClient: riverClient}
}

func (s *service) Routes() chi.Router {
	r := chi.NewRouter()

	r.Use(auth.EnsureAdmin)
	r.Use(background.Set(gradient.Admin))
	r.Get("/", s.page)
	r.Route("/quizzes", quizzes.NewService(s.Queries).Router)
	r.Route("/messages", messages.NewService(s.Queries).Router)

	if s.RiverClient != nil {
		logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
		RiverUIHandler, _ = riverui.NewHandler(&riverui.HandlerOpts{
			Endpoints: riverui.NewEndpoints(s.RiverClient, nil),
			Logger:    logger,
			Prefix:    "/admin/river",
		})
		r.Mount("/river/", RiverUIHandler)
	}

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

	layout.Layout(l, "Admin", adminPage(allUsers)).Render(w)
}

func adminPage(users []api.User) g.Node {
	return h.Div(
		h.A(h.Href("/admin/quizzes"), g.Text("quizzes")),
		h.A(h.Href("/admin/messages"), g.Text("messages")),
		h.Hr(),
		g.Map(users, func(u api.User) g.Node {
			return h.Div(
				h.A(h.Href(fmt.Sprintf("/u/%d", u.ID)), g.Text(u.Username)),
			)
		}),
	)
}
