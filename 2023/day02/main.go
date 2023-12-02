package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Solution struct {
	values []string
	cubes  []string
}

type Cubes struct {
	id    int
	blue  int
	green int
	red   int
}

func main() {
	CubeConundrum("example")
	CubeConundrum("input")
}

func CubeConundrum(file string) {
	exampleBytes, _ := os.ReadFile(fmt.Sprintf("%s.txt", file))
	example := strings.Split(string(exampleBytes), "\n")

	var result int

	for _, element := range example {

		solution := &Solution{}
		batch := strings.FieldsFunc(element, Split)
		str := fmt.Sprintf("%s", batch)

		reNumbers := regexp.MustCompile("[0-9]+")
		reString := regexp.MustCompile("[a-zA-Z]+")

		solution.cubes = reString.FindAllString(str, -1)
		solution.values = reNumbers.FindAllString(str, -1)

		cubes := CubeCalculation(solution)
		if isValid(cubes, &Cubes{id: 0, red: 12, green: 13, blue: 14}) {
			result = result + cubes.id
		}
	}

	fmt.Printf("%s: %d", file, result)
}

func CubeCalculation(solution *Solution) *Cubes {
	cubes := &Cubes{}

	for index, element := range solution.cubes {
		cube := element
		value, _ := strconv.Atoi(solution.values[index])

		switch cube {
		case "Game":
			cubes.id = value
		case "blue":
			cubes.blue = max(cubes.blue, value)
		case "green":
			cubes.green = max(cubes.green, value)
		case "red":
			cubes.red = max(cubes.red, value)
		}
	}

	return cubes

}

func isValid(cubes *Cubes, valid *Cubes) bool {
	if cubes.blue <= valid.blue && cubes.green <= valid.green && cubes.red <= valid.red {
		return true
	}

	return false
}

func Split(r rune) bool {
	return r == ';' || r == ',' || r == ':'
}

func SplitPart2(r rune) bool {
	return r == ';' || r == ':'
}
