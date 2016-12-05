package player

import (
	. "github.com/go-number-1-fan/tic-tac-toe/board"
	. "github.com/go-number-1-fan/tic-tac-toe/ui"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEasyComputerPlayer_CanReturnAValidMove(t *testing.T) {
	player := EasyComputerPlayer{"X", "Computer"}
	board := Board{
		X, O, X,
		X, O, O,
		O, X, E}
	assert.Equal(t, player.GetMove(board, MockUI{-1}), 8)
}

func TestEasyComputerPlayer_CanReturnAnotherValidMove(t *testing.T) {
	player := EasyComputerPlayer{"X", "Computer"}
	board := Board{
		X, O, X,
		X, O, O,
		O, X, E}
	assert.Equal(t, player.GetMove(board, MockUI{-1}), 8)
}

func TestEasyComputerPlayer_CanReturnAValidMarker(t *testing.T) {
	player := EasyComputerPlayer{"X", "Computer"}
	assert.Equal(t, player.GetMarker(), "X")
}

func TestEasyComputerPlayer_CanReturnAValidName(t *testing.T) {
	player := EasyComputerPlayer{"X", "Computer"}
	assert.Equal(t, player.GetName(), "Computer")
}
