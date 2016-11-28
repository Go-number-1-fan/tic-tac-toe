package player

import . "github.com/go-number-1-fan/tic-tac-toe/marker"

type Player interface {
	GetMove() int
	GetMarker() Marker
}
