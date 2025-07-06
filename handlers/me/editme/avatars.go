package editme

import (
	_ "embed"
	"fmt"
	"html/template"
	"net/http"
	"oj/api"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type service struct {
	Conn    *pgxpool.Pool
	Queries *api.Queries
}

func NewService(q *api.Queries, conn *pgxpool.Pool) *service {
	return &service{Queries: q, Conn: conn}
}

//go:embed "avatars.gohtml"
var AvatarContent string

var avatarsTemplate = template.Must(template.New("avatars").Parse(AvatarContent))

func (s *service) GetAvatars(w http.ResponseWriter, r *http.Request) {
	const count = 44

	ctx := r.Context()
	user := auth.FromContext(ctx)

	urls := []string{user.AvatarURL}

	for i := 0; i < count; i += 1 {
		base := "https://api.dicebear.com/7.x/avataaars/svg?seed=%s"
		seed := fmt.Sprintf("%s-%d-%d", time.Now().Format(time.DateOnly), user.ID, i)
		url := fmt.Sprintf(base, seed)
		if url != urls[0] {
			urls = append(urls, url)
		}
	}

	render.ExecuteNamed(w, avatarsTemplate, "avatars", struct{ URLs []string }{urls})
}

func (s *service) PutAvatar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := auth.FromContext(ctx)
	newAvatarURL := r.FormValue("URL")

	user, err := s.Queries.UpdateAvatar(ctx, api.UpdateAvatarParams{
		ID:        user.ID,
		AvatarURL: newAvatarURL,
	})
	if err != nil {
		render.Error(w, fmt.Errorf("UpdateAvatar: %w", err), http.StatusInternalServerError)
		return
	}
	render.ExecuteNamed(w, avatarsTemplate, "changeable-avatar", struct{ User api.User }{user})
}
