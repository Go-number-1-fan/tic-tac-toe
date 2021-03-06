package player

import (
	. "github.com/go-number-1-fan/tic-tac-toe/board"
	. "github.com/go-number-1-fan/tic-tac-toe/ui"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlayer_CanReturnAValidMove(t *testing.T) {
	player := HumanPlayer{"X", "Player1"}
	assert.Equal(t, player.GetMove(EmptyBoard(), MockUI{2}), 2)
}

func TestPlayer_CanReturnAValidMarker(t *testing.T) {
	player := HumanPlayer{"X", "Player2"}
	assert.Equal(t, player.GetMarker(), "X")
}

func TestPlayer_CanReturnAValidName(t *testing.T) {
	player := HumanPlayer{"X", "Player2"}
	assert.Equal(t, player.GetName(), "Player2")
}
