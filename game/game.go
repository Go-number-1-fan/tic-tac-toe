package game

import . "github.com/go-number-1-fan/tic-tac-toe/board"
import . "github.com/go-number-1-fan/tic-tac-toe/player"
import . "github.com/go-number-1-fan/tic-tac-toe/referee"
import . "github.com/go-number-1-fan/tic-tac-toe/ui"

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

func (game Game) PlayGame() {
	game.ui.DisplayWelcomeMessage()
	for game.status == Continue {
		game = game.takeTurn()
	}
}

func (game Game) takeTurn() Game {
	currentPlayer := game.players[0]
	game.ui.DisplayBoard(game.board)
	playerMove := currentPlayer.GetMove(game.board, game.ui)
	game.board = game.board.MakeMove(playerMove, game.getCurrentPlayerMarker())

	game.status = game.referee.GetGameStatus(game.board)

	switch game.status {
	case Continue:
		game = game.swapPlayers()
	case Tie:
		game.ui.DisplayBoard(game.board)
		game.ui.DisplayTieMessage()
	default:
		game.ui.DisplayBoard(game.board)
		game.ui.DisplayWinMessage(currentPlayer.GetMarker())
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
