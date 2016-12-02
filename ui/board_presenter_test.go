package ui

import "github.com/stretchr/testify/assert"
import . "github.com/go-number-1-fan/tic-tac-toe/board"
import "strings"
import "testing"

func TestConsoleUI_CanGetTheNumberOfCharactersPerRow(t *testing.T) {
	assert.Equal(t, 13, getCharsPerRow(3))
}

func TestConsoleUI_CanGetTheHorizontalDividerString(t *testing.T) {
	assert.Equal(t, " -------------", getHorizontalDividerString(13))
}

func TestConsoleUI_CanGetBoardRowString(t *testing.T) {
	stringBoard := []string{
		"X", "O", "3",
		"O", "O", "6",
		"7", "8", "9",
	}

	assert.Equal(t, " | X | O | 3 | ", getBoardRowString(stringBoard, 0, 3))
	assert.Equal(t, " | O | O | 6 | ", getBoardRowString(stringBoard, 3, 3))
	assert.Equal(t, " | 7 | 8 | 9 | ", getBoardRowString(stringBoard, 6, 3))
}

func TestConsoleUI_CanGetTheDisplayBoardFromAEmptyBoard(t *testing.T) {
	expectedBoard :=
		"\n -------------\n" +
			" | 1 | 2 | 3 | \n" +
			" -------------\n" +
			" | 4 | 5 | 6 | \n" +
			" -------------\n" +
			" | 7 | 8 | 9 | \n" +
			" -------------\n\n"
	assert.Equal(t, expectedBoard, getDisplayBoard(EmptyBoard().StringBoard()))
}

func TestConsoleUI_CanGetTheDisplayBoardFromANonEmptyBoard(t *testing.T) {
	board := EmptyBoard().MakeMove(0, X).MakeMove(8, O).MakeMove(4, X).StringBoard()
	expectedBoard :=
		"\n -------------\n" +
			" | X | 2 | 3 | \n" +
			" -------------\n" +
			" | 4 | X | 6 | \n" +
			" -------------\n" +
			" | 7 | 8 | O | \n" +
			" -------------\n\n"
	assert.Equal(t, expectedBoard, getDisplayBoard(board))
}

func TestConsoleUI_CanGetTheDisplayBoardFromAFullBoard(t *testing.T) {
	board := Board{
		X, X, O,
		O, O, X,
		X, X, O,
	}.StringBoard()
	expectedBoard :=
		"\n -------------\n" +
			" | X | X | O | \n" +
			" -------------\n" +
			" | O | O | X | \n" +
			" -------------\n" +
			" | X | X | O | \n" +
			" -------------\n\n"
	assert.Equal(t, expectedBoard, getDisplayBoard(board))
}

func TestConsoleUI_CanGetTheHumanSelectMessage(t *testing.T) {
	assert.True(t, strings.Contains(getHumanSelectMessage("X"), "X"))
	assert.True(t, strings.Contains(getHumanSelectMessage("O"), "O"))
	assert.True(t, strings.Contains(getHumanSelectMessage("BLAH"), "BLAH"))
}
