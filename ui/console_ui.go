package ui

import . "github.com/go-number-1-fan/tic-tac-toe/board"

type ConsoleUI struct {
	Input  Input
	Output Output
}

func getCharsPerRow(rowLength int) int {
	charsPerRow := (rowLength * 3) + (rowLength + 1)
	return charsPerRow
}

func getDisplayBoard(board Board) string {
	rowLength := board.RowLength()
	charsPerRow := getCharsPerRow(rowLength)
	boardString := ""
	stringBoard := board.StringBoard()
	for i := 0; i < len(board); {
		for j := 0; j < charsPerRow; j++ {
			boardString = boardString + "-"
		}
		boardString = boardString + "\n"
		boardString = boardString + "| " + stringBoard[i] + " | " + stringBoard[i+1] + " | " + stringBoard[i+2] + " |\n"
		i = i + rowLength
	}
	return boardString
}

func (ui ConsoleUI) DisplayBoard(board Board) {
	ui.Output.Write(getDisplayBoard(board))
}

func (ui ConsoleUI) GetValidMove(board Board) int {
	ui.Output.Write("Please select an open spot by the index:\n")
	selected := ui.Input.ReadInt()
	for !board.IsMoveOpen(selected) {
		ui.Output.Write("Your selection is Invalid!!\n")
		ui.Output.Write("Please select an open spot of the board by the index:\n")
		selected = ui.Input.ReadInt()
	}
	return selected
}
