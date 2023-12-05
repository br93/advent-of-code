package main

import (
	"fmt"
	"math"
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

	solution.loopMap()
	solution.minimumLocation()

	fmt.Println(solution.location)
}

func (s *Solution) loopMap() {

	maps := s.almanac.maps
	for {
		for _, element := range maps[0] {
			for index, seed := range s.almanac.seeds {
				min := element.src
				max := element.src + element.len - 1

				if seed == s.almanac.seeds[index] && (seed >= min && seed <= max) {
					s.almanac.seeds[index] = seed - element.src + element.dest
				}
			}
		}

		copy(maps[0:], maps[0+1:])
		maps[len(maps)-1] = nil
		maps = maps[:len(maps)-1]

		if len(maps) == 0 {
			break
		}
	}

}

func (s *Solution) minimumLocation() {
	location := math.MaxInt

	for _, seed := range s.almanac.seeds {
		location = min(location, seed)
	}

	s.location = location
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
