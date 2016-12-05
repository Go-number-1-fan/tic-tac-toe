package main

import (
	. "github.com/go-number-1-fan/tic-tac-toe/game"
	. "github.com/go-number-1-fan/tic-tac-toe/ui"
)

func main() {
	ui := CreateConsoleUI(ConsoleInput{}, ConsoleOutput{})
	ui.Clear()
	ui.DisplayWelcomeMessage()

	gameBuilder := GameBuilder{ui}
	game := gameBuilder.BuildGame()

	playAgain := true
	for playAgain {
		game.PlayGame()
		playAgain = ui.GetPlayAgainSelection()
		game = game.ResetGame()
	}
}
