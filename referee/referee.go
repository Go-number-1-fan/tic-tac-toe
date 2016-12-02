package referee

import . "github.com/go-number-1-fan/tic-tac-toe/board"

type WinableCombination []Marker

type Referee interface {
	GetGameStatus(board Board) GameStatus
}
