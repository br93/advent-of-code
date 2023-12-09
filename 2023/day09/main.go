package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Solution struct {
	input   string
	report  []string
	answers []int
}

const Day = 9

func main() {
	solution := &Solution{}
	solution.day9("example")
	solution.day9("input")
}

func (s *Solution) day9(file string) {
	s.input = readFile(file)
	s.report = s.getReport()

	s.answers = make([]int, 2)
	for _, line := range s.report {
		part1, part2 := s.simulate(line)
		s.answers[0] += part1
		s.answers[1] += part2
	}

	s.print(file)
}

func (s *Solution) getReport() []string {
	input := s.input
	return strings.Split(input, "\n")
}

func (s *Solution) simulate(str string) (int, int) {
	slice := [][]int{toIntSlice(str)}

	for !isAllZeros(slice[len(slice)-1]) {
		aux := make([]int, len(slice[len(slice)-1])-1)
		for index := range aux {
			aux[index] = slice[len(slice)-1][index+1] - slice[len(slice)-1][index]
		}
		slice = append(slice, aux)
	}

	return s.extrapolation(slice)
}

func (s *Solution) extrapolation(slice [][]int) (int, int) {
	var part1, part2 int
	for i := len(slice) - 1; i >= 0; i-- {
		part1 += slice[i][len(slice[i])-1]
		part2 = slice[i][0] - part2
	}

	return part1, part2
}

func (s *Solution) print(f string) {
	fmt.Printf("Day %d - %s : (%d) - (%d)\n", Day, f, s.answers[0], s.answers[1])
}

func readFile(file string) string {
	fileBytes, err := os.ReadFile(fmt.Sprintf("%s.txt", file))

	if err != nil {
		log.Fatal(err)
	}

	return string(fileBytes)
}

func toInt(str string) int {
	num, err := strconv.Atoi(str)

	if err != nil {
		log.Fatal(err)
	}

	return num
}

func toIntSlice(str string) []int {
	fields := strings.Fields(str)
	result := make([]int, len(fields))

	for index, element := range fields {
		num := toInt(element)
		result[index] = num
	}
	return result
}

func isAllZeros(nums []int) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}
