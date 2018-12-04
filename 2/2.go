package main

import (
	"fmt"
	"strings"

	"github.com/jaydenivanovic/advent-of-code-2018/utils"
)

func partOne() {
	lines := utils.ReadLinesFromFile("./2.txt")

	two := 0
	three := 0
	for _, line := range lines {
		charCounter := make(map[rune]int)
		for _, character := range line {
			_, present := charCounter[character]
			if present {
				charCounter[character] += 1
			} else {
				charCounter[character] = 1
			}
		}
		foundTwo := false
		foundThree := false
		for _, count := range charCounter {
			if count == 2 && !foundTwo {
				two += 1
				foundTwo = true
			} else if count == 3 && !foundThree {
				three += 1
				foundThree = true
			}
		}
	}

	fmt.Println(two * three)
}

func validBoxPair(a string, b string) bool {
	const MissMax = 2

	length := len(a)
	i := 0
	misses := 0
	for i < length && misses < MissMax {
		if a[i] != b[i] {
			misses += 1
		}
		i += 1
	}

	return misses < MissMax
}

func getCommonCharacters(a string, b string) string {
	builder := strings.Builder{}

	for i, _ := range a {
		if a[i] == b[i] {
			builder.WriteByte(a[i])
		}
	}

	return builder.String()
}

func partTwo() {
	lines := utils.ReadLinesFromFile("./2.txt")

	for i, line := range lines {
		for _, innerLine := range lines[i+1:] {
			if validBoxPair(line, innerLine) {
				fmt.Println(getCommonCharacters(line, innerLine))
			}
		}
	}
}

func main() {
	partOne()
	partTwo()
}
