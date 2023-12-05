package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Solution struct {
	input    []string
	almanac  Almanac
	location int
}

type Almanac struct {
	seeds []int
	maps  [][]Map
}

type Map struct {
	dest int
	src  int
	len  int
}

func main() {
	solution := Solution{}
	solution.day5("example")

	//fmt.Println(solution.almanac.seeds)
	fmt.Printf("0: %d", solution.almanac.seeds)
	solution.loopMap()
	fmt.Printf("1: %d", solution.almanac.seeds)

}

func (s *Solution) loopMap() {

	maps := s.almanac.maps
	flag := true

	for index, seed := range s.almanac.seeds {
		for _, element1 := range maps {
			for _, element2 := range element1 {
				dest := element2.dest
				src := element2.src
				len := element2.len

				min := src
				max := src + len - 1

				if !(seed < min || seed > max) && flag {
					s.almanac.seeds[index] = seed - src + dest
					flag = false
				}
			}
			flag = true
		}
	}

}

func (a *Almanac) setSeeds(input []string) {

	for _, element := range strings.Fields(strings.Split(input[0], "seeds: ")[1]) {
		num, _ := strconv.Atoi(element)
		a.seeds = append(a.seeds, num)
	}
}

func (a *Almanac) setMaps(input []string) {
	input = input[1:]

	for _, element := range input {
		a.maps = append(a.maps, createMap(element))
	}

}

func createMap(value string) []Map {
	values := strings.Split(value, "\n")
	values = values[1:]

	maps := make([]Map, 0)

	for _, element := range values {
		if element != "" {
			fields := strings.Fields(element)
			maps = append(maps, Map{
				dest: intField(fields[0]),
				src:  intField(fields[1]),
				len:  intField(fields[2]),
			})
		}
	}

	return maps
}

func intField(field string) int {
	num, _ := strconv.Atoi(field)

	return num
}

func (s *Solution) day5(file string) {
	s.input = readFile(file)
	almanac := s.createAlmanac()
	s.setAlmanac(almanac)
}

func (s *Solution) createAlmanac() *Almanac {
	almanac := &Almanac{}
	almanac.setSeeds(s.input)
	almanac.setMaps(s.input)

	return almanac
}

func (s *Solution) setAlmanac(almanac *Almanac) {
	s.almanac = *almanac
}

func readFile(file string) []string {
	fileBytes, err := os.ReadFile(fmt.Sprintf("%s.txt", file))

	if err != nil {
		panic(err)
	}

	str := string(fileBytes)
	return strings.Split(str, "\n\n")
}
