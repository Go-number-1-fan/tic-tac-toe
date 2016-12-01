package player

import "github.com/stretchr/testify/assert"
import . "github.com/go-number-1-fan/tic-tac-toe/board"
import . "github.com/go-number-1-fan/tic-tac-toe/ui"
import . "github.com/go-number-1-fan/tic-tac-toe/referee"
import "testing"

func TestHardComputerPlayer_CanTakeTheCornerOfAnOpenBoard(t *testing.T) {
	player := HardComputerPlayer{"X", StandardReferee{}}
	board := Board{
		E, E, E,
		E, E, E,
		E, E, E}

	computerMove := player.GetMove(board, MockUI{-1})
	assert.Equal(t, 0, computerMove)
}

func TestHardComputerPlayer_CanWinGivenTheChance(t *testing.T) {
	player := HardComputerPlayer{"X", StandardReferee{}}
	board := Board{
		X, X, E,
		E, O, E,
		E, O, E}

	computerMove := player.GetMove(board, MockUI{-1})
	assert.Equal(t, 2, computerMove)
}

func TestHardComputerPlayer_CanBlockAWinGivenTheChance(t *testing.T) {
	player := HardComputerPlayer{"X", StandardReferee{}}
	board := Board{
		O, O, E,
		E, X, E,
		E, X, E}

	computerMove := player.GetMove(board, MockUI{-1})
	assert.Equal(t, 2, computerMove)
}

func TestHardComputerPlayer_CanTakeTheMiddleIfITakeACorner(t *testing.T) {
	player := HardComputerPlayer{"O", StandardReferee{}}
	board := Board{
		X, E, E,
		E, E, E,
		E, E, E}

	computerMove := player.GetMove(board, MockUI{-1})
	assert.Equal(t, 4, computerMove)
}

func TestHardComputerPlayer_TakesTheLastSpotOnAnAlmostFullBoard(t *testing.T) {
	player := HardComputerPlayer{"X", StandardReferee{}}
	board := Board{
		X, O, X,
		X, O, O,
		O, X, E}

	computerMove := player.GetMove(board, MockUI{-1})
	assert.Equal(t, 8, computerMove)
}
