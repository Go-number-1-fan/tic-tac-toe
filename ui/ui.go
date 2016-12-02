package ui

type UI interface {
	DisplayBoard(board []string, player1Marker string, player2Marker string)
	DisplayTieMessage()
	DisplayComputerThinkingMessage(computerName string)
	DisplayWelcomeMessage()
	DisplayWinMessage(winner string)
	GetValidMove(openSpots []int, marker string) int
	GetRefereeSelection() RefereeTypeSelection
	GetPlayerNameSelection(playerNumber string) string
	GetPlayerTypeSelection(playerName string, playerMarker string) PlayerTypeSelection
	GetPlayerMarkerSelection(playerName string, playerNumber string) string
	GetPlayAgainSelection() bool
	Clear()
}
