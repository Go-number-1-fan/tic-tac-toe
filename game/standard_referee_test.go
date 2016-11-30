package game

import "github.com/stretchr/testify/assert"
import . "github.com/go-number-1-fan/tic-tac-toe/board"
import "testing"

func TestReferee_ReturnsTheHorizontalWinsOfAEmptyBoard(t *testing.T) {
	board := Board{
		E, E, E,
		E, E, E,
		E, E, E,
	}

	expectedWins := []WinableCombination{
		WinableCombination{E, E, E},
		WinableCombination{E, E, E},
		WinableCombination{E, E, E}}

	assert.Equal(t, expectedWins, getHorizontalWins(board))
}

func TestReferee_ReturnsTheHorizontalWinsOfABoard(t *testing.T) {
	board := Board{
		X, E, X,
		E, X, E,
		E, O, O,
	}

	expectedWins := []WinableCombination{
		WinableCombination{X, E, X},
		WinableCombination{E, X, E},
		WinableCombination{E, O, O}}

	assert.Equal(t, expectedWins, getHorizontalWins(board))
}

func TestReferee_ReturnsTrueIfAWinningCombinationHasAllX(t *testing.T) {
	winningCombination := WinableCombination{X, X, X}
	assert.True(t, checkWinableCombinationForWin(X, winningCombination))
}

func TestReferee_ReturnsTrueIfAWinningCombinationHasAllO(t *testing.T) {
	winningCombination := WinableCombination{O, O, O}
	assert.True(t, checkWinableCombinationForWin(O, winningCombination))
}

func TestReferee_ReturnsFalseIfAWinningCombinationHasAllOButChecksForX(t *testing.T) {
	winningCombination := WinableCombination{O, O, O}
	assert.False(t, checkWinableCombinationForWin(X, winningCombination))
}

func TestReferee_ReturnsFalseIfAWinningCombinationHasAllXButChecksForO(t *testing.T) {
	winningCombination := WinableCombination{X, X, X}
	assert.False(t, checkWinableCombinationForWin(O, winningCombination))
}

func TestReferee_ReturnsFalseIfAWinningCombinationHasAllE(t *testing.T) {
	winningCombination := WinableCombination{E, E, E}

	assert.False(t, checkWinableCombinationForWin(O, winningCombination))
	assert.False(t, checkWinableCombinationForWin(X, winningCombination))
}

func TestReferee_ReturnsWinP1_IfXWinsHorizontally(t *testing.T) {
	ref := StandardReferee{}
	board := Board{
		X, X, X,
		E, E, E,
		E, E, E,
	}
	assert.Equal(t, WinP1, ref.GetGameStatus(board))
}
