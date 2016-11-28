package game

import "github.com/stretchr/testify/assert"
import . "github.com/go-number-1-fan/tic-tac-toe/marker"
import . "github.com/go-number-1-fan/tic-tac-toe/player"
import . "github.com/go-number-1-fan/tic-tac-toe/mocks"
import "testing"

func TestGame_HasABoard(t *testing.T) {
	assert.Equal(t, EmptyBoard(), Game{EmptyBoard(), []Player{MockPlayer{X, 0}, MockPlayer{O, 1}}}.board)
}

func TestGame_CanMakeAMoveOnABoard(t *testing.T) {
	game := Game{EmptyBoard(), []Player{MockPlayer{X, 0}, MockPlayer{O, 1}}}
	game = game.takeTurn()
	assert.NotEqual(t, EmptyBoard(), game.board)
}

func TestGame_CanMakeAMoveAtASelectedLocation(t *testing.T) {
	game := Game{EmptyBoard(), []Player{MockPlayer{X, 4}, MockPlayer{O, 1}}}
	game = game.takeTurn()
	assert.Equal(t, X, game.board[4])
}

func TestGame_CanSwitchThePlayersAtTheEndOfTheTurn(t *testing.T) {
	currentPlayer := MockPlayer{X, 4}
	nextPlayer := MockPlayer{O, 1}
	game := Game{EmptyBoard(), []Player{currentPlayer, nextPlayer}}
	updatedGame := game.takeTurn()
	assert.Equal(t, nextPlayer, updatedGame.players[0])
}

func TestGame_PlaysWithDifferentMarkersEachTurn(t *testing.T) {
	currentPlayer := MockPlayer{X, 4}
	nextPlayer := MockPlayer{O, 1}
	game := Game{EmptyBoard(), []Player{currentPlayer, nextPlayer}}

	updatedGame := game.takeTurn()
	updatedGame = game.takeTurn()

	assert.Equal(t, X, updatedGame.board[4])
	assert.Equal(t, O, updatedGame.board[1])
}

func TestGame_CanCreateAGameWithTheGivenAttributes(t *testing.T) {
	board := Board{
		X, E, E,
		E, E, E,
		E, E, E,
	}
	firstPlayer := MockPlayer{X, 5}
	secondPlayer := MockPlayer{O, 2}

	game := CreateGame(board, firstPlayer, secondPlayer)

	assert.Equal(t, board, game.board)
	assert.Equal(t, firstPlayer, game.players[0])
	assert.Equal(t, secondPlayer, game.players[1])

}
