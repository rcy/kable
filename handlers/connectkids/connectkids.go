package connectkids

import (
	_ "embed"
	"net/http"
	"oj/api"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
	"oj/services/reachable"
	"oj/worker"
	"strconv"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type service struct {
	Conn    *pgxpool.Conn
	Queries *api.Queries
}

func NewService(conn *pgxpool.Conn, queries *api.Queries) *service {
	return &service{
		Conn:    conn,
		Queries: queries,
	}
}

var (
	//go:embed connectkids.gohtml
	pageContent string
	t           = layout.MustParse(pageContent)
)

func (s *service) KidConnect(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(ctx)

	connections, err := reachable.ReachableKids(ctx, s.Queries, l.User.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Execute(w, t, struct {
		Layout      layout.Data
		Connections []api.GetConnectionRow
	}{
		Layout:      l,
		Connections: connections,
	})
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
	err = pgxscan.Get(ctx, s.Conn, &friendID, `insert into friends(a_id, b_id, b_role) values(?,?,'friend') returning id`, currentUser.ID, user.ID)
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

	render.ExecuteNamed(w, t, "connection", connection)
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

	render.ExecuteNamed(w, t, "connection", connection)
}
