package chess

import (
	_ "embed"
	"errors"
	"fmt"
	"io"
	"net/http"
	"oj/api"
	"oj/internal/link"
	"oj/internal/middleware/auth"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/notnil/chess"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type GameState struct {
	Board      [8][8]*UISquare
	Game       *chess.Game
	ValidMoves []*chess.Move
	MatchID    int64
}

type UISquare struct {
	SVGPiece string
	Selected bool
	Action   string
	Dot      bool
	Rank     int
	File     int
	MatchID  int64
}

func (s UISquare) Occupied() bool {
	return s.SVGPiece != ""
}

func (s UISquare) Render(w io.Writer) error {
	white := (s.File+s.Rank)%2 == 0

	var background string = "rgba(100,100,100,1)" // default black
	if s.Selected {
		background = "orange"
	} else if white {
		background = "rgba(255,255,255,1)"
	}

	style := "height: 100%;"
	style += fmt.Sprintf("background-color: %s;", background)

	return h.Div(
		h.Style(style),
		g.If(s.Action != "", g.Group{
			g.Attr("hx-post", link.ChessMatch(s.MatchID, s.Action)),
			g.Attr("hx-target", "closest .board"),
		}),

		// occupied square
		g.If(s.Occupied(),
			h.Div(g.If(s.Dot, h.Style("background: radial-gradient(rgba(0,0,0,0) 80%, orange 80%)")),
				h.Img(h.Width("100%"), h.Src(s.SVGPiece)))),

		// empty square
		g.If(!s.Occupied() && s.Dot, h.Div(
			h.Style("display:flex; height: 100%; align-items: center; justify-content: center"),
			h.SVG(
				g.Attr("viewBox", "0 0 100 100"),
				g.Attr("xmlns", "http://www.w3.org/2000/svg"),
				g.El("circle",
					g.Attr("cx", "50"),
					g.Attr("cy", "50"),
					g.Attr("r", "20"),
					g.Attr("fill", "orange"),
				),
			),
		)),
	).Render(w)
}

func (s GameState) Render(w io.Writer) error {
	rows := make([][]g.Node, 0, 8)
	for rank := 0; rank < 8; rank += 1 {
		row := make([]g.Node, 0, 8)
		for file := 0; file < 8; file += 1 {
			row = append(row, s.Board[rank][file])
		}
		rows = append(rows, row)
	}

	return h.Div(
		h.Class("board"),
		h.Div(h.Style("height:100%; width:100%; display:flex; flex-direction:column"),
			g.Map(rows, func(row []g.Node) g.Node {
				return h.Div(h.Style("flex:1; display:flex"),
					g.Map(row, func(square g.Node) g.Node {
						return h.Div(h.Style("flex:1"), square)
					}))
			})),
	).Render(w)
}

func chessPageNode(gameState *GameState) g.Node {
	return h.Div(
		h.H1(g.Text("chess club")),
		h.Div(h.ID("board-container"), h.Style("height: 80vh; width: 80vh;"),
			gameState,
		))
}

func gameStateFromMatch(match api.ChessMatch) (*GameState, error) {
	reader := strings.NewReader(match.Pgn)
	fn, err := chess.PGN(reader)
	if err != nil {
		return nil, err
	}
	game := chess.NewGame(fn)

	pos := game.Position()

	svgPiece := [13]string{
		"", // empty piece
		"/assets/chess/wK.svg",
		"/assets/chess/wQ.svg",
		"/assets/chess/wR.svg",
		"/assets/chess/wB.svg",
		"/assets/chess/wN.svg",
		"/assets/chess/wP.svg",
		"/assets/chess/bK.svg",
		"/assets/chess/bQ.svg",
		"/assets/chess/bR.svg",
		"/assets/chess/bB.svg",
		"/assets/chess/bN.svg",
		"/assets/chess/bP.svg",
	}

	state := GameState{Game: game}

	squareMap := pos.Board().SquareMap()

	for i := 0; i < 64; i++ {
		piece := squareMap[chess.Square(i)]
		rank := 7 - i/8
		file := i % 8
		state.Board[rank][file] = &UISquare{
			SVGPiece: svgPiece[piece],
			Rank:     rank,
			File:     file,
			MatchID:  match.ID,
			Action:   fmt.Sprintf("select?rank=%d&file=%d", rank, file),
		}
	}

	return &state, nil
}

