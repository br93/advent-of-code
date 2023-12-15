package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Solution struct {
	input   string
	answers []int
}

const (
	DAY     = 15
	EXAMPLE = "example"
	INPUT   = "input"
)

func main() {

	solution := &Solution{}
	solution.day15(EXAMPLE)
	solution.day15(INPUT)

}

func (s *Solution) day15(file string) {
	s.input = readFile(file)
	s.answers = s.solvePuzzle()
	s.print(file)
}

func (s *Solution) solvePuzzle() []int {

	solution := []int{s.calculate(), 0}
	return solution
}

func (s *Solution) calculate() int {
	sequence := strings.Split(s.input, ",")
	var values []int
	for _, element := range sequence {
		values = append(values, s.hash(element))
	}

	return sumArray(values...)

}

func sumArray(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum = sum + num
	}

	return sum
}

func (s *Solution) hash(str string) int {
	value := 0

	for _, char := range str {
		value = ((value + int(char)) * 17) % 256
	}

	return value
}

func (s *Solution) print(file string) {
	fmt.Printf("Day %d - %s : (%d) - (%d)\n", DAY, file, s.answers[0], s.answers[1])
}

func readFile(file string) string {
	fileBytes, err := os.ReadFile(fmt.Sprintf("%s.txt", file))
	readError(file, err)

	return string(fileBytes)
}

func readError(file string, err error) {
	if err != nil {
		if file == "input" {
			advent := fmt.Sprintf("%s %s %s", "check adventofcode.com for", file, "file")
			log.Fatal(fmt.Errorf("%s\n%s", err, advent))
		}

		log.Fatal(err)
	}
}
