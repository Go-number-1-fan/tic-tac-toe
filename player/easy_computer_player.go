package player

import . "github.com/go-number-1-fan/tic-tac-toe/board"
import . "github.com/go-number-1-fan/tic-tac-toe/ui"
import "math/rand"

type EasyComputerPlayer struct {
	Marker string
}

func (computer EasyComputerPlayer) GetMove(board Board, ui UI) int {
	emptySpots := board.EmptySpots()
	emptySpotSelection := rand.Intn(len(emptySpots))
	ui.DisplayComputerThinkingMessage()
	return emptySpots[emptySpotSelection]
}

func (player EasyComputerPlayer) GetMarker() string {
	return player.Marker
}
