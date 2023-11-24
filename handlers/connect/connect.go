package connect

import (
	_ "embed"
	"net/http"
	"oj/api"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
	"oj/models/users"
	"oj/worker"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var (
	//go:embed connect.gohtml
	pageContent string

	//go:embed connection.gohtml
	ConnectionContent string

	t = layout.MustParse(pageContent, ConnectionContent)
)

type Connection struct {
	users.User
	RoleIn  string `db:"role_in"`
	RoleOut string `db:"role_out"`
}

func (f Connection) Status() string {
	if f.RoleOut == "" {
		if f.RoleIn == "" {
			return "none"
		} else {
			return "request received"
		}
	} else {
		if f.RoleIn == "" {
			return "request sent"
		} else {
			return "connected"
		}
	}
}

func Connect(w http.ResponseWriter, r *http.Request) {
	lay := layout.FromContext(r.Context())

	var connections []Connection
	err := db.DB.Select(&connections, `
select u.*,
       case
           when f1.a_id = $1 then f1.b_role
           else ""
       end as role_out,
       case
           when f2.b_id = $1 then f2.b_role
           else ""
       end as role_in
from users u
left join friends f1 on f1.b_id = u.id and f1.a_id = $1
left join friends f2 on f2.a_id = u.id and f2.b_id = $1
where
  u.id != $1
and
  is_parent = 1
order by role_in desc
limit 128;
`, lay.User.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Execute(w, t, struct {
		Layout      layout.Data
		Connections []Connection
	}{
		Layout:      lay,
		Connections: connections,
	})
}

func PutParentFriend(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	currentUser := auth.FromContext(ctx)
	queries := api.New(db.DB)
	userID, _ := strconv.Atoi(chi.URLParam(r, "userID"))

	var user users.User
	err := db.DB.Get(&user, `select * from users where id = $1`, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if !user.IsParent {
		http.Error(w, "not a parent", http.StatusBadRequest)
		return
	}

	result, err := db.DB.Exec(`insert into friends(a_id, b_id, b_role) values(?,?,'friend')`, currentUser.ID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	friendID, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	go worker.NotifyFriend(friendID)

	connection, err := queries.GetConnection(ctx, api.GetConnectionParams{AID: currentUser.ID, ID: int64(userID)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("HX-Trigger", "connectionChange")
	render.ExecuteNamed(w, t, "connection", connection)
}

func DeleteParentFriend(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	currentUser := auth.FromContext(ctx)
	queries := api.New(db.DB)

	userID, _ := strconv.Atoi(chi.URLParam(r, "userID"))

	var user users.User
	err := db.DB.Get(&user, `select * from users where id = $1`, userID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if !user.IsParent {
		render.Error(w, "not a parent", http.StatusBadRequest)
		return
	}

	_, err = db.DB.Exec(`delete from friends where a_id = $1 and b_id = $2`, currentUser.ID, userID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	connection, err := queries.GetConnection(ctx, api.GetConnectionParams{AID: currentUser.ID, ID: int64(userID)})
	if err != nil {
		render.Error(w, "xxx"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("HX-Trigger", "connectionChange")
	render.ExecuteNamed(w, t, "connection", connection)
}
