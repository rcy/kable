package middleware

import (
	"net/http"
	"oj/models/users"
	"time"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("kh_session")
		if err != nil {
			login(w, r)
			return
		}

		user, err := users.FromSessionKey(cookie.Value)
		if err != nil {
			login(w, r)
			return
		}

		ctx := users.NewContext(r.Context(), user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Save the current path in a cookie and redirect to welcome page
func login(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "redirect",
		Value:   r.URL.Path,
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour)})

	http.Redirect(w, r, "/welcome", http.StatusSeeOther)
}
