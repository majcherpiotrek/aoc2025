package day_04

import "fmt"

type Coordinates struct {
	Row int
	Col int
}

type GridElement struct {
	Value      rune
	Coords     Coordinates
	Neighbours []Coordinates
}

func parseInput(input *[]string) [][]GridElement {
	parsedGrid := make([][]GridElement, len(*input))

	for i, row := range *input {
		prevRow := i - 1
		nextRow := i + 1

		parsedGrid[i] = make([]GridElement, len(row))

		for k, element := range row {

			gridElement := GridElement{
				Value: element,
				Coords: Coordinates{
					Row: i,
					Col: k,
				},
				Neighbours: make([]Coordinates, 0, 8),
			}

			if gridElement.Value != '@' {
				parsedGrid[i][k] = gridElement
				continue
			}

			prevColumn := k - 1
			nextColumn := k + 1

			// fmt.Printf("Searching for neighbours for row %d, col %d\n", i, k)

			for i2 := prevRow; i2 <= nextRow; i2++ {
				if i2 < 0 || i2 >= len(*input) {
					continue
				}
				row2 := (*input)[i2]
				for k2 := prevColumn; k2 <= nextColumn; k2++ {
					if k2 < 0 || k2 >= len(row2) {
						continue
					}

					if i2 == i && k2 == k {
						continue
					}

					if row2[k2] == '@' {
						gridElement.Neighbours = append(gridElement.Neighbours, Coordinates{
							Row: i2,
							Col: k2,
						})
					}
				}
			}

			parsedGrid[i][k] = gridElement

		}
	}

	return parsedGrid
}

func Part1(input *[]string) (int, error) {
	grid := parseInput(input)

	accessibleGridElements := 0

	for _, row := range grid {
		for _, element := range row {
			// fmt.Printf("len of neighbours for row %d col %d = %d\n", i, k, len(element.Neighbours))
			if element.Value == '@' && len(element.Neighbours) < 4 {
				accessibleGridElements++
			}
		}
	}

	return accessibleGridElements, nil
}

func Part2(input *[]string) (int, error) {

	return -1, fmt.Errorf("")
}
