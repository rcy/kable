package chess

import (
	"errors"
	"fmt"
	"oj/api"
	"strings"

	"github.com/notnil/chess"
)

type GameState struct {
	Board      [8][8]*UISquare
	Game       *chess.Game
	ValidMoves []*chess.Move
	MatchID    int64
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
