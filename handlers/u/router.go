package u

import (
	"oj/api"

	"github.com/go-chi/chi/v5"
)

type service struct {
	Queries *api.Queries
}

func NewService(q *api.Queries) *service {
	return &service{Queries: q}
}

func (s *service) Router(r chi.Router) {
	r.Get("/", s.Page)
	r.Get("/postcard", s.Page)
}
