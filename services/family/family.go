package family

import (
	"context"
	"oj/api"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

// Create a new kid user with username
func CreateKid(ctx context.Context, queries *api.Queries, parentID int64, username string) (*api.User, error) {
	// tx, err := db.DB.Beginx()
	// if err != nil {
	// 	return nil, err
	// }
	// defer tx.Rollback()

	kid, err := queries.CreateUser(ctx, username)
	if err != nil {
		return nil, err
	}

	_, err = queries.CreateKidParent(ctx, api.CreateKidParentParams{
		KidID:    kid.ID,
		ParentID: parentID,
	})
	if err != nil {
		return nil, err
	}

	_, err = queries.CreateFriend(ctx, api.CreateFriendParams{
		AID:   parentID,
		BID:   kid.ID,
		BRole: "child",
	})
	if err != nil {
		return nil, err
	}

	_, err = queries.CreateFriend(ctx, api.CreateFriendParams{
		AID:   kid.ID,
		BID:   parentID,
		BRole: "parent",
	})
	if err != nil {
		return nil, err
	}

	// err = tx.Commit()
	// if err != nil {
	// 	return nil, err
	// }

	return &kid, nil
}

func FindOrCreateParentByEmail(ctx context.Context, queries *api.Queries, email string) (api.User, error) {
	nullableEmail := pgtype.Text{String: email, Valid: true}

	user, err := queries.UserByEmail(ctx, nullableEmail)
	if err != nil {
		if err == pgx.ErrNoRows {
			// we don't have a username here, so use the email, they can change it later
			return queries.CreateParent(ctx, api.CreateParentParams{
				Email:    nullableEmail,
				Username: email,
			})
		}
		return api.User{}, err
	}
	return user, nil
}
