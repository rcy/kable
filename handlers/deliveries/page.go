package deliveries

import (
	"database/sql"
	_ "embed"
	"fmt"
	"net/http"
	"oj/api"
	"oj/handlers/layout"
	"oj/handlers/render"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type service struct {
	Conn    *pgxpool.Pool
	Queries *api.Queries
}

func NewService(q *api.Queries, conn *pgxpool.Pool) *service {
	return &service{Queries: q, Conn: conn}
}

var (
	//go:embed "page.gohtml"
	pageContent  string
	pageTemplate = layout.MustParse(pageContent)
)

func (s *service) Page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(r.Context())

	deliveryID, _ := strconv.Atoi(chi.URLParam(r, "deliveryID"))
	delivery, err := s.Queries.Delivery(ctx, int64(deliveryID))
	if err != nil {
		if err == sql.ErrNoRows {
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

	render.Execute(w, pageTemplate, struct {
		Layout          layout.Data
		LogoutActionURL string
		Delivery        api.Delivery
	}{
		Layout:          l,
		LogoutActionURL: fmt.Sprintf("%d/logout", delivery.ID),
		Delivery:        delivery,
	})
}

// Logout and redirect back to delivery page to recheck current user
func (s *service) Logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(r.Context())

	deliveryID, _ := strconv.Atoi(chi.URLParam(r, "deliveryID"))
	delivery, err := s.Queries.Delivery(ctx, int64(deliveryID))
	if err != nil {
		if err == sql.ErrNoRows {
			render.Error(w, fmt.Errorf("Delivery: %w", err), http.StatusNotFound)
		} else {
			render.Error(w, fmt.Errorf("Delivery: %w", err), http.StatusInternalServerError)
		}
		return
	}

	_, err = s.Conn.Exec(ctx, `delete from sessions where user_id = ?`, l.User.ID)
	if err != nil {
		render.Error(w, fmt.Errorf("delete from sessions: %w", err), http.StatusInternalServerError)
		return
	}

	url := fmt.Sprintf("/deliveries/%d", delivery.ID)
	http.Redirect(w, r, url, http.StatusSeeOther)
}
