package UI

type UI interface {
	GetValidMove(availableSpots []int) int
}
