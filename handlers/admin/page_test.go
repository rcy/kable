package admin

import (
	"net/http/httptest"
	"oj/api"
	"oj/internal/middleware/auth"
	"testing"
)

func TestRouter(t *testing.T) {
	t.Skip("db related test")

	routes := NewService(nil, nil).Routes()

	for _, tc := range []struct {
		name           string
		path           string
		user           api.User
		wantStatusCode int
	}{
		{
			name:           "non admin user",
			path:           "/",
			user:           api.User{Admin: false},
			wantStatusCode: 401,
		},
		{
			name:           "admin user",
			path:           "/",
			user:           api.User{Admin: true},
			wantStatusCode: 200,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", tc.path, nil)
			ctx := auth.NewContext(req.Context(), tc.user)
			routes.ServeHTTP(w, req.WithContext(ctx))
			resp := w.Result()

			if resp.StatusCode != tc.wantStatusCode {
				t.Errorf("bad status code, expected %d, got %d", tc.wantStatusCode, resp.StatusCode)
			}
		})
	}
}
