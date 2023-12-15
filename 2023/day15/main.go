package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
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
	//solution.day15(INPUT)

}

func (s *Solution) day15(file string) {
	s.input = readFile(file)
	s.answers = s.solvePuzzle()
	s.print(file)
}

func (s *Solution) solvePuzzle() []int {

	part1, part2 := s.calculate()
	solution := []int{part1, part2}
	return solution
}

func (s *Solution) calculate() (int, int) {
	sequence := strings.Split(s.input, ",")
	var values []int

	boxes := make([][]string, 256)
	lens := make(map[string]int, 256)

	for _, element := range sequence {
		values = append(values, s.hash(element))
		if strings.Contains(element, "-") {
			label := element[:len(element)-1]
			index := s.hash(label)
			if slices.Contains(boxes[index], label) {
				slices.Delete(boxes[index], slices.Index(boxes[index], label), slices.Index(boxes[index], label))
			}

		} else {
			equals := strings.Index(element, "=")
			label := element[:equals]
			index := s.hash(label)
			if !slices.Contains(boxes[index], label) {
				boxes[index] = append(boxes[index], label)
			}

			lens[label] = toInt(element[equals+1:])
		}
	}

	power := 0

	for indexBox, box := range boxes {
		for indexLabel, label := range box {
			power = power + (indexBox+1)*(indexLabel+1)*lens[label]
		}
	}
	return sumArray(values...), power

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

func toInt(str string) int {
	num, err := strconv.Atoi(str)

	if err != nil {
		log.Fatal(err)
	}

	return num
}
