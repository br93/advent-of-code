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
	step    string
	network map[string][]string
	answers []int
}

const Day = 9

func main() {
	solution := &Solution{}
	solution.day8("example1")
	solution.day8("example2")
	solution.day8("input")
}

func (s *Solution) day8(f string) {
	s.input = readFile(f)
	s.step, s.network = s.formatInput()
	s.answers = make([]int, 2)

	s.part1().part2().print(f)
}

func (s *Solution) print(f string) {
	fmt.Printf("Day %d - %s : (%d) - (%d)\n", Day, f, s.answers[0], s.answers[1])
}

func (s *Solution) formatInput() (string, map[string][]string) {
	input := s.input
	input = replacer(input)

	array := strings.Split(input, "\n\n")
	step := array[0]

	array = strings.Split(array[1], "\n")
	network := make(map[string][]string)

	for _, element := range array {
		aux := strings.Split(element, " ")

		var value []string
		value = append(value, aux[1], aux[2])
		network[aux[0]] = value
	}

	return step, network

}

func (s *Solution) part1() *Solution {
	s.answers[0] = s.steps("AAA")
	return s
}

func (s *Solution) part2() *Solution {

	network := s.network
	var positions []string
	var counts []int

	for index := range network {
		if index[2] == 'A' {
			positions = append(positions, index)
		}
	}

	for _, position := range positions {
		counts = append(counts, s.steps(position))
	}

	var answer int

	if len(counts) > 1 {
		answer = 1

		for i := 0; i < len(counts); i++ {
			answer = lcm(answer, counts[i])
		}
	} else {
		answer = counts[0]
	}

	s.answers[1] = answer

	return s
}

func (s *Solution) steps(pos string) int {
	var count int

	step := s.step
	network := s.network

	for pos[2] != 'Z' {
		count++

		if step[0] == 'L' {
			pos = network[pos][0]
		} else {
			pos = network[pos][1]
		}

		step = step[1:] + string(step[0])
	}

	return count
}

func replacer(str string) string {
	replacer := strings.NewReplacer(
		" = ", " ",
		"(", "",
		",", "",
		")", "")
	return replacer.Replace(str)
}

func readFile(f string) string {
	bytes, err := os.ReadFile(fmt.Sprintf("%s.txt", f))

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

func gcd(a, b int) int {

	for b != 0 {
		aux := a
		a = b
		b = aux % b
	}

	return a

}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
