package background

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"oj/api"
	"oj/element/gradient"
)

func ForUser(ctx context.Context, queries *api.Queries, userID int64) (*gradient.Gradient, error) {
	gradientRow, err := queries.UserGradient(ctx, userID)
	if errors.Is(err, sql.ErrNoRows) {
		return &gradient.Default, nil
	}
	if err != nil {
		return nil, fmt.Errorf("UserGradient: %w", err)
	}

	return &gradientRow.Gradient, nil
}
