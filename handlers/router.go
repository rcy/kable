package handlers

import (
	"errors"
	"log"
	"net/http"
	"oj/api"
	"oj/handlers/admin"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/handlers/welcome"
	"oj/internal/app"
	"oj/internal/middleware/auth"
	"oj/internal/middleware/become"
	"oj/internal/middleware/redirect"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Router(conn *pgxpool.Pool) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)

	middleware.DefaultLogger = middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: log.New(os.Stdout, "", log.LstdFlags), NoColor: true})
	r.Use(middleware.Logger)

	queries := api.New(conn)

	// authenticated routes
	r.Route("/", func(r chi.Router) {
		r.Use(auth.NewService(queries).Provider)
		r.Use(become.NewService(queries).Provider)
		r.Use(redirect.Redirect)
		r.Use(layout.NewService(queries, conn).Provider)
		r.Mount("/", app.Service{Conn: conn, Queries: queries}.Routes())
		r.Mount("/admin", admin.NewService(queries, conn).Routes())
	})

	// non authenticated routes
	r.Group(func(r chi.Router) {
		s := welcome.NewService(queries, conn)
		r.Route("/welcome", s.Route)
	})

	r.Get("/.well-known/*", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	// serve static files
	fs := http.FileServer(http.Dir("assets"))
	r.Handle("/assets/*", http.StripPrefix("/assets", fs))

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "assets/favicon.ico")
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		render.Error(w, errors.New("Page not found"), 404)
	})
	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		render.Error(w, errors.New("Method not allowed"), 405)
	})

	return r
}
