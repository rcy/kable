package compose

import (
	"fmt"
	"log"
	"net/http"
	"oj/api"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
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
	r.Post("/", s.post)
}

func (s *service) page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(r.Context())

	connections, err := s.Queries.GetConnections(ctx, l.User.ID)
	if err != nil {
		render.Error(w, fmt.Errorf("GetConnections: %w", err), http.StatusInternalServerError)
		return
	}

	layout.Layout(l, "Compose Postcard", composePage(connections)).Render(w)
}

func composePage(connections []api.User) g.Node {
	return h.Div(
		h.H1(g.Text("compose new postcard")),
		h.Form(h.Method("post"),
			h.Label(g.Attr("for", "recepient_field"), g.Text("To")),
			h.Div(h.Class("nes-select"),
				g.El("select", g.Attr("required", ""), h.Name("recipient"), h.ID("recipient_field"),
					g.El("option", h.Value(""), g.Attr("disabled", ""), g.Attr("selected", ""), g.Attr("hidden", ""), g.Text("Select...")),
					g.Map(connections, func(u api.User) g.Node {
						return g.El("option", h.Value(fmt.Sprint(u.ID)), g.Text(u.Username))
					}),
				),
			),
			h.Div(h.Class("nes-field"),
				h.Label(g.Attr("for", "subject_field"), g.Text("Subject")),
				h.Input(g.Attr("required", ""), h.Name("subject"), h.Type("text"), h.ID("subject_field"), h.Class("nes-input")),
			),
			h.Label(g.Attr("for", "body_field"), g.Text("Message")),
			h.Textarea(g.Attr("required", ""), g.Attr("rows", "10"), h.Name("body"), h.ID("body_field"), h.Class("nes-textarea")),
			h.Button(h.Class("nes-btn primary"), g.Text("send")),
			h.A(h.Href("inbox"), h.Class("nes-btn"), g.Text("cancel")),
		),
	)
}

func (s *service) post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sender := auth.FromContext(ctx)

	recipient, _ := strconv.Atoi(r.FormValue("recipient"))
	params := api.CreatePostcardParams{
		Sender:    sender.ID,
		Recipient: int64(recipient),
		Subject:   r.FormValue("subject"),
		Body:      r.FormValue("body"),
		State:     "queued",
	}

	log.Print("postcard", params)

	_, err := s.Queries.CreatePostcard(ctx, params)
	if err != nil {
		render.Error(w, fmt.Errorf("CreatePostcard: %w", err), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/postoffice", http.StatusSeeOther)
}
