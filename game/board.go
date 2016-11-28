package game

import . "github.com/go-number-1-fan/tic-tac-toe/marker"

type Board []Marker

func EmptyBoard() Board {
	return Board{
		E, E, E,
		E, E, E,
		E, E, E,
	}
}

func (board Board) MakeMove(spot int, marker Marker) Board {
	board[spot] = marker
	return board
}

func (board Board) emptyAt(spot int) bool {
	return board[spot] == E
}

func (board Board) EmptySpots() []int {
	var emptySpots []int
	for spot := range board {
		if board.emptyAt(spot) {
			emptySpots = append(emptySpots, spot)
		}
	}
	return emptySpots
}

func (board Board) CountOccupiedSpots() int {
	return len(board) - len(board.EmptySpots())
}
