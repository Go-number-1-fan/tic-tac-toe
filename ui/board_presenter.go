package ui

import "math"

func getCharsPerRow(rowLength int) int {
	charsPerRow := (rowLength * 3) + (rowLength + 1)
	return charsPerRow
}

func getHorizontalDividerString(charsPerRow int) string {
	boardString := " "
	for j := 0; j < charsPerRow; j++ {
		boardString = boardString + HorizontalDividerString
	}
	return boardString
}

func getBoardRowString(stringBoard []string, boardRowStartIndex int, rowLength int) string {
	boardRow := ""
	for boardColumn := 0; boardColumn < rowLength; boardColumn++ {
		boardRow = boardRow + VerticalDividerString + stringBoard[boardRowStartIndex+boardColumn]
	}
	boardRow = boardRow + VerticalDividerString
	return boardRow
}

func getDisplayBoard(stringBoard []string) string {
	rowLength := boardRowLength(stringBoard)
	charsPerRow := getCharsPerRow(rowLength)
	boardString := NewLineString
	for rowStartIndex := 0; rowStartIndex < len(stringBoard); {
		boardString = boardString + getHorizontalDividerString(charsPerRow) + NewLineString
		boardString = boardString + getBoardRowString(stringBoard, rowStartIndex, rowLength) + NewLineString
		rowStartIndex = rowStartIndex + rowLength
	}
	boardString = boardString + getHorizontalDividerString(charsPerRow) + NewLineString + NewLineString
	return boardString
}

func boardRowLength(stringBoard []string) int {
	sqrt := math.Sqrt(float64(len(stringBoard)))
	return int(sqrt)
}
