package chat

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"oj/api"
	"oj/handlers/eventsource"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
	"oj/md"
	"oj/services/room"
	"oj/templatehelpers"
	"slices"
	"strconv"

	"github.com/alexandrevicenzi/go-sse"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type Resource struct {
	Queries *api.Queries
	Conn    *pgxpool.Pool
}

func NewService(q *api.Queries, conn *pgxpool.Pool) *Resource {
	return &Resource{Queries: q, Conn: conn}
}

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
	slices.Reverse(records)

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
	l.BackgroundGradient = pageUser.Gradient

	layout.Layout(l, pageUser.Username, chatPage(l, pageUser, room.ID, records)).Render(w)
}

func (rs Resource) updateDeliveries(roomID, userID int64) error {
	log.Printf("UPDATE DELIVERIES %d %v", userID, rs.Conn)
	_, err := rs.Conn.Exec(context.TODO(), `update deliveries set sent_at = now() where sent_at is null and room_id = $1 and recipient_id = $2`, roomID, userID)
	if err != nil {
		return fmt.Errorf("Exec: %w", err)
	}
	log.Printf("UPDATE DELIVERIES %d...done", userID)

	eventsource.SSE.SendMessage(
		fmt.Sprintf("/es/user-%d", userID),
		sse.NewMessage("", "simple", "USER_UPDATE"))

	return err
}

func chatPage(l layout.Data, pageUser api.User, roomID int64, records []api.RecentRoomMessagesRow) g.Node {
	return g.Group{
		h.StyleEl(g.Raw(`
.hide-scrollbar::-webkit-scrollbar { display: none; }
.hide-scrollbar { -ms-overflow-style: none; scrollbar-width: none; }
`)),
		chatInitEventSourceEl(l.User.ID, roomID),
		chatEnsureInteractionEl(),
		h.Div(h.Style("height:100%; display:flex; flex-direction:column"),
			h.A(h.Href(fmt.Sprintf("/u/%d", pageUser.ID)), h.Style("z-index:1000; text-decoration: none;"),
				h.Section(h.Class("ghost"), h.Style("display:flex; align-items: top; gap:1em; box-shadow: 0 8px rgba(0,0,0,.5);"),
					h.Img(h.Src(pageUser.Avatar.URL()), h.Width("80px"), h.Height("80px")),
					h.Div(
						g.Text("chatting with"),
						h.H1(g.Text(pageUser.Username)),
					),
				),
			),
			h.Div(h.Class("hide-scrollbar"), h.Style("flex:1; overflow-y: scroll; display:flex; flex-direction: column-reverse"),
				chatMessagesEl(pageUser, l.User.ID, records),
			),
			h.Div(chatInputEl(roomID)),
		),
	}
}

func chatMessagesEl(pageUser api.User, currentUserID int64, records []api.RecentRoomMessagesRow) g.Node {
	return h.Div(
		g.Attr("hx-get", fmt.Sprintf("/u/%d/chat", pageUser.ID)),
		g.Attr("hx-trigger", "NEW_MESSAGE from:body, ROOM_CONNECTED from:body"),
		h.ID(fmt.Sprintf("messages-%d", pageUser.ID)),
		g.Attr("hx-select", fmt.Sprintf("#messages-%d", pageUser.ID)),
		g.Attr("hx-swap", "outerHTML"),
		g.Map(records, func(row api.RecentRoomMessagesRow) g.Node {
			if row.Message.SenderID == currentUserID {
				return chatMessageMineEl(row)
			}
			return chatMessageOtherEl(row)
		}),
	)
}

