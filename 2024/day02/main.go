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
	input []string
	part1 int
	part2 int
}

const (
	DAY     = 02
	EXAMPLE = "example"
	INPUT   = "input"
)

func main() {
	solution := &Solution{}
	solution.day02(EXAMPLE)
	solution.day02(INPUT)
}

func (s *Solution) day02(file string) {
	s.input = getInput(file)
	s.part1 = s.safeReports()
	s.print(file)
}

func (s *Solution) safeReports() int {

	var safe int

	for _, element := range s.input {

		report := toSliceInt(strings.Split(element, " "))
		asc := toSliceInt(strings.Split(element, " "))
		desc := toSliceInt(strings.Split(element, " "))

		slices.Sort(asc)
		slices.Sort(desc)
		slices.Reverse(desc)

		if slices.Equal(report, asc) || slices.Equal(report, desc) {
			safe = safe + s.checkUnsafe(report)
		}
	}

	return safe
}

func (s *Solution) checkUnsafe(num []int) int {
	for index := range num {

		if index < len(num)-1 {
			diff := abs(num[index] - num[index+1])
			if diff > 3 || diff == 0 {
				return 0
			}
		}

	}
	return 1
}

func (s *Solution) print(file string) {
	fmt.Printf("Day %d - %s : (part 1: %d) (part 2: %d)\n", DAY, file, s.part1, s.part2)
}

func getInput(file string) []string {
	str := readFile(file)
	return strings.Split(str, "\n")
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

func toSliceInt(str []string) []int {
	var intSlice []int

	for _, element := range str {
		intSlice = append(intSlice, toInt(element))
	}

	return intSlice
}

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}
