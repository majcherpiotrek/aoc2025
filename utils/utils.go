package utils

import "strings"

func Abs(i int) int {
	if i < 0 {
		return -1 * i
	}

	return i
}

func PadCellValue(str string, size int) string {
	space := size - len(str)
	leftPad := space / 2
	rightPad := max(size-leftPad-len(str), 0)

	return strings.Repeat(" ", leftPad) + str + strings.Repeat(" ", rightPad)
}
