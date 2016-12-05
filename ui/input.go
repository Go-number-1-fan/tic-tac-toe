package ui

type Input interface {
	ReadInt(out Output) int
	ReadStringOfLengthWithDefault(length int, def string) string
}
