package game

import . "github.com/go-number-1-fan/tic-tac-toe/marker"
import . "github.com/go-number-1-fan/tic-tac-toe/player"

type Game struct {
	board   Board
	players []Player
}

func CreateGame(board Board, player1 Player, player2 Player) Game {
	game := Game{
		board,
		[]Player{player1, player2},
	}
	return game
}

func (game Game) playGame() {
}

func (game Game) swapPlayers() Game {
	game.players[0], game.players[1] = game.players[1], game.players[0]
	return game
}

func (game Game) takeTurn() Game {
	currentPlayer := game.players[0]
	game.board = game.board.MakeMove(currentPlayer.GetMove(), game.getCurrentPlayerMarker())
	game = game.swapPlayers()
	return game
}

func (game Game) getCurrentPlayerMarker() Marker {
	occupiedSpots := game.board.CountOccupiedSpots()
	if occupiedSpots%2 == 0 {
		return X
	} else {
		return O
	}
}
