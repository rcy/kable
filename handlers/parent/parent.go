package parent

import (
	"database/sql"
	_ "embed"
	"log"
	"net/http"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
	"oj/models/users"

	"github.com/go-chi/chi/v5"
)

var (
	//go:embed parent.gohtml
	pageContent string

	t = layout.MustParse(pageContent)
)

func Index(w http.ResponseWriter, r *http.Request) {
	l := layout.FromContext(r.Context())

	kids, err := users.Kids(l.User)
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}

	err = t.Execute(w, struct {
		Layout layout.Data
		User   users.User
		Kids   []users.User
	}{
		Layout: l,
		User:   l.User,
		Kids:   kids,
	})
	if err != nil {
		render.Error(w, err.Error(), 500)
	}
}

func CreateKid(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := auth.FromContext(ctx)
	err := r.ParseForm()
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}
	username := r.PostForm.Get("username")

	kid, err := users.FindByUsername(username)
	if err != nil && err != sql.ErrNoRows {
		render.Error(w, err.Error(), 500)
		return
	}
	if kid != nil {
		render.Error(w, "username taken", http.StatusConflict)
		return
	}

	kid, err = users.CreateKid(user, username)
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}

	log.Printf("kid: %v", kid)
	http.Redirect(w, r, "/parent", http.StatusSeeOther)
}

func DeleteKid(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	tx, err := db.DB.Beginx()
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
delete from kids_parents where kid_id = ?;
delete from bios where user_id = ?;
delete from deliveries where sender_id = ? or recipient_id = ?;
delete from gradients where user_id = ?;
delete from kids_codes where user_id = ?;
delete from messages where sender_id = ?;
delete from room_users where user_id = ?;
delete from sessions where user_id = ?;
delete from users where id = ?;
`, userID, userID, userID, userID, userID, userID, userID, userID, userID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tx.Commit()
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func LogoutKid(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	_, err := db.DB.Exec(`delete from sessions where user_id = ?`, userID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
