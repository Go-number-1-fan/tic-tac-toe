package player

import . "github.com/go-number-1-fan/tic-tac-toe/board"
import . "github.com/go-number-1-fan/tic-tac-toe/ui"
import "math/rand"

type ComputerPlayer struct {
	Marker string
}

func (player ComputerPlayer) GetMove(board Board, ui UI) int {
	emptySpots := board.EmptySpots()
	emptySpotSelection := rand.Intn(len(emptySpots))
	return emptySpots[emptySpotSelection]
}

func (player ComputerPlayer) GetMarker() string {
	return player.Marker
}
