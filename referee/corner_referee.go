package referee

import . "github.com/go-number-1-fan/tic-tac-toe/board"

type CornerReferee struct{}

func (referee CornerReferee) GetGameStatus(board Board) GameStatus {
	winableCombination := getCornerWins(board)
	switch {
	case checkMarkerForAllButOne("X", winableCombination):
		return WinP1
	case checkMarkerForAllButOne("O", winableCombination):
		return WinP2
	case checkAllCornersOccupied(winableCombination):
		return Tie
	default:
		return Continue
	}

}

func checkMarkerForAllButOne(marker Marker, winableCombination WinableCombination) bool {
	incorrectCount := 0
	for _, winableMarker := range winableCombination {
		if winableMarker != marker {
			incorrectCount = incorrectCount + 1
		}
	}
	return incorrectCount <= 1
}

func checkAllCornersOccupied(winableCombination WinableCombination) bool {
	occupiedCorners := 0
	for _, winableMarker := range winableCombination {
		if winableMarker != E {
			occupiedCorners = occupiedCorners + 1
		}
	}
	return occupiedCorners == 4
}

func getCornerWins(board Board) WinableCombination {
	rowLength := board.RowLength()
	corner1 := board.MarkerAt(0)
	corner2 := board.MarkerAt(rowLength - 1)
	corner3 := board.MarkerAt(rowLength * (rowLength - 1))
	corner4 := board.MarkerAt(rowLength*rowLength - 1)
	return WinableCombination{corner1, corner2, corner3, corner4}
}
