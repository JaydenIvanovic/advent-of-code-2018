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

	return lines
}
