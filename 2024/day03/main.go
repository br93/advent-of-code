package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Solution struct {
	input string
	part1 int
	part2 int
}

const (
	DAY     = 03
	EXAMPLE = "example"
	INPUT   = "input"
)

func main() {
	solution := &Solution{}
	solution.day03(EXAMPLE)
	solution.day03(INPUT)
}

func (s *Solution) print(file string) {
	fmt.Printf("Day %d - %s : (part 1: %d) (part 2: %d)\n", DAY, file, s.part1, s.part2)
}

func (s *Solution) day03(file string) {
	s.input = readFile(file)
	s.part1 = s.multiplyInstructions(s.getValidInput())
	s.part2 = s.getActiveValidInput()
	s.print(file)
}

func (s *Solution) getActiveValidInput() int {
	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	dontRegex := regexp.MustCompile(`don't\(\)`)
	doRegex := regexp.MustCompile(`do\(\)`)

	segments := regex.Split(s.input, -1)
	matches := regex.FindAllStringSubmatch(s.input, -1)

	ignoreCommand := false
	var sum int

	for index, element := range matches {
		if dontRegex.MatchString(segments[index]) {
			ignoreCommand = true
		}

		if doRegex.MatchString(segments[index]) {
			ignoreCommand = false
		}

		if !ignoreCommand {
			sum = sum + (toInt(element[1]) * toInt(element[2]))
		}
	}

	return sum
}

func (s *Solution) getValidInput() [][]string {
	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	return regex.FindAllStringSubmatch(s.input, -1)
}

func (s *Solution) multiplyInstructions(section [][]string) int {
	var sum int

	for _, element := range section {
		sum = sum + (toInt(element[1]) * toInt(element[2]))
	}

	return sum
}

func readFile(file string) string {
	bytes, err := os.ReadFile(fmt.Sprintf("%s.txt", file))

	if err != nil {
		log.Fatal(err)
	}

	return string(bytes)
}

func toInt(str string) int {
	num, err := strconv.Atoi(str)

	if err != nil {
		log.Fatal(err)
	}

	return num
}
