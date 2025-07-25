package app

import (
	"net/http"
	"oj/api"

	//"oj/handlers/admin/quizzes"
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
	"oj/handlers/fun/notebook"
	"oj/handlers/header"
	"oj/handlers/humans"
	"oj/handlers/me"
	"oj/handlers/me/editme"
	"oj/handlers/postoffice"
	"oj/handlers/u"
	"oj/handlers/u/quizzes/attempt"
	"oj/handlers/u/quizzes/create"
	"oj/handlers/u/quizzes/show"
	"oj/handlers/u/quizzes/view"
	"oj/internal/ai"
	"oj/internal/resources/parent"
	"oj/internal/resources/stickers"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jmoiron/sqlx"
)

type Service struct {
	DB      *sqlx.DB
	Conn    *pgxpool.Pool
	Queries *api.Queries
}

func NewService(q *api.Queries, conn *pgxpool.Pool) *Service {
	return &Service{Queries: q, Conn: conn}
}

func (rs Service) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/me", http.StatusFound)
	})

	r.Get("/header", header.Header)

	r.Mount("/parent", parent.Resource{DB: rs.DB, Queries: rs.Queries}.Routes())

	r.Mount("/es", eventsource.SSE)

	r.Group(func(r chi.Router) {
		s := me.NewService(rs.Queries)
		r.Get("/me", s.Page)
	})

	r.Group(func(r chi.Router) {
		s := editme.NewService(rs.Queries, rs.Conn)
		r.Get("/me/edit", s.MyPageEdit)
		r.Post("/me/edit", s.Post)
		r.Get("/avatars", s.GetAvatars)
		r.Put("/avatar", s.PutAvatar)
	})

	r.Group(func(r chi.Router) {
		s := humans.NewService(rs.Queries)
		r.Get("/me/humans", s.Page)
	})
	r.Group(func(r chi.Router) {
		s := family.NewService(rs.Queries)
		r.Get("/me/family", s.Page)
	})
	r.Group(func(r chi.Router) {
		s := friends.NewService(rs.Queries)
		r.Get("/me/friends", s.Page)
	})
	r.Get("/fun", fun.Page)
	r.Get("/fun/gradients", gradients.Index)
	r.Post("/fun/gradients/picker", gradients.Picker)
	r.Post("/fun/gradients/set-background", gradients.NewService(rs.Queries, rs.Conn).SetBackground)

	r.Mount("/stickers", stickers.Resource{Conn: rs.Conn}.Routes())

	r.Get("/fun/chess", chess.Page)
	r.Get("/fun/chess/select/{rank}/{file}", chess.Select)
	r.Get("/fun/chess/unselect", chess.Unselect)
	//r.Get("/fun/chess/select/{r1}/{f1}/{r2}/{f2}", chess.Move)

	r.Route("/u/{userID}", func(r chi.Router) {
		r.Route("/", u.NewService(rs.Queries).Router)

		r.Route("/quizzes", func(r chi.Router) {
			r.Route("/create", create.NewService(rs.Queries).Router)
			r.Route("/{quizID}", func(r chi.Router) {
				r.Route("/", show.NewService(rs.Queries).Router)
				r.Route("/view", view.NewService(rs.Queries).Router)
				r.Route("/attempts/{attemptID}", func(r chi.Router) {
					s := attempt.NewService(rs.Queries)
					r.Get("/", s.Page)
					r.Get("/done", s.CompletedPage)
					r.Post("/question/{questionID}/response", s.PostResponse)
				})
			})
		})
	})

	r.Group(func(r chi.Router) {
		s := notebook.NewService(rs.Queries)
		r.Get("/fun/notes", s.Page)
		r.Post("/fun/notes", s.Post)
		r.Put("/fun/notes/{noteID}", s.Put)
		r.Delete("/fun/notes/{noteID}", s.Delete)
		r.Post("/fun/notes/from-chat/{messageID}", s.PostFromChat)
	})

	r.Mount("/bots", bots.Resource{Model: rs.Queries, AI: ai.New().Client}.Routes())

	//r.Route("/u/{userID}", u.NewService(rs.Queries).Router)

	r.Group(func(r chi.Router) {
		s := chat.NewService(rs.Queries, rs.Conn)
		r.Get("/u/{userID}/chat", s.Page)
		r.Post("/chat/messages", s.PostChatMessage)
	})

	r.Group(func(r chi.Router) {
		s := connect.NewService(rs.Queries, rs.Conn)
		r.Get("/connect", s.Connect)
		r.Put("/connect/friend/{userID}", s.PutParentFriend)
		r.Delete("/connect/friend/{userID}", s.DeleteParentFriend)
	})

	r.Group(func(r chi.Router) {
		s := connectkids.NewService(rs.Conn, rs.Queries)
		r.Get("/connectkids", s.KidConnect)
		r.Put("/connectkids/friend/{userID}", s.PutKidFriend)
		r.Delete("/connectkids/friend/{userID}", s.DeleteKidFriend)
	})

	r.Group(func(r chi.Router) {
		s := deliveries.NewService(rs.Queries, rs.Conn)
		r.Get("/deliveries/{deliveryID}", s.Page)
		r.Get("/delivery/{deliveryID}", s.Page) // temporary
		r.Post("/deliveries/{deliveryID}/logout", s.Logout)
	})

	r.Group(func(r chi.Router) {
		s := postoffice.NewService(rs.Queries)
		r.Route("/postoffice", s.Router)
	})

	return r
}
