package game

import (
	. "github.com/go-number-1-fan/tic-tac-toe/board"
	. "github.com/go-number-1-fan/tic-tac-toe/player"
	. "github.com/go-number-1-fan/tic-tac-toe/referee"
	. "github.com/go-number-1-fan/tic-tac-toe/ui"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGame_HasABoard(t *testing.T) {
	assert.Equal(t, EmptyBoard(), CreateGame(MockUI{0}, EmptyBoard(), MockPlayer{"X", "Player1", 0}, MockPlayer{"O", "Player2", 1}, MockReferee{Continue}).board)
}

func TestGame_CanMakeAMoveOnABoard(t *testing.T) {
	game := CreateGame(MockUI{0}, EmptyBoard(), MockPlayer{"X", "Player1", 0}, MockPlayer{"O", "Player2", 1}, MockReferee{Continue})
	game = game.takeTurn()
	assert.NotEqual(t, EmptyBoard(), game.board)
}

func TestGame_CanMakeAMoveAtASelectedLocation(t *testing.T) {
	game := CreateGame(MockUI{0}, EmptyBoard(), MockPlayer{"X", "Player1", 4}, MockPlayer{"O", "Player2", 1}, MockReferee{Continue})
	game = game.takeTurn()
	assert.Equal(t, X, game.board[4])
}

func TestGame_CanSwitchThePlayersAtTheEndOfTheTurn(t *testing.T) {
	currentPlayer := MockPlayer{"X", "Player1", 4}
	nextPlayer := MockPlayer{"O", "Player2", 1}
	game := CreateGame(MockUI{0}, EmptyBoard(), currentPlayer, nextPlayer, MockReferee{Continue})
	updatedGame := game.takeTurn()
	assert.Equal(t, nextPlayer, updatedGame.players[0])
}

func TestGame_DoesNotSwitchThePlayersAtTheEndOfTheTurn_IfTheresAWinner(t *testing.T) {
	currentPlayer := MockPlayer{"X", "Player1", 4}
	nextPlayer := MockPlayer{"O", "Player2", 1}
	game := CreateGame(MockUI{0}, EmptyBoard(), currentPlayer, nextPlayer, MockReferee{WinP1})
	updatedGame := game.takeTurn()
	assert.NotEqual(t, nextPlayer, updatedGame.players[0])
}

func TestGame_DoesNotSwitchThePlayersAtTheEndOfTheTurn_IfTheresATie(t *testing.T) {
	currentPlayer := MockPlayer{"X", "Player1", 4}
	nextPlayer := MockPlayer{"O", "Player2", 1}
	game := CreateGame(MockUI{0}, EmptyBoard(), currentPlayer, nextPlayer, MockReferee{Tie})
	updatedGame := game.takeTurn()
	assert.NotEqual(t, nextPlayer, updatedGame.players[0])
}

func TestGame_EndsIfTheStatusIsWin(t *testing.T) {
	currentPlayer := MockPlayer{"X", "Player1", 4}
	nextPlayer := MockPlayer{"O", "Player2", 1}
	board := Board{
		E, X, O,
		E, E, O,
		E, X, E,
	}
	game := CreateGame(MockUI{0}, board, currentPlayer, nextPlayer, MockReferee{WinP1})
	updatedGame := game.PlayGame()
	assert.Equal(t, X, updatedGame.board[4])
	assert.NotEqual(t, nextPlayer, updatedGame.players[0])
}

func TestGame_PlaysWithDifferentMarkersEachTurn(t *testing.T) {
	currentPlayer := MockPlayer{"X", "Player1", 4}
	nextPlayer := MockPlayer{"O", "Player2", 1}
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
	firstPlayer := MockPlayer{"X", "Player1", 5}
	secondPlayer := MockPlayer{"O", "Player2", 2}

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
	firstPlayer := MockPlayer{"X", "Player1", 5}
	secondPlayer := MockPlayer{"O", "Player2", 2}

	game := CreateGame(MockUI{1}, board, firstPlayer, secondPlayer, MockReferee{Continue})

	assert.Equal(t, X, game.getCurrentPlayerMarker())
}

func TestGame_CanGetTheCurrentPlayersMarker_ForO(t *testing.T) {
	board := Board{
		E, E, E,
		E, E, X,
		E, E, E,
	}
	firstPlayer := MockPlayer{"X", "Player1", 5}
	secondPlayer := MockPlayer{"O", "Player2", 2}

	game := CreateGame(MockUI{1}, board, firstPlayer, secondPlayer, MockReferee{Continue})

	assert.Equal(t, O, game.getCurrentPlayerMarker())
}

func TestGame_CanTakeATurnWithAComputer(t *testing.T) {
	board := Board{
		X, O, X,
		X, O, E,
		O, X, E,
	}
	firstPlayer := MockPlayer{"X", "Player1", 5}
	secondPlayer := EasyComputerPlayer{"O", "Computer2"}

	game := CreateGame(MockUI{1}, board, secondPlayer, firstPlayer, StandardReferee{})
	updatedGame := game.takeTurn()
	assert.True(t, updatedGame.board[5] == O || updatedGame.board[8] == O)
}

func TestGame_CanPlayAnEntireGameWithComputers(t *testing.T) {
	board := Board{
		E, E, E,
		E, E, E,
		E, E, E,
	}
	firstPlayer := EasyComputerPlayer{"X", "Computer1"}
	secondPlayer := EasyComputerPlayer{"O", "Computer2"}

	game := CreateGame(MockUI{1}, board, firstPlayer, secondPlayer, StandardReferee{})
	updatedGame := game.PlayGame()
	assert.NotEqual(t, Continue, updatedGame.status)
}

func TestGame_CanResetTheGameToPlayAgain(t *testing.T) {
	board := Board{
		E, E, E,
		E, E, E,
		E, E, E,
	}
	firstPlayer := EasyComputerPlayer{"X", "Computer1"}
	secondPlayer := EasyComputerPlayer{"O", "Computer2"}

	game := CreateGame(MockUI{1}, board, firstPlayer, secondPlayer, StandardReferee{})
	updatedGame := game.PlayGame()
	assert.NotEqual(t, Continue, updatedGame.status)
	assert.True(t, len(updatedGame.board.EmptySpots()) < 6)
	clearedGame := updatedGame.ResetGame()
	assert.Equal(t, Continue, clearedGame.status)
	assert.Equal(t, 9, len(clearedGame.board.EmptySpots()))
}
