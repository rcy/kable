package family

import (
	_ "embed"
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
	pageContent    string
	MyPageTemplate = layout.MustParse(pageContent, me.CardContent)
)

func (s *service) Page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(r.Context())

	family, err := s.Queries.GetFamilyWithGradient(ctx, l.User.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	d := struct {
		Layout layout.Data
		User   api.User
		Family []api.GetFamilyWithGradientRow
	}{
		Layout: l,
		User:   l.User,
		Family: family,
	}

	render.Execute(w, MyPageTemplate, d)
}
