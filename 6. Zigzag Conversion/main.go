package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([]strings.Builder, numRows)

	var rowIdx int
	var direction int = 1
	for _, letter := range s {
		rows[rowIdx].WriteRune(letter)

		if rowIdx == numRows-1 {
			direction = -1
		}
		if rowIdx == 0 {
			direction = 1
		}
		rowIdx += direction
	}

	var rowStrings []string
	for _, row := range rows {
		rowStrings = append(rowStrings, row.String())
	}

	return strings.Join(rowStrings, "")
}
