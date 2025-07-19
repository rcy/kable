package u

import (
	_ "embed"
	"fmt"
	"net/http"
	"oj/api"
	"oj/handlers/connect"
	"oj/handlers/layout"
	"oj/handlers/me"
	"oj/handlers/render"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func (s *service) Page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(r.Context())

	userID, _ := strconv.Atoi(chi.URLParam(r, "userID"))
	pageUser, err := s.Queries.UserByID(ctx, int64(userID))
	if err != nil {
		if err == pgx.ErrNoRows {
			render.Error(w, fmt.Errorf("UserByID: %w", err), http.StatusNotFound)
			return
		}
		render.Error(w, fmt.Errorf("UserByID: %w", err), http.StatusInternalServerError)
		return
	}

	if l.User.ID == pageUser.ID {
		http.Redirect(w, r, "/me", http.StatusFound)
		return
	}

	// override layout gradient to show the page user's not the request user's
	l.BackgroundGradient = pageUser.Gradient

	connection, err := s.Queries.GetConnection(ctx, api.GetConnectionParams{
		AID: l.User.ID,
		ID:  pageUser.ID,
	})
	if err != nil {
		render.Error(w, fmt.Errorf("GetConnection: %w", err), http.StatusInternalServerError)
		return
	}

	connected := connection.RoleIn != "" && connection.RoleOut != ""

	layout.Layout(l,
		l.User.Username,
		h.Div(
			g.Attr("hx-get", fmt.Sprintf("/u/%d", pageUser.ID)),
			g.Attr("hx-trigger", "connectionChange from:body"),
			g.Attr("hx-target", "body"),
			h.Style("display:flex;flex-direction:column;gap:1em;"),
			g.If(connect.ConnectionStatus(connection.RoleIn, connection.RoleOut) != "connected",
				connect.ConnectionEl(connection.User, connection.RoleIn, connection.RoleOut)),
			me.ProfileEl(pageUser, false),
			g.If(connected, h.A(
				h.Class("nes-btn is-success"),
				h.Href(fmt.Sprintf("/u/%d/chat", pageUser.ID)),
				g.Text("Chat"),
			)),
		)).Render(w)
}
