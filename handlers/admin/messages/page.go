package messages

import (
	_ "embed"
	"fmt"
	"net/http"
	"oj/api"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/templatehelpers"
	"strconv"

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
	r.Get("/", s.page)
	r.Delete("/{messageID}", s.deleteMessage)
}

func (s *service) page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	l := layout.FromContext(r.Context())

	messages, err := s.Queries.AdminRecentMessages(ctx)
	if err != nil {
		render.Error(w, fmt.Errorf("AdminRecentMessages: %w", err), http.StatusInternalServerError)
		return
	}

	layout.Layout(l, "recent messages",
		h.Div(
			h.Class("nes-container ghost"),
			h.H1(
				g.Text("recent messages"),
			),
			h.Table(
				h.Class("nes-table is-bordered is-centered"),
				h.THead(
					h.Tr(
						h.Th(
							g.Text("id"),
						),
						h.Th(
							g.Text("r"),
						),
						h.Th(
							g.Text("sender"),
						),
						h.Th(
							g.Text("message body"),
						),
						h.Th(
							g.Text("sent"),
						),
						h.Th(),
					),
				),
				h.TBody(
					g.Map(messages, func(msg api.AdminRecentMessagesRow) g.Node {
						return messageRowEl(msg.Message, msg.User)
					}),
				),
			)),
	).Render(w)
}

func messageRowEl(message api.Message, sender api.User) g.Node {
	return h.Tr(
		h.Td(g.Text(fmt.Sprint(message.ID))),
		h.Td(g.Text(fmt.Sprint(message.RoomID))),
		h.Td(h.Style("text-wrap:wrap"), g.Text(fmt.Sprintf("%s:%d", sender.Username, sender.ID))),
		h.Td(g.Text(message.Body)),
		h.Td(g.Text(templatehelpers.Ago(message.CreatedAt))),
		h.Td(
			h.Button(
				g.Attr("hx-delete", "/admin/messages/{{.ID}}"),
				g.Attr("hx-confirm", fmt.Sprintf("delete %s?", message.Body)),
				g.Attr("hx-target", "closest tr"),
				g.Attr("hx-swap", "outerHTML"),
				h.Class("nes-btn"),
				g.Text("delete"),
			),
		),
	)
}

func (s *service) deleteMessage(w http.ResponseWriter, r *http.Request) {
	messageID, _ := strconv.Atoi(chi.URLParam(r, "messageID"))
	ctx := r.Context()

	message, err := s.Queries.AdminDeleteMessage(ctx, int64(messageID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	messageRowEl(message, api.User{}).Render(w)
}
