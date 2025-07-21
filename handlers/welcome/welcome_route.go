package welcome

import (
	cryptorand "crypto/rand"
	_ "embed"
	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	mathrand "math/rand"
	"net/http"
	"oj/api"
	"oj/handlers/render"
	"oj/services/email"
	"oj/services/family"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type service struct {
	Conn    *pgxpool.Pool
	Queries *api.Queries
}

func NewService(q *api.Queries, conn *pgxpool.Pool) *service {
	return &service{Queries: q, Conn: conn}
}

func (s *service) Route(r chi.Router) {
	r.Get("/", welcome)

	r.Get("/parents", welcomeParents)
	r.Post("/parents/email", s.emailRegisterAction)
	r.Get("/parents/code", parentsCode)
	r.Post("/parents/code", s.parentsCodeAction)

	r.Get("/kids", welcomeKids)
	r.Post("/kids/username", s.kidsUsernameAction)
	r.Get("/kids/code", kidsCode)
	r.Post("/kids/code", s.kidsCodeAction)

	r.Get("/signout", signout)
}

//go:embed layout.gohtml
var layoutContent string

func mustLayout(content string) *template.Template {
	return template.Must(template.New("").Parse(layoutContent + content))
}

//go:embed welcome.gohtml
var welcomeContent string
var welcomeTemplate = mustLayout(welcomeContent)

func welcome(w http.ResponseWriter, r *http.Request) {
	err := welcomeTemplate.Execute(w, nil)
	if err != nil {
		render.Error(w, fmt.Errorf("welcomeTemplate.Execute: %w", err), 500)
		return
	}
}

//go:embed welcome_kids.gohtml
var welcomeKidsContent string
var welcomeKidsTemplate = mustLayout(welcomeKidsContent)

func welcomeKids(w http.ResponseWriter, r *http.Request) {
	err := welcomeKidsTemplate.Execute(w, struct{ Error string }{""})
	if err != nil {
		render.Error(w, fmt.Errorf("welcomeKidsTemplate.Execute: %w", err), 500)
		return
	}
}

//go:embed welcome_parents.gohtml
var welcomeParentsContent string
var welcomeParentsTemplate = mustLayout(welcomeParentsContent)

func welcomeParents(w http.ResponseWriter, r *http.Request) {
	err := welcomeParentsTemplate.Execute(w, nil)
	if err != nil {
		render.Error(w, fmt.Errorf("welcomeKidsTemplate.Execute: %w", err), 500)
		return
	}
}

func signout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "kh_session", Path: "/", Expires: time.Now().Add(-time.Hour)})
	http.Redirect(w, r, "/", http.StatusFound)
}

func generateDigitCode() string {
	code := ""
	for i := 0; i < 4; i++ {
		digit := mathrand.Intn(10)
		code += fmt.Sprint(digit)
	}

	return code
}

func (s *service) emailRegisterAction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	address := r.FormValue("email")
	if address == "" {
		http.Redirect(w, r, "/welcome/parents", http.StatusSeeOther)
		return
	}

	// store generated code in pending registrations table along with email
	nonce, err := generateSecureString(32)
	if err != nil {
		render.Error(w, fmt.Errorf("generateSecureString: %w", err), 500)
		return
	}
	code := generateDigitCode()
	_, err = s.Conn.Exec(ctx, "insert into codes(nonce, email, code) values($1, $2, $3)", nonce, address, code)
	if err != nil {
		render.Error(w, fmt.Errorf("insert into codes: %w", err), 500)
		return
	}

	http.SetCookie(w, &http.Cookie{Name: "kh_nonce", Value: nonce, Path: "/", Expires: time.Now().Add(time.Hour)})

	// email code to user
	err = email.Send(
		fmt.Sprintf("Parent sign in code: %s", code),
		fmt.Sprintf("Your JuiceBox verification code is %s", code),
		address)
	if err != nil {
		render.Error(w, fmt.Errorf("Error emailing code: %w", err), 500)
		return
	}

	// redirect to page to input code
	http.Redirect(w, r, "/welcome/parents/code", http.StatusSeeOther)
}