var ErrSquareNotOccupied = errors.New("no piece on square")

func (s *GameState) selectSquare(rank, file int) error {
	var selected *UISquare

	for _, row := range s.Board {
		for _, square := range row {
			if square.Rank == rank && square.File == file {
				if square.Occupied() {
					selected = square
					square.Selected = true
					square.Action = "unselect"
				}
			}
		}
	}

	if selected == nil {
		return ErrSquareNotOccupied
	}

	// add the dots to the places the peice on the selected square can move to
	moves := s.Game.ValidMoves()

	selectedSquare := chess.Square((7-selected.Rank)*8 + selected.File)
	fmt.Println("DEBUGX L6yq 1")

	for _, move := range moves {
		if move.S1() == selectedSquare {
			s.ValidMoves = append(s.ValidMoves, move)
		}
	}

	for _, move := range s.ValidMoves {
		target := move.S2()
		square := s.Board[7-target/8][target%8]
		square.Dot = true
		square.Action = fmt.Sprintf("move?s1=%d&s2=%d", move.S1(), move.S2())
	}

	return nil
}

func (s *service) HandleSelect(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	currentUser := auth.FromContext(ctx)
	matchID, _ := strconv.Atoi(chi.URLParam(r, "matchID"))
	match, err := s.Queries.ChessMatchByID(ctx, int64(matchID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	state, err := gameStateFromMatch(match)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	currentUserColor, err := userMatchColor(currentUser, match)
	if err != nil {
		http.Error(w, "userMatchColor: "+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(state.Game.Position().Turn(), currentUserColor)
	if state.Game.Position().Turn() != currentUserColor {
		state.Render(w)
		return
	}

	rank, _ := strconv.Atoi(r.FormValue("rank"))
	file, _ := strconv.Atoi(r.FormValue("file"))

	err = state.selectSquare(rank, file)
	if err != nil {
		if !errors.Is(err, ErrSquareNotOccupied) {
			http.Error(w, "selectSquare: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	state.Render(w)
}

func userMatchColor(user api.User, match api.ChessMatch) (chess.Color, error) {
	if user.ID == match.BlackUserID {
		return chess.Black, nil
	}
	if user.ID == match.WhiteUserID {
		return chess.White, nil
	}
	return chess.NoColor, fmt.Errorf("current user not part of match")
}

func (s *service) HandleDeselect(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	matchID, _ := strconv.Atoi(chi.URLParam(r, "matchID"))
	match, err := s.Queries.ChessMatchByID(ctx, int64(matchID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	state, err := gameStateFromMatch(match)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	state.Render(w)
}

func (s *service) HandleMove(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	currentUser := auth.FromContext(ctx)
	matchID, _ := strconv.Atoi(chi.URLParam(r, "matchID"))
	match, err := s.Queries.ChessMatchByID(ctx, int64(matchID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	state, err := gameStateFromMatch(match)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	currentUserColor, err := userMatchColor(currentUser, match)
	if err != nil {
		http.Error(w, "userMatchColor: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if state.Game.Position().Turn() != currentUserColor {
		http.Error(w, "not your turn in this match", http.StatusBadRequest)
		return
	}

	s1, _ := strconv.Atoi(r.FormValue("s1"))
	s2, _ := strconv.Atoi(r.FormValue("s2"))

	fmt.Println("BEFORE:\n", state.Game.Position().Board().Draw())

	for _, move := range state.Game.ValidMoves() {
		if int(move.S1()) == s1 && int(move.S2()) == s2 {
			err := state.Game.Move(move)
			if err != nil {
				http.Error(w, "move error: "+err.Error(), http.StatusInternalServerError)
				return
			}
			break
		}
	}

	fmt.Println("AFTER:\n", state.Game.Position().Board().Draw())

	fmt.Println("STRING: ", state.Game.String())

	match, err = s.Queries.UpdateChessMatchPGN(ctx, api.UpdateChessMatchPGNParams{
		ID:  match.ID,
		Pgn: strings.TrimSpace(state.Game.String()),
	})
	if err != nil {
		http.Error(w, "UpdateChessMatchPGN: "+err.Error(), http.StatusInternalServerError)
		return
	}

	state, err = gameStateFromMatch(match)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	state.Render(w)
}
