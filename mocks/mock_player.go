package mocks

import . "github.com/go-number-1-fan/tic-tac-toe/marker"

type MockPlayer struct {
	Marker Marker
	Move   int
}

func (player MockPlayer) GetMove() int {
	return player.Move
}

func (player MockPlayer) GetMarker() Marker {
	return player.Marker
}
