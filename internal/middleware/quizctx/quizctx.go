package quizctx

import (
	"context"
	"fmt"
	"net/http"
	"oj/api"
	"oj/handlers/render"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

type service struct {
	Queries *api.Queries
}

func NewService(q *api.Queries) *service {
	return &service{Queries: q}
}

type contextKey int

const quizContextKey contextKey = iota

func Value(ctx context.Context) api.Quiz {
	return ctx.Value(quizContextKey).(api.Quiz)
}

func (s *service) Provider(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		quizID, _ := strconv.Atoi(chi.URLParam(r, "quizID"))
		quiz, err := s.Queries.Quiz(ctx, int64(quizID))
		if err != nil {
			if err == pgx.ErrNoRows {
				render.NotFound(w)
				return
			}
			render.Error(w, fmt.Errorf("Quiz: %w", err), http.StatusInternalServerError)
			return
		}
		ctx = context.WithValue(ctx, quizContextKey, quiz)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
