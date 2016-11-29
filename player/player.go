package player

import . "github.com/go-number-1-fan/tic-tac-toe/board"

type Player interface {
	GetMove(board Board) int
	GetMarker() Marker
}
