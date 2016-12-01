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
			fmt.Println(NotANumberMessage)
		}
	}
	return number
}

func (input ConsoleInput) ReadStringOfLengthWithDefault(length int, def string) string {
	var str string
	var hasReadError bool
	str, hasReadError = readAndTrimLine()
	hasError := hasReadError || len(str) == 0
	if hasError {
		return def
	}
	if len(str) > length {
		str = str[:length]
	}
	return str
}
