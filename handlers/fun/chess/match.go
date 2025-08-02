package chess

import (
	"context"
	"fmt"
	"net/http"
	"oj/api"
	"oj/handlers/layout"
	"oj/internal/link"
	"oj/internal/middleware/auth"
	"strconv"

	"github.com/go-chi/chi/v5"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func (s *service) HandleMatch(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(ctx)
	matchID, _ := strconv.Atoi(chi.URLParam(r, "matchID"))
	match, err := s.Queries.ChessMatchByID(ctx, int64(matchID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	board, err := gameStateFromMatch(match)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	currentUser := auth.FromContext(ctx)
	opponent, err := matchOpponent(ctx, s.Queries, match, currentUser.ID)

	uiGameState := UIGameState{gameState: board, flip: l.User.ID == match.BlackUserID}
	elMatchID := fmt.Sprintf("match-%d", match.ID)

	layout.Layout(l, "chess club",
		h.Div(h.ID(elMatchID),
			g.Attr("hx-get", link.ChessMatch(match.ID)),
			g.Attr("hx-trigger", "USER_UPDATE from:body"),
			g.Attr("hx-target", "#"+elMatchID),
			g.Attr("hx-select", "#"+elMatchID),
			g.Attr("hx-swap", "outerHTML"),
			h.H4(h.Style("padding-left: 1em; background:white; foreground:black"),
				g.Text(fmt.Sprintf("%s v %s", opponent.Username, currentUser.Username))),
			h.Div(h.Style("display:flex; gap:1em"),
				h.Div(h.ID("board-container"), h.Style("height: 80vh; width: 80vh;"),
					uiGameState,
				),
				h.Div(h.Style("flex:1"), g.Text(uiGameState.gameState.Game.String())),
			),
		),
	).Render(w)
}

func matchOpponent(ctx context.Context, qtx *api.Queries, match api.ChessMatch, userID int64) (*api.User, error) {
	opponentID := match.WhiteUserID
	if userID == match.WhiteUserID {
		opponentID = match.BlackUserID
	}
	opponent, err := qtx.UserByID(ctx, opponentID)
	if err != nil {
		return nil, fmt.Errorf("UserByID: %w", err)
	}
	return &opponent, nil
}
