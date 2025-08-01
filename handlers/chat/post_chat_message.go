package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"oj/handlers/eventsource"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
	"oj/worker"
	"strconv"
	"strings"
	"time"

	"github.com/alexandrevicenzi/go-sse"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (rs Resource) PostChatMessage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := auth.FromContext(ctx)
	roomID, err := strconv.Atoi(r.FormValue("roomID"))
	if err != nil {
		render.Error(w, fmt.Errorf("Atoi: %w", err), 500)
		return
	}
	body := r.FormValue("body")

	if strings.TrimSpace(body) != "" {
		err = rs.PostMessage(ctx, int64(roomID), user.ID, body)
		if err != nil {
			render.Error(w, fmt.Errorf("postMessage: %w", err), 500)
			return
		}
	}

	render.ExecuteNamed(w, pageTemplate, "chatInput", struct{ RoomID int }{RoomID: roomID})
}

type RoomUser struct {
	ID        int64
	CreatedAt time.Time `db:"created_at"`
	RoomID    int64     `db:"room_id"`
	UserID    int64     `db:"user_id"`
	Email     *string   `db:"email"`
}

func (rs Resource) PostMessage(ctx context.Context, roomID, senderID int64, body string) error {
	var roomUsers []RoomUser

	tx, err := rs.Conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = rs.Queries.UserByID(ctx, senderID)
	if err != nil {
		return err
	}

	// get the users of the room
	err = pgxscan.Select(ctx, tx, &roomUsers, `select room_users.*, users.email from room_users join users on room_users.user_id = users.id where room_id = $1`, roomID)
	if err != nil {
		return err
	}

	// create the message
	var messageID int64
	err = pgxscan.Get(ctx, tx, &messageID, `insert into messages(room_id, sender_id, body) values($1,$2,$3) returning id`, roomID, senderID, body)
	if err != nil {
		return err
	}

	// create deliveries for each user in the room
	var deliveryIDs []int64
	for _, roomUser := range roomUsers {
		var deliveryID int64
		err = pgxscan.Get(ctx, tx, &deliveryID, `insert into deliveries(message_id, room_id, sender_id, recipient_id) values($1,$2,$3,$4) returning id`, messageID, roomID, senderID, roomUser.UserID)
		if err != nil {
			return err
		}

		deliveryIDs = append(deliveryIDs, deliveryID)
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}

	// send notifications after the transaction has been committed

	for _, deliveryID := range deliveryIDs {
		worker.NotifyDelivery(deliveryID)
	}

	data, err := json.Marshal(map[string]interface{}{
		"senderID": fmt.Sprint(senderID),
	})
	if err != nil {
		return err
	}

	eventsource.SSE.SendMessage(
		fmt.Sprintf("/es/room-%d", roomID),
		sse.NewMessage("", string(data), "NEW_MESSAGE"))

	go func() {
		time.Sleep(time.Second)
		for _, roomUser := range roomUsers {
			if roomUser.UserID == senderID {
				continue
			}

			eventsource.SSE.SendMessage(
				fmt.Sprintf("/es/user-%d", roomUser.UserID),
				sse.NewMessage("", "simple", "USER_UPDATE"))
		}
	}()

	return nil
}
