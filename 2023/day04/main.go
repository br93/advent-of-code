package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
)

type Solution struct {
	cards Cards
}

type Cards struct {
	owned   []string
	winning []string
	values  []int
	copies  []int
}

type Regex struct {
	number *regexp.Regexp
}

func main() {

	solution := &Solution{}
	solution.scratchCards("4", "input")
}

func (s *Solution) getCards(file string) {
	fileBytes, err := os.ReadFile(fmt.Sprintf("%s.txt", file))

	if err != nil {
		panic(err)
	}

	str := string(fileBytes)
	str = strings.ReplaceAll(str, "  ", " 0")
	cards := strings.Split(str, "\n")

	regex := &Regex{
		number: regexp.MustCompile("[0-9]+"),
	}

	for _, element := range cards {
		wn := strings.Split(element, "|")
		owned := strings.Split(wn[0], ":")
		card := regex.number.FindAllString(owned[1], -1)
		s.cards.owned = append(s.cards.owned, strings.Join(card, " "))
		s.cards.winning = append(s.cards.winning, wn[1])
	}
}

func (s *Solution) points() float64 {
	var value float64
	owned := s.cards.owned
	winning := s.cards.winning

	for index, element := range owned {
		cards := strings.Split(element, " ")
		points := s.comparingCards(cards, winning[index], index)
		value = value + points
	}

	return value
}

func (s *Solution) comparingCards(cards []string, wn string, index int) float64 {
	var matches float64

	for _, element := range cards {
		if strings.Contains(wn, element) {
			matches = matches + 1
		}
	}

	if matches > 0 {
		s.cards.values[index] = int(matches)
		return math.Pow(2, matches-1)
	}

	return 0
}

func (s *Solution) printSolution(day string, points float64, copies int) {
	fmt.Printf("Day %s : points: %.0f, copies: %d\n", day, points, copies)
}

func (s *Solution) scratchCards(day, file string) {
	s.getCards(file)
	s.firstHand()
	points := s.points()
	copies := s.processingCopies()
	s.printSolution(day, points, copies)
}

func (s *Solution) firstHand() {
	for range s.cards.owned {
		s.cards.values = append(s.cards.values, 0)
		s.cards.copies = append(s.cards.copies, 1)
	}
}

func (s *Solution) processingCopies() int {

	for index, value := range s.cards.values {
		for copy := 0; copy < s.cards.copies[index]; copy++ {
			for add := 0; add < value; add++ {
				s.cards.copies[index+add+1] = s.cards.copies[index+add+1] + 1
			}
		}
	}

	return s.sumCopies()
}

func (s *Solution) sumCopies() int {
	var sum int
	for _, element := range s.cards.copies {
		sum = sum + element
	}

	return sum
}
