package player

import (
	. "github.com/go-number-1-fan/tic-tac-toe/board"
	. "github.com/go-number-1-fan/tic-tac-toe/ui"
)

type MockPlayer struct {
	Marker string
	Name   string
	Move   int
}

func (player MockPlayer) GetMove(board Board, ui UI) int {
	return player.Move
}

func (player MockPlayer) GetMarker() string {
	return player.Marker
}

func (player MockPlayer) GetName() string {
	return player.Name
}
