package mocks

import . "github.com/go-number-1-fan/tic-tac-toe/board"
import . "github.com/go-number-1-fan/tic-tac-toe/ui"

type MockPlayer struct {
	Marker Marker
	Move   int
}

func (player MockPlayer) GetMove(board Board, ui UI) int {
	return player.Move
}

func (player MockPlayer) GetMarker() Marker {
	return player.Marker
}