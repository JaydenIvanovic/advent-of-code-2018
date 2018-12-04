package utils

import (
	"io/ioutil"
	"log"
	"strings"
)

func ReadLinesFromFile(fileName string) []string {
	rawData, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatalln("Unable to read file")
	}

	data := RawBytesToString(rawData)

	lines := strings.Split(data, "\n")

	return RemoveEmptyStrings(lines)
}

func RemoveEmptyStrings(lines []string) []string {
	cleanedStrings := make([]string, 0, len(lines))
	for _, line := range lines {
		if line != "" {
			cleanedStrings = append(cleanedStrings, line)
		}
	}
	return cleanedStrings
}
