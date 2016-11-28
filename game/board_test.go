package game

import "github.com/stretchr/testify/assert"
import . "github.com/go-number-1-fan/tic-tac-toe/marker"
import "testing"

func TestEmptyBoard_HasALenghtOfNine(t *testing.T) {
	board := EmptyBoard()
	assert.Equal(t, 9, len(board))
}

func TestBoard_CanMakeAMoveAtAnIndex(t *testing.T) {
	board := EmptyBoard()
	board.MakeMove(0, X)

	assert.Equal(t, X, board[0])
}

func TestBoard_CanTellYouIfItsEmptyAtIndex(t *testing.T) {
	board := EmptyBoard()
	board = board.MakeMove(0, X)
	assert.True(t, board.emptyAt(4))
}

func TestBoard_CanTellYouIfItsNotEmptyAtIndex(t *testing.T) {
	board := EmptyBoard()
	board = board.MakeMove(0, X)
	assert.False(t, board.emptyAt(0))
}

func TestBoard_CanGiveAListOfAllEmptySpots(t *testing.T) {
	board := EmptyBoard()

	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8}, board.EmptySpots())
}

func TestBoard_CanGiveAListOfAllEmptySpotsWhenSomeSpotsAreOccupied(t *testing.T) {

	board := EmptyBoard()
	board = board.MakeMove(0, X).MakeMove(1, O)

	assert.Equal(t, []int{2, 3, 4, 5, 6, 7, 8}, board.EmptySpots())
}

func TestBoard_CanCountTheNumberOfOccupiedSpacesForAnEmptyBoard(t *testing.T) {
	board := EmptyBoard()
	numberOfOccupiedSpaces := board.CountOccupiedSpots()

	assert.Equal(t, 0, numberOfOccupiedSpaces)
}

func TestBoard_CanCountTheNumberOfOccupiedSpacesForABoardWithSpotsOccupied(t *testing.T) {
	board := EmptyBoard()
	board = board.MakeMove(0, X).MakeMove(1, O)
	numberOfOccupiedSpaces := board.CountOccupiedSpots()

	assert.Equal(t, 2, numberOfOccupiedSpaces)
}

func TestBoard_CanCountTheNumberOfOccupiedSpacesForABoardWithAFullBoard(t *testing.T) {
	board := Board{
		X, X, X,
		X, X, X,
		X, X, X,
	}
	numberOfOccupiedSpaces := board.CountOccupiedSpots()

	assert.Equal(t, 9, numberOfOccupiedSpaces)
}
