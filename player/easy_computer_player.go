package player

import (
	. "github.com/go-number-1-fan/tic-tac-toe/board"
	. "github.com/go-number-1-fan/tic-tac-toe/ui"
	"math/rand"
	"time"
)

type EasyComputerPlayer struct {
	Marker string
	Name   string
}

func (computer EasyComputerPlayer) GetMove(board Board, ui UI) int {
	emptySpots := board.EmptySpots()
	rand.Seed(int64(time.Now().UnixNano()))
	emptySpotSelection := rand.Intn(len(emptySpots))
	ui.DisplayComputerThinkingMessage(computer.Name)
	return emptySpots[emptySpotSelection]
}

func (computer EasyComputerPlayer) GetMarker() string {
	return computer.Marker
}

func (computer EasyComputerPlayer) GetName() string {
	return computer.Name
}
