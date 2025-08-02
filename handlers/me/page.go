package me

import (
	_ "embed"
	"fmt"
	"net/http"
	"oj/api"
	"oj/avatar"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/text"
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

	friends, err := s.Queries.GetConnections(ctx, l.User.ID)
	if err != nil {
		render.Error(w, fmt.Errorf("GetConnections: %w", err), http.StatusInternalServerError)
		return
	}

	quizzes, err := s.Queries.AllUserQuizzes(ctx, l.User.ID)
	if err != nil {
		render.Error(w, fmt.Errorf("GetConnections: %w", err), http.StatusInternalServerError)
		return
	}

	layout.Layout(l,
		l.User.Username,
		h.Div(h.Style("display:flex;flex-direction:column;gap:1em;"),
			h.Div(h.Style("display:flex; gap: 1em; justify-content: flex-end"),
				g.If(l.User.IsParent,
					g.Group{
						h.A(h.Href("/parent"), h.Class("nes-btn"), g.Text("Family")),
						h.A(h.Href("/connect"), h.Class("nes-btn"), g.Text("Friends")),
					},
				),
				g.If(!l.User.IsParent,
					h.A(h.Href("/connectkids"), h.Class("nes-btn"), g.Text("Friends")),
				),
			),

			h.Section(ProfileEl(l.User, true)),

			h.Section(
				h.Style("display:flex; flex-direction:column; gap: 1em"),
				g.Map(unreadUsers, func(friend api.UsersWithUnreadCountsRow) g.Node {
					return unreadFriend(friend)
				}),
			),

			h.Section(
				h.Div(h.Style("display:flex; flex-wrap: wrap; gap: 38px"),
					g.Map(friends, func(friend api.User) g.Node {
						return h.A(h.Href(fmt.Sprintf("/u/%d/chat", friend.ID)),
							UserCard(friend, true))
					})),
			),

			QuizzesEl(l.User.ID, quizzes),
			h.Div(
				h.Style("margin-bottom: 50vh"),
			),

			h.Section(
				h.Style("display:flex; gap: 1em"),
				h.A(
					g.Attr("onclick", "return confirm('really logout?')"),
					h.Href("/welcome/signout"),
					h.Class("nes-btn"),
					g.Text("Logout"),
				),
				g.If(l.User.Admin,
					h.A(
						h.Href("/admin"),
						h.Class("nes-btn is-error"),
						g.Text("Admin"),
					),
				),
			),
		)).Render(w)
}

func ProfileEl(user api.User, canEdit bool) g.Node {
	return h.Div(
		h.ID("card"),
		h.Class("nes-container ghost"),
		g.Attr("hx-swap", "outerHTML"),
		h.Div(
			h.Style("display: flex; align-items: top; gap: 1em;  justify-content: space-between"),
			h.Div(
				h.Style("display: flex; align-items: top; gap: 1em;"),
				UserCard(user, true),
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
			g.If(canEdit, h.Div(h.A(
				h.Href("/me/edit"),
				h.Class("nes-btn is-primary"),
				g.Text("Edit My Profile"),
			))),
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
					h.Src(friend.Avatar.URL()),
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

func UserCard(friend api.User, withUsername bool) g.Node {
	return h.Figure(
		h.Div(h.Style("background: "+string(friend.Gradient.Render())),
			h.Img(h.Width("128px"), h.Src(friend.Avatar.URL()))),
		g.If(withUsername,
			h.Div(h.Style("background: black; color: white"),
				g.Text(text.Shorten(friend.Username, 8)))))
}

func IfElse(condition bool, n1 g.Node, n2 g.Node) g.Node {
	if condition {
		return n1
	}
	return n2
}

func QuizCard(quiz api.Quiz) g.Node {
	avi := avatar.New(fmt.Sprint(quiz.ID), avatar.IconsStyle)

	return IfElse(quiz.Published,
		h.Div(h.Style("display:flex; justify-content:space-between"),
			h.Div(h.Style("display:flex; gap:1em; align-items: center;"),
				h.Img(h.Width("32px"), h.Src(avi.URL())),
				h.A(
					h.Href(fmt.Sprintf("/u/%d/quizzes/%d/view", quiz.UserID, quiz.ID)),
					g.Text(quiz.Name),
				)),
		),
		// else
		h.Div(h.Style("display:flex; justify-content:space-between"),
			h.Div(h.Style("display:flex; gap:1em; align-items: center;"),
				h.A(
					h.Href(fmt.Sprintf("/u/%d/quizzes/%d", quiz.UserID, quiz.ID)),
					g.Text("Edit "+quiz.Name),
				)),
		),
	)
}

func QuizzesEl(userID int64, quizzes []api.Quiz) g.Node {
	return h.Section(h.ID("quizzes"),
		h.Div(h.Class("nes-container ghost"), h.Style("display:flex; flex-direction:column; gap:1em"),
			h.Div(h.Style("display:flex; justify-content:space-between"),
				h.H1(g.Text("Quizzes")),
				g.If(userID != 0,
					h.A(h.Class("nes-btn"), h.Href(fmt.Sprintf("/u/%d/quizzes/create", userID)), g.Text("Create Quiz")))),
			g.Map(quizzes, func(quiz api.Quiz) g.Node {
				return QuizCard(quiz)
			})),
	)
}
