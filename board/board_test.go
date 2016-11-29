package board

import "github.com/stretchr/testify/assert"
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

func TestBoard_CanTellYouIfAMoveIsInTheListOfEmptySpots(t *testing.T) {
	board := EmptyBoard()
	assert.True(t, board.IsMoveOpen(0))
}

func TestBoard_CanTellYouIfAMoveIsNotInTheListOfEmptySpots(t *testing.T) {
	board := EmptyBoard().MakeMove(0, X)
	assert.False(t, board.IsMoveOpen(0))
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

func TestBoard_CanGetTheLengthOfOneOfItsRows(t *testing.T) {
	board := EmptyBoard()
	assert.Equal(t, 3, board.RowLength())
}

func TestBoard_CanGetTheLengthOfOneOfItsRowsForALargeBoard(t *testing.T) {
	board := Board{
		E, E, E, E,
		E, E, E, E,
		E, E, E, E,
		E, E, E, E,
	}
	assert.Equal(t, 4, board.RowLength())
}

func TestBoard_CanReturnItselfAsAStringAndReplaceTheOpenSpotsWithTheIndex(t *testing.T) {
	board := EmptyBoard()
	expectedStringBoard := []string{
		"0", "1", "2",
		"3", "4", "5",
		"6", "7", "8",
	}
	assert.Equal(t, expectedStringBoard, board.StringBoard())
}

func TestBoard_CanReturnItselfAsAStringAndReplaceOccupiedSpotsWithTheMarker(t *testing.T) {
	board := EmptyBoard().MakeMove(1, X).MakeMove(8, O)
	expectedStringBoard := []string{
		"0", "X", "2",
		"3", "4", "5",
		"6", "7", "O",
	}
	assert.Equal(t, expectedStringBoard, board.StringBoard())
}
