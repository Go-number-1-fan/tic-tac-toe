package ui

import . "github.com/go-number-1-fan/tic-tac-toe/board"

type UI interface {
	DisplayBoard(board Board)
	DisplayTieMessage()
	DisplayWelcomeMessage()
	DisplayWinMessage(winner string)
	GetValidMove(board Board, marker string) int
}
