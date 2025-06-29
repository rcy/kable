package u

import (
	"database/sql"
	_ "embed"
	"fmt"
	"net/http"
	"oj/api"
	"oj/handlers/connect"
	"oj/handlers/layout"
	"oj/handlers/me"
	"oj/handlers/render"
	"oj/services/background"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var (
	//go:embed page.gohtml
	pageContent string

	pageTemplate = layout.MustParse(pageContent, me.CardContent, connect.ConnectionContent)
)

func (s *service) Page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(r.Context())

	userID, _ := strconv.Atoi(chi.URLParam(r, "userID"))
	pageUser, err := s.Queries.UserByID(ctx, int64(userID))
	if err != nil {
		if err == sql.ErrNoRows {
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

	ug, err := background.ForUser(ctx, s.Queries, pageUser.ID)
	if err != nil {
		render.Error(w, fmt.Errorf("ForUser: %w", err), http.StatusInternalServerError)
		return
	}
	// override layout gradient to show the page user's not the request user's
	l.BackgroundGradient = *ug

	connection, err := s.Queries.GetConnection(ctx, api.GetConnectionParams{
		AID: l.User.ID,
		ID:  pageUser.ID,
	})
	if err != nil {
		render.Error(w, fmt.Errorf("GetConnection: %w", err), http.StatusInternalServerError)
		return
	}

	connected := connection.RoleIn != "" && connection.RoleOut != ""

	d := struct {
		Layout     layout.Data
		User       api.User
		Connection api.GetConnectionRow
		Connected  bool
	}{
		Layout:     l,
		User:       pageUser,
		Connection: connection,
		Connected:  connected,
	}

	render.Execute(w, pageTemplate, d)
}
