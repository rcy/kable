package welcome

import (
	cryptorand "crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	mathrand "math/rand"
	"net/http"
	"oj/api"
	"oj/handlers/render"
	"oj/services/email"
	"oj/services/family"
	"os"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
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

func welcomeLayout(main g.Node) g.Node {
	return h.HTML(
		h.Head(
			h.Script(h.Src("https://cdn.jsdelivr.net/npm/alpinejs@3.x.x/dist/cdn.min.js"), g.Attr("defer", "")),
			h.Link(h.Href("https://unpkg.com/nes.css@latest/css/nes.min.css"), h.Rel("stylesheet")),
			h.Link(h.Href("https://fonts.googleapis.com/css?family=Press+Start+2P"), h.Rel("stylesheet")),
			h.Meta(h.Charset("utf-8")),
			h.Meta(h.Name("viewport"), h.Content("width=device-width, initial-scale=1")),
			h.StyleEl(g.Raw(`html, body, pre, code, kbd, samp { font-family: 'Press Start 2P'; }`)),
		),
		h.Body(g.Attr("x-data", ""),
			h.Div(h.Style("display:flex; flex-direction: column; gap:2em"),
				h.Header(h.Style("background: rgba(0,0,0,.9); color:white;"),
					h.Div(h.Style("max-width: 960px; margin: 0 auto;"),
						h.Div(h.Style("padding: 1em 0 1em"),
							h.Div(h.Style("display:flex;gap:4px;align-items:center;justify-content:space-between"),
								h.A(h.Href("/"), h.Style("display:flex; align-items: center; gap:6px; color: inherit; text-decoration: none"),
									h.I(h.Class("nes-icon coin")),
									h.H3(h.Style("margin:0"), g.Text("Kable")),
								),
							),
						),
					),
				),
				h.Main(h.Style("flex:auto; height:100%; overflow: auto;"),
					h.Div(h.Style("max-width: 960px; margin: 0 auto;"),
						h.H1(g.Text("Welcome to Kable!")),
						h.P(g.Text("Kable is a place for kids to play and socialize with friends and family.")),
						main,
					),
				),
			),
		),
	)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	welcomeLayout(welcomeMain()).Render(w)
}

func welcomeMain() g.Node {
	return h.Div(h.Style("display:flex; flex-direction:column; gap:3em"),
		h.Div(h.Class("nes-container"),
			h.H2(g.Text("Kids Zone")),
			h.I(h.Class("nes-ash")),
			h.P(g.Text("Use this button if you have an account")),
			h.A(h.Class("nes-btn is-success"), h.Href("/welcome/kids"), g.Text("I am a Kid!")),
		),
		h.Div(h.Class("nes-container"),
			h.H2(g.Text("Parents")),
			h.P(g.Text("A parent account is required to use Kable.")),
			h.P(g.Text("The first step is to create an account using your email address.")),
			h.P(g.Text("Once logged in you will be able to create managed accounts for your children.")),
			h.A(h.Class("nes-btn is-primary"), h.Href("/welcome/parents"), g.Text("Sign In")),
		),
	)
}

func welcomeParents(w http.ResponseWriter, r *http.Request) {
	welcomeLayout(welcomeParentsMain()).Render(w)
}

func welcomeParentsMain() g.Node {
	return h.Div(h.Style("display:flex; flex-direction:column;gap:2em"),
		h.Div(h.Class("nes-container"),
			h.H2(g.Text("Parent Sign In "), h.Small(g.Text("Step 1: Email"))),
			h.P(g.Text("New and returning parents sign in here.")),
			h.P(h.Class("nes-text is-error"), g.Text("New sign ups are temporarily disabled.  Only existing accounts can sign in.")),
			h.Form(h.Method("post"), h.Action("/welcome/parents/email"),
				h.Label(
					g.Text("Email"),
					h.Input(g.Attr("placeholder", "human@example.com"), h.Class("nes-input"), h.Name("email"), h.Type("email"), g.Attr("required", "")),
				),
				h.Button(h.Class("nes-btn is-primary"), h.Type("submit"), g.Text("Submit")),
			),
		),
		h.Div(h.Class("box kid"),
			g.Text("If you are looking for the kids login, "),
			h.A(h.Class("nes-text is-success"), h.Href("/welcome/kids"), g.Text("click here!")),
		),
	)
}

func welcomeParentsMainWithError(errMsg string) g.Node {
	return h.Div(h.Style("display:flex; flex-direction:column;gap:2em"),
		h.Div(h.Class("nes-container"),
			h.H2(g.Text("Parent Sign In "), h.Small(g.Text("Step 1: Email"))),
			h.P(g.Text("New and returning parents sign in here.")),
			h.P(h.Class("nes-text is-error"), g.Text("New sign ups are temporarily disabled.  Only existing accounts can sign in.")),
			h.Form(h.Method("post"), h.Action("/welcome/parents/email"),
				h.Label(
					g.Text("Email"),
					h.Input(g.Attr("placeholder", "human@example.com"), h.Class("nes-input"), h.Name("email"), h.Type("email"), g.Attr("required", "")),
				),
				h.Div(h.Class("nes-text is-error"), g.Text(errMsg)),
				h.Button(h.Class("nes-btn is-primary"), h.Type("submit"), g.Text("Submit")),
			),
		),
		h.Div(h.Class("box kid"),
			g.Text("If you are looking for the kids login, "),
			h.A(h.Class("nes-text is-success"), h.Href("/welcome/kids"), g.Text("click here!")),
		),
	)
}

func parentsCode(w http.ResponseWriter, r *http.Request) {
	welcomeLayout(welcomeParentsCodeMain("")).Render(w)
}

func welcomeParentsCodeMain(errMsg string) g.Node {
	return h.Div(h.Style("display:flex; flex-direction:column;gap:2em"),
		h.Div(h.Class("nes-container"),
			h.H2(g.Text("Parent Sign In "), h.Small(g.Text("Step 2: Code"))),
			h.P(g.Text("A verification code was sent to your email.")),
			h.Form(h.Method("post"), h.Action("/welcome/parents/code"),
				h.Div(
					h.Label(
						g.Text("4 Digit Code"),
						h.Input(h.Class("nes-input"),
							g.Attr("placeholder", "Enter code..."),
							h.Name("code"),
							h.Type("text"),
							g.Attr("required", ""),
							g.Attr("x-on:keydown", "$refs.error.setHTML('')"),
						),
						h.Div(h.Class("nes-text is-error"), g.Attr("x-ref", "error"), g.Text(errMsg)),
					),
				),
				h.Button(h.Class("nes-btn is-primary"), h.Type("submit"), g.Text("Submit")),
			),
			h.P(g.Text("Didn't get a code?  "), h.A(h.Href("/welcome"), g.Text("Start over"))),
		),
	)
}

func welcomeKids(w http.ResponseWriter, r *http.Request) {
	welcomeLayout(welcomeKidsMain("")).Render(w)
}

func welcomeKidsMain(errMsg string) g.Node {
	return h.Div(h.ID("welcomeKidsSection"), h.Style("display:flex; flex-direction:column;gap:2em"),
		h.Div(h.Class("nes-container"),
			h.H2(g.Text("Kids Login - Step 1: Username")),
			h.I(h.Class("nes-pokeball")),
			h.Form(h.Method("post"), h.Action("/welcome/kids/username"),
				h.Div(
					h.Label(
						g.Text("Username"),
						h.Input(h.Class("nes-input"),
							g.Attr("placeholder", "Type your username..."),
							h.Name("username"),
							h.Type("text"),
							g.Attr("required", ""),
							g.Attr("x-on:keydown", "$refs.error.setHTML('')"),
						),
						h.Div(h.Class("nes-text is-error"), g.Attr("x-ref", "error"), g.Text(errMsg)),
					),
				),
				h.Button(h.Class("nes-btn is-success"), h.Type("submit"), g.Text("Submit")),
			),
		),
		h.Div(h.Class("nes-container is-dark"),
			h.Strong(h.Class("block titlebar"), g.Text("What if I am new?")),
			h.P(h.Style("margin-top:1em"),
				g.Text("If this is your first time here and you do not have a username, you will need to have a "),
				h.A(h.Href("/welcome/parents"), g.Text("parent register")),
				g.Text(" and create an account for you."),
			),
		),
	)
}

func kidsCode(w http.ResponseWriter, r *http.Request) {
	welcomeLayout(welcomeKidsCodeMain("")).Render(w)
}

func welcomeKidsCodeMain(errMsg string) g.Node {
	return h.Div(h.Class("nes-container"),
		h.H2(g.Text("Kids Login - Step 2: Code")),
		h.I(h.Class("nes-squirtle")),
		h.P(g.Text("Please enter the four numbers sent to your parent's email")),
		h.Form(h.Method("post"), h.Action("/welcome/kids/code"),
			h.Div(
				h.Label(
					g.Text("4 Digit Code"),
					h.Input(h.Class("nes-input"),
						g.Attr("placeholder", "Type code..."),
						h.Name("code"),
						h.Type("text"),
						g.Attr("required", ""),
						g.Attr("x-on:keydown", "$refs.error.setHTML('')"),
					),
					h.Div(h.Class("nes-text is-error"), g.Attr("x-ref", "error"), g.Text(errMsg)),
				),
			),
			h.Button(h.Class("nes-btn is-success"), h.Type("submit"), g.Text("Submit")),
		),
	)
}

func signout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "kh_session", Path: "/", Expires: time.Now().Add(-time.Hour)})
	http.Redirect(w, r, "/", http.StatusFound)
}

