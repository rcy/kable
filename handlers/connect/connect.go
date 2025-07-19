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

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type service struct {
	Conn    *pgxpool.Pool
	Queries *api.Queries
}

func NewService(q *api.Queries, conn *pgxpool.Pool) *service {
	return &service{Queries: q, Conn: conn}
}

func (s *service) Connect(w http.ResponseWriter, r *http.Request) {
	l := layout.FromContext(r.Context())

	connections, err := s.Queries.GetCurrentAndPotentialParentConnections(r.Context(), l.User.ID)
	if err != nil {
		render.Error(w, fmt.Errorf("GetCurrentAndPotentialParentConnections: %w", err), http.StatusInternalServerError)
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
						g.Text("Connect with Other Parents"),
					),
					h.P(
						g.Text("Once you are connected with another parent, your children will be able to connect with eachother."),
					),
				),
				h.Div(
					h.Style("display:flex; flex-direction:column; gap:1em"),
					g.Map(connections, func(c api.GetCurrentAndPotentialParentConnectionsRow) g.Node {
						if ConnectionStatus(c.RoleIn, c.RoleOut) != "connection" {
							return ConnectionEl(c.User, c.RoleIn, c.RoleOut)
						}
						return h.Div()
					}),
				),
			),
		)).Render(w)
}

func ConnectionStatus(roleIn string, roleOut string) string {
	if roleOut == "" {
		if roleIn == "" {
			return "none"
		} else {
			return "request received"
		}
	} else {
		if roleIn == "" {
			return "request sent"
		} else {
			return "connected"
		}
	}

}

func ConnectionEl(user api.User, roleIn string, roleOut string) g.Node {
	var linkText string
	var button g.Node

	switch ConnectionStatus(roleIn, roleOut) {
	case "none":
		linkText = user.Username
		button = h.Button(
			h.Class("nes-btn is-primary"),
			g.Attr("hx-put", fmt.Sprintf("/connect/friend/%d", user.ID)),
			g.Attr("hx-target", "closest .hx-connection"),
			g.Attr("hx-swap", "outerHTML"),
			h.Div(
				h.Style("display:flex; gap: 1em; align-items: center"),
				h.I(
					h.Class("nes-icon is-small user"),
				),
				h.Span(
					g.Text("Add Friend"),
				),
			),
		)
	case "connected":
		linkText = fmt.Sprintf("%s is your %s", user.Username, roleOut)
		button = h.Button(
			h.Class("nes-btn"),
			g.Attr("hx-delete", fmt.Sprintf("/connect/friend/%d", user.ID)),
			g.Attr("hx-confirm", fmt.Sprintf("Do you want to unfriend %s?", user.Username)),
			g.Attr("hx-target", "closest .hx-connection"),
			g.Attr("hx-swap", "outerHTML"),
			h.Div(
				h.Style("display:flex; gap: 1em; align-items: center"),
				h.Span(
					g.Text("Un"+roleOut),
				),
			),
		)
	case "request received":
		linkText = fmt.Sprintf("%s sent you a friend request", user.Username)
		button = h.Button(
			h.Class("nes-btn is-success"),
			g.Attr("hx-put", fmt.Sprintf("/connect/friend/%d", user.ID)),
			g.Attr("hx-target", "closest .hx-connection"),
			g.Attr("hx-swap", "outerHTML"),
			h.Div(
				h.Style("display:flex; gap: 1em; align-items: center"),
				h.I(
					h.Class("nes-icon is-small check"),
				),
				h.Span(
					g.Text("Accept Request"),
				),
			),
		)
	case "request sent":
		linkText = fmt.Sprintf("you sent a friend request to %s", user.Username)
		button = h.Button(
			h.Class("nes-btn"),
			g.Attr("hx-delete", fmt.Sprintf("/connect/friend/%d", user.ID)),
			g.Attr("hx-target", "closest .hx-connection"),
			g.Attr("hx-swap", "outerHTML"),
			h.Div(
				h.Style("display:flex; gap: 1em; align-items: center"),
				h.I(
					h.Class("nes-icon is-small times"),
				),
				h.Span(
					g.Text("Cancel Request"),
				),
			),
		)
	default:
		linkText = "THIS IS WEIRD"
	}

	return h.Div(
		h.Class("nes-container ghost hx-connection"),
		h.Style("display:flex; gap:1em; justify-content: space-between"),
		h.A(
			h.Href(fmt.Sprintf("/u/%d", user.ID)),
			g.Text(linkText),
		),
		button,
	)
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
	ConnectionEl(connection.User, connection.RoleIn, connection.RoleOut).Render(w)
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
	ConnectionEl(connection.User, connection.RoleIn, connection.RoleOut).Render(w)
}
