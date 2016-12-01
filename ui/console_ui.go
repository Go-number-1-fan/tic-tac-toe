package ui

import . "github.com/go-number-1-fan/tic-tac-toe/board"

type ConsoleUI struct {
	Input  Input
	Output Output
}

func (ui ConsoleUI) DisplayWelcomeMessage() {
	ui.Output.Write(WelcomeMessage)
}

func (ui ConsoleUI) DisplayTieMessage() {
	ui.Output.Write(TieMessage)
}

func (ui ConsoleUI) DisplayComputerThinkingMessage() {
	ui.Output.Write(ComputerThinkingMessage)
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
	for !board.IsMoveOpen(selected) {
		ui.Output.Write(MoveNotValidMessage)
		ui.Output.Write(selectMessage)
		selected = ui.Input.ReadInt()
	}
	return selected
}
