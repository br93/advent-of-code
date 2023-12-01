package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	fileBytes, _ := os.ReadFile("input.txt")
	examples := strings.Split(string(fileBytes), "\n")

	keyBytes, _ := os.ReadFile("key.txt")
	keys := strings.Split(string(keyBytes), "\n")

	m := make(map[string]string)
	for index, element := range keys {
		m[element] = fmt.Sprintf("%d", index+1)
	}

	var sum int

	for _, element := range examples {

		for _, key := range keys {
			for strings.Contains(element, key) {
				str0 := fmt.Sprintf("%c", key[0])
				strF := fmt.Sprintf("%c", key[len(key)-1])
				element = strings.Replace(element, key, str0+m[key]+strF, -1)
			}
		}

		re := regexp.MustCompile("[0-9]")
		strings := re.FindAllString(element, -1)
		value := strings[0] + strings[len(strings)-1]

		num, _ := strconv.Atoi(value)
		sum = sum + num
	}

	fmt.Println(sum)
}
