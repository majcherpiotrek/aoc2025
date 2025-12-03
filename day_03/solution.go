package day_03

import (
	"fmt"
	"strconv"
)

func Part1(input *[]string) (int, error) {
	banks, err := parseInput(input)

	if err != nil {
		return -1, err
	}

	sum := 0

	for _, bank := range banks {
		b := findLargestNumberInBank(2, &bank)
		sum += b
	}

	return sum, nil
}

func Part2(input *[]string) (int, error) {

	return -1, fmt.Errorf("not implementd")
}

func findLargestNumberInBank(numLength int, batteries *[]int) int {
	searchStart := 0
	num := make([]int, numLength)

	for digitPosition := range numLength {
		maxNum := -1
		maxNumIndex := -1
		digitsLeft := numLength - digitPosition
		lastPossibleIndexOfFirstDigit := len(*batteries) - digitsLeft

		for i := searchStart; i <= lastPossibleIndexOfFirstDigit; i++ {
			n := (*batteries)[i]

			if n > maxNum {
				maxNum = n
				maxNumIndex = i
			}
		}

		if maxNum == -1 || maxNumIndex == -1 {
			panic("No max number found")
		}

		num[digitPosition] = maxNum

		searchStart = maxNumIndex + 1
	}

	numStr := ""

	for _, n := range num {
		numStr = fmt.Sprintf("%s%d", numStr, n)
	}

	res, err := strconv.Atoi(numStr)

	if err != nil {
		panic(err)
	}

	return res
}

func parseInput(input *[]string) ([][]int, error) {
	res := make([][]int, len(*input))
	for i, line := range *input {
		resLine := make([]int, len(line))

		for j, digitStr := range line {
			digit, err := strconv.Atoi(string(digitStr))

			if err != nil {
				return [][]int{}, err
			}

			resLine[j] = digit
		}

		res[i] = resLine

	}
	return res, nil
}
