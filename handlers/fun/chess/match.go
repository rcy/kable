package chess

import (
	"fmt"
	"net/http"
	"oj/api"
	"oj/handlers/layout"
	"oj/internal/link"
	"strconv"

	"github.com/go-chi/chi/v5"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func (s *service) HandleMatch(w http.ResponseWriter, r *http.Request) {
	l := layout.FromContext(r.Context())
	ctx := r.Context()
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

	layout.Layout(l, "chess club", chessPageNode(match, board)).Render(w)
}

func chessPageNode(match api.ChessMatch, gameState *GameState) g.Node {
	elMatchID := fmt.Sprintf("match-%d", match.ID)
	return h.Div(h.ID(elMatchID),
		g.Attr("hx-get", link.ChessMatch(match.ID)),
		g.Attr("hx-trigger", "USER_UPDATE from:body"),
		g.Attr("hx-target", "#"+elMatchID),
		g.Attr("hx-select", "#"+elMatchID),
		g.Attr("hx-swap", "outerHTML"),
		h.H1(g.Text("chess club")),
		h.Div(h.ID("board-container"), h.Style("height: 80vh; width: 80vh;"),
			gameState,
		))
}
