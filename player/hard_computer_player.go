package player

import (
	. "github.com/go-number-1-fan/tic-tac-toe/board"
	. "github.com/go-number-1-fan/tic-tac-toe/referee"
	. "github.com/go-number-1-fan/tic-tac-toe/ui"
)

type HardComputerPlayer struct {
	Marker  string
	Name    string
	Referee Referee
}

type ScoredSpot struct {
	Spot  int
	Score int
}

func CreateHardComputerPlayer(marker string, name string, ref Referee) HardComputerPlayer {
	return HardComputerPlayer{marker, name, ref}
}

func (computer HardComputerPlayer) GetMarker() string {
	return computer.Marker
}

func (computer HardComputerPlayer) GetName() string {
	return computer.Name
}

func (computer HardComputerPlayer) GetMove(board Board, ui UI) int {
	var scoredSpots []ScoredSpot
	var depth int
	ui.DisplayComputerThinkingMessage(computer.Name)
	channel := make(chan ScoredSpot)

	emptySpots := board.EmptySpots()
	for _, emptySpot := range emptySpots {
		go func(emptySpot int) {
			currentMarker := getCurrentMarker(board)
			boardCopy := append(Board(nil), board...)
			boardCopy = boardCopy.MakeMove(emptySpot, currentMarker)
			nextMarker := flipMarker(currentMarker)
			computer.channelNegamax(emptySpot, boardCopy, nextMarker, depth+1, channel)
			boardCopy[emptySpot] = E
		}(emptySpot)
	}

	for range board.EmptySpots() {
		scoredSpot := <-channel
		scoredSpots = append(scoredSpots, scoredSpot)
	}

	maxSpot, _ := getMaxs(scoredSpots)

	return maxSpot
}

func (computer HardComputerPlayer) channelNegamax(originalSpot int, board Board, currentMarker Marker, depth int, c chan ScoredSpot) {
	var scoredSpot ScoredSpot
	scoredSpot.Spot = originalSpot
	scoredSpot.Score = -1 * computer.scoreSpot(board, currentMarker, depth)
	c <- scoredSpot
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
	_, maxScore := getMaxs(scoredSpots)
	return maxScore
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
