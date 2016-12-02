package ui

import . "github.com/go-number-1-fan/tic-tac-toe/board"

type UI interface {
	DisplayBoard(board Board, player1Marker string, player2Marker string)
	DisplayTieMessage()
	DisplayComputerThinkingMessage(computerName string)
	DisplayWelcomeMessage()
	DisplayWinMessage(winner string)
	GetValidMove(board Board, marker string) int
	GetRefereeSelection() RefereeTypeSelection
	GetPlayerNameSelection(playerNumber string) string
	GetPlayerTypeSelection(playerName string, playerMarker string) PlayerTypeSelection
	GetPlayerMarkerSelection(playerName string, playerNumber string) string
	GetPlayAgainSelection() bool
}
