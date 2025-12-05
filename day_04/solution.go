package day_04

import "fmt"

type Coordinates struct {
	Row int
	Col int
}

type GridElement struct {
	Value           rune
	Coords          Coordinates
	NumOfNeighbours int
}

func printGrid(grid *[][]GridElement) {
	fmt.Printf("\n")
	for _, row := range *grid {
		rowToPrint := ""
		for _, element := range row {
			rowToPrint = fmt.Sprintf("%s%s", rowToPrint, string(element.Value))
		}

		fmt.Printf("%s\n", rowToPrint)
	}
	fmt.Printf("\n")
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
				NumOfNeighbours: 0,
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
						gridElement.NumOfNeighbours += 1
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
			if element.Value == '@' && element.NumOfNeighbours < 4 {
				accessibleGridElements++
			}
		}
	}

	return accessibleGridElements, nil
}

func Part2(input *[]string) (int, error) {
	grid := parseInput(input)

	// printGrid(&grid)

	step := 0

	rollsRemoved := 0

	for true {
		step++

		accessibleElements := make([]Coordinates, 0)
		for _, row := range grid {
			for _, element := range row {
				if element.Value == '@' && element.NumOfNeighbours < 4 {
					accessibleElements = append(accessibleElements, element.Coords)
					// fmt.Printf("Accessible element in %d,%d - %s, neighbours - %d\n", element.Coords.Row, element.Coords.Col, string(element.Value), element.NumOfNeighbours)
				} else {

					// fmt.Printf("NOT accessible element in %d,%d - %s, neighbours - %d\n", element.Coords.Row, element.Coords.Col, string(element.Value), element.NumOfNeighbours)
				}
			}
		}

		if len(accessibleElements) == 0 {
			break
		}

		// fmt.Printf("Found %d accessible elements in step %d\n", len(accessibleElements), step)

		for _, coords := range accessibleElements {
			grid[coords.Row][coords.Col].Value = 'x'
			rollsRemoved++

			// fmt.Printf("Removing at %d,%d\n", coords.Row, coords.Col)
			for i := coords.Row - 1; i <= coords.Row+1; i++ {
				if i < 0 || i >= len(grid) {
					continue
				}
				for k := coords.Col - 1; k <= coords.Col+1; k++ {
					if k < 0 || k >= len(grid[i]) {
						continue
					}

					if i == coords.Row && k == coords.Col {
						continue
					}

					// fmt.Printf("Checking neighbour %d,%d with value %s, neighbours %d\n", i, k, string(grid[i][k].Value), grid[i][k].NumOfNeighbours)

					if grid[i][k].Value == '@' {
						grid[i][k].NumOfNeighbours -= 1
					}

					// fmt.Printf("Processed neighbour %d,%d with value %s, neighbours %d\n\n", i, k, string(grid[i][k].Value), grid[i][k].NumOfNeighbours)
				}
			}
		}

		// printGrid(&grid)
	}

	return rollsRemoved, nil
}
