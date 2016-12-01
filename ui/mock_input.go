package ui

type MockInput struct {
	MockConsoleIntInput    int
	MockConsoleStringInput string
}

func (input MockInput) ReadInt() int {
	return input.MockConsoleIntInput
}

func (input MockInput) ReadStringOfLengthWithDefault(length int, def string) string {
	return input.MockConsoleStringInput
}
