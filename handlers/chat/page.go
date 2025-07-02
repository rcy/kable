package chat

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"oj/api"
	"oj/handlers/eventsource"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
	"oj/services/background"
	"oj/services/room"
	"strconv"

	"github.com/alexandrevicenzi/go-sse"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Resource struct {
	Queries *api.Queries
	Conn    *pgxpool.Pool
}

func NewService(q *api.Queries, conn *pgxpool.Pool) *Resource {
	return &Resource{Queries: q, Conn: conn}
}

var (
	//go:embed page.gohtml
	pageContent  string
	pageTemplate = layout.MustParse(pageContent)
)

func (rs Resource) Page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := auth.FromContext(ctx)

	pageUserID, _ := strconv.Atoi(chi.URLParam(r, "userID"))
	pageUser, err := rs.Queries.UserByID(ctx, int64(pageUserID))
	if err != nil {
		if err == pgx.ErrNoRows {
			render.Error(w, fmt.Errorf("UserByID: %w", err), 404)
			return
		}
	}
	ug, err := background.ForUser(ctx, rs.Queries, pageUser.ID)
	if err != nil {
		render.Error(w, fmt.Errorf("ForUser: %w", err), 500)
		return
	}

	room, err := room.FindOrCreateByUserIDs(ctx, rs.Conn, rs.Queries, user.ID, pageUser.ID)
	if err != nil {
		render.Error(w, fmt.Errorf("FindOrCreateByUserIDs: %w", err), 500)
		return
	}

	records, err := rs.Queries.RecentRoomMessages(ctx, room.ID)
	if err != nil {
		render.Error(w, fmt.Errorf("RecentRoomMessages: %w", err), 500)
		return
	}

	err = rs.updateDeliveries(room.ID, user.ID)
	if err != nil {
		render.Error(w, fmt.Errorf("updateDeliveries: %w", err), 500)
		return
	}

	// get the layout after the deliveries have been updated to ensure unread count is correct
	l, err := layout.NewService(rs.Queries, rs.Conn).FromUser(ctx, user)
	if err != nil {
		render.Error(w, fmt.Errorf("FromUser: %w", err), 500)
		return
	}
	// override layout gradient to show the page user's not the request user's
	l.BackgroundGradient = *ug

	pd := struct {
		Layout   layout.Data
		User     api.User
		RoomID   int64
		Messages []api.RecentRoomMessagesRow
	}{
		Layout:   l,
		User:     pageUser,
		RoomID:   room.ID,
		Messages: records,
	}

	render.Execute(w, pageTemplate, pd)
}

func (rs Resource) updateDeliveries(roomID, userID int64) error {
	log.Printf("UPDATE DELIVERIES %d %v", userID, rs.Conn)
	_, err := rs.Conn.Exec(context.TODO(), `update deliveries set sent_at = now() where sent_at is null and room_id = $1 and recipient_id = $2`, roomID, userID)
	if err != nil {
		return fmt.Errorf("Exec", err)
	}
	log.Printf("UPDATE DELIVERIES %d...done", userID)

	eventsource.SSE.SendMessage(
		fmt.Sprintf("/es/user-%d", userID),
		sse.NewMessage("", "simple", "USER_UPDATE"))

	return err
}
