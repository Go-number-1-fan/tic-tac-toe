package board

import "math"
import "strconv"

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

func (board Board) MarkerAt(spot int) Marker {
	return board[spot]
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

func (board Board) RowLength() int {
	sqrt := math.Sqrt(float64(len(board)))
	return int(sqrt)
}

func (board Board) StringBoard() []string {
	var stringBoard []string
	for spot, marker := range board {
		switch marker {
		case X:
			stringBoard = append(stringBoard, "X")
		case O:
			stringBoard = append(stringBoard, "O")
		default:
			stringBoard = append(stringBoard, strconv.Itoa(spot))
		}
	}
	return stringBoard
}
