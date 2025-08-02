package u

import (
	_ "embed"
	"fmt"
	"net/http"
	"oj/api"
	"oj/handlers/connect"
	"oj/handlers/layout"
	"oj/handlers/me"
	"oj/handlers/render"
	"oj/internal/link"
	"oj/internal/middleware/auth"
	"oj/internal/text"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/notnil/chess"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func (s *service) Page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(r.Context())

	pageUserID, _ := strconv.Atoi(chi.URLParam(r, "userID"))
	pageUser, err := s.Queries.UserByID(ctx, int64(pageUserID))
	if err != nil {
		if err == pgx.ErrNoRows {
			render.Error(w, fmt.Errorf("UserByID: %w", err), http.StatusNotFound)
			return
		}
		render.Error(w, fmt.Errorf("UserByID: %w", err), http.StatusInternalServerError)
		return
	}

	if l.User.ID == pageUser.ID {
		http.Redirect(w, r, "/me", http.StatusFound)
		return
	}

	// override layout gradient to show the page user's not the request user's
	l.BackgroundGradient = pageUser.Gradient

	connection, err := s.Queries.GetConnection(ctx, api.GetConnectionParams{
		AID: l.User.ID,
		ID:  pageUser.ID,
	})
	if err != nil {
		render.Error(w, fmt.Errorf("GetConnection: %w", err), http.StatusInternalServerError)
		return
	}

	connected := connection.RoleIn != "" && connection.RoleOut != ""

	quizzes, err := s.Queries.PublishedUserQuizzes(ctx, pageUser.ID)
	if err != nil {
		render.Error(w, fmt.Errorf("GetConnections: %w", err), http.StatusInternalServerError)
		return
	}

	chessMatches, err := s.Queries.ChessMatchesBetweenUsers(ctx, api.ChessMatchesBetweenUsersParams{
		User1ID: l.User.ID,
		User2ID: pageUser.ID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	layout.Layout(l,
		l.User.Username,
		h.Div(
			g.Attr("hx-get", fmt.Sprintf("/u/%d", pageUser.ID)),
			g.Attr("hx-trigger", "connectionChange from:body"),
			g.Attr("hx-target", "body"),
			h.Style("display:flex;flex-direction:column;gap:1em;"),
			g.If(connect.ConnectionStatus(connection.RoleIn, connection.RoleOut) != "connected",
				connect.ConnectionEl(connection.User, connection.RoleIn, connection.RoleOut)),
			me.ProfileEl(pageUser, false),
			g.If(connected, h.A(
				h.Class("nes-btn is-success"),
				h.Href(fmt.Sprintf("/u/%d/chat", pageUser.ID)),
				g.Text(fmt.Sprintf("Chat with %s", text.Shorten(pageUser.Username, 8))),
			)),
			// chess
			g.If(connected, chessButton(l.User, pageUser, chessMatches)),

			g.If(len(quizzes) > 0, me.QuizzesEl(0, quizzes)),
		)).Render(w)
}

func chessButton(currentUser api.User, user api.User, matches []api.ChessMatch) g.Node {
	if len(matches) == 0 {
		// if !currentUser.Admin {
		// 	return g.Group{}
		// }
		return h.Button(
			g.Attr("hx-post", link.User(user.ID, "chess-challenge")),
			h.Class("nes-btn"), g.Text(fmt.Sprintf("Challenge %s to a game of chess", user.Username)))
	}

	// limit to 1 match only for now
	match := matches[0]

	return h.A(
		h.Class("nes-btn is-success"), g.Text("View chess match"),
		h.Href(link.ChessMatch(match.ID)))
}

func (s *service) HandleChessChallenge(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	currentUser := auth.FromContext(ctx)
	pageUserID, _ := strconv.Atoi(chi.URLParam(r, "userID"))

	game := chess.NewGame()

	pgn, err := game.MarshalText()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	match, err := s.Queries.CreateChessMatch(ctx, api.CreateChessMatchParams{
		WhiteUserID: currentUser.ID,
		BlackUserID: int64(pageUserID),
		Pgn:         string(pgn),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("HX-Redirect", link.ChessMatch(match.ID))
}
