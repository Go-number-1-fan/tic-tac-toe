package player

import "github.com/stretchr/testify/assert"
import . "github.com/go-number-1-fan/tic-tac-toe/board"
import . "github.com/go-number-1-fan/tic-tac-toe/ui"
import . "github.com/go-number-1-fan/tic-tac-toe/referee"
import "testing"

func TestHardComputerPlayer_CanReturnAValidMarker(t *testing.T) {
	player := CreateHardComputerPlayer("X", "Computer", StandardReferee{})
	assert.Equal(t, player.GetMarker(), "X")
}

func TestHardComputerPlayer_CanReturnAValidName(t *testing.T) {
	player := HardComputerPlayer{"X", "Computer", StandardReferee{}}
	assert.Equal(t, player.GetName(), "Computer")
}

func TestHardComputerPlayer_CanTakeTheCornerOfAnOpenBoard(t *testing.T) {
	player := HardComputerPlayer{"X", "Computer", StandardReferee{}}
	board := Board{
		E, E, E,
		E, E, E,
		E, E, E}

	computerMove := player.GetMove(board, MockUI{-1})
	assert.Equal(t, 0, computerMove)
}

func TestHardComputerPlayer_CanWinGivenTheChance(t *testing.T) {
	player := HardComputerPlayer{"X", "Computer", StandardReferee{}}
	board := Board{
		X, X, E,
		E, O, E,
		E, O, E}

	computerMove := player.GetMove(board, MockUI{-1})
	assert.Equal(t, 2, computerMove)
}

func TestHardComputerPlayer_CanBlockAWinGivenTheChance(t *testing.T) {
	player := HardComputerPlayer{"X", "Computer", StandardReferee{}}
	board := Board{
		O, O, E,
		E, X, E,
		E, X, E}

	computerMove := player.GetMove(board, MockUI{-1})
	assert.Equal(t, 2, computerMove)
}

func TestHardComputerPlayer_CanTakeTheMiddleIfITakeACorner(t *testing.T) {
	player := HardComputerPlayer{"O", "Computer", StandardReferee{}}
	board := Board{
		X, E, E,
		E, E, E,
		E, E, E}

	computerMove := player.GetMove(board, MockUI{-1})
	assert.Equal(t, 4, computerMove)
}

func TestHardComputerPlayer_TakesTheLastSpotOnAnAlmostFullBoard(t *testing.T) {
	player := HardComputerPlayer{"X", "Computer", StandardReferee{}}
	board := Board{
		X, O, X,
		X, O, O,
		O, X, E}

	computerMove := player.GetMove(board, MockUI{-1})
	assert.Equal(t, 8, computerMove)
}

func TestHardComputerPlayer_CornerReferee_TakesACornerOnAnOpenBoard(t *testing.T) {
	player := HardComputerPlayer{"X", "Computer", CornerReferee{}}
	board := Board{
		E, E, E,
		E, E, E,
		E, E, E}

	computerMove := player.GetMove(board, MockUI{-1})
	assert.Equal(t, 0, computerMove)
}

func TestHardComputerPlayer_CornerReferee_StopsAWin(t *testing.T) {
	player := HardComputerPlayer{"O", "Computer", CornerReferee{}}
	board := Board{
		X, E, O,
		E, E, E,
		X, E, E}

	computerMove := player.GetMove(board, MockUI{-1})
	assert.Equal(t, 8, computerMove)
}

func TestHardComputerPlayer_CornerReferee_WinsIfItCan(t *testing.T) {
	player := HardComputerPlayer{"O", "Computer", CornerReferee{}}
	board := Board{
		X, E, O,
		X, E, X,
		O, E, E}

	computerMove := player.GetMove(board, MockUI{-1})
	assert.Equal(t, 8, computerMove)
}
