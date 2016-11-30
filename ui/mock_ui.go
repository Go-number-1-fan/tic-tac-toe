package ui

import . "github.com/go-number-1-fan/tic-tac-toe/board"

type MockUI struct {
	MockUserInput int
}

func (ui MockUI) DisplayWelcomeMessage() {

}

func (ui MockUI) GetValidMove(board Board, marker string) int {
	return ui.MockUserInput
}

func (ui MockUI) DisplayBoard(board Board) {
}