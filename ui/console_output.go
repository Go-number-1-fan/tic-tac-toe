package ui

import "fmt"
import "runtime"
import "os"
import "os/exec"

type ConsoleOutput struct{}

func (output ConsoleOutput) Write(str string) {
	fmt.Print(str)
}

func (output ConsoleOutput) Clear() {
	operatingSys := runtime.GOOS
	switch operatingSys {
	case "windows":
		cmd := exec.Command("cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
