package humans

import (
	_ "embed"
	"fmt"
	"net/http"
	"oj/api"
	"oj/handlers/layout"
	"oj/handlers/me"
	"oj/handlers/render"
)

type service struct {
	Queries *api.Queries
}

func NewService(q *api.Queries) *service {
	return &service{Queries: q}
}

var (
	//go:embed "page.gohtml"
	pageContent string

	pageTemplate = layout.MustParse(pageContent, me.CardContent)
)

func (s *service) Page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(r.Context())

	friends, err := s.Queries.GetConnections(ctx, l.User.ID)
	if err != nil {
		render.Error(w, fmt.Errorf("GetConnections: %w", err), http.StatusInternalServerError)
		return
	}

	d := struct {
		Layout  layout.Data
		User    api.User
		Friends []api.User
	}{
		Layout:  l,
		User:    l.User,
		Friends: friends,
	}

	render.Execute(w, pageTemplate, d)
}
