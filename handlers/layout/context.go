package layout

import (
	"context"
	"fmt"
	"net/http"
	"oj/api"
	"oj/handlers/render"
	"oj/internal/middleware/auth"

	"github.com/jackc/pgx/v5/pgxpool"
)

type service struct {
	Conn    *pgxpool.Conn
	Queries *api.Queries
}

func NewService(q *api.Queries, conn *pgxpool.Conn) *service {
	return &service{Queries: q, Conn: conn}
}

type contextKey int

const layoutContextKey contextKey = iota

func FromContext(ctx context.Context) Data {
	value := ctx.Value(layoutContextKey)
	if value == nil {
		return Data{}
	}
	return value.(Data)
}

func NewContext(ctx context.Context, data Data) context.Context {
	return context.WithValue(ctx, layoutContextKey, data)
}

func (s *service) Provider(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := auth.FromContext(ctx)
		data, err := s.FromUser(ctx, user)
		if err != nil {
			render.Error(w, fmt.Errorf("s.FromUser: %w", err), http.StatusInternalServerError)
			return
		}
		ctx = NewContext(ctx, data)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
