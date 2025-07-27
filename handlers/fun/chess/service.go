package chess

import (
	"oj/api"

	"github.com/jackc/pgx/v5/pgxpool"
)

type service struct {
	Conn    *pgxpool.Pool
	Queries *api.Queries
}

func NewService(q *api.Queries, conn *pgxpool.Pool) *service {
	return &service{Queries: q, Conn: conn}
}
