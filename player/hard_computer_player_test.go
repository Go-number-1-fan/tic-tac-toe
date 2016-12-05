package player

import (
	. "github.com/go-number-1-fan/tic-tac-toe/board"
	. "github.com/go-number-1-fan/tic-tac-toe/referee"
	. "github.com/go-number-1-fan/tic-tac-toe/ui"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHardComputerPlayer_CanReturnAValidMarker(t *testing.T) {
	player := CreateHardComputerPlayer("X", "Computer", StandardReferee{})
	assert.Equal(t, player.GetMarker(), "X")
}

func TestHardComputerPlayer_CanReturnAValidName(t *testing.T) {
	player := HardComputerPlayer{"X", "Computer", StandardReferee{}}
	assert.Equal(t, player.GetName(), "Computer")
}

func TestHardComputerPlayer_NeverLosesAGame_StandardRef_GoingSecond(t *testing.T) {
	board := EmptyBoard()
	computer := CreateHardComputerPlayer("O", "Computer", StandardReferee{})
	playerWon := playAllGamesVsComputer(board, X, computer, false, t)
	assert.False(t, playerWon)
}

func TestHardComputerPlayer_NeverLosesAGame_StandardRef_GoingFirst(t *testing.T) {
	board := EmptyBoard()
	computer := CreateHardComputerPlayer("O", "Computer", StandardReferee{})
	board = board.MakeMove(computer.GetMove(board, MockUI{-1}), X)
	playerWon := playAllGamesVsComputer(board, O, computer, true, t)
	assert.False(t, playerWon)
}

func TestHardComputerPlayer_NeverLosesAGame_CornerRef_GoingSecond(t *testing.T) {
	board := EmptyBoard()
	computer := CreateHardComputerPlayer("O", "Computer", CornerReferee{})
	playerWon := playAllGamesVsComputer(board, X, computer, false, t)
	assert.False(t, playerWon)
}

func TestHardComputerPlayer_NeverLosesAGame_CornerRef_GoingFirst(t *testing.T) {
	board := EmptyBoard()
	computer := CreateHardComputerPlayer("O", "Computer", CornerReferee{})
	board = board.MakeMove(computer.GetMove(board, MockUI{-1}), X)
	playerWon := playAllGamesVsComputer(board, O, computer, true, t)
	assert.False(t, playerWon)
}

func playAllGamesVsComputer(board Board, currentMarker Marker, computer HardComputerPlayer, computerP1 bool, t *testing.T) bool {
	playerWon := false
	currentGameStatus := computer.Referee.GetGameStatus(board)

	switch currentGameStatus {
	case WinP1:
		return !computerP1
	case WinP2:
		return computerP1
	case Tie:
		return false
	default:
		nextMarker := flipMarker(currentMarker)
		if computerP1 {
			if currentMarker == X {
				computerMove := computer.GetMove(board, MockUI{-1})
				updatedBoard := append(Board(nil), board...).MakeMove(computerMove, currentMarker)
				playerWon = playerWon || playAllGamesVsComputer(updatedBoard, nextMarker, computer, computerP1, t)
			} else {
				emptySpots := board.EmptySpots()
				for _, emptySpot := range emptySpots {
					updatedBoard := append(Board(nil), board...).MakeMove(emptySpot, currentMarker)
					playerWon = playerWon || playAllGamesVsComputer(updatedBoard, nextMarker, computer, computerP1, t)
					updatedBoard = updatedBoard.MakeMove(emptySpot, E)
				}
			}
		} else {
			if currentMarker == X {
				emptySpots := board.EmptySpots()
				for _, emptySpot := range emptySpots {
					updatedBoard := append(Board(nil), board...).MakeMove(emptySpot, currentMarker)
					playerWon = playerWon || playAllGamesVsComputer(updatedBoard, nextMarker, computer, computerP1, t)
					updatedBoard = updatedBoard.MakeMove(emptySpot, E)
				}
			} else {
				computerMove := computer.GetMove(board, MockUI{-1})
				updatedBoard := append(Board(nil), board...).MakeMove(computerMove, currentMarker)
				playerWon = playerWon || playAllGamesVsComputer(updatedBoard, nextMarker, computer, computerP1, t)
			}
		}
	}
	return playerWon
}
