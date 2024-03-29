package editme

import (
	_ "embed"
	"fmt"
	"html/template"
	"net/http"
	"oj/api"
	"oj/db"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
	"time"
)

//go:embed "avatars.gohtml"
var AvatarContent string

var avatarsTemplate = template.Must(template.New("avatars").Parse(AvatarContent))

func GetAvatars(w http.ResponseWriter, r *http.Request) {
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

func PutAvatar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	queries := api.New(db.DB)
	user := auth.FromContext(ctx)
	newAvatarURL := r.FormValue("URL")

	user, err := queries.UpdateAvatar(ctx, api.UpdateAvatarParams{
		ID:        user.ID,
		AvatarURL: newAvatarURL,
	})
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	render.ExecuteNamed(w, avatarsTemplate, "changeable-avatar", struct{ User api.User }{user})
}
