package game

import . "github.com/go-number-1-fan/tic-tac-toe/board"

type WinableCombination []Marker

type StandardReferee struct{}

func (referee StandardReferee) GetGameStatus(board Board) GameStatus {
	allWinableCombinations := getAllWinableCombinations(board)
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

func getAllWinableCombinations(board Board) []WinableCombination {
	var allWinableCombinations []WinableCombination
	allWinableCombinations = append(allWinableCombinations, getHorizontalWins(board)...)
	allWinableCombinations = append(allWinableCombinations, getVerticalWins(board)...)
	allWinableCombinations = append(allWinableCombinations, getDiagonalWins(board)...)
	return allWinableCombinations
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

func getVerticalWins(board Board) []WinableCombination {
	rowLength := board.RowLength()
	horizontalWins := getHorizontalWins(board)
	var verticalWins []WinableCombination

	for column := 0; column < rowLength; column++ {
		var verticalWin WinableCombination

		for _, horizontalWin := range horizontalWins {
			verticalWin = append(verticalWin, horizontalWin[column])
		}
		verticalWins = append(verticalWins, verticalWin)
	}
	return verticalWins
}

func getDiagonalWins(board Board) []WinableCombination {
	rowLength := board.RowLength()
	horizontalWins := getHorizontalWins(board)
	var diagonalWins []WinableCombination
	var diagonalWin1 WinableCombination
	var diagonalWin2 WinableCombination
	for start := 0; start < rowLength; start++ {
		diagonalWin1 = append(diagonalWin1, horizontalWins[start][start])
		diagonalWin2 = append(diagonalWin2, horizontalWins[start][(rowLength-1)-start])
	}
	diagonalWins = append(diagonalWins, diagonalWin1)
	diagonalWins = append(diagonalWins, diagonalWin2)
	return diagonalWins
}
