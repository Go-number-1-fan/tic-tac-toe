package player

import "github.com/stretchr/testify/assert"
import . "github.com/go-number-1-fan/tic-tac-toe/board"
import . "github.com/go-number-1-fan/tic-tac-toe/mocks"
import "testing"

func TestPlayer_CanReturnAValidMove(t *testing.T) {
	player := HumanPlayer{"X"}
	assert.Equal(t, player.GetMove(EmptyBoard(), MockUI{2}), 2)
}

func TestPlayer_CanReturnAValidMarker(t *testing.T) {
	player := HumanPlayer{"X"}
	assert.Equal(t, player.GetMarker(), "X")
}
