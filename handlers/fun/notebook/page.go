package notebook

import (
	"database/sql"
	_ "embed"
	"fmt"
	"net/http"
	"oj/api"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type service struct {
	Queries *api.Queries
}

func NewService(q *api.Queries) *service {
	return &service{Queries: q}
}

var (
	//go:embed page.gohtml
	pageContent  string
	pageTemplate = layout.MustParse(pageContent)
)

func (s *service) Page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(ctx)

	allNotes, err := s.Queries.UserNotes(ctx, l.User.ID)
	if err != nil && err != sql.ErrNoRows {
		render.Error(w, fmt.Errorf("UserNotes: %w", err), http.StatusInternalServerError)
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

func (s *service) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := auth.FromContext(ctx)
	note, err := s.Queries.CreateNote(ctx, api.CreateNoteParams{
		OwnerID: user.ID,
	})
	if err != nil {
		render.Error(w, fmt.Errorf("CreateNote: %w", err), http.StatusInternalServerError)
		return
	}
	render.ExecuteNamed(w, pageTemplate, "note", note)
}

func (s *service) PostFromChat(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := auth.FromContext(ctx)

	messageID, _ := strconv.Atoi(chi.URLParam(r, "messageID"))

	msg, err := s.Queries.MessageByID(ctx, int64(messageID))
	if err != nil {
		render.Error(w, fmt.Errorf("MessageByID: %w", err), http.StatusInternalServerError)
		return
	}

	sender, err := s.Queries.UserByID(ctx, msg.SenderID)
	if err != nil {
		render.Error(w, fmt.Errorf("UserByID: %w", err), http.StatusInternalServerError)
		return
	}

	body := fmt.Sprintf("From: %s\n%s", sender.Username, msg.Body)

	note, err := s.Queries.CreateNote(ctx, api.CreateNoteParams{
		OwnerID: user.ID,
		Body:    body,
	})
	if err != nil {
		render.Error(w, fmt.Errorf("CreateNote: %w", err), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("saved note %d", note.ID)))
}

func (s *service) Put(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("body")
	ctx := r.Context()
	noteID, _ := strconv.Atoi(chi.URLParam(r, "noteID"))
	user := auth.FromContext(ctx)
	_, err := s.Queries.UpdateNote(ctx, api.UpdateNoteParams{
		ID:      int64(noteID),
		OwnerID: user.ID,
		Body:    body,
	})
	if err != nil {
		render.Error(w, fmt.Errorf("UpdateNote: %w", err), http.StatusInternalServerError)
		return
	}
}

func (s *service) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	noteID, _ := strconv.Atoi(chi.URLParam(r, "noteID"))
	user := auth.FromContext(ctx)
	err := s.Queries.DeleteNote(ctx, api.DeleteNoteParams{
		ID:      int64(noteID),
		OwnerID: user.ID,
	})
	if err != nil {
		render.Error(w, fmt.Errorf("DeleteNote: %w", err), http.StatusInternalServerError)
		return
	}
}
