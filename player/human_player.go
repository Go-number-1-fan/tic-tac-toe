package player

import . "github.com/go-number-1-fan/tic-tac-toe/board"
import . "github.com/go-number-1-fan/tic-tac-toe/ui"

type HumanPlayer struct {
	Marker Marker
}

func (player HumanPlayer) GetMove(board Board, ui UI) int {
	return ui.GetValidMove(board)
}

func (player HumanPlayer) GetMarker() Marker {
	return player.Marker
}
