package me

import (
	_ "embed"
	"fmt"
	"net/http"
	"oj/api"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/md"
	"oj/templatehelpers"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type service struct {
	Queries *api.Queries
}

func NewService(q *api.Queries) *service {
	return &service{Queries: q}
}

var (
	//go:embed card.gohtml
	CardContent string
)

func (s *service) Page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(ctx)

	unreadUsers, err := s.Queries.UsersWithUnreadCounts(ctx, l.User.ID)
	if err != nil {
		render.Error(w, fmt.Errorf("UsersWithUnreadCounts: %w", err), http.StatusInternalServerError)
		return
	}

	layout.Layout(l,
		l.User.Username,
		h.Div(h.Style("display:flex;flex-direction:column;gap:1em;"),
			h.Header(
				h.Style("display:flex; justify-content: space-between; align-items: center"),
				h.H1(
					h.Style("text-shadow: -1px 0 white, 0 1px white, 1px 0 white, 0 -1px white;"),
					g.Text("My Page"),
				),
				h.A(
					h.Href("/me/edit"),
					h.Class("nes-btn is-primary"),
					g.Text("Edit My Profile"),
				),
			),

			h.Section(
				card(l.User),
			),

			h.Section(
				h.A(
					g.Attr("onclick", "return confirm('really logout?')"),
					h.Href("/welcome/signout"),
					h.Class("nes-btn"),
					g.Text("Logout"),
				),
				g.If(l.User.Admin,
					h.A(
						h.Href("/admin"),
						h.Class("nes-btn"),
						g.Text("Admin"),
					),
				),
			),

			h.Section(
				h.Style("display:flex; flex-direction:column; gap: 1em"),
				g.Map(unreadUsers, func(friend api.UsersWithUnreadCountsRow) g.Node {
					return unreadFriend(friend)
				}),
			),
			h.Div(
				h.Style("margin-bottom: 50vh"),
			),
		)).Render(w)
}

func card(user api.User) g.Node {
	return h.Div(
		h.ID("card"),
		h.Class("nes-container ghost"),
		g.Attr("hx-swap", "outerHTML"),
		h.Div(
			h.Style("display: flex; align-items: top; gap: 1em;  justify-content: space-between"),
			h.Div(
				h.Style("display: flex; align-items: top; gap: 1em;"),
				h.Figure(
					h.Style("width:80px; height:80px;"),
					h.Img(
						h.Src(user.AvatarURL),
					),
				),
				h.Div(
					h.Style("display: flex; flex-direction: column"),
					h.H1(
						g.Text(user.Username),
					),
					h.Div(
						g.Text("Joined "+templatehelpers.FromNow(user.CreatedAt.Time)+" ago"),
					),
				),
			),
			g.If(user.IsParent,
				h.Div(
					h.Class("nes-badge"),
					h.Span(
						h.Class("is-primary"),
						g.Text("parent"),
					),
				),
			),
		),
		about(user),
	)
}

func about(user api.User) g.Node {
	return g.Group{
		h.H3(
			g.Text("About me"),
		),
		h.Div(
			g.If(user.Bio == "",
				h.P(
					h.Class("nes-text is-disabled"),
					g.Text("nothing here yet"),
				),
			),
			g.If(user.Bio != "", markdown(user.Bio)),
		),
	}
}

func markdown(text string) g.Node {
	return g.Raw(string(md.RenderString(text)))
}

func unreadFriend(friend api.UsersWithUnreadCountsRow) g.Node {
	return h.Div(
		h.Style("padding: 1em; border: 4px solid black; background: red;"),
		h.Div(
			h.Class("nes-container"),
			h.Style("display:flex; flex-direction:column; gap: 1em; background: rgba(255,255,255,.9)"),
			h.A(
				h.Href(fmt.Sprintf("/u/%d/chat", friend.ID)),
				h.Style("display:flex; gap:1em; color: inherit"),
				h.Img(
					h.Src(friend.AvatarURL),
					h.Height("80px"),
					h.Width("80px"),
				),
				h.Div(
					h.Style("display:flex;flex-direction:column"),
					h.H2(
						g.Text(friend.Username),
					),
					h.Div(
						h.Style("display:flex; gap:1em"),
						g.If(friend.UnreadCount == 1,
							h.P(
								h.Class("nes-text is-error"),
								g.Text(fmt.Sprintf("%d unread message", friend.UnreadCount)),
							),
						),
						g.If(friend.UnreadCount > 1,
							h.P(
								h.Class("nes-text is-error"),
								g.Text(fmt.Sprintf("%d unread messages", friend.UnreadCount)),
							),
						),
					),
				),
			),
		),
	)
}
