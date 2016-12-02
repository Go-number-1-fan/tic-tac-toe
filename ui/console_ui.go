package ui

import . "github.com/go-number-1-fan/tic-tac-toe/board"
import "time"
import "strings"

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

func getDisplayBoard(board Board, player1Marker string, player2Marker string) string {
	rowLength := board.RowLength()
	charsPerRow := getCharsPerRow(rowLength)
	boardString := NewLineString
	stringBoard := board.StringBoard(player1Marker, player2Marker)
	for rowStartIndex := 0; rowStartIndex < len(board); {
		boardString = boardString + getHorizontalDividerString(charsPerRow) + NewLineString
		boardString = boardString + getBoardRowString(stringBoard, rowStartIndex, rowLength) + NewLineString
		rowStartIndex = rowStartIndex + rowLength
	}
	boardString = boardString + getHorizontalDividerString(charsPerRow) + NewLineString + NewLineString
	return boardString
}

func (ui ConsoleUI) DisplayBoard(board Board, player1Marker string, player2Marker string) {
	ui.Clear()
	ui.Output.Write(getDisplayBoard(board, player1Marker, player2Marker))
}

func getHumanSelectMessage(marker string) string {
	return marker + SelectASpotMessage
}

func (ui ConsoleUI) GetValidMove(board Board, marker string) int {
	selectMessage := getHumanSelectMessage(marker)
	ui.Output.Write(selectMessage)
	selected := ui.Input.ReadInt()
	for !ui.Validator.IsValid(selected-1, board.EmptySpots()) {
		ui.Output.Write(NotValidMessage)
		ui.Output.Write(selectMessage)
		selected = ui.Input.ReadInt()
	}
	return selected - 1
}

func (ui ConsoleUI) GetRefereeSelection() RefereeTypeSelection {
	ui.Output.Write(RefereeSelectMessage)
	selected := ui.Input.ReadInt()
	for !ui.Validator.IsValid(selected, []int{1, 2}) {
		ui.Output.Write(RefereeSelectMessage)
		ui.Output.Write(NotValidMessage)
		selected = ui.Input.ReadInt()
	}
	return RefereeTypeSelection(selected)
}

func (ui ConsoleUI) GetPlayerTypeSelection(playerName string, playerMarker string) PlayerTypeSelection {
	ui.Output.Write(NewLineString + NewLineString + playerName + VerticalDividerString + playerMarker + NewLineString)
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
	nameSelectionMessage := NameSelectMessage + playerNumber + ":" + NewLineString
	ui.Output.Write(nameSelectionMessage)
	selection := ui.Input.ReadStringOfLengthWithDefault(DefaultMaxNameLength, DefaultPlayerName+playerNumber)
	return selection
}

func (ui ConsoleUI) GetPlayerMarkerSelection(playerName string, playerNumber string) string {
	ui.Output.Write(MarkerSelectMessage + playerName + ":" + NewLineString)
	var def string
	if playerNumber == "1" {
		def = DefaultPlayer1Marker
	} else {
		def = DefaultPlayer2Marker
	}
	selection := ui.Input.ReadStringOfLengthWithDefault(1, def)
	return selection
}

func (ui ConsoleUI) GetPlayAgainSelection() bool {
	ui.Output.Write(PlayAgainSelectMessage)
	selection := strings.ToLower(ui.Input.ReadStringOfLengthWithDefault(1, "n"))
	return selection == "y"
}

func (ui ConsoleUI) Clear() {
	ui.Output.Clear()
}
