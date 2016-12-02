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
	assert.Equal(t, expectedBoard, getDisplayBoard(EmptyBoard(), "X", "O"))
}

func TestConsoleUI_CanGetTheDisplayBoardFromANonEmptyBoard(t *testing.T) {
	board := EmptyBoard().MakeMove(0, X).MakeMove(8, O).MakeMove(4, X)
	expectedBoard :=
		"\n -------------\n" +
			" | X | 2 | 3 | \n" +
			" -------------\n" +
			" | 4 | X | 6 | \n" +
			" -------------\n" +
			" | 7 | 8 | O | \n" +
			" -------------\n\n"
	assert.Equal(t, expectedBoard, getDisplayBoard(board, "X", "O"))
}

func TestConsoleUI_CanGetTheDisplayBoardFromAFullBoard(t *testing.T) {
	board := Board{
		X, X, O,
		O, O, X,
		X, X, O,
	}
	expectedBoard :=
		"\n -------------\n" +
			" | X | X | O | \n" +
			" -------------\n" +
			" | O | O | X | \n" +
			" -------------\n" +
			" | X | X | O | \n" +
			" -------------\n\n"
	assert.Equal(t, expectedBoard, getDisplayBoard(board, "X", "O"))
}

func TestConsoleUI_CanGetTheDisplayBoardFromAFullBoardWithOtherMarkers(t *testing.T) {
	board := Board{
		X, X, O,
		O, O, X,
		X, X, O,
	}
	expectedBoard :=
		"\n -------------\n" +
			" | A | A | B | \n" +
			" -------------\n" +
			" | B | B | A | \n" +
			" -------------\n" +
			" | A | A | B | \n" +
			" -------------\n\n"
	assert.Equal(t, expectedBoard, getDisplayBoard(board, "A", "B"))
}

func TestConsoleUI_CanGetTheHumanSelectMessage(t *testing.T) {
	assert.True(t, strings.Contains(getHumanSelectMessage("X"), "X"))
	assert.True(t, strings.Contains(getHumanSelectMessage("O"), "O"))
	assert.True(t, strings.Contains(getHumanSelectMessage("BLAH"), "BLAH"))
}

func ExampleWelcomeMessage() {
	ui := ConsoleUI{MockInput{0, ""}, ConsoleOutput{}, ConsoleInputValidator{}}
	ui.DisplayWelcomeMessage()
	// Output:
	//Welcome To Tic Tac Toe!
	//Follow the instructions and enter requested information when prompted!
	//
	//Most Importantly... Have fun!
}

func ExampleGetRefereeSelection() {
	ui := ConsoleUI{MockInput{2, ""}, ConsoleOutput{}, ConsoleInputValidator{}}
	ui.GetRefereeSelection()
	// Output:
	//Please Select a Ruleset:
	//	1. Standard Ruleset - 3 in a row, Horizontal/Vertical/Diagonal
	//	2. Corner Ruleset - Place a Marker in 3 of the 4 corners

}

func ExampleGetPlayerTypeSelection() {
	ui := ConsoleUI{MockInput{2, ""}, ConsoleOutput{}, ConsoleInputValidator{}}
	ui.GetPlayerTypeSelection("Player1", "X")
	// Output:
	//	Player1 | X
	//Please Select a Player Type:
	//	1. Human Player
	//	2. Easy Computer Player
	//	3. Hard Computer Player
}

func ExampleGetPlayerNameSelection() {
	ui := ConsoleUI{MockInput{2, "Player1"}, ConsoleOutput{}, ConsoleInputValidator{}}
	ui.GetPlayerNameSelection("1")
	// Output:
	//Please Select a Name for Player1:
}

func ExampleGetPlayerMarkerSelection() {
	ui := ConsoleUI{MockInput{2, "O"}, ConsoleOutput{}, ConsoleInputValidator{}}
	ui.GetPlayerMarkerSelection("Tom", "2")
	// Output:
	//Please Select a Marker for Tom:
}

func ExampleGetPlayAgainSelection() {
	ui := ConsoleUI{MockInput{2, "O"}, ConsoleOutput{}, ConsoleInputValidator{}}
	ui.GetPlayAgainSelection()
	// Output:
	//Would you like to play again? (y | n):
}

func ExampleCanClearTheBoard() {
	ui := ConsoleUI{MockInput{2, "O"}, MockOutput{}, ConsoleInputValidator{}}
	ui.Clear()
	// Output:
	//CLEARED
}
