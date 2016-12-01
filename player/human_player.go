package player

import . "github.com/go-number-1-fan/tic-tac-toe/board"
import . "github.com/go-number-1-fan/tic-tac-toe/ui"

type HumanPlayer struct {
	Marker string
	Name   string
}

func (player HumanPlayer) GetMove(board Board, ui UI) int {
	return ui.GetValidMove(board, player.Name)
}

func (player HumanPlayer) GetMarker() string {
	return player.Marker
}

func (player HumanPlayer) GetName() string {
	return player.Name
}
