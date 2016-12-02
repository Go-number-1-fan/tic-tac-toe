package ui

import "fmt"

type MockOutput struct{}

func (output MockOutput) Write(str string) {}

func (output MockOutput) Clear() {
	fmt.Println("CLEARED")
}
