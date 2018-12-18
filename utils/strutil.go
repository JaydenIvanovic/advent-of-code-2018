package utils

import (
	"log"
	"strconv"
	"strings"
)

func RawBytesToString(data []byte) string {
	b := strings.Builder{}

	_, err := b.Write(data)
	if err != nil {
		log.Fatalln("Failed converting raw byte data to string")
	}

	return b.String()
}

func StrToInt(i string) int {
	num, err := strconv.Atoi(strings.Trim(i, " "))

	if err != nil {
		panic(err)
	}

	return num
}
