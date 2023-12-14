package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Solution struct {
	input   string
	answers []int
}

const (
	DAY     = 14
	EXAMPLE = "example"
	INPUT   = "input"
)

func main() {

	solution := &Solution{}
	solution.day14(EXAMPLE)
	solution.day14(INPUT)

}

func (s *Solution) day14(file string) {
	s.input = readFile(file)
	s.answers = s.solvePuzzle()
	s.print(file)
}

func (s *Solution) solvePuzzle() []int {

	solution := []int{s.calculate(), 0}
	return solution
}

func (s *Solution) calculate() int {
	grid := strings.Split(s.input, "\n")
	transposed := transpose(grid)

	for index, row := range transposed {
		rocks := strings.Split(row, "#")

		for index, rock := range rocks {
			runes := []rune(rock)

			sort.Slice(runes, func(a, b int) bool {
				return runes[a] > runes[b]
			})

			rocks[index] = string(runes)
		}

		transposed[index] = strings.Join(rocks, "#")
	}

	grid = transpose(transposed)

	count := 0
	for index, row := range grid {
		count = count + strings.Count(row, "O")*(len(grid)-index)
	}

	return count

}

func (s *Solution) print(file string) {
	fmt.Printf("Day %d - %s : (%d) - (%d)\n", DAY, file, s.answers[0], s.answers[1])
}

func readFile(file string) string {
	fileBytes, err := os.ReadFile(fmt.Sprintf("%s.txt", file))

	if err != nil {
		if file == "input" {
			advent := fmt.Sprintf("%s %s %s", "check adventofcode.com for", file, "file")
			log.Fatal(fmt.Errorf("%s\n%s", err, advent))
		}

		log.Fatal(err)
	}

	return string(fileBytes)
}

func transpose(grid []string) []string {
	var transposed []string

	for i := 0; i < len(grid[0]); i++ {
		var col []byte
		for _, row := range grid {
			col = append(col, row[i])
		}

		transposed = append(transposed, string(col))
	}

	return transposed
}