func generateDigitCode() string {
	if os.Getenv("AUTH_CODE") != "" {
		return os.Getenv("AUTH_CODE")
	}
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

	_, err := s.Queries.UserByEmail(ctx, pgtype.Text{String: address, Valid: true})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			welcomeLayout(welcomeParentsMainWithError("Sign ups are temporarily disabled.  If you already have an account, please double check your email address.")).Render(w)
			return
		}
		render.Error(w, fmt.Errorf("UserByEmail: %w", err), 500)
		return
	}

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

	err = email.Send(
		fmt.Sprintf("Parent sign in code: %s", code),
		fmt.Sprintf("Your Kable verification code is %s", code),
		address)
	if err != nil {
		render.Error(w, fmt.Errorf("Error emailing code: %w", err), 500)
		return
	}

	http.Redirect(w, r, "/welcome/parents/code", http.StatusSeeOther)
}

func (s *service) kidsUsernameAction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	username := r.FormValue("username")

	user, err := s.Queries.UserByUsername(r.Context(), username)
	if err != nil {
		if err == pgx.ErrNoRows {
			welcomeLayout(welcomeKidsMain("User not found")).Render(w)
			return
		}
		render.Error(w, fmt.Errorf("UserByUsername: %w", err), http.StatusInternalServerError)
		return
	}

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
			fmt.Sprintf("Your child, %s, is trying to login to Kable.  The verification code is %s.",
				username, code),
			parent.Email.String)
		if err != nil {
			render.Error(w, fmt.Errorf("email.Send: %w", err), http.StatusInternalServerError)
		}
	}

	http.Redirect(w, r, "/welcome/kids/code", http.StatusSeeOther)
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

	err = pgxscan.Get(ctx, s.Conn, &userID, "select user_id from kids_codes where nonce = $1 and code = $2", nonce, code)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			render.Error(w, fmt.Errorf("select user_id from kids_codes: %w", err), 500)
			return
		}
	}

	if userID != 0 {
		log.Println("code is good")
		user, err := s.Queries.UserByID(ctx, userID)
		if err != nil {
			render.Error(w, fmt.Errorf("UserByID: %w", err), 500)
			return
		}
		log.Printf("user %v", user)
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
		http.SetCookie(w, &http.Cookie{Name: "kh_session", Value: key, Path: "/", Expires: time.Now().Add(365 * 24 * time.Hour)})
		http.SetCookie(w, &http.Cookie{Name: "kh_nonce", Path: "/", Expires: time.Now().Add(-time.Hour)})
		http.Redirect(w, r, "/", 303)
	} else {
		log.Println("code is bad")
		welcomeLayout(welcomeKidsCodeMain("bad code, try again")).Render(w)
	}
}

func (s *service) parentsCodeAction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var emailAddr string

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

	err = pgxscan.Get(ctx, s.Conn, &emailAddr, "select email from codes where nonce = $1 and code = $2", nonce, code)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			render.Error(w, fmt.Errorf("select email from codes: %w", err), 500)
			return
		}
	}

	if emailAddr != "" {
		log.Println("code is good")
		user, err := family.FindOrCreateParentByEmail(ctx, s.Queries, emailAddr)
		if err != nil {
			render.Error(w, fmt.Errorf("FindOrCreateParentByEmail: %w", err), 500)
			return
		}
		log.Printf("user %v", user)
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
		http.SetCookie(w, &http.Cookie{Name: "kh_session", Value: key, Path: "/", Expires: time.Now().Add(365 * 30 * 24 * time.Hour)})
		http.SetCookie(w, &http.Cookie{Name: "kh_nonce", Path: "/", Expires: time.Now().Add(-time.Hour)})
		http.Redirect(w, r, "/", 303)
	} else {
		log.Println("code is bad")
		welcomeLayout(welcomeParentsCodeMain("bad code, try again")).Render(w)
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
