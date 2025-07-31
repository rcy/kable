package connectkids

import (
	_ "embed"
	"fmt"
	"net/http"
	"oj/api"
	"oj/handlers/connect"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
	"oj/services/reachable"
	"oj/worker"
	"strconv"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type service struct {
	Conn    *pgxpool.Pool
	Queries *api.Queries
}

func NewService(conn *pgxpool.Pool, queries *api.Queries) *service {
	return &service{
		Conn:    conn,
		Queries: queries,
	}
}

func (s *service) KidConnect(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(ctx)

	kids, err := reachable.ReachableKids(ctx, s.Queries, l.User.ID)
	if err != nil {
		render.Error(w, fmt.Errorf("ReachableKids: %w", err), http.StatusInternalServerError)
		return
	}

	layout.Layout(l,
		"Connect",
		h.Div(
			h.StyleEl(g.Raw(".htmx-request { opacity: .5; }")),
			h.Div(
				h.Style("display:flex;flex-direction:column; gap:1em"),
				h.Div(
					h.Class("nes-container ghost"),
					h.H1(
						g.Text("Connect with Other Kids"),
					),
					h.P(
						g.Text("You can connect with other kids when your parents are also connected."),
					),
				),
				h.Div(
					h.Style("display:flex; flex-direction:column; gap:1em"),
					g.Map(kids, func(c api.GetConnectionRow) g.Node {
						if connect.ConnectionStatus(c.RoleIn, c.RoleOut) != "connection" {
							return connect.ConnectionEl(c.User, c.RoleIn, c.RoleOut)
						}
						return h.Div()
					}),
				),
			),
		)).Render(w)
}

func (s *service) PutKidFriend(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	currentUser := auth.FromContext(ctx)
	userID, _ := strconv.Atoi(chi.URLParam(r, "userID"))

	user, err := s.Queries.UserByID(ctx, int64(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var friendID int64
	err = pgxscan.Get(ctx, s.Conn, &friendID, `insert into friends(a_id, b_id, b_role) values($1,$2,'friend') returning id`, currentUser.ID, user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	go worker.NotifyKidFriend(friendID)

	connection, err := s.Queries.GetConnection(ctx, api.GetConnectionParams{
		AID: currentUser.ID,
		ID:  user.ID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("HX-Trigger", "connectionChange")
	connect.ConnectionEl(connection.User, connection.RoleIn, connection.RoleOut).Render(w)
}

func (s *service) DeleteKidFriend(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	currentUser := auth.FromContext(ctx)
	userID, _ := strconv.Atoi(chi.URLParam(r, "userID"))

	user, err := s.Queries.UserByID(ctx, int64(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	_, err = s.Conn.Exec(ctx, `delete from friends where a_id = $1 and b_id = $2`, currentUser.ID, user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	connection, err := s.Queries.GetConnection(ctx, api.GetConnectionParams{
		AID: currentUser.ID,
		ID:  user.ID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("HX-Trigger", "connectionChange")
	connect.ConnectionEl(connection.User, connection.RoleIn, connection.RoleOut).Render(w)
}
