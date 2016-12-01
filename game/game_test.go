package game

import "github.com/stretchr/testify/assert"
import . "github.com/go-number-1-fan/tic-tac-toe/board"
import . "github.com/go-number-1-fan/tic-tac-toe/player"
import . "github.com/go-number-1-fan/tic-tac-toe/referee"
import . "github.com/go-number-1-fan/tic-tac-toe/ui"
import "testing"

func TestGame_HasABoard(t *testing.T) {
	assert.Equal(t, EmptyBoard(), CreateGame(MockUI{0}, EmptyBoard(), MockPlayer{"X", 0}, MockPlayer{"O", 1}, MockReferee{Continue}).board)
}

func TestGame_CanMakeAMoveOnABoard(t *testing.T) {
	game := CreateGame(MockUI{0}, EmptyBoard(), MockPlayer{"X", 0}, MockPlayer{"O", 1}, MockReferee{Continue})
	game = game.takeTurn()
	assert.NotEqual(t, EmptyBoard(), game.board)
}

func TestGame_CanMakeAMoveAtASelectedLocation(t *testing.T) {
	game := CreateGame(MockUI{0}, EmptyBoard(), MockPlayer{"X", 4}, MockPlayer{"O", 1}, MockReferee{Continue})
	game = game.takeTurn()
	assert.Equal(t, X, game.board[4])
}

func TestGame_CanSwitchThePlayersAtTheEndOfTheTurn(t *testing.T) {
	currentPlayer := MockPlayer{"X", 4}
	nextPlayer := MockPlayer{"O", 1}
	game := CreateGame(MockUI{0}, EmptyBoard(), currentPlayer, nextPlayer, MockReferee{Continue})
	updatedGame := game.takeTurn()
	assert.Equal(t, nextPlayer, updatedGame.players[0])
}

func TestGame_DoesNotSwitchThePlayersAtTheEndOfTheTurn_IfTheresAWinner(t *testing.T) {
	currentPlayer := MockPlayer{"X", 4}
	nextPlayer := MockPlayer{"O", 1}
	game := CreateGame(MockUI{0}, EmptyBoard(), currentPlayer, nextPlayer, MockReferee{WinP1})
	updatedGame := game.takeTurn()
	assert.NotEqual(t, nextPlayer, updatedGame.players[0])
}

func TestGame_DoesNotSwitchThePlayersAtTheEndOfTheTurn_IfTheresATie(t *testing.T) {
	currentPlayer := MockPlayer{"X", 4}
	nextPlayer := MockPlayer{"O", 1}
	game := CreateGame(MockUI{0}, EmptyBoard(), currentPlayer, nextPlayer, MockReferee{Tie})
	updatedGame := game.takeTurn()
	assert.NotEqual(t, nextPlayer, updatedGame.players[0])
}

func TestGame_EndsIfTheStatusIsWin(t *testing.T) {
	currentPlayer := MockPlayer{"X", 4}
	nextPlayer := MockPlayer{"O", 1}
	game := CreateGame(MockUI{0}, EmptyBoard(), currentPlayer, nextPlayer, MockReferee{Tie})
	updatedGame := game.takeTurn()
	assert.NotEqual(t, nextPlayer, updatedGame.players[0])
}

func TestGame_PlaysWithDifferentMarkersEachTurn(t *testing.T) {
	currentPlayer := MockPlayer{"X", 4}
	nextPlayer := MockPlayer{"O", 1}
	game := CreateGame(MockUI{0}, EmptyBoard(), currentPlayer, nextPlayer, MockReferee{Continue})

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
	firstPlayer := MockPlayer{"X", 5}
	secondPlayer := MockPlayer{"O", 2}

	game := CreateGame(MockUI{1}, board, firstPlayer, secondPlayer, MockReferee{Continue})

	assert.Equal(t, board, game.board)
	assert.Equal(t, firstPlayer, game.players[0])
	assert.Equal(t, secondPlayer, game.players[1])
}

func TestGame_CanGetTheCurrentPlayersMarker_ForX(t *testing.T) {
	board := Board{
		E, E, E,
		E, E, E,
		E, E, E,
	}
	firstPlayer := MockPlayer{"X", 5}
	secondPlayer := MockPlayer{"O", 2}

	game := CreateGame(MockUI{1}, board, firstPlayer, secondPlayer, MockReferee{Continue})

	assert.Equal(t, X, game.getCurrentPlayerMarker())
}

func TestGame_CanGetTheCurrentPlayersMarker_ForO(t *testing.T) {
	board := Board{
		E, E, E,
		E, E, X,
		E, E, E,
	}
	firstPlayer := MockPlayer{"X", 5}
	secondPlayer := MockPlayer{"O", 2}

	game := CreateGame(MockUI{1}, board, firstPlayer, secondPlayer, MockReferee{Continue})

	assert.Equal(t, O, game.getCurrentPlayerMarker())
}

// Computer Test

func TestGame_CanTakeATurnWithAComputer(t *testing.T) {
	board := Board{
		X, O, X,
		X, O, E,
		O, X, E,
	}
	firstPlayer := MockPlayer{"X", 5}
	secondPlayer := EasyComputerPlayer{"O"}

	game := CreateGame(MockUI{1}, board, secondPlayer, firstPlayer, StandardReferee{})
	updatedGame := game.takeTurn()
	assert.True(t, updatedGame.board[5] == O || updatedGame.board[8] == O)
}
