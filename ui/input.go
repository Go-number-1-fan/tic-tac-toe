package ui

type Input interface {
	ReadInt() int
	ReadStringOfLengthWithDefault(length int, def string) string
}
