package game

import (
	. "github.com/go-number-1-fan/tic-tac-toe/board"
	. "github.com/go-number-1-fan/tic-tac-toe/player"
	. "github.com/go-number-1-fan/tic-tac-toe/referee"
	. "github.com/go-number-1-fan/tic-tac-toe/ui"
)

type Game struct {
	ui      UI
	board   Board
	players []Player
	referee Referee
	status  GameStatus
}

func CreateGame(ui UI, board Board, player1 Player, player2 Player, referee Referee) Game {
	defaultGameStatus := Continue
	game := Game{
		ui,
		board,
		[]Player{player1, player2},
		referee,
		defaultGameStatus,
	}
	return game
}

func (game Game) ResetGame() Game {
	game.status = Continue
	player2, player1 := game.getPlayersInOrder()
	game.players = []Player{player1, player2}
	game.board = EmptyBoard()
	return game
}

func (game Game) PlayGame() Game {
	for game.status == Continue {
		game = game.takeTurn()
	}
	return game
}

func (game Game) takeTurn() Game {
	currentPlayer := game.players[0]
	player1, player2 := game.getPlayersInOrder()
	game.ui.DisplayBoard(game.board.StringBoard(), player1.GetMarker(), player2.GetMarker())
	playerMove := currentPlayer.GetMove(game.board, game.ui)
	game.board = game.board.MakeMove(playerMove, game.getCurrentPlayerMarker())

	game.status = game.referee.GetGameStatus(game.board)

	switch game.status {
	case Continue:
		game = game.swapPlayers()
	case Tie:
		game.ui.DisplayBoard(game.board.StringBoard(), player1.GetMarker(), player2.GetMarker())
		game.ui.DisplayTieMessage()
	default:
		game.ui.DisplayBoard(game.board.StringBoard(), player1.GetMarker(), player2.GetMarker())
		game.ui.DisplayWinMessage(currentPlayer.GetName())
	}
	return game
}

func (game Game) swapPlayers() Game {
	game.players[0], game.players[1] = game.players[1], game.players[0]
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

func (game Game) getPlayersInOrder() (Player, Player) {
	occupiedSpots := game.board.CountOccupiedSpots()
	var player1, player2 Player
	if occupiedSpots%2 == 0 {
		player1, player2 = game.players[0], game.players[1]
	} else {
		player1, player2 = game.players[1], game.players[0]
	}
	return player1, player2
}
