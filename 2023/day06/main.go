package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Race struct {
	time     int
	distance int
}

type Solution struct {
	input   string
	race    []Race
	answers []int
}

const Day = 5

func main() {
	solution := Solution{}
	solution.day5("input")
	solution.print()
}

func (s *Solution) day5(file string) {
	s.answers = make([]int, 0)

	s.setRaceArray("input", "1")
	s.boatRace()

	s.setRaceArray("input", "2")
	s.boatRace()
}

func (s *Solution) setRaceArray(file, part string) {
	s.input = readFile(fmt.Sprintf("%s.txt", file))
	if part == "2" {
		s.input = strings.ReplaceAll(s.input, " ", "")
	}

	array := strings.Split(s.input, "\n")
	reNumbers := regexp.MustCompile("[0-9]+")

	time := reNumbers.FindAllString(array[0], -1)
	distance := reNumbers.FindAllString(array[1], -1)

	race := make([]Race, 0)

	for index := range time {
		race = append(race, Race{
			time:     convertToInt(time[index]),
			distance: convertToInt(distance[index])})
	}

	s.race = race
}

func (s *Solution) boatRace() {

	races := s.race

	var count int
	wins := make([]int, 0)

	for _, element := range races {
		count = 0

		for speed := 0; speed < element.time; speed++ {
			distance := speed * (element.time - speed)
			if distance > element.distance {
				count++
			}

		}
		wins = append(wins, count)
	}

	s.answers = append(s.answers, combination(wins))
}

func (s *Solution) print() {
	fmt.Printf("Day %d : answers: %d and %d\n", Day, s.answers[0], s.answers[1])
}

func combination(wins []int) int {
	answer := 1

	for _, element := range wins {
		answer *= element
	}

	return answer
}

func readFile(file string) string {
	fileBytes, err := os.ReadFile(file)

	if err != nil {
		log.Panic(err)
	}

	return string(fileBytes)
}

func convertToInt(str string) int {
	num, err := strconv.Atoi(str)

	if err != nil {
		log.Panic(err)
	}

	return num
}
