package chess

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"oj/handlers/layout"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/notnil/chess"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type Square struct {
	SVGPiece string
	Selected bool
	Action   string
	Dot      bool
	Rank     int
	File     int
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
		g.Attr("hx-get", "/fun/chess/"+s.Action),
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

type Board [8][8]Square

type GameBoard struct {
	Board      Board
	ValidMoves []*chess.Move
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

var game *chess.Game = chess.NewGame()

func Page(w http.ResponseWriter, r *http.Request) {
	l := layout.FromContext(r.Context())

	for i := 0; i < 1; i += 1 {
		moves := game.ValidMoves()
		if len(moves) > 0 {
			move := moves[rand.Intn(len(moves))]
			game.Move(move)
		}
	}

	log.Println(game.Position().Board().Draw())

	gameBoard := gameBoard(game, nil)

	layout.Layout(l, "chess club", chessPageEl(gameBoard)).Render(w)
}

func chessPageEl(gameboard GameBoard) g.Node {
	return h.Div(
		h.H1(g.Text("chess club")),
		h.Div(h.ID("board-container"), h.Style("height: 80vh; width: 80vh;"),
			gameboard,
		))
}

func gameBoard(game *chess.Game, position *Position) GameBoard {
	pos := game.Position()
	gb := gameBoardFromSquareMap(pos.Board().SquareMap(), position)

	moves := game.ValidMoves()
	if position == nil {
		gb.ValidMoves = moves
	} else {
		selectedSquare := chess.Square((7-position.rank)*8 + position.file)
		for _, move := range moves {
			if move.S1() == selectedSquare {
				gb.ValidMoves = append(gb.ValidMoves, move)
			}
		}

		for _, move := range gb.ValidMoves {
			target := move.S2()
			square := &gb.Board[7-target/8][target%8]
			square.Dot = true
			square.Action = "move/" + move.String()
		}
	}

	return gb
}

type Position struct {
	rank int
	file int
}

func gameBoardFromSquareMap(squareMap map[chess.Square]chess.Piece, selected *Position) GameBoard {
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

	var gb GameBoard

	for i := 0; i < 64; i += 1 {
		piece := squareMap[chess.Square(i)]
		rank := 7 - i/8
		file := i % 8
		square := &gb.Board[rank][file]
		square.SVGPiece = svgPiece[piece]

		square.Rank = rank
		square.File = file

		if selected != nil && selected.rank == rank && selected.file == file {
			square.Action = "unselect"
			square.Selected = true
		} else {
			square.Action = fmt.Sprintf("select/%d/%d", rank, file)
		}
	}

	return gb
}

func Select(w http.ResponseWriter, r *http.Request) {
	rank, _ := strconv.Atoi(chi.URLParam(r, "rank"))
	file, _ := strconv.Atoi(chi.URLParam(r, "file"))
	gameBoard(game, &Position{rank, file}).Render(w)
}

func Unselect(w http.ResponseWriter, r *http.Request) {
	gameBoard(game, nil).Render(w)
}
