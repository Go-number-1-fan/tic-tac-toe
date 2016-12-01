package main

import . "github.com/go-number-1-fan/tic-tac-toe/ui"
import . "github.com/go-number-1-fan/tic-tac-toe/game"
import . "github.com/go-number-1-fan/tic-tac-toe/referee"

func main() {
	ui := CreateConsoleUI(ConsoleInput{}, ConsoleOutput{})
	ref := StandardReferee{}
	ui.DisplayWelcomeMessage()

	gameBuilder := GameBuilder{ui, ref}
	game := gameBuilder.BuildGame()

	game.PlayGame()
}
