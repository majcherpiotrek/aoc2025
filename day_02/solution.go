package day_02

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type IntSet map[int]struct{}

func (is IntSet) Add(i int) {
	is[i] = struct{}{}
}

func (is IntSet) AddAll(ints []int) {
	for _, i := range ints {
		is.Add(i)
	}
}

func (is IntSet) Merge(otherSet IntSet) {
	for k := range otherSet {
		is.Add(k)
	}
}

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
	ranges, err := parseRanges(input)

	if err != nil {
		return -1, err
	}

	invalidIds := make(IntSet)

	for _, r := range ranges {
		fmt.Printf("Range %d - %d\n", r.Min, r.Max)
		minStr := strconv.Itoa(r.Min)
		lenMin := len(minStr)
		maxStr := strconv.Itoa(r.Max)
		lenMax := len(maxStr)

		idsForRange := make(IntSet)

		if lenMin == lenMax {
			idsForRange.AddAll(generateAllPatterns(lenMin, r.Min, r.Max))
		} else {
			diff := lenMax - lenMin

			for i := 0; i <= diff; i++ {
				fmt.Printf("for diff %d/%d\n", i, diff)
				l := lenMin + i
				partialMin := max(r.Min, int(math.Pow10(l-1)))
				partialMax := min(r.Max, int(math.Pow10(l)-1))

				idsForRange.AddAll(generateAllPatterns(l, partialMin, partialMax))
			}

		}

		fmt.Printf("IDs for range %d - %d:\n", r.Min, r.Max)
		for _, id := range idsForRange {
			fmt.Printf("%d\n", id)
		}

		invalidIds.Merge(idsForRange)
		fmt.Printf("\n")
	}

	sum := 0

	fmt.Printf("Invalid ids\n")
	for id := range invalidIds {
		fmt.Printf("%d\n", id)
		sum += id
	}

	return sum, nil
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

func getAllPossiblePatternSizesForLength(length int) []int {
	if length < 2 {
		return []int{}
	}

	if length == 2 {
		return []int{1}
	}

	patternSizes := []int{1}

	patternSize := 2

	for patternSize < length {
		if length%patternSize == 0 {
			patternSizes = append(patternSizes, patternSize)
		}
		patternSize++
	}

	return patternSizes
}

func generateAllPatterns(length int, min int, max int) []int {
	fmt.Printf("generateAllPatterns(length = %d, min = %d, max = %d)\n", length, min, max)
	possiblePatternSizes := getAllPossiblePatternSizesForLength(length)
	patterns := []int{}

	fmt.Printf("possiblePatternSizes=%v\n", possiblePatternSizes)

	for _, size := range possiblePatternSizes {
		numOfPatternsInLength := length / size

		patternMax := int(math.Pow10(size)) - 1
		patternMin := int(math.Pow10(size - 1))
		pattern := patternMin

		for pattern <= patternMax {
			idStr := strings.Repeat(strconv.Itoa(pattern), numOfPatternsInLength)
			id, err := strconv.Atoi(idStr)

			if err != nil {
				fmt.Printf("this should not happen, invalid num from pattern %s\n", idStr)
				continue
			}

			if id > max {
				break
			}

			if id >= min {
				patterns = append(patterns, id)
			}

			pattern++
		}

	}

	return patterns
}
