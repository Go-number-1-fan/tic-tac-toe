package mocks

type MockInput struct {
	MockConsoleInput int
}

func (input MockInput) ReadInt() int {
	return input.MockConsoleInput
}
