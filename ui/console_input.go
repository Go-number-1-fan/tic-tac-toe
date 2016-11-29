package ui

import "bufio"
import "os"
import "strconv"
import "fmt"
import "strings"

type ConsoleInput struct{}

func readAndTrimLine() (string, bool) {
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	line = strings.TrimSpace(line)
	hasError := err != nil
	return line, hasError
}

func convertStringToInt(stringNum string) (int, bool) {
	number, err := strconv.Atoi(stringNum)
	hasError := err != nil
	return number, hasError
}

func (input ConsoleInput) ReadInt() int {
	var number int
	var hasReadError, hasParseError bool
	var line string
	hasError := true
	for hasError {
		line, hasReadError = readAndTrimLine()
		number, hasParseError = convertStringToInt(line)
		hasError = hasReadError || hasParseError
		if hasError {
			fmt.Println("Please enter a valid integer: ")
		}
	}
	return number
}