//go:embed welcome_parents_code.gohtml
var welcomeParentsCodeContent string
var welcomeParentsCodeTemplate = mustLayout(welcomeParentsCodeContent)

func parentsCode(w http.ResponseWriter, r *http.Request) {
	err := welcomeParentsCodeTemplate.Execute(w, struct{ Error string }{""})
	if err != nil {
		render.Error(w, fmt.Errorf("welcomeParentsCodeTemplate.Execute: %w", err), 500)
		return
	}
}

func (s *service) kidsUsernameAction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	username := r.FormValue("username")

	user, err := s.Queries.UserByUsername(r.Context(), username)
	if err != nil {
		if err == pgx.ErrNoRows {
			err = welcomeKidsTemplate.Execute(w, struct{ Error string }{"User not found"})
			if err != nil {
				render.Error(w, fmt.Errorf("welcomeKidsTemplate.Execute: %w", err), http.StatusInternalServerError)
			}
			return
		}
		render.Error(w, fmt.Errorf("UserByUsername: %w", err), http.StatusInternalServerError)
		return
	}

	// store generated code in pending registrations table along with email
	nonce, err := generateSecureString(32)
	if err != nil {
		render.Error(w, fmt.Errorf("generateSecureString: %w", err), 500)
		return
	}
	code := generateDigitCode()
	_, err = s.Conn.Exec(ctx, "insert into kids_codes(nonce, user_id, code) values($1, $2, $3)", nonce, user.ID, code)
	if err != nil {
		render.Error(w, fmt.Errorf("insert into kids_codes: %w", err), 500)
		return
	}

	http.SetCookie(w, &http.Cookie{Name: "kh_nonce", Value: nonce, Path: "/", Expires: time.Now().Add(time.Hour)})

	// email code to kids parent(s)
	parents, err := s.Queries.ParentsByKidID(ctx, user.ID)
	if err != nil {
		render.Error(w, fmt.Errorf("ParentsByKidID: %w", err), 500)
		return
	}

	if len(parents) == 0 {
		render.Error(w, fmt.Errorf("No parents"), 500)
		return
	}

	for _, parent := range parents {
		err = email.Send(
			fmt.Sprintf("Code for %s is %s", username, code),
			fmt.Sprintf("Your child, %s, is trying to login to JuiceBox.  The verification code is %s.",
				username, code),
			parent.Email.String)
		if err != nil {
			render.Error(w, fmt.Errorf("email.Send: %w", err), http.StatusInternalServerError)
		}
	}

	// redirect to page to input code
	http.Redirect(w, r, "/welcome/kids/code", http.StatusSeeOther)
}

//go:embed welcome_kids_code.gohtml
var welcomeKidsCodeContent string
var welcomeKidsCodeTemplate = mustLayout(welcomeKidsCodeContent)

func kidsCode(w http.ResponseWriter, r *http.Request) {
	err := welcomeKidsCodeTemplate.Execute(w, struct{ Error string }{""})
	if err != nil {
		render.Error(w, fmt.Errorf("welcomKidsCodeTemplate.Execute: %w", err), 500)
	}
}

