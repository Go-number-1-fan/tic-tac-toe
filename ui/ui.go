package ui

import . "github.com/go-number-1-fan/tic-tac-toe/board"

type UI interface {
	DisplayBoard(board Board)
	DisplayTieMessage()
	DisplayComputerThinkingMessage()
	DisplayWelcomeMessage()
	DisplayWinMessage(winner string)
	GetValidMove(board Board, marker string) int
	GetPlayerTypeSelection(playerName string) PlayerTypeSelection
}

type PlayerTypeSelection int

const (
	HumanSelected        PlayerTypeSelection = 1
	EasyComputerSelected PlayerTypeSelection = 2
	HardComputerSelected PlayerTypeSelection = 3
)
