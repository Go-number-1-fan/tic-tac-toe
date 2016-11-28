package player

import "github.com/stretchr/testify/assert"
import . "github.com/go-number-1-fan/tic-tac-toe/marker"
import . "github.com/go-number-1-fan/tic-tac-toe/mocks"
import "testing"

func TestPlayer_CanReturnAValidMove(t *testing.T) {
	player := HumanPlayer{X, MockUI{2}}
	emptySpots := []int{2, 3, 4, 5, 6}
	assert.Equal(t, player.GetMove(emptySpots), 2)
}

func TestPlayer_CanReturnAValidMarker(t *testing.T) {
	player := HumanPlayer{X, MockUI{2}}
	assert.Equal(t, player.GetMarker(), X)
}
