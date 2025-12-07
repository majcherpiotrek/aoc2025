package day_05

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type IdsRange struct {
	Min int
	Max int
}

func parseInput(input *[]string) ([]IdsRange, []int, error) {
	parseRanges := true
	freshIds := make([]IdsRange, 0)
	ids := make([]int, 0)

	for i, row := range *input {
		if len(row) == 0 {
			parseRanges = false
			continue
		}

		if parseRanges {
			rangeStr := strings.Split(row, "-")

			if len(rangeStr) != 2 {
				return nil, []int{}, fmt.Errorf("invalid input row %d - range of ids expected, received %s", i, row)
			}

			rangeStart, err := strconv.Atoi(rangeStr[0])
			if err != nil {
				return nil, []int{}, err
			}
			rangeEnd, err := strconv.Atoi(rangeStr[1])
			if err != nil {
				return nil, []int{}, err
			}

			newRange := IdsRange{
				Min: rangeStart,
				Max: rangeEnd,
			}

			freshIds = append(freshIds, newRange)

			slices.SortFunc(freshIds, func(a IdsRange, b IdsRange) int {
				return a.Min - b.Min
			})
		} else {
			id, err := strconv.Atoi(row)
			if err != nil {
				return nil, []int{}, err
			}

			ids = append(ids, id)
		}
	}

	return freshIds, ids, nil
}

func Part1(input *[]string) (int, error) {
	freshIds, ids, err := parseInput(input)

	if err != nil {
		return -1, err
	}

	if len(freshIds) < 1 {
		return -1, fmt.Errorf("no fresh id ranges")
	}

	fresh := 0

	slices.Sort(ids)

	mergedIds := mergeRanges(&freshIds)

	for _, id := range ids {
		for _, idsRange := range mergedIds {
			if id >= idsRange.Min && id <= idsRange.Max {
				fresh++
				break
			}
		}
	}

	return fresh, nil
}

func Part2(input *[]string) (int, error) {
	freshIds, _, err := parseInput(input)

	if err != nil {
		return -1, err
	}

	if len(freshIds) < 1 {
		return -1, fmt.Errorf("no fresh id ranges")
	}

	mergedFreshIds := mergeRanges(&freshIds)
	numOfAllFreshIds := 0

	for _, mergedRange := range mergedFreshIds {
		numOfIdsInRange := mergedRange.Max - mergedRange.Min + 1
		numOfAllFreshIds += numOfIdsInRange
	}

	return numOfAllFreshIds, nil

}

func mergeRanges(ranges *[]IdsRange) []IdsRange {
	merged := make([]IdsRange, 0)

	for _, currentRange := range *ranges {
		var lastMergedRange *IdsRange

		if len(merged) > 0 {
			lastMergedRange = &merged[len(merged)-1]
		}

		if lastMergedRange != nil {
			if currentRange.Min > lastMergedRange.Max+1 {
				merged = append(merged, currentRange)
			} else if currentRange.Max > lastMergedRange.Max {
				lastMergedRange.Max = currentRange.Max
			}
		} else {
			merged = append(merged, currentRange)
		}
	}

	return merged
}
