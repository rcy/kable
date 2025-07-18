package editme

import (
	_ "embed"
	"fmt"
	"net/http"
	"oj/api"
	"oj/avatar"
	"oj/handlers/me"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type service struct {
	Conn    *pgxpool.Pool
	Queries *api.Queries
}

func NewService(q *api.Queries, conn *pgxpool.Pool) *service {
	return &service{Queries: q, Conn: conn}
}

func (s *service) GetAvatars(w http.ResponseWriter, r *http.Request) {
	const count = 44

	ctx := r.Context()
	user := auth.FromContext(ctx)

	avis := []avatar.Avatar{user.Avatar}

	for i := 0; i < count; i += 1 {
		seed := fmt.Sprintf("%s-%d-%d", time.Now().Format(time.DateOnly), user.ID, i)
		if seed != avis[0].Seed {
			avis = append(avis, avatar.New(seed, avatar.DefaultStyle))
		}
	}

	h.Div(
		h.ID("avatar-list"),
		h.Style("display: flex; flex-wrap: wrap; gap: 1em"),
		g.Map(avis, func(avi avatar.Avatar) g.Node {
			return h.Div(
				g.Attr("hx-put", "/avatar"),
				g.Attr("hx-swap", "outerHTML"),
				g.Attr("hx-target", "#avatar-list"),
				g.Attr("hx-vals", fmt.Sprintf(`{"Seed":"%s"}`, avi.Seed)),
				h.Figure(
					h.Div(h.Style("background: "+string(user.Gradient.Render())),
						h.Img(
							h.Width("128px"),
							h.Height("128px"),
							h.Src(avi.URL()),
						))))
		})).Render(w)
}

func (s *service) PutAvatar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := auth.FromContext(ctx)
	seed := r.FormValue("Seed")

	// TODO: use same style as existing
	avi := avatar.New(seed, avatar.DefaultStyle)

	user, err := s.Queries.UpdateUserAvatar(ctx, api.UpdateUserAvatarParams{
		ID:     user.ID,
		Avatar: avi,
	})
	if err != nil {
		render.Error(w, fmt.Errorf("UpdateAvatar: %w", err), http.StatusInternalServerError)
		return
	}

	changeableAvatar(user).Render(w)
}

func changeableAvatar(user api.User) g.Node {
	id := fmt.Sprintf("avatar-%d", user.ID)
	return h.Div(
		h.ID(id),
		h.Div(h.Style("display:flex; gap: 1em"),
			h.A(
				me.UserCard(user, false),
				//h.Class("nes-btn"),
				h.Href("#"),
				g.Attr("hx-get", "/avatars"),
				g.Attr("hx-target", "#"+id),
			),
			// h.A(
			// 	h.Div(h.Style("width: 128px; height: 128px; background: "+string(user.Gradient.Render()))),
			// 	h.Href("#"),
			// 	g.Attr("hx-get", "/fun/gradients"),
			// 	g.Attr("hx-target", "#"+id)),
		),
	)
}
