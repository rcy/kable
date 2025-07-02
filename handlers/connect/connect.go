package connect

import (
	_ "embed"
	"fmt"
	"net/http"
	"oj/api"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
	"oj/worker"
	"strconv"

	"github.com/georgysavva/scany/v2/pgxscan"
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

var (
	//go:embed connect.gohtml
	pageContent string

	//go:embed connection.gohtml
	ConnectionContent string

	t = layout.MustParse(pageContent, ConnectionContent)
)

func (s *service) Connect(w http.ResponseWriter, r *http.Request) {
	lay := layout.FromContext(r.Context())

	connections, err := s.Queries.GetCurrentAndPotentialParentConnections(r.Context(), lay.User.ID)
	if err != nil {
		render.Error(w, fmt.Errorf("GetCurrentAndPotentialParentConnections: %w", err), http.StatusInternalServerError)
		return
	}

	render.Execute(w, t, struct {
		Layout      layout.Data
		Connections []api.GetCurrentAndPotentialParentConnectionsRow
	}{
		Layout:      lay,
		Connections: connections,
	})
}

func (s *service) PutParentFriend(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	currentUser := auth.FromContext(ctx)
	userID, _ := strconv.Atoi(chi.URLParam(r, "userID"))

	user, err := s.Queries.ParentByID(ctx, int64(userID))
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
	go worker.NotifyFriend(friendID)

	connection, err := s.Queries.GetConnection(ctx, api.GetConnectionParams{
		AID: currentUser.ID,
		ID:  user.ID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("HX-Trigger", "connectionChange")
	render.ExecuteNamed(w, t, "connection", connection)
}

func (s *service) DeleteParentFriend(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	currentUser := auth.FromContext(ctx)

	userID, _ := strconv.Atoi(chi.URLParam(r, "userID"))

	user, err := s.Queries.ParentByID(ctx, int64(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	_, err = s.Conn.Exec(ctx, `delete from friends where a_id = $1 and b_id = $2`, currentUser.ID, user.ID)
	if err != nil {
		render.Error(w, fmt.Errorf("delete from friends: %w", err), http.StatusInternalServerError)
		return
	}

	connection, err := s.Queries.GetConnection(ctx, api.GetConnectionParams{
		AID: currentUser.ID,
		ID:  int64(user.ID)},
	)
	if err != nil {
		render.Error(w, fmt.Errorf("GetConnection: %w", err), http.StatusInternalServerError)
		return
	}

	w.Header().Add("HX-Trigger", "connectionChange")
	render.ExecuteNamed(w, t, "connection", connection)
}
