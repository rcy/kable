package chess

import (
	"net/http"
	"oj/handlers/layout"
	"strconv"

	"github.com/go-chi/chi/v5"
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

	board, err := gameBoardFromMatch(match, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	layout.Layout(l, "chess club", chessPageEl(board)).Render(w)
}
