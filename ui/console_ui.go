package ui

import (
	"strings"
	"time"
)

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

func (ui ConsoleUI) DisplayBoard(board []string, player1Marker string, player2Marker string) {
	ui.Clear()
	stringBoard := getCustomizedStringBoard(board, player1Marker, player2Marker)
	ui.Output.Write(getDisplayBoard(stringBoard))
}

func getCustomizedStringBoard(stringBoard []string, player1Marker string, player2Marker string) []string {
	var customizedStringBoard []string
	for _, marker := range stringBoard {
		switch marker {
		case "X":
			customizedStringBoard = append(customizedStringBoard, DefaultPlayer1Color+player1Marker+DefaultPlayerColorClose)
		case "O":
			customizedStringBoard = append(customizedStringBoard, DefaultPlayer2Color+player2Marker+DefaultPlayerColorClose)
		default:
			customizedStringBoard = append(customizedStringBoard, marker)
		}
	}
	return customizedStringBoard
}

func getHumanSelectMessage(marker string) string {
	return marker + SelectASpotMessage
}

func (ui ConsoleUI) GetValidMove(openSpots []int, marker string) int {
	selectMessage := getHumanSelectMessage(marker)
	ui.Output.Write(selectMessage)
	selected := ui.Input.ReadInt(ui.Output)
	for !ui.Validator.IsValid(selected-1, openSpots) {
		ui.Output.Write(NotValidMessage)
		ui.Output.Write(selectMessage)
		selected = ui.Input.ReadInt(ui.Output)
	}
	return selected - 1
}

func (ui ConsoleUI) GetRefereeSelection() RefereeTypeSelection {
	ui.Output.Write(RefereeSelectMessage)
	selected := ui.Input.ReadInt(ui.Output)
	for !ui.Validator.IsValid(selected, []int{1, 2}) {
		ui.Output.Write(RefereeSelectMessage)
		ui.Output.Write(NotValidMessage)
		selected = ui.Input.ReadInt(ui.Output)
	}
	return RefereeTypeSelection(selected)
}

func (ui ConsoleUI) GetPlayerTypeSelection(playerName string, playerMarker string) PlayerTypeSelection {
	ui.Output.Write(NewLineString + NewLineString + playerName + VerticalDividerString + playerMarker + NewLineString)
	ui.Output.Write(PlayerTypeSelectMessage)
	selected := ui.Input.ReadInt(ui.Output)
	for !ui.Validator.IsValid(selected, []int{1, 2, 3}) {
		ui.Output.Write(PlayerTypeSelectMessage)
		ui.Output.Write(NotValidMessage)
		selected = ui.Input.ReadInt(ui.Output)
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
