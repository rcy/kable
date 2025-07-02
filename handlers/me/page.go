package me

import (
	_ "embed"
	"fmt"
	"net/http"
	"oj/api"
	"oj/handlers/layout"
	"oj/handlers/render"
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

	//go:embed page.gohtml
	pageContent string

	pageTemplate = layout.MustParse(pageContent, CardContent)
)

func (s *service) Page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	l := layout.FromContext(r.Context())

	unreadUsers, err := s.Queries.UsersWithUnreadCounts(ctx, l.User.ID)
	if err != nil {
		render.Error(w, fmt.Errorf("UsersWithUnreadCounts: %w", err), http.StatusInternalServerError)
		return
	}

	d := struct {
		Layout      layout.Data
		User        api.User
		UnreadUsers []api.UsersWithUnreadCountsRow
	}{
		Layout:      l,
		User:        l.User,
		UnreadUsers: unreadUsers,
	}

	render.Execute(w, pageTemplate, d)
}
