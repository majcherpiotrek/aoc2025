package day_01

import (
	"aoc2025/utils"
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

		nextPosition, _ := calculateNextPosition(current, rotationsNum)

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

func calculateNextPosition(currentPosition int, rotations int) (int, int) {
	if rotations == 0 {
		return currentPosition, 0
	}

	crossZero := utils.Abs(rotations / 100)
	rest := rotations % 100
	end := currentPosition + rest

	if rotations > 0 && end >= 100 {
		crossZero += 1
	}

	if rotations < 0 && currentPosition > 0 && end <= 0 {
		crossZero += 1
	}

	if end < 0 {
		return 100 + (end % 100), crossZero
	}

	return end % 100, crossZero
}

func Part2(input *[]string) (int, error) {
	current := 50
	stopsAt0 := 0

	fmt.Printf("|---------|---------|---------|---------|\n")
	fmt.Printf("|   pos   |   cli   |   res   |  cross  |\n")
	fmt.Printf("|---------|---------|---------|---------|\n")

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

		nextPosition, cross0 := calculateNextPosition(current, rotationsNum)

		stopsAt0 += cross0

		pos := strconv.Itoa(current)
		cli := strconv.Itoa(rotationsNum)
		res := strconv.Itoa(nextPosition)
		cross := strconv.Itoa(cross0)

		fmt.Printf("|%s|%s|%s|%s|\n", padCellValue(pos, 9), padCellValue(cli, 9), padCellValue(res, 9), padCellValue(cross, 9))
		fmt.Printf("|---------|---------|---------|---------|\n")

		current = nextPosition
	}

	return stopsAt0, nil
}
