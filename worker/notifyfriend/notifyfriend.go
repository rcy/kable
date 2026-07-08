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

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/riverqueue/river"
)

type NotifyFriendArgs struct {
	ID int64 `json:"id"`
}

func (NotifyFriendArgs) Kind() string { return "notify_friend" }

type Worker struct {
	river.WorkerDefaults[NotifyFriendArgs]
	Queries *api.Queries
	Conn    *pgxpool.Pool
}

func (w *Worker) Work(ctx context.Context, job *river.Job[NotifyFriendArgs]) error {
	log.Printf("handleNotifyFriend job id: %d, friend id: %d", job.ID, job.Args.ID)

	var friend struct {
		ID          int64
		CreatedAt   time.Time `db:"created_at"`
		AID         int64     `db:"a_id"`
		BID         int64     `db:"b_id"`
		Email       string    `db:"email"`
		Username    string    `db:"username"`
		TargetEmail string    `db:"target_email"`
	}

	err := pgxscan.Get(ctx, w.Conn, &friend, `
select
  f.id, f.created_at,
  a.id a_id, a.email, a.username,
  b.id b_id, b.email target_email
from friends f
join users a on a.id = f.a_id
join users b on b.id = f.b_id
where f.id = $1
`, job.Args.ID)
	if err != nil {
		return err
	}

	var mutualID int64
	err = pgxscan.Get(ctx, w.Conn, &mutualID, `select id from friends where a_id = $1 and b_id = $2`, friend.BID, friend.AID)
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