func (s *service) kidsCodeAction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var userID int64

	cookie, err := r.Cookie("kh_nonce")
	if err != nil {
		if err != http.ErrNoCookie {
			log.Printf("weird error 792pR3LQagv5ej3Xi %s", err)
		}
		http.Redirect(w, r, "/welcome/parents", 303)
		return
	}

	nonce := cookie.Value
	code := r.FormValue("code")

	// look up code
	// XXX fetch by id alone, compare code, and add retry count
	err = pgxscan.Get(ctx, s.Conn, &userID, "select user_id from kids_codes where nonce = $1 and code = $2", nonce, code)
	if err != nil {
		if err != pgx.ErrNoRows {
			render.Error(w, fmt.Errorf("select user_id from kids_codes: %w", err), 500)
			return
		}
	}

	if userID != 0 {
		log.Println("code is good")
		// found email, code is good
		// create user if not exists
		user, err := s.Queries.UserByID(ctx, userID)
		if err != nil {
			render.Error(w, fmt.Errorf("UserByID: %w", err), 500)
			return
		}
		log.Printf("user %v", user)
		// create a new session
		key, err := generateSecureString(32)
		if err != nil {
			render.Error(w, fmt.Errorf("generateSecureString: %w", err), 500)
			return
		}
		_, err = s.Conn.Exec(ctx, "insert into sessions(key, user_id) values($1, $2)", key, user.ID)
		if err != nil {
			render.Error(w, fmt.Errorf("error creating session: %w", err), 500)
			return
		}
		// set session cookie
		http.SetCookie(w, &http.Cookie{Name: "kh_session", Value: key, Path: "/", Expires: time.Now().Add(365 * 24 * time.Hour)})
		// clear nonce cookie
		http.SetCookie(w, &http.Cookie{Name: "kh_nonce", Path: "/", Expires: time.Now().Add(-time.Hour)})

		// redirect
		http.Redirect(w, r, "/", 303)
	} else {
		log.Println("code is bad")
		// code is bad
		err = welcomeKidsCodeTemplate.Execute(w, struct{ Error string }{"bad code, try again"})
		if err != nil {
			render.Error(w, fmt.Errorf("Execute: %w", err), 500)
			return
		}
	}
}

func (s *service) parentsCodeAction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var email string

	cookie, err := r.Cookie("kh_nonce")
	if err != nil {
		if err != http.ErrNoCookie {
			log.Printf("weird error 792pR3LQagv5ej3Xi %s", err)
		}
		http.Redirect(w, r, "/welcome/parents", 303)
		return
	}

	nonce := cookie.Value
	code := r.FormValue("code")

	// look up code
	// XXX fetch by id alone, compare code, and add retry count
	err = pgxscan.Get(ctx, s.Conn, &email, "select email from codes where nonce = $1 and code = $2", nonce, code)
	if err != nil {
		if err != pgx.ErrNoRows {
			render.Error(w, fmt.Errorf("select email from codes: %w", err), 500)
			return
		}
	}

	if email != "" {
		log.Println("code is good")
		// found email, code is good
		// create user if not exists
		user, err := family.FindOrCreateParentByEmail(ctx, s.Queries, email)
		if err != nil {
			render.Error(w, fmt.Errorf("FindOrCreateParentByEmail: %w", err), 500)
			return
		}
		log.Printf("user %v", user)
		// create a new session
		key, err := generateSecureString(32)
		if err != nil {
			render.Error(w, fmt.Errorf("generateSecureString: %w", err), 500)
			return
		}
		_, err = s.Conn.Exec(ctx, "insert into sessions(key, user_id) values($1, $2)", key, user.ID)
		if err != nil {
			render.Error(w, fmt.Errorf("error creating session: %w", err), 500)
			return
		}
		// set session cookie
		http.SetCookie(w, &http.Cookie{Name: "kh_session", Value: key, Path: "/", Expires: time.Now().Add(365 * 30 * 24 * time.Hour)})
		// clear nonce cookie
		http.SetCookie(w, &http.Cookie{Name: "kh_nonce", Path: "/", Expires: time.Now().Add(-time.Hour)})

		// redirect
		http.Redirect(w, r, "/", 303)
	} else {
		// code is bad
		log.Println("code is bad")
		err := welcomeParentsCodeTemplate.Execute(w, struct{ Error string }{"bad code, try again"})
		if err != nil {
			render.Error(w, fmt.Errorf("Execute: %w", err), 500)
		}
		//http.Redirect(w, r, "/welcome/parents/code?retry", 303)

	}
}

func generateSecureString(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := cryptorand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	randomString := base64.RawURLEncoding.EncodeToString(randomBytes)
	return randomString, nil
}
