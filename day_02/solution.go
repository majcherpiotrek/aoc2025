package day_02

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Part1(input *[]string) (int, error) {

	ranges, err := parseRanges(input)

	if err != nil {
		return -1, err
	}

	var invalidIds []int

	for _, r := range ranges {
		fmt.Printf("Range %d - %d\n", r.Min, r.Max)
		minStr := strconv.Itoa(r.Min)
		lenMin := len(minStr)
		maxStr := strconv.Itoa(r.Max)
		lenMax := len(maxStr)

		if lenMin%2 != 0 && lenMax%2 != 0 {
			continue
		}

		current := r.Min

		if lenMin%2 != 0 {
			current = int(math.Pow10(lenMax - 1))
		}

		actualMax := r.Max

		if lenMax%2 != 0 {
			actualMax = int(math.Pow10(lenMax-1)) - 1
		}

		fmt.Printf("Actual range %d - %d\n", current, actualMax)

		for current <= actualMax {
			a, b, err := splitToPair(current)

			if err != nil {
				current++
				continue
			}

			if a == b {
				invalidIds = append(invalidIds, current)
			}

			current++
		}
	}

	sum := 0

	for _, id := range invalidIds {
		sum += id
	}

	return sum, nil
}

func Part2(input *[]string) (int, error) {

	return -1, fmt.Errorf("not implemented yet")
}

type Range struct {
	Min int
	Max int
}

func parseRanges(input *[]string) ([]Range, error) {
	if len(*input) != 1 {
		return []Range{}, fmt.Errorf("Exactly one row expected")
	}

	line := (*input)[0]

	rangesStr := strings.Split(line, ",")
	ranges := make([]Range, len(rangesStr))

	for i, rangeStr := range rangesStr {

		minMaxStr := strings.Split(rangeStr, "-")

		if len(minMaxStr) != 2 {
			return []Range{}, fmt.Errorf("Invalid range at %d", i)
		}

		min, err := strconv.Atoi(minMaxStr[0])

		if err != nil {
			return []Range{}, err
		}

		max, err := strconv.Atoi(minMaxStr[1])

		if err != nil {
			return []Range{}, err
		}

		ranges[i] = Range{
			Min: min,
			Max: max,
		}
	}

	return ranges, nil
}

func splitToPair(num int) (int, int, error) {
	numStr := strconv.Itoa(num)
	l := len(numStr)

	if l == 0 || l%2 != 0 {
		return -1, -1, fmt.Errorf("Not splittable")
	}

	mid := l / 2

	a := numStr[:mid]
	b := numStr[mid:]

	aNum, err := strconv.Atoi(a)

	if err != nil {
		return -1, -1, err
	}

	bNum, err := strconv.Atoi(b)
	if err != nil {
		return -1, -1, err
	}

	return aNum, bNum, nil
}
