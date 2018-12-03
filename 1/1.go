package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func rawBytesToString(data []byte) string {
	b := strings.Builder{}

	_, err := b.Write(data)
	if err != nil {
		log.Fatalln("Failed converting raw byte data to string")
	}

	return b.String()
}

func applyFrequencyChange(frequency int, delta string) int {
	num, err := strconv.Atoi(delta[1:])

	if err != nil {
		log.Println("Failed to parse frequency change to int:", delta)
	}

	if delta[0] == '+' {
		return frequency + num
	} else if delta[0] == '-' {
		return frequency - num
	} else {
		log.Println("Could not apply frequency change:", delta)
	}

	return 0
}

func main() {
	frequencyChangesRaw, err := ioutil.ReadFile("./1.txt")

	if err != nil {
		log.Fatalln("Error reading file")
	}

	frequency := 0
	frequencyChanges := strings.Split(rawBytesToString(frequencyChangesRaw), "\n")
	for _, frequencyChange := range frequencyChanges {
		if frequencyChange != "" {
			frequency = applyFrequencyChange(frequency, frequencyChange)
		}
	}

	fmt.Println("Calculated frequency:", frequency)
}
