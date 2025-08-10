package parent

import (
	_ "embed"
	"errors"
	"fmt"
	"log"
	"net/http"
	"oj/api"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/link"
	"oj/internal/middleware/auth"
	"oj/services/family"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jmoiron/sqlx"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type Resource struct {
	DB      *sqlx.DB
	Queries *api.Queries
}

func (rs Resource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", rs.index)
	r.Post("/kids", rs.createKid)
	r.Delete("/kids/{userID}", rs.deleteKid)
	r.Post("/kids/{userID}/logout", rs.logoutKid)
	return r
}

func (rs Resource) index(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(ctx)

	kids, err := rs.Queries.KidsByParentID(ctx, l.User.ID)
	if err != nil {
		render.Error(w, fmt.Errorf("KidsByParentID: %w", err), 500)
		return
	}

	layout.Layout(l, "parent",
		h.Div(
			h.Style("display:flex; flex-direction: column; gap:1em; margin-bottom: 50%"),
			h.Div(
				h.Class("nes-container is-dark"),
				h.Style("display:flex; flex-direction:column; gap:1em"),
				h.H1(
					g.Text("Hello parent!"),
					h.Small(
						g.Text(l.User.Email.String),
					),
				),
				h.P(
					g.Text("Here you can add managed accounts for your kids."),
				),
				h.P(
					g.Text("You are the manager for these accounts.  You can remove them and\n        all associated data at any time."),
				),
				h.P(
					g.Text("Choose a unique username for your child.  It can contain their name, but doesn't have to.  They will be able to change it to whatever they want when they login."),
				),
				h.Div(
					h.Class("nes-container is-dark"),
					h.Form(
						h.Action("/parent/kids"),
						h.Method("post"),
						h.Label(
							g.Text("Child's Username"),
							h.Input(
								h.Class("nes-input"),
								h.Type("text"),
								h.Name("username"),
							),
						),
						h.Button(
							h.Class("nes-btn is-primary"),
							g.Text("Add Kid"),
						),
					),
				),
				h.P(
					g.Text(fmt.Sprintf("Kids login with their username and a one time code that will be emailed to %s.",
						l.User.Email)),
				),
			),
			g.Map(kids, func(kid api.User) g.Node {
				return h.Div(
					h.Class("nes-container ghost kid"),
					h.Div(
						h.Style("display:flex; justify-content:space-between"),
						h.A(
							h.Href(link.User(l.User.ID)),
							h.Style("display: flex; gap:1em"),
							h.Img(
								h.Width("100"),
								h.Src(kid.Avatar.URL()),
							),
							h.Div(
								h.Style("display:flex; flex-direction: column"),
								h.H2(
									g.Text(fmt.Sprintf("username: %s", kid.Username)),
								),
								h.Div(
									g.Text(fmt.Sprintf("Joined %s", kid.CreatedAt.Time.Format(time.DateOnly))),
								),
							),
						),
						h.Div(
							h.Button(
								h.Class("nes-btn is-error"),
								g.Attr("hx-delete", link.ParentKid(kid.ID)),
								g.Attr("hx-confirm", fmt.Sprintf("Permanently delete %s and all associated data?", kid.Username)),
								g.Attr("hx-target", "closest .kid"),
								g.Attr("hx-swap", "outerHTML"),
								g.Text("delete"),
							),
						),
					),
				)
			}),
		)).Render(w)
}

func (rs Resource) createKid(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := auth.FromContext(ctx)

	err := r.ParseForm()
	if err != nil {
		render.Error(w, fmt.Errorf("ParseForm: %w", err), 500)
		return
	}
	username := r.PostForm.Get("username")

	_, err = rs.Queries.UserByUsername(ctx, username)
	if errors.Is(err, pgx.ErrNoRows) {
		kid, err := family.CreateKid(ctx, rs.Queries, user.ID, username)
		if err != nil {
			render.Error(w, fmt.Errorf("CreateKid: %w", err), 500)
			return
		}
		log.Printf("kid: %v", kid)
		http.Redirect(w, r, "/parent", http.StatusSeeOther)
		return
	}

	if err != nil {
		render.Error(w, fmt.Errorf("UserByUsername: %w", err), 500)
		return
	}

	render.Error(w, fmt.Errorf("username taken"), http.StatusConflict)
	return
}

func (rs Resource) deleteKid(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	tx, err := rs.DB.Beginx()
	if err != nil {
		render.Error(w, fmt.Errorf("Beginx: %w", err), http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
delete from kids_parents where kid_id = ?;
delete from bios where user_id = ?;
delete from deliveries where sender_id = ? or recipient_id = ?;
delete from gradients where user_id = ?;
delete from kids_codes where user_id = ?;
delete from messages where sender_id = ?;
delete from room_users where user_id = ?;
delete from sessions where user_id = ?;
delete from users where id = ?;
`, userID, userID, userID, userID, userID, userID, userID, userID, userID)
	if err != nil {
		render.Error(w, fmt.Errorf("delete delete delete: %w", err), http.StatusInternalServerError)
		return
	}

	err = tx.Commit()
	if err != nil {
		render.Error(w, fmt.Errorf("Commit: %w", err), http.StatusInternalServerError)
		return
	}
}

func (rs Resource) logoutKid(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	_, err := rs.DB.Exec(`delete from sessions where user_id = $1`, userID)
	if err != nil {
		render.Error(w, fmt.Errorf("delete from sessions: %w", err), http.StatusInternalServerError)
		return
	}
}
