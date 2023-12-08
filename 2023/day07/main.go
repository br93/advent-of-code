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
	joker   []Hand
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
	s.hands, _ = s.getHands()

	s.answers[0] = s.evaluate(s.hands)
	//s.answers[1] = s.evaluate(s.joker)

}

func (*Solution) readFile(file string) string {
	fileBytes, err := os.ReadFile(fmt.Sprintf("%s.txt", file))

	if err != nil {
		log.Fatal(err)
	}

	return string(fileBytes)
}

func (s *Solution) getHands() ([]Hand, []Hand) {
	s.hands, s.joker = s.getHandFromInput()
	s.convertHand("part1")
	//s.convertHand("part2")

	hands := s.hands
	//joker := s.joker

	hands, _ = s.setPairValues(hands)
	hands = sortCardsByValue(hands)
	//joker = sortCardsByValue(joker)

	return hands, nil

}

func (s *Solution) convertHand(step string) {
	for index := range s.hands {
		if step == "part1" {
			s.hands[index].cards = convertCards(s.hands[index].cards, step)
		} else {
			s.joker[index].cards = convertCards(s.joker[index].cards, step)
		}

	}
}

func convertCards(cards, step string) string {
	s := strings.Split(cards, "")

	for index, element := range s {
		if step == "part1" {
			s[index] = strings.Map(rules, element)
		} else {
			s[index] = strings.Map(joker, element)
		}

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

func joker(r rune) rune {
	switch r {
	case 'B':
		return '0'
	default:
		return r
	}
}

func (s *Solution) getHandFromInput() ([]Hand, []Hand) {
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

	return hands, hands
}

func (s *Solution) setPairValues(hands []Hand) ([]Hand, []Hand) {

	//joker := s.joker

	for index := range hands {
		hands[index].value = pairStrenght(hands[index].cards)
		//joker[index].value = s.checkJoker(joker[index].cards, pairStrenght(joker[index].cards))
	}

	return hands, nil

}

func (s *Solution) checkJoker(cards string, value int) int {

	count := strings.Count(cards, "0")

	switch count {
	case 5:
		return count + 1
	case 4:
		return count + 2
	case 3:
		return value + count + 2
	case 2:
		return value + count + 1
	case 1:
		{
			if value == 5 || value == 0 {
				return value + count
			}

			return value + 2
		}
	default:
		return value

	}

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
		if char != '0' {
			counts[char]++
		}
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

func (s *Solution) evaluate(hands []Hand) int {
	var value int

	for index, element := range hands {
		value = value + (index+1)*element.bid
	}
	return value
}
