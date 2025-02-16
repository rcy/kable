package notebook

import (
	"database/sql"
	_ "embed"
	"fmt"
	"net/http"
	"oj/api"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var (
	//go:embed page.gohtml
	pageContent  string
	pageTemplate = layout.MustParse(pageContent)
)

func Page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(ctx)
	queries := api.New(db.DB)

	allNotes, err := queries.UserNotes(ctx, l.User.ID)
	if err != nil && err != sql.ErrNoRows {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Execute(w, pageTemplate, struct {
		Layout layout.Data
		Notes  []api.Note
	}{
		Layout: l,
		Notes:  allNotes,
	})
}

func Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := auth.FromContext(ctx)
	queries := api.New(db.DB)
	note, err := queries.CreateNote(ctx, api.CreateNoteParams{
		OwnerID: user.ID,
	})
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	render.ExecuteNamed(w, pageTemplate, "note", note)
}

func PostFromChat(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := auth.FromContext(ctx)
	queries := api.New(db.DB)

	messageID, _ := strconv.Atoi(chi.URLParam(r, "messageID"))

	msg, err := queries.MessageByID(ctx, int64(messageID))
	if err != nil {
		render.Error(w, "messageByID: "+err.Error(), http.StatusInternalServerError)
		return
	}

	sender, err := queries.UserByID(ctx, msg.SenderID)
	if err != nil {
		render.Error(w, "UserByID: "+err.Error(), http.StatusInternalServerError)
		return
	}

	body := fmt.Sprintf("From: %s\n%s", sender.Username, msg.Body)

	note, err := queries.CreateNote(ctx, api.CreateNoteParams{
		OwnerID: user.ID,
		Body:    body,
	})
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("saved note %d", note.ID)))
}

func Put(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("body")
	ctx := r.Context()
	noteID, _ := strconv.Atoi(chi.URLParam(r, "noteID"))
	user := auth.FromContext(ctx)
	queries := api.New(db.DB)
	_, err := queries.UpdateNote(ctx, api.UpdateNoteParams{
		ID:      int64(noteID),
		OwnerID: user.ID,
		Body:    body,
	})
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	noteID, _ := strconv.Atoi(chi.URLParam(r, "noteID"))
	user := auth.FromContext(ctx)
	queries := api.New(db.DB)
	err := queries.DeleteNote(ctx, api.DeleteNoteParams{
		ID:      int64(noteID),
		OwnerID: user.ID,
	})
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
