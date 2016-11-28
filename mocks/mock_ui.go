package mocks

type MockUI struct {
	MockUserInput int
}

func (ui MockUI) GetValidMove(availableSpots []int) int {
	return ui.MockUserInput
}
