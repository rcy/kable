package auth

import (
	"fmt"
	"net/http"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
)

func EnsureAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := auth.FromContext(r.Context())
		if !user.Admin {
			render.Error(w, fmt.Errorf("auth.FromContext"), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
