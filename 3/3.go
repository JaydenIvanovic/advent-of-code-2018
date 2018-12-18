package main

import (
	"fmt"
	"strings"

	"github.com/jaydenivanovic/advent-of-code-2018/utils"
)

type FabricPlan struct {
	Id      string
	OffsetX int
	OffsetY int
	Width   int
	Height  int
}

func (fp *FabricPlan) CreateKey(x int, y int) string {
	return fmt.Sprintf("%d-%d", fp.OffsetX+x, fp.OffsetY+y)
}

var fabricOverlapMap map[string]int

func partOne() {
	lines := utils.ReadLinesFromFile("./puzzle_input.txt")

	for _, line := range lines {
		processLine(line)
	}

	overlapCount := 0
	for _, count := range fabricOverlapMap {
		if count >= 2 {
			overlapCount++
		}
	}

	fmt.Println(overlapCount)
}

func processLine(line string) {
	// build a map by creating keys from the x and y values
	// for example (given this line): #1 @ 808,550: 12x22
	// keys can be created as so 808-550, 809-550, 810-550 ... 820-570, 820-571, 820-572
	// when these keys are found we can increment the integer counter (the maps value)
	// this means we only need to iterate through the list once
	fabricPlan := createFabricPlan(line)
	for i := 0; i < fabricPlan.Width; i++ {
		for j := 0; j < fabricPlan.Height; j++ {
			key := fabricPlan.CreateKey(i, j)
			_, present := fabricOverlapMap[key]
			if present {
				fabricOverlapMap[key] += 1
			} else {
				fabricOverlapMap[key] = 1
			}
		}
	}
}

func createFabricPlan(planInput string) *FabricPlan {
	s := strings.Split(planInput, "@")

	id := s[0]

	s = strings.Split(s[1], ":")

	offsetPair := strings.Split(s[0], ",")
	widthHeightPair := strings.Split(s[1], "x")

	return &FabricPlan{
		Id:      id,
		OffsetX: utils.StrToInt(offsetPair[0]),
		OffsetY: utils.StrToInt(offsetPair[1]),
		Width:   utils.StrToInt(widthHeightPair[0]),
		Height:  utils.StrToInt(widthHeightPair[1]),
	}
}

func main() {
	fabricOverlapMap = make(map[string]int)
	partOne()
}
