package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Solution struct {
	input   string
	hands   []Hand
	answers []int
}

type Hand struct {
	cards string
	bid   int
	value int
}

const Day = 7

func main() {
	solution := &Solution{}
	solution.day07("input")
	solution.print()
}

func (s *Solution) day07(file string) {
	s.answers = make([]int, 2)

	s.input = s.readFile(file)
	s.hands = s.getHands()

	s.answers[0] = s.evaluate()
}

func (*Solution) readFile(file string) string {
	fileBytes, err := os.ReadFile(fmt.Sprintf("%s.txt", file))

	if err != nil {
		log.Fatal(err)
	}

	return string(fileBytes)
}

func (s *Solution) getHands() []Hand {
	s.hands = s.getHandFromInput()
	s.convertHand()

	hands := s.hands

	hands = s.setPairValues(hands)
	hands = sortCardsByValue(hands)

	return hands

}

func (s *Solution) convertHand() {
	for index := range s.hands {
		s.hands[index].cards = convertCards(s.hands[index].cards)
	}
}

func convertCards(cards string) string {
	s := strings.Split(cards, "")

	for index, element := range s {
		s[index] = strings.Map(rules, element)
	}

	return strings.Join(s, "")
}

func rules(r rune) rune {
	switch r {
	case 'T':
		return 'A'
	case 'J':
		return 'B'
	case 'Q':
		return 'C'
	case 'K':
		return 'D'
	case 'A':
		return 'E'
	default:
		return r
	}
}

func (s *Solution) getHandFromInput() []Hand {
	input := strings.Split(s.input, "\n")

	var cards string
	var bid int

	var hands []Hand

	for _, element := range input {
		array := strings.Split(element, " ")
		cards = array[0]
		bid = toInt(array[1])

		hand := Hand{
			cards: cards,
			bid:   bid,
		}

		hands = append(hands, hand)
	}

	return hands
}

func (s *Solution) setPairValues(hands []Hand) []Hand {

	for index := range hands {
		hands[index].value = pairStrenght(hands[index].cards)
	}

	return hands

}

func sortCardsByValue(hands []Hand) []Hand {

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].value == hands[j].value {
			return hands[i].cards < hands[j].cards
		}
		return hands[i].value < hands[j].value
	})

	return hands
}

func (s *Solution) print() {
	fmt.Printf("DAY %d : (%d)  -  (%d)", Day, s.answers[0], s.answers[1])
}

func toInt(str string) int {
	num, err := strconv.Atoi(str)

	if err != nil {
		log.Fatal(err)
	}

	return num
}

func pairStrenght(card string) int {

	counts := make(map[rune]int)
	var values []int

	for _, char := range card {
		counts[char]++
	}

	for _, count := range counts {
		values = append(values, count)
	}

	if slices.Contains(values, 5) {
		return 6
	}

	if slices.Contains(values, 4) {
		return 5
	}

	if slices.Contains(values, 3) {
		if slices.Contains(values, 2) {
			return 4
		}
		return 3
	}

	if slices.Contains(values, 2) {
		index := slices.Index(values, 2)
		slices.Replace(values, index, index+1, 10)
		if slices.Contains(values, 2) {
			return 2
		}
		return 1
	}

	return 0
}

func (s *Solution) evaluate() int {
	hands := s.hands
	var value int

	for index, element := range hands {
		value = value + (index+1)*element.bid
	}
	return value
}
