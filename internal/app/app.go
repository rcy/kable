package app

import (
	"net/http"
	"oj/handlers/bots"
	"oj/handlers/chat"
	"oj/handlers/connect"
	"oj/handlers/connectkids"
	"oj/handlers/deliveries"
	"oj/handlers/eventsource"
	"oj/handlers/family"
	"oj/handlers/friends"
	"oj/handlers/fun"
	"oj/handlers/fun/chess"
	"oj/handlers/fun/gradients"
	"oj/handlers/fun/quizzes"
	"oj/handlers/fun/quizzes/attempt"
	"oj/handlers/fun/quizzes/attempt/completed"
	"oj/handlers/fun/quizzes/quiz"
	"oj/handlers/header"
	"oj/handlers/humans"
	"oj/handlers/me"
	"oj/handlers/me/editme"
	"oj/handlers/parent"
	"oj/handlers/postoffice"
	"oj/handlers/u"
	"oj/internal/resources/stickers"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

type Resource struct {
	DB *sqlx.DB
}

func (rs Resource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/me", http.StatusFound)
	})

	r.Get("/header", header.Header)

	r.Get("/parent", parent.Index)
	r.Post("/parent/kids", parent.CreateKid)
	r.Delete("/parent/kids/{userID}", parent.DeleteKid)
	r.Post("/parent/kids/{userID}/logout", parent.LogoutKid)

	r.Post("/chat/messages", chat.PostChatMessage)
	r.Mount("/es", eventsource.SSE)

	r.Get("/me", me.Page)
	r.Get("/me/edit", editme.MyPageEdit)
	r.Post("/me/edit", editme.Post)
	r.Get("/avatars", editme.GetAvatars)
	r.Put("/avatar", editme.PutAvatar)

	r.Get("/me/humans", humans.Page)
	r.Get("/me/family", family.Page)
	r.Get("/me/friends", friends.Page)

	r.Get("/fun", fun.Page)
	r.Get("/fun/gradients", gradients.Index)
	r.Post("/fun/gradients/picker", gradients.Picker)
	r.Post("/fun/gradients/set-background", gradients.SetBackground)

	r.Mount("/stickers", stickers.Resource{DB: rs.DB}.Routes())

	r.Get("/fun/chess", chess.Page)
	r.Get("/fun/chess/select/{rank}/{file}", chess.Select)
	r.Get("/fun/chess/unselect", chess.Unselect)
	//r.Get("/fun/chess/select/{r1}/{f1}/{r2}/{f2}", chess.Move)

	r.Get("/fun/quizzes", quizzes.Page)
	r.Route("/fun/quizzes/{quizID}", quiz.Router)
	r.Get("/fun/quizzes/attempts/{attemptID}", attempt.Page)
	r.Get("/fun/quizzes/attempts/{attemptID}/done", completed.Page)
	r.Post("/fun/quizzes/attempts/{attemptID}/question/{questionID}/response", attempt.PostResponse)

	r.Route("/bots", bots.Router)

	r.Route("/u/{userID}", u.Router)
	r.Get("/u/{userID}/chat", chat.Page)

	r.Get("/connect", connect.Connect)
	r.Put("/connect/friend/{userID}", connect.PutParentFriend)
	r.Delete("/connect/friend/{userID}", connect.DeleteParentFriend)

	r.Get("/connectkids", connectkids.KidConnect)
	r.Put("/connectkids/friend/{userID}", connectkids.PutKidFriend)
	r.Delete("/connectkids/friend/{userID}", connectkids.DeleteKidFriend)

	r.Get("/deliveries/{deliveryID}", deliveries.Page)
	r.Get("/delivery/{deliveryID}", deliveries.Page) // temporary
	r.Post("/deliveries/{deliveryID}/logout", deliveries.Logout)

	r.Route("/postoffice", postoffice.Router)

	return r
}
