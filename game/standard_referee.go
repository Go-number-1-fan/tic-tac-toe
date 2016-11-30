package game

import . "github.com/go-number-1-fan/tic-tac-toe/board"

type WinableCombination []Marker

type StandardReferee struct{}

func (referee StandardReferee) GetGameStatus(board Board) GameStatus {
	var allWinableCombinations []WinableCombination
	allWinableCombinations = append(allWinableCombinations, getHorizontalWins(board)...)

	for _, winableCombination := range allWinableCombinations {
		switch {
		case checkWinableCombinationForWin(X, winableCombination):
			return WinP1
		case checkWinableCombinationForWin(O, winableCombination):
			return WinP2
		}
	}
	return Continue
}

func checkWinableCombinationForWin(marker Marker, winableCombination WinableCombination) bool {
	for _, winableMarker := range winableCombination {
		if winableMarker != marker {
			return false
		}
	}
	return true
}

func getHorizontalWins(board Board) []WinableCombination {
	rowLength := board.RowLength()
	var wins []WinableCombination
	for rowStartIndex := 0; rowStartIndex < len(board); {
		wins = append(wins, getHorizontalWin(board, rowStartIndex, rowLength))
		rowStartIndex = rowStartIndex + rowLength
	}
	return wins
}

func getHorizontalWin(board Board, boardRowStartIndex int, rowLength int) WinableCombination {
	var win WinableCombination
	for boardColumn := 0; boardColumn < rowLength; boardColumn++ {
		win = append(win, board.MarkerAt(boardRowStartIndex+boardColumn))
	}
	return win
}
