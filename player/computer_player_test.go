package player

import "github.com/stretchr/testify/assert"
import . "github.com/go-number-1-fan/tic-tac-toe/board"
import . "github.com/go-number-1-fan/tic-tac-toe/ui"
import "testing"

func TestComputerPlayer_CanReturnAValidMove(t *testing.T) {
	player := ComputerPlayer{"X"}
	board := Board{
		X, O, X,
		X, O, O,
		O, X, E}
	assert.Equal(t, player.GetMove(board, MockUI{-1}), 8)
}

func TestComputerPlayer_CanReturnAnotherValidMove(t *testing.T) {
	player := ComputerPlayer{"X"}
	board := Board{
		X, O, X,
		X, O, O,
		O, X, E}
	assert.Equal(t, player.GetMove(board, MockUI{-1}), 8)
}

func TestComputerPlayer_CanReturnAValidMarker(t *testing.T) {
	player := ComputerPlayer{"X"}
	assert.Equal(t, player.GetMarker(), "X")
}
