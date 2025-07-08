package editme

import (
	_ "embed"
	"fmt"
	"net/http"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/middleware/auth"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func (s *service) MyPageEdit(w http.ResponseWriter, r *http.Request) {
	l := layout.FromContext(r.Context())

	layout.Layout(l,
		l.User.Username,
		h.Section(
			h.Class("nes-container is-dark"),
			g.Attr("hx-swap", "outerHTML"),
			g.Attr("hx-target", "closest section"),
			h.H3(
				g.Text("Edit My Profile"),
			),
			h.Div(
				h.Style("margin-bottom: 2em"),
				h.Div(
					h.Style("display:flex;width:100%; align-items:center; justify-content:space-between; gap: 1em"),
					changeableAvatar(l.User),
				),
			),
			h.Form(
				h.Method("post"),
				h.Style("display: flex; flex-direction: column; gap: 2em"),
				h.Div(
					h.Class("nes-field"),
					h.Label(
						h.For("username-field"),
						g.Text("My Username"),
					),
					h.Input(
						h.ID("username-field"),
						h.Class("nes-input"),
						h.Name("username"),
						h.Type("text"),
						h.Value(l.User.Username),
					),
				),
				h.Div(
					h.Class("nes-field"),
					h.Label(
						h.For("username-field"),
						g.Text("About Me"),
					),
					h.Textarea(
						h.Name("bio"),
						h.Class("nes-textarea"),
						h.Placeholder("Write something about yourself..."),
						h.Rows("10"),
						g.Text(l.User.Bio),
					),
				),
				h.Div(
					h.Style("display:flex; justify-content:space-between;"),
					h.Button(
						h.Type("submit"),
						h.Class("nes-btn is-primary"),
						g.Text("save"),
					),
					h.A(
						h.Href("/me"),
						h.Class("nes-btn"),
						g.Text("cancel"),
					),
				),
			),
		)).Render(w)
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
