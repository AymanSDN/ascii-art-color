package utils

import (
	"log"
	"os"
	"strings"
)

func Readfile(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	if filename == "ascii/thinkertoy.txt" {
		return strings.Split(string(data), "\r\n")
	}
	return strings.Split(string(data), "\n")
}