func chatMessageOtherEl(row api.RecentRoomMessagesRow) g.Node {
	return h.Div(h.Style("display:flex; align-items:end"),
		h.A(h.Href(fmt.Sprintf("/u/%d", row.User.ID)),
			h.Img(h.Style("margin-bottom: 8px"), h.Src(row.User.Avatar.URL()), h.Width("80"), h.Height("80")),
		),
		h.Div(h.Class("nes-balloon from-left"), h.Style("overflow-wrap: break-word; max-width: 70%"),
			h.P(g.Raw(string(md.RenderString(row.Message.Body)))),
			h.Div(h.Style("display:flex; gap:1em; font-size: .8em; float:right; color: gray"),
				g.Text(templatehelpers.Ago(row.Message.CreatedAt)),
				h.Div(h.A(h.Href("#"),
					g.Attr("hx-swap", "outerHTML"),
					g.Attr("hx-post", fmt.Sprintf("/fun/notes/from-chat/%d", row.Message.ID)),
					g.Text("save"),
				)),
			),
		),
		h.Div(h.Style("flex:1")),
	)
}

func chatMessageMineEl(row api.RecentRoomMessagesRow) g.Node {
	return h.Div(h.Style("display:flex; align-items:end"),
		h.Div(h.Style("flex:1")),
		h.Div(h.Class("nes-balloon from-right"), h.Style("overflow-wrap: break-word; max-width: 70%"),
			h.P(g.Raw(string(md.RenderString(row.Message.Body)))),
			h.Div(h.Style("display:flex; gap:1em; font-size: .8em; float:right; color: gray"),
				g.Text(templatehelpers.Ago(row.Message.CreatedAt)),
				h.Div(h.A(h.Href("#"),
					g.Attr("hx-swap", "outerHTML"),
					g.Attr("hx-post", fmt.Sprintf("/fun/notes/from-chat/%d", row.Message.ID)),
					g.Text("save"),
				)),
			),
		),
		h.Img(h.Style("margin-bottom: 8px"), h.Src(row.User.Avatar.URL()), h.Width("80"), h.Height("80")),
	)
}

func chatInputEl(roomID int64) g.Node {
	return h.Form(h.ID("form2"),
		g.Attr("autocomplete", "off"),
		g.Attr("hx-post", "/chat/messages"),
		g.Attr("hx-swap", "outerHTML"),
		g.Attr("_", "on submit add @disabled to #inp"),
		h.Input(h.Name("roomID"), h.Value(fmt.Sprint(roomID)), h.Type("hidden")),
		h.Div(h.Style("display:flex"),
			h.Input(h.ID("inp"), h.Name("body"), h.Type("text"), g.Attr("autofocus", ""), g.Attr("required", ""), h.Class("nes-input"), g.Attr("placeholder", "Type a message")),
			h.Button(h.Class("nes-btn is-primary"), g.Text("Send")),
		),
	)
}

func chatInitEventSourceEl(userID, roomID int64) g.Node {
	return h.Script(g.Raw(fmt.Sprintf(`
const userID = "%d";
(function(){
  const es = new EventSource("/es/room-%d");
  es.addEventListener("NEW_MESSAGE", (e) => {
    const data = JSON.parse(e.data)
    if (data.senderID !== userID) {
      document.getElementById("beeper").play()
    }
    document.getElementsByTagName('body')[0].dispatchEvent(new Event("NEW_MESSAGE"))
  });
  es.addEventListener("open", (e) => {
    console.log("The connection has been established.");
    document.getElementsByTagName('body')[0].dispatchEvent(new Event("ROOM_CONNECTED"))
  });
  es.addEventListener("error", (e) => {
    console.log("An error occurred while attempting to connect.");
  });
  window.addEventListener('beforeunload', function() { es.close(); });
})();`, userID, roomID)))
}

func chatEnsureInteractionEl() g.Node {
	return h.Section(
		h.Script(g.Raw(`
function testBeep() { return (new Audio()).play() }
window.onload = async function() {
  try { await testBeep() } catch(e) {
    document.getElementById('force-interaction').showModal()
  }
}`)),
		g.El("dialog", h.Class("nes-dialog"), h.ID("force-interaction"),
			h.Form(h.Method("dialog"),
				h.P(h.Class("title"), g.Text("Hello!")),
				h.P(g.Text("Click button to enable beeps")),
				h.Menu(h.Class("dialog-menu"),
					h.Button(h.Class("nes-btn is-primary"), g.Attr("onclick", "testBeep()"), g.Text("OK")),
				),
			),
		),
	)
}
