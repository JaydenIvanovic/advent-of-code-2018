package utils

import (
	"log"
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
