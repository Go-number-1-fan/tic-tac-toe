package ui

import "fmt"

type ConsoleOutput struct{}

func (output ConsoleOutput) Write(str string) {
	fmt.Print(str)
}
