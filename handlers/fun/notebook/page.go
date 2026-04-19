package notebook

import (
	"fmt"
	"net/http"
	"oj/api"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type service struct {
	Queries *api.Queries
}

func NewService(q *api.Queries) *service {
	return &service{Queries: q}
}

func (s *service) Page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(ctx)

	allNotes, err := s.Queries.UserNotes(ctx, l.User.ID)
	if err != nil && err != pgx.ErrNoRows {
		render.Error(w, fmt.Errorf("UserNotes: %w", err), http.StatusInternalServerError)
		return
	}

	layout.Layout(l, "Notebook", notebookPage(allNotes)).Render(w)
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
	noteEl(note).Render(w)
}

func notebookPage(notes []api.Note) g.Node {
	return h.Div(h.Style("display:flex; flex-direction: column; gap: 1em"),
		h.Div(h.Class("nes-container ghost"),
			h.H1(g.Text("Notebook")),
		),
		h.Div(h.Style("display:flex; justify-content:flex-start"),
			h.Button(h.Class("nes-btn is-success"),
				g.Attr("hx-post", "/fun/notes"),
				g.Attr("hx-swap", "afterbegin"),
				g.Attr("hx-target", "#notes"),
				g.Text("new note"),
			),
		),
		h.Div(h.ID("notes"), h.Style("display:flex; flex-direction:column; gap:1em"),
			g.Map(notes, func(note api.Note) g.Node { return noteEl(note) }),
		),
	)
}

func noteEl(note api.Note) g.Node {
	return h.Div(h.Style("display:flex; gap:1em"), h.Class("note"),
		h.Div(h.Style("flex:1"),
			h.Textarea(
				h.Class("nes-textarea"),
				g.Attr("placeholder", "type your note here"),
				g.Attr("rows", "5"),
				g.Attr("hx-put", fmt.Sprintf("/fun/notes/%d", note.ID)),
				g.Attr("hx-trigger", "keyup changed delay:500ms"),
				h.Name("body"),
				g.Text(note.Body),
			),
		),
		h.Div(
			h.Button(h.Class("nes-btn is-error"),
				g.Attr("hx-delete", fmt.Sprintf("/fun/notes/%d", note.ID)),
				g.Attr("hx-target", "closest .note"),
				g.Attr("hx-swap", "outerHTML"),
				g.Attr("hx-confirm", "really delete this note?"),
				g.Text("-"),
			),
		),
	)
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
