package main

import . "github.com/go-number-1-fan/tic-tac-toe/ui"
import . "github.com/go-number-1-fan/tic-tac-toe/board"
import . "github.com/go-number-1-fan/tic-tac-toe/game"
import . "github.com/go-number-1-fan/tic-tac-toe/player"

func main() {
	ui := ConsoleUI{ConsoleInput{}, ConsoleOutput{}}
	game := CreateGame(ui, EmptyBoard(), HumanPlayer{X, ui}, HumanPlayer{O, ui})
	game.PlayGame()
}
