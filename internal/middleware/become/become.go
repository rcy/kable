package become

import (
	"context"
	"net/http"
	"oj/api"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
)

type service struct {
	Queries *api.Queries
}

func NewService(q *api.Queries) *service {
	return &service{Queries: q}
}

func (s *service) Provider(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		becomeUser, err := s.getUser(ctx)
		if err != nil {
			if err == auth.ErrNotAuthorized {
				render.Error(w, "getUser:"+err.Error(), http.StatusUnauthorized)
				return
			}
			render.Error(w, "getUser:"+err.Error(), http.StatusInternalServerError)
			return
		}

		if becomeUser != nil {
			ctx = auth.NewContext(ctx, *becomeUser)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func (s *service) getUser(ctx context.Context) (*api.User, error) {
	user := auth.FromContext(ctx)
	if !user.BecomeUserID.Valid {
		return nil, nil
	}
	if !user.Admin {
		return nil, auth.ErrNotAuthorized
	}
	becomeUser, err := s.Queries.UserByID(ctx, user.BecomeUserID.Int64)
	return &becomeUser, err
}
