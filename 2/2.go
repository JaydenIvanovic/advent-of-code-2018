package main

import (
	"fmt"

	"github.com/jaydenivanovic/advent-of-code-2018/utils"
)

func main() {
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
