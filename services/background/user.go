package background

import (
	"context"
	"errors"
	"fmt"
	"oj/api"
	"oj/element/gradient"

	"github.com/jackc/pgx/v5"
)

func ForUser(ctx context.Context, queries *api.Queries, userID int64) (*gradient.Gradient, error) {
	gradientRow, err := queries.UserGradient(ctx, userID)
	if errors.Is(err, pgx.ErrNoRows) {
		return &gradient.Default, nil
	}
	if err != nil {
		return nil, fmt.Errorf("UserGradient: %w", err)
	}

	return &gradientRow.Gradient, nil
}
