package game

import . "github.com/go-number-1-fan/tic-tac-toe/board"

type MockReferee struct {
	MockGameStatus GameStatus
}

func (referee MockReferee) GetGameStatus(board Board) GameStatus {
	return referee.MockGameStatus
}
