package messages

import (
	_ "embed"
	"fmt"
	"net/http"
	"oj/api"
	"oj/handlers/layout"
	"oj/handlers/render"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type service struct {
	Queries *api.Queries
}

func NewService(q *api.Queries) *service {
	return &service{Queries: q}
}

func (s *service) Router(r chi.Router) {
	r.Get("/", s.page)
	r.Delete("/{messageID}", s.deleteMessage)
}

var (
	//go:embed page.gohtml
	pageContent  string
	pageTemplate = layout.MustParse(pageContent, pageContent)
)

func (s *service) page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	l := layout.FromContext(r.Context())

	messages, err := s.Queries.AdminRecentMessages(ctx)
	if err != nil {
		render.Error(w, fmt.Errorf("AdminRecentMessages: %w", err), http.StatusInternalServerError)
		return
	}

	render.Execute(w, pageTemplate, struct {
		Layout   layout.Data
		Messages []api.AdminRecentMessagesRow
	}{
		Layout:   l,
		Messages: messages,
	})
}

func (s *service) deleteMessage(w http.ResponseWriter, r *http.Request) {
	messageID, _ := strconv.Atoi(chi.URLParam(r, "messageID"))
	ctx := r.Context()

	message, err := s.Queries.AdminDeleteMessage(ctx, int64(messageID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.ExecuteNamed(w, pageTemplate, "message-row", message)
}
