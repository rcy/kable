package room

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"oj/api"

	"github.com/jackc/pgx/v5/pgxpool"
)

func FindOrCreateByUserIDs(ctx context.Context, conn *pgxpool.Pool, model *api.Queries, id1, id2 int64) (*api.Room, error) {
	key := makeRoomKey(id1, id2)

	room, err := model.RoomByKey(ctx, key)
	if err != nil {
		if err == sql.ErrNoRows {
			return build(ctx, conn, model, key, id1, id2)
		}
		return nil, err
	}
	return &room, nil
}

// Create room and add users
func build(ctx context.Context, conn *pgxpool.Pool, model *api.Queries, key string, userID1, userID2 int64) (*api.Room, error) {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)
	txModel := model.WithTx(tx)

	room, err := txModel.CreateRoom(ctx, key)
	if err != nil {
		return nil, err
	}

	for _, userID := range []int64{userID1, userID2} {
		_, err = txModel.CreateRoomUser(ctx, api.CreateRoomUserParams{
			RoomID: room.ID,
			UserID: userID,
		})
		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return &room, nil
}

// Generate a string to be used as a roomKey given 2 user IDs
func makeRoomKey(id1, id2 int64) string {
	min := int64(math.Min(float64(id1), float64(id2)))
	max := int64(math.Max(float64(id1), float64(id2)))

	return fmt.Sprintf("dm-%d-%d", min, max)
}
