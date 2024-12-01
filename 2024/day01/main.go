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
	input string
	part1 int
	part2 int
}

const (
	DAY     = 01
	EXAMPLE = "example"
	INPUT   = "input"
)

func main() {
	solution := &Solution{}
	solution.day01(EXAMPLE)
	solution.day01(INPUT)
}

func (s *Solution) day01(file string) {
	s.input = readFile(file)
	s.part1 = s.distance()
	s.part2 = s.similarity()
	s.print(file)
}

func (s *Solution) distance() int {

	var firstList []int
	var secondList []int

	var array []string = strings.Split(s.input, "\n")

	for _, element := range array {
		var content = strings.Split(element, "   ")
		firstList = append(firstList, toInt(content[0]))
		secondList = append(secondList, toInt(content[1]))
	}

	slices.Sort(firstList)
	slices.Sort(secondList)

	var distance int

	for index := range firstList {
		greater, smaller := max(firstList[index], secondList[index]), min(firstList[index], secondList[index])
		distance = distance + (greater - smaller)
	}

	return distance
}

func (s *Solution) similarity() int {
	var firstList []string
	var secondList []string

	var array []string = strings.Split(s.input, "\n")

	for _, element := range array {
		var content = strings.Split(element, "   ")
		firstList = append(firstList, content[0])
		secondList = append(secondList, content[1])
	}

	var similarity int

	for _, element := range firstList {
		similarity = similarity + toInt(element)*count(secondList, func(x string) bool {
			return strings.Contains(x, string(element))
		})
	}

	return similarity

}

func (s *Solution) print(file string) {
	fmt.Printf("Day %d - %s : (part 1: %d) (part 2: %d)\n", DAY, file, s.part1, s.part2)
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

func count[T any](slice []T, f func(T) bool) int {
	count := 0
	for _, s := range slice {
		if f(s) {
			count++
		}
	}
	return count
}
