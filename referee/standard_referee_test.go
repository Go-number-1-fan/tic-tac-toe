package referee

import (
	. "github.com/go-number-1-fan/tic-tac-toe/board"
	"github.com/stretchr/testify/assert"
	"testing"
)

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

func TestReferee_ReturnsWinP2_IfOWinsHorizontally(t *testing.T) {
	ref := StandardReferee{}
	board := Board{
		O, O, O,
		E, E, E,
		E, E, E,
	}
	assert.Equal(t, WinP2, ref.GetGameStatus(board))
}

func TestReferee_ReturnsWinP2_IfOWinsHorizontally_InADifferentRow(t *testing.T) {
	ref := StandardReferee{}
	board := Board{
		E, E, E,
		E, E, E,
		O, O, O,
	}
	assert.Equal(t, WinP2, ref.GetGameStatus(board))
}

func TestReferee_ReturnsWinP1_IfXWinsHorizontally_InADifferentRow(t *testing.T) {
	ref := StandardReferee{}
	board := Board{
		E, E, E,
		E, E, E,
		X, X, X,
	}
	assert.Equal(t, WinP1, ref.GetGameStatus(board))
}

func TestReferee_ReturnsTheVerticalWinsOfAnEmptyBoard(t *testing.T) {
	board := Board{
		E, E, E,
		E, E, E,
		E, E, E,
	}

	expectedWins := []WinableCombination{
		WinableCombination{E, E, E},
		WinableCombination{E, E, E},
		WinableCombination{E, E, E}}

	assert.Equal(t, expectedWins, getVerticalWins(board))
}

func TestReferee_ReturnsTheVerticalWinsOfABoard(t *testing.T) {
	board := Board{
		X, E, X,
		E, X, E,
		E, O, O,
	}

	expectedWins := []WinableCombination{
		WinableCombination{X, E, E},
		WinableCombination{E, X, O},
		WinableCombination{X, E, O}}

	assert.Equal(t, expectedWins, getVerticalWins(board))
}

func TestReferee_ReturnsWinP1_IfXWinsVertically(t *testing.T) {
	ref := StandardReferee{}
	board := Board{
		X, E, E,
		X, E, E,
		X, E, E,
	}
	assert.Equal(t, WinP1, ref.GetGameStatus(board))
}

func TestReferee_ReturnsWinP2_IfOWinsVertically(t *testing.T) {
	ref := StandardReferee{}
	board := Board{
		O, E, E,
		O, E, E,
		O, E, E,
	}
	assert.Equal(t, WinP2, ref.GetGameStatus(board))
}

func TestReferee_ReturnsWinP2_IfOWinsVertically_InADifferentRow(t *testing.T) {
	ref := StandardReferee{}
	board := Board{
		E, E, O,
		E, E, O,
		E, E, O,
	}
	assert.Equal(t, WinP2, ref.GetGameStatus(board))
}

func TestReferee_ReturnsWinP1_IfXWinsVertically_InADifferentRow(t *testing.T) {
	ref := StandardReferee{}
	board := Board{
		E, E, X,
		E, E, X,
		E, E, X,
	}
	assert.Equal(t, WinP1, ref.GetGameStatus(board))
}

func TestReferee_GetAllTheDiagonalWinsForAnEmptyBoard(t *testing.T) {
	board := EmptyBoard()
	expectedWins := []WinableCombination{
		WinableCombination{E, E, E},
		WinableCombination{E, E, E},
	}

	assert.Equal(t, expectedWins, getDiagonalWins(board))
}

func TestReferee_GetAllTheDiagonalWinsForABoard(t *testing.T) {
	board := Board{
		E, E, X,
		E, E, X,
		E, E, X,
	}
	expectedWins := []WinableCombination{
		WinableCombination{E, E, X},
		WinableCombination{X, E, E},
	}

	assert.Equal(t, expectedWins, getDiagonalWins(board))
}

func TestReferee_GetAllTheDiagonalWinsForAFullBoard(t *testing.T) {
	board := Board{
		O, E, X,
		E, X, X,
		E, E, X,
	}
	expectedWins := []WinableCombination{
		WinableCombination{O, X, X},
		WinableCombination{X, X, E},
	}

	assert.Equal(t, expectedWins, getDiagonalWins(board))
}

func TestReferee_ReturnsP1WinsIfXWinsDiagonally(t *testing.T) {
	ref := StandardReferee{}
	board := Board{
		X, E, O,
		E, X, X,
		E, E, X,
	}

	assert.Equal(t, WinP1, ref.GetGameStatus(board))
}

func TestReferee_ReturnsP1WinsIfOWinsDiagonally(t *testing.T) {
	ref := StandardReferee{}
	board := Board{
		X, E, O,
		E, O, X,
		O, E, X,
	}

	assert.Equal(t, WinP2, ref.GetGameStatus(board))
}

func TestReferee_ReturnsContinueIfNoWin(t *testing.T) {
	ref := StandardReferee{}
	board := Board{
		X, E, O,
		E, E, X,
		O, E, X,
	}

	assert.Equal(t, Continue, ref.GetGameStatus(board))
}

func TestReferee_ReturnsTieIfBoardIsFull(t *testing.T) {
	ref := StandardReferee{}
	board := Board{
		X, O, O,
		O, X, X,
		O, X, O,
	}

	assert.Equal(t, Tie, ref.GetGameStatus(board))
}

func TestReferee_ReturnsP1WinIfBoardIsFullAndXWins(t *testing.T) {
	ref := StandardReferee{}
	board := Board{
		X, O, O,
		O, X, X,
		O, X, X,
	}

	assert.Equal(t, WinP1, ref.GetGameStatus(board))
}

func TestReferee_GetsAllWinningCombinations(t *testing.T) {
	board := Board{
		X, O, E,
		E, X, O,
		O, E, X,
	}

	expectedWins := []WinableCombination{
		WinableCombination{X, O, E},
		WinableCombination{E, X, O},
		WinableCombination{O, E, X},
		WinableCombination{X, E, O},
		WinableCombination{O, X, E},
		WinableCombination{E, O, X},
		WinableCombination{X, X, X},
		WinableCombination{E, X, O},
	}

	assert.Equal(t, expectedWins, getAllWinableCombinations(board))
}
