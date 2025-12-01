package day_01

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(input *[]string) (int, error) {
	current := 50
	stopsAt0 := 0

	fmt.Printf("_______________________________\n")
	fmt.Printf("|   pos   |   cli   |   res   |\n")
	fmt.Printf("|---------|---------|---------|\n")

	for i, line := range *input {
		direction := line[:1]
		rotations := line[1:]

		rotationsNum, err := strconv.Atoi(rotations)

		if err != nil {
			return -1, fmt.Errorf("Invalid input at line %d", i)
		}

		if direction == "L" {
			rotationsNum = -1 * rotationsNum
		}

		nextPosition := calculateNextPosition(current, rotationsNum)

		if nextPosition == 0 {
			stopsAt0++
		}

		pos := strconv.Itoa(current)
		cli := strconv.Itoa(rotationsNum)
		res := strconv.Itoa(nextPosition)

		fmt.Printf("|%s|%s|%s|\n", padCellValue(pos, 9), padCellValue(cli, 9), padCellValue(res, 9))
		fmt.Printf("|---------|---------|---------|\n")

		current = nextPosition
	}

	return stopsAt0, nil
}

func padCellValue(str string, size int) string {
	space := size - len(str)
	leftPad := space / 2
	rightPad := max(size-leftPad-len(str), 0)

	return strings.Repeat(" ", leftPad) + str + strings.Repeat(" ", rightPad)
}

func calculateNextPosition(currentPosition int, rotations int) int {
	partial := currentPosition + rotations
	res := partial % 100

	if res < 0 {
		return 100 + res
	}

	return res
}

func Part2(input *[]string) (int, error) {

	return -1, fmt.Errorf("not implemented")
}
