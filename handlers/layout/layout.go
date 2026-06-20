package layout

import (
	"context"
	"fmt"
	"oj/api"
	"oj/gradient"

	"github.com/georgysavva/scany/v2/pgxscan"
)

type Data struct {
	User               api.User
	BackgroundGradient gradient.Gradient
	UnreadCount        int
}

func (s *service) FromUser(ctx context.Context, user api.User) (Data, error) {
	var unreadCount int
	err := pgxscan.Get(ctx, s.Conn, &unreadCount, `select count(*) from deliveries where recipient_id = $1 and sent_at is null`, user.ID)
	if err != nil {
		return Data{}, fmt.Errorf("pgxscan: %w", err)
	}

	return Data{
		User:               user,
		BackgroundGradient: user.Gradient,
		UnreadCount:        unreadCount,
	}, nil
}
