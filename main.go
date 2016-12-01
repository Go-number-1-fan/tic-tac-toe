package main

import . "github.com/go-number-1-fan/tic-tac-toe/ui"
import . "github.com/go-number-1-fan/tic-tac-toe/board"
import . "github.com/go-number-1-fan/tic-tac-toe/game"
import . "github.com/go-number-1-fan/tic-tac-toe/player"
import . "github.com/go-number-1-fan/tic-tac-toe/referee"

func main() {
	ui := ConsoleUI{ConsoleInput{}, ConsoleOutput{}}
	ref := StandardReferee{}
	game := CreateGame(ui, EmptyBoard(), HumanPlayer{"X"}, HardComputerPlayer{"O", ref}, ref)
	game.PlayGame()
}
