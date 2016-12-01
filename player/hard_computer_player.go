package player

import . "github.com/go-number-1-fan/tic-tac-toe/board"
import . "github.com/go-number-1-fan/tic-tac-toe/ui"
import . "github.com/go-number-1-fan/tic-tac-toe/referee"

type HardComputerPlayer struct {
	Marker  string
	Referee Referee
}
type ScoredSpot struct {
	Spot  int
	Score int
}

func CreateHardComputerPlayer(marker string, ref Referee) HardComputerPlayer {
	return HardComputerPlayer{marker, ref}
}

func (player HardComputerPlayer) GetMove(board Board, ui UI) int {
	ui.DisplayComputerThinkingMessage()
	return player.scoreSpot(board, getCurrentMarker(board), 0)
}

func (computer HardComputerPlayer) scoreSpot(board Board, currentMarker Marker, depth int) int {
	var scoredSpots []ScoredSpot
	currentGameStatus := computer.Referee.GetGameStatus(board)

	switch currentGameStatus {
	case WinP1:
		return (-10 + depth)
	case WinP2:
		return (-10 + depth)
	case Tie:
		return 0
	default:
		emptySpots := board.EmptySpots()
		for _, emptySpot := range emptySpots {
			board := board.MakeMove(emptySpot, currentMarker)
			nextMarker := flipMarker(currentMarker)
			score := -1 * (computer.scoreSpot(board, nextMarker, depth+1))
			scoredSpots = append(scoredSpots, ScoredSpot{emptySpot, score})
			board[emptySpot] = E
		}
	}
	maxSpot, maxScore := getMaxs(scoredSpots)
	if depth == 0 {
		return maxSpot
	} else {
		return maxScore
	}
}

func (player HardComputerPlayer) GetMarker() string {
	return player.Marker
}

func flipMarker(currentMarker Marker) Marker {

	if currentMarker == X {
		return O
	} else {
		return X
	}
}

func getMaxs(scoredSpots []ScoredSpot) (int, int) {
	var maxScore int = -11
	var maxSpot int = 0
	for _, scoredSpot := range scoredSpots {
		if scoredSpot.Score > maxScore {
			maxScore = scoredSpot.Score
			maxSpot = scoredSpot.Spot
		}
	}
	return maxSpot, maxScore
}

func getCurrentMarker(board Board) Marker {
	if board.CountOccupiedSpots()%2 == 0 {
		return X
	}
	return O
}
