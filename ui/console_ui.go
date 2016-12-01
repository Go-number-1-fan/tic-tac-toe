package ui

import . "github.com/go-number-1-fan/tic-tac-toe/board"
import "time"

type ConsoleUI struct {
	Input     Input
	Output    Output
	Validator ConsoleInputValidator
}

func CreateConsoleUI(input Input, output Output) ConsoleUI {
	return ConsoleUI{input, output, ConsoleInputValidator{}}
}

func (ui ConsoleUI) DisplayWelcomeMessage() {
	ui.Output.Write(WelcomeMessage)
}

func (ui ConsoleUI) DisplayTieMessage() {
	ui.Output.Write(TieMessage)
}

func (ui ConsoleUI) DisplayComputerThinkingMessage(computerName string) {
	ui.Output.Write(computerName + ComputerThinkingMessage)
	time.Sleep(2 * time.Second)
}

func (ui ConsoleUI) DisplayWinMessage(winner string) {
	ui.Output.Write(winner + WinMessage)
}

func getCharsPerRow(rowLength int) int {
	charsPerRow := (rowLength * 3) + (rowLength + 1)
	return charsPerRow
}

func getHorizontalDividerString(charsPerRow int) string {
	boardString := " "
	for j := 0; j < charsPerRow; j++ {
		boardString = boardString + HorizontalDividerString
	}
	return boardString
}

func getBoardRowString(stringBoard []string, boardRowStartIndex int, rowLength int) string {
	boardRow := ""
	for boardColumn := 0; boardColumn < rowLength; boardColumn++ {
		boardRow = boardRow + VerticalDividerString + stringBoard[boardRowStartIndex+boardColumn]
	}
	boardRow = boardRow + VerticalDividerString
	return boardRow
}

func getDisplayBoard(board Board) string {
	rowLength := board.RowLength()
	charsPerRow := getCharsPerRow(rowLength)
	boardString := NewLineString
	stringBoard := board.StringBoard()
	for rowStartIndex := 0; rowStartIndex < len(board); {
		boardString = boardString + getHorizontalDividerString(charsPerRow) + NewLineString
		boardString = boardString + getBoardRowString(stringBoard, rowStartIndex, rowLength) + NewLineString
		rowStartIndex = rowStartIndex + rowLength
	}
	return boardString
}

func (ui ConsoleUI) DisplayBoard(board Board) {
	ui.Output.Write(getDisplayBoard(board))
}

func getHumanSelectMessage(marker string) string {
	return marker + SelectASpotMessage
}

func (ui ConsoleUI) GetValidMove(board Board, marker string) int {
	selectMessage := getHumanSelectMessage(marker)
	ui.Output.Write(selectMessage)
	selected := ui.Input.ReadInt()
	for !ui.Validator.IsValid(selected, board.EmptySpots()) {
		ui.Output.Write(NotValidMessage)
		ui.Output.Write(selectMessage)
		selected = ui.Input.ReadInt()
	}
	return selected
}

func (ui ConsoleUI) GetPlayerTypeSelection(playerName string) PlayerTypeSelection {
	ui.Output.Write(NewLineString + NewLineString + playerName + NewLineString)
	ui.Output.Write(PlayerTypeSelectMessage)
	selected := ui.Input.ReadInt()
	for !ui.Validator.IsValid(selected, []int{1, 2, 3}) {
		ui.Output.Write(PlayerTypeSelectMessage)
		ui.Output.Write(NotValidMessage)
		selected = ui.Input.ReadInt()
	}
	return PlayerTypeSelection(selected)
}

func (ui ConsoleUI) GetPlayerNameSelection(playerNumber string) string {
	nameSelectionMessage := NewLineString + NameSelectMessage + playerNumber + ":" + NewLineString
	ui.Output.Write(nameSelectionMessage)
	selection := ui.Input.ReadStringOfLengthWithDefault(15, "Player"+playerNumber)
	return selection
}
