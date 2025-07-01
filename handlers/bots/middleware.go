package bots

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"oj/api"
	"oj/handlers/render"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

func (rs Resource) provideBot(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		botID, _ := strconv.Atoi(chi.URLParam(r, "botID"))
		bot, err := rs.Model.Bot(ctx, int64(botID))
		if errors.Is(err, pgx.ErrNoRows) {
			render.NotFound(w)
			return
		}
		if err != nil {
			render.Error(w, fmt.Errorf("Model.Bot: %w", err), http.StatusInternalServerError)
			return
		}

		ctx = context.WithValue(ctx, "bot", bot)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func botFromContext(ctx context.Context) api.Bot {
	return ctx.Value("bot").(api.Bot)
}
