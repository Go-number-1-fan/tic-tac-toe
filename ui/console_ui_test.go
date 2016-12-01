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
		"X", "O", "2",
		"O", "O", "5",
		"6", "7", "8",
	}

	assert.Equal(t, " | X | O | 2 | ", getBoardRowString(stringBoard, 0, 3))
	assert.Equal(t, " | O | O | 5 | ", getBoardRowString(stringBoard, 3, 3))
	assert.Equal(t, " | 6 | 7 | 8 | ", getBoardRowString(stringBoard, 6, 3))
}

func TestConsoleUI_CanGetTheDisplayBoardFromAEmptyBoard(t *testing.T) {
	expectedBoard :=
		"\n -------------\n" +
			" | 0 | 1 | 2 | \n" +
			" -------------\n" +
			" | 3 | 4 | 5 | \n" +
			" -------------\n" +
			" | 6 | 7 | 8 | \n"
	assert.Equal(t, expectedBoard, getDisplayBoard(EmptyBoard()))
}

func TestConsoleUI_CanGetTheDisplayBoardFromANonEmptyBoard(t *testing.T) {
	board := EmptyBoard().MakeMove(0, X).MakeMove(8, O).MakeMove(4, X)
	expectedBoard :=
		"\n -------------\n" +
			" | X | 1 | 2 | \n" +
			" -------------\n" +
			" | 3 | X | 5 | \n" +
			" -------------\n" +
			" | 6 | 7 | O | \n"
	assert.Equal(t, expectedBoard, getDisplayBoard(board))
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
			" | X | X | O | \n"
	assert.Equal(t, expectedBoard, getDisplayBoard(board))
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

func ExampleGetPlayerTypeSelection() {
	ui := ConsoleUI{MockInput{2, ""}, ConsoleOutput{}, ConsoleInputValidator{}}
	ui.GetPlayerTypeSelection("Player1")
	// Output:
	//	Player1
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
