package ui

import . "github.com/go-number-1-fan/tic-tac-toe/board"

type UI interface {
	DisplayWelcomeMessage()
	GetValidMove(board Board) int
	DisplayBoard(board Board)
}
