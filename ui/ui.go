package ui

import . "github.com/go-number-1-fan/tic-tac-toe/board"

type UI interface {
	DisplayBoard(board Board)
	DisplayTieMessage()
	DisplayComputerThinkingMessage()
	DisplayWelcomeMessage()
	DisplayWinMessage(winner string)
	GetValidMove(board Board, marker string) int
}
