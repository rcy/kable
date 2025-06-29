package editme

import (
	_ "embed"
	"fmt"
	"net/http"
	"oj/api"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
)

var (
	//go:embed my_page_edit.gohtml
	pageContent        string
	myPageEditTemplate = layout.MustParse(pageContent, AvatarContent)
)

func (s *service) MyPageEdit(w http.ResponseWriter, r *http.Request) {
	l := layout.FromContext(r.Context())

	d := struct {
		Layout layout.Data
		User   api.User
	}{
		Layout: l,
		User:   l.User,
	}

	render.Execute(w, myPageEditTemplate, d)
}

func (s *service) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := auth.FromContext(ctx)
	username := r.FormValue("username")
	bio := r.FormValue("bio")

	_, err := s.Conn.Exec(ctx, "update users set username=$1, bio=$2 where id=$3", username, bio, user.ID)
	if err != nil {
		render.Error(w, fmt.Errorf("update users: %w", err), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/me", http.StatusSeeOther)
}
