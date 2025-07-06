package notifyfriend

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/url"
	"oj/api"
	"oj/app"
	"oj/services/email"
	"time"

	"github.com/acaloiaro/neoq/jobs"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type service struct {
	Conn    *pgxpool.Pool
	Queries *api.Queries
}

func NewService(q *api.Queries, conn *pgxpool.Pool) *service {
	return &service{Queries: q, Conn: conn}
}

func (s *service) Handle(ctx context.Context) error {
	j, err := jobs.FromContext(ctx)
	if err != nil {
		return err
	}
	log.Printf("handleNotifyFriend job id: %d, payload: %v", j.ID, j.Payload)

	var friend struct {
		ID          int64
		CreatedAt   time.Time `db:"created_at"`
		AID         int64     `db:"a_id"`
		BID         int64     `db:"b_id"`
		Email       string    `db:"email"`
		Username    string    `db:"username"`
		TargetEmail string    `db:"target_email"`
	}

	err = pgxscan.Get(ctx, s.Conn, &friend, `
select
  f.id, f.created_at,
  a.id a_id, a.email, a.username,
  b.id b_id, b.email target_email
from friends f
join users a on a.id = f.a_id
join users b on b.id = f.b_id
where f.id = $1
`, j.Payload["id"])
	if err != nil {
		return err
	}

	var mutualID int64
	err = pgxscan.Get(ctx, s.Conn, &mutualID, `select id from friends where a_id = $1 and b_id = $2`, friend.BID, friend.AID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return err
	}

	var link url.URL
	var subject, emailBody string

	if mutualID != 0 {
		link = app.AbsoluteURL(url.URL{Path: fmt.Sprintf("/u/%d", friend.AID)})
		subject = fmt.Sprintf("%s accepted your friend request", friend.Username)
		emailBody = fmt.Sprintf("click here to view %s", link.String())
	} else {
		link = app.AbsoluteURL(url.URL{Path: "/connect"})
		subject = fmt.Sprintf("%s sent you a friend request", friend.Username)
		emailBody = fmt.Sprintf("click here to accept %s", link.String())
	}

	return email.Send(subject, emailBody, friend.TargetEmail)
}
