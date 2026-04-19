package deliveries

import (
	"fmt"
	"net/http"
	"oj/api"
	"oj/handlers/layout"
	"oj/handlers/render"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
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

func (s *service) Page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(r.Context())

	deliveryID, _ := strconv.Atoi(chi.URLParam(r, "deliveryID"))
	delivery, err := s.Queries.Delivery(ctx, int64(deliveryID))
	if err != nil {
		if err == pgx.ErrNoRows {
			render.Error(w, fmt.Errorf("Delivery not found: %w", err), http.StatusNotFound)
		} else {
			render.Error(w, fmt.Errorf("Delivery: %w", err), http.StatusInternalServerError)
		}
		return
	}

	if delivery.RecipientID == l.User.ID {
		url := fmt.Sprintf("/u/%d/chat", delivery.SenderID)
		http.Redirect(w, r, url, http.StatusSeeOther)
		return
	}

	layout.Layout(l, "Delivery", deliveryPage(l.User.Username, delivery)).Render(w)
}

func deliveryPage(username string, delivery api.Delivery) g.Node {
	logoutURL := fmt.Sprintf("%d/logout", delivery.ID)
	return g.El("dialog", g.Attr("open", ""), h.Class("nes-dialog"),
		h.Form(h.Method("post"), h.Action(logoutURL),
			h.P(g.Text("You are currently logged in as "+username+".")),
			h.P(g.Text("This content is for a different user!")),
			h.Div(h.Style("display:flex; justify-content: space-between"),
				h.Button(h.Class("nes-btn is-primary"), g.Text("Switch User")),
				h.A(h.Href("/"), h.Class("nes-btn"), g.Text("Cancel")),
			),
		),
	)
}

// Logout and redirect back to delivery page to recheck current user
func (s *service) Logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(r.Context())

	deliveryID, _ := strconv.Atoi(chi.URLParam(r, "deliveryID"))
	delivery, err := s.Queries.Delivery(ctx, int64(deliveryID))
	if err != nil {
		if err == pgx.ErrNoRows {
			render.Error(w, fmt.Errorf("Delivery: %w", err), http.StatusNotFound)
		} else {
			render.Error(w, fmt.Errorf("Delivery: %w", err), http.StatusInternalServerError)
		}
		return
	}

	_, err = s.Conn.Exec(ctx, `delete from sessions where user_id = $1`, l.User.ID)
	if err != nil {
		render.Error(w, fmt.Errorf("delete from sessions: %w", err), http.StatusInternalServerError)
		return
	}

	url := fmt.Sprintf("/deliveries/%d", delivery.ID)
	http.Redirect(w, r, url, http.StatusSeeOther)
}
