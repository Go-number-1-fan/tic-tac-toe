package player

import . "github.com/go-number-1-fan/tic-tac-toe/marker"
import . "github.com/go-number-1-fan/tic-tac-toe/UI"

type HumanPlayer struct {
	Marker Marker
	UI     UI
}

func (player HumanPlayer) GetMove(availableSpots []int) int {
	return player.UI.GetValidMove(availableSpots)
}

func (player HumanPlayer) GetMarker() Marker {
	return player.Marker
}
