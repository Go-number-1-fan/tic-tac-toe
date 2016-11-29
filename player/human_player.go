package player

import . "github.com/go-number-1-fan/tic-tac-toe/board"
import . "github.com/go-number-1-fan/tic-tac-toe/ui"

type HumanPlayer struct {
	Marker Marker
	UI     UI
}

func (player HumanPlayer) GetMove(board Board) int {
	return player.UI.GetValidMove(board)
}

func (player HumanPlayer) GetMarker() Marker {
	return player.Marker
}
