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

	minId := freshIds[0].Min
	maxId := freshIds[len(freshIds)-1].Max

	fresh := 0

	loops := 0

	slices.Sort(ids)

	for _, id := range ids {
		if id < minId && id > maxId {
			continue
		}

		for _, idsRange := range freshIds {
			loops++
			if id >= idsRange.Min && id <= idsRange.Max {
				fresh++
				break
			}
		}
	}

	return fresh, nil
}

func Part2(input *[]string) (int, error) {

	return -1, fmt.Errorf("")
}
