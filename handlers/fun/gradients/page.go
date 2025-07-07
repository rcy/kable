package gradients

import (
	_ "embed"
	"fmt"
	"net/http"
	"oj/api"
	"oj/gradient"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/middleware/auth"

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
	pageContent string
	t           = layout.MustParse(pageContent)
)

func Index(w http.ResponseWriter, r *http.Request) {
	l := layout.FromContext(r.Context())

	render.Execute(w, t, struct {
		Layout   layout.Data
		Gradient gradient.Gradient
	}{
		Layout:   l,
		Gradient: l.BackgroundGradient,
	})
}

func Picker(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		render.Error(w, fmt.Errorf("ParseForm: %w", err), 500)
	}

	g, err := gradient.NewFromURLValues(r.PostForm)
	if err != nil {
		render.Error(w, fmt.Errorf("gradientFromUrlValues: %w", err), 500)
		return
	}

	t.ExecuteTemplate(w, "picker", struct {
		Gradient gradient.Gradient
	}{
		Gradient: g,
	})
}

func (s *service) SetBackground(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := auth.FromContext(ctx)

	err := r.ParseForm()
	if err != nil {
		render.Error(w, fmt.Errorf("ParseForm: %w", err), 500)
	}

	g, err := gradient.NewFromURLValues(r.PostForm)
	if err != nil {
		render.Error(w, fmt.Errorf("gradientFromUrlValues: %w", err), 500)
		return
	}

	tx, err := s.Conn.Begin(ctx)
	if err != nil {
		render.Error(w, err, http.StatusInternalServerError)
		return
	}
	defer tx.Rollback(ctx)

	qtx := s.Queries.WithTx(tx)
	_, err = qtx.InsertGradient(ctx, api.InsertGradientParams{
		UserID:   user.ID,
		Gradient: g,
	})
	if err != nil {
		render.Error(w, fmt.Errorf("InsertGradient: %w", err), 500)
		return
	}

	_, err = qtx.UpdateUserGradient(ctx, api.UpdateUserGradientParams{
		UserID:   user.ID,
		Gradient: g,
	})
	if err != nil {
		render.Error(w, fmt.Errorf("UpdateUserGradient: %w", err), 500)
		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		render.Error(w, fmt.Errorf("Commit: %w", err), 500)
		return
	}

	style := fmt.Sprintf("body { background: %s; }", g.Render())

	w.Write([]byte(style))
}
