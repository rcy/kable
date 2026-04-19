package postoffice

import (
	"fmt"
	"net/http"
	"oj/api"
	"oj/handlers/layout"
	"oj/handlers/postoffice/compose"
	"oj/handlers/render"

	"github.com/go-chi/chi/v5"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type service struct {
	Queries *api.Queries
}

func NewService(q *api.Queries) *service {
	return &service{Queries: q}
}

func (s *service) Router(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/postoffice/inbox", http.StatusSeeOther)
	})

	r.Get("/inbox", s.page)
	r.Route("/compose", compose.NewService(s.Queries).Router)
}

func (s *service) page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(r.Context())

	received, err := s.Queries.UserPostcardsReceived(ctx, l.User.ID)
	if err != nil {
		render.Error(w, fmt.Errorf("UserPostcardsReceived: %w", err), http.StatusInternalServerError)
		return
	}

	sent, err := s.Queries.UserPostcardsSent(ctx, l.User.ID)
	if err != nil {
		render.Error(w, fmt.Errorf("UserPostcardsSent: %w", err), http.StatusInternalServerError)
		return
	}

	layout.Layout(l, "Post Office", postOfficePage(sent, received)).Render(w)
}

func postOfficePage(sent []api.UserPostcardsSentRow, received []api.UserPostcardsReceivedRow) g.Node {
	return h.Div(
		h.H1(g.Text("welcome to the post office")),
		h.Div(h.Style("display:flex; justify-content:space-between"),
			h.Div(),
			h.A(h.Href("compose"), h.Class("nes-btn is-primary"), g.Text("new postcard")),
		),
		g.Map(sent, func(p api.UserPostcardsSentRow) g.Node {
			return h.Div(
				h.Img(h.Height("24"), h.Width("24"), h.Src(p.Avatar.URL())),
				g.Text(" "+p.Username+" "+p.Subject),
			)
		}),
		h.H2(g.Text("Inbox")),
		g.Map(received, func(p api.UserPostcardsReceivedRow) g.Node {
			return h.Div(
				h.Img(h.Height("24"), h.Width("24"), h.Src(p.Avatar.URL())),
				g.Text(" "+p.Username+" "+p.Subject+" ("+p.State+")"),
			)
		}),
	)
}
