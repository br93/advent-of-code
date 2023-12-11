package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Solution struct {
	input   string
	grid    [][]rune
	answers []int
}

const (
	DAY     = 11
	EXAMPLE = "example"
	INPUT   = "input"
	PART1   = 2
	PART2   = 1_000_000
)

func main() {

	solution := &Solution{}
	solution.day11(EXAMPLE)
	solution.day11(INPUT)
}

func (s *Solution) day11(file string) {
	input := readFile(file)

	if input != "" {
		s.input = input
		s.grid = s.makeGrid()
		s.answers = s.solvePuzzle()

		s.print(file)

	} else {
		fileError(file)
	}

}

func (s *Solution) makeGrid() [][]rune {
	lines := strings.Split(strings.TrimSpace(s.input), "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	return grid
}

func (s *Solution) solvePuzzle() []int {

	solution := []int{s.calculate(PART1), s.calculate(PART2)}
	return solution
}

func (s *Solution) calculate(factor int) int {
	galaxies := s.findGalaxies()
	emptyRows := s.findEmptyRows()
	emptyColumns := s.findEmptyColumns()

	sum := 0
	for i := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			xMin, xMax := minMax(galaxies[i][0], galaxies[j][0])
			yMin, yMax := minMax(galaxies[i][1], galaxies[j][1])

			expandX := countEmpty(emptyRows, xMin, xMax) * (factor - 1)
			expandY := countEmpty(emptyColumns, yMin, yMax) * (factor - 1)

			distanceX := xMax + expandX - xMin
			distanceY := yMax + expandY - yMin

			sum = sum + distanceX + distanceY
		}
	}
	return sum
}

func (s *Solution) findGalaxies() [][]int {

	grid := s.grid
	var galaxies [][]int

	for i, row := range grid {
		for j, element := range row {
			if element == '#' {
				galaxies = append(galaxies, []int{i, j})
			}
		}
	}

	return galaxies
}

func (s *Solution) findEmptyRows() []int {
	var empty []int

	grid := s.grid
	for index, row := range grid {
		isEmpty := true
		for _, element := range row {
			if element == '#' {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			empty = append(empty, index)
		}
	}

	return empty
}

func (s *Solution) findEmptyColumns() []int {
	var empty []int

	grid := s.grid

	for i := 0; i < len(grid[0]); i++ {
		isEmpty := true
		for j := 0; j < len(grid); j++ {
			if grid[j][i] == '#' {
				isEmpty = false
				break
			}
		}

		if isEmpty {
			empty = append(empty, i)
		}
	}

	return empty
}

func minMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func countEmpty(arr []int, start, end int) int {
	var count int
	for _, x := range arr {
		if start <= x && x <= end {
			count++
		}
	}
	return count
}

func (s *Solution) print(file string) {
	fmt.Printf("Day %d - %s : (%d) - (%d)\n", DAY, file, s.answers[0], s.answers[1])
}

func readFile(file string) string {
	fileBytes, err := os.ReadFile(fmt.Sprintf("%s.txt", file))

	if err != nil {
		return ""
	}

	return string(fileBytes)
}

func fileError(file string) {
	var advent string

	if file == "input" {
		advent = fmt.Sprintf("%s %s %s", "check adventofcode.com for", file, "file")
	}

	error := fmt.Errorf("%s.txt not found %s", file, advent)
	log.Println(error)
}
