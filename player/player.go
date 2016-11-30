package player

import . "github.com/go-number-1-fan/tic-tac-toe/board"
import . "github.com/go-number-1-fan/tic-tac-toe/ui"

type Player interface {
	GetMove(board Board, ui UI) int
	GetMarker() Marker
}
