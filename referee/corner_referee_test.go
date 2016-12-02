package referee

import "github.com/stretchr/testify/assert"
import . "github.com/go-number-1-fan/tic-tac-toe/board"
import "testing"

func TestCornerReferee_getCornerWins_ReturnsTheCornersOfAEmptyBoard(t *testing.T) {
	board := Board{
		E, E, E,
		E, E, E,
		E, E, E,
	}

	expectedWin := WinableCombination{E, E, E, E}

	assert.Equal(t, expectedWin, getCornerWins(board))
}

func TestCornerReferee_getCornerWins_ReturnsTheCornersOfABoardWithOne(t *testing.T) {
	board := Board{
		X, E, E,
		E, E, E,
		E, E, E,
	}

	expectedWin := WinableCombination{X, E, E, E}

	assert.Equal(t, expectedWin, getCornerWins(board))
}

func TestCornerReferee_getCornerWins_ReturnsTheCornersOfABoardWithTwo(t *testing.T) {
	board := Board{
		X, E, O,
		E, E, E,
		E, E, E,
	}

	expectedWin := WinableCombination{X, O, E, E}

	assert.Equal(t, expectedWin, getCornerWins(board))
}

func TestCornerReferee_getCornerWins_ReturnsTheCornersOfABoardWithThree(t *testing.T) {
	board := Board{
		X, E, O,
		E, E, E,
		X, E, E,
	}

	expectedWin := WinableCombination{X, O, X, E}

	assert.Equal(t, expectedWin, getCornerWins(board))
}

func TestCornerReferee_getCornerWins_ReturnsTheCornersOfABoardWithFour(t *testing.T) {
	board := Board{
		X, E, O,
		E, E, E,
		X, E, O,
	}

	expectedWin := WinableCombination{X, O, X, O}

	assert.Equal(t, expectedWin, getCornerWins(board))
}

func TestCornerReferee_getCornerWins_ReturnsTheCornersOfABoardWithFour_LargeBoard(t *testing.T) {
	board := Board{
		X, E, E, O,
		E, E, E, E,
		E, E, E, E,
		X, E, E, O,
	}

	expectedWin := WinableCombination{X, O, X, O}

	assert.Equal(t, expectedWin, getCornerWins(board))
}

func TestCornerReferee_CheckMarkerForAllButOne_ReturnsTrueIfAtLeast3MarkersAreTheSameForX(t *testing.T) {
	expectedWin := WinableCombination{X, O, X, X}
	assert.True(t, checkMarkerForAllButOne(X, expectedWin))
}

func TestCornerReferee_CheckMarkerForAllButOne_ReturnsTrueIfAtLeast3MarkersAreTheSameForO(t *testing.T) {
	expectedWin := WinableCombination{X, O, O, O}

	assert.True(t, checkMarkerForAllButOne(O, expectedWin))
}

func TestCornerReferee_CheckMarkerForAllButOne_ReturnsFalseIfCheckingTheWrongMarker(t *testing.T) {
	expectedWin := WinableCombination{X, O, O, O}

	assert.False(t, checkMarkerForAllButOne(X, expectedWin))
}

func TestCornerReferee_CheckMarkerForAllButOne_ReturnsFalseForBothIfTheresNoWin(t *testing.T) {
	expectedWin := WinableCombination{E, E, E, E}

	assert.False(t, checkMarkerForAllButOne(X, expectedWin))
	assert.False(t, checkMarkerForAllButOne(O, expectedWin))
}

func TestCornerReferee_checkAllCornersOccupied_ReturnsFalseWhenTheresAtLeastOneOpenCorner(t *testing.T) {
	expectedWin := WinableCombination{E, X, O, X}

	assert.False(t, checkAllCornersOccupied(expectedWin))
}

func TestCornerReferee_checkAllCornersOccupied_ReturnsTrueWhenTheresNoOpenCorners(t *testing.T) {
	expectedWin := WinableCombination{O, X, O, X}

	assert.True(t, checkAllCornersOccupied(expectedWin))
}

func TestCornerReferee_ReturnsWinP1WhenXWins(t *testing.T) {
	ref := CornerReferee{}
	board := Board{
		X, E, X,
		E, E, E,
		X, E, O,
	}

	assert.Equal(t, WinP1, ref.GetGameStatus(board))
}

func TestCornerReferee_ReturnsWinP2WhenOWins(t *testing.T) {
	ref := CornerReferee{}
	board := Board{
		X, E, O,
		E, E, E,
		O, E, O,
	}

	assert.Equal(t, WinP2, ref.GetGameStatus(board))
}

func TestCornerReferee_ReturnsTieWhenNoWinsAndCornersFull(t *testing.T) {
	ref := CornerReferee{}
	board := Board{
		X, E, X,
		E, E, E,
		O, E, O,
	}

	assert.Equal(t, Tie, ref.GetGameStatus(board))
}

func TestCornerReferee_ReturnsContinueWhenTheresNoWinAndStillCornersOpen(t *testing.T) {
	ref := CornerReferee{}
	board := Board{
		E, E, X,
		E, E, E,
		O, E, O,
	}

	assert.Equal(t, Continue, ref.GetGameStatus(board))
}
