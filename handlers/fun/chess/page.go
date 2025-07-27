package chess

import (
	_ "embed"
	"fmt"
	"io"
	"net/http"
	"oj/api"
	"oj/internal/link"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/notnil/chess"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type Board [8][8]Square

type GameBoard struct {
	Board      Board
	Game       *chess.Game
	ValidMoves []*chess.Move
	MatchID    int64
}

type Square struct {
	SVGPiece string
	Selected bool
	Action   string
	Dot      bool
	Rank     int
	File     int
	MatchID  int64
}

func (s Square) Occupied() bool {
	return s.SVGPiece != ""
}

func (s Square) Render(w io.Writer) error {
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
		g.Attr("hx-post", link.ChessMatch(s.MatchID, s.Action)),
		g.Attr("hx-target", "closest .board"),

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

func (gb GameBoard) Render(w io.Writer) error {
	rows := make([][]g.Node, 0, 8)
	for rank := 0; rank < 8; rank += 1 {
		row := make([]g.Node, 0, 8)
		for file := 0; file < 8; file += 1 {
			row = append(row, gb.Board[rank][file])
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

func chessPageEl(gameboard *GameBoard) g.Node {
	return h.Div(
		h.H1(g.Text("chess club")),
		h.Div(h.ID("board-container"), h.Style("height: 80vh; width: 80vh;"),
			gameboard,
		))
}

func gameBoardFromMatch(match api.ChessMatch, selectedSquare *Square) (*GameBoard, error) {
	reader := strings.NewReader(match.Pgn)
	fn, err := chess.PGN(reader)
	if err != nil {
		return nil, err
	}
	game := chess.NewGame(fn)

	pos := game.Position()
	gb := gameBoardFromSquareMap(match.ID, game, pos.Board().SquareMap(), selectedSquare)

	moves := game.ValidMoves()
	if selectedSquare == nil {
		gb.ValidMoves = moves
	} else {
		selectedSquare := chess.Square((7-selectedSquare.Rank)*8 + selectedSquare.File)
		for _, move := range moves {
			if move.S1() == selectedSquare {
				gb.ValidMoves = append(gb.ValidMoves, move)
			}
		}

		for _, move := range gb.ValidMoves {
			target := move.S2()
			square := &gb.Board[7-target/8][target%8]
			square.Dot = true
			square.Action = fmt.Sprintf("move?s1=%d&s2=%d", move.S1(), move.S2())
		}
	}

	return &gb, nil
}
func gameBoardFromSquareMap(matchID int64, game *chess.Game, squareMap map[chess.Square]chess.Piece, selected *Square) GameBoard {
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

	gb := GameBoard{Game: game}

	for i := 0; i < 64; i += 1 {
		piece := squareMap[chess.Square(i)]
		rank := 7 - i/8
		file := i % 8
		square := &gb.Board[rank][file]
		square.SVGPiece = svgPiece[piece]

		square.Rank = rank
		square.File = file
		square.MatchID = matchID

		if selected != nil && selected.Rank == rank && selected.File == file {
			square.Action = "unselect"
			square.Selected = true
		} else {
			square.Action = fmt.Sprintf("select/%d/%d", rank, file)
		}
	}

	return gb
}

func (s *service) HandleSelect(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	matchID, _ := strconv.Atoi(chi.URLParam(r, "matchID"))
	match, err := s.Queries.ChessMatchByID(ctx, int64(matchID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rank, _ := strconv.Atoi(chi.URLParam(r, "rank"))
	file, _ := strconv.Atoi(chi.URLParam(r, "file"))

	board, err := gameBoardFromMatch(match, &Square{Rank: rank, File: file})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	board.Render(w)
}

func (s *service) HandleDeselect(w http.ResponseWriter, r *http.Request) {
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

	board.Render(w)
}

func (s *service) HandleMove(w http.ResponseWriter, r *http.Request) {
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

	s1, _ := strconv.Atoi(r.FormValue("s1"))
	s2, _ := strconv.Atoi(r.FormValue("s2"))

	fmt.Println("BEFORE:\n", board.Game.Position().Board().Draw())

	for _, move := range board.Game.ValidMoves() {
		if int(move.S1()) == s1 && int(move.S2()) == s2 {
			err := board.Game.Move(move)
			if err != nil {
				http.Error(w, "move error: "+err.Error(), http.StatusInternalServerError)
				return
			}
			break
		}
	}

	fmt.Println("AFTER:\n", board.Game.Position().Board().Draw())

	fmt.Println("STRING: ", board.Game.String())

	match, err = s.Queries.UpdateChessMatchPGN(ctx, api.UpdateChessMatchPGNParams{
		ID:  match.ID,
		Pgn: strings.TrimSpace(board.Game.String()),
	})
	if err != nil {
		http.Error(w, "UpdateChessMatchPGN: "+err.Error(), http.StatusInternalServerError)
		return
	}

	board, err = gameBoardFromMatch(match, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	board.Render(w)
}
