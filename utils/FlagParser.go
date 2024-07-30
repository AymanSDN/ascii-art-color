package utils

import (
	"fmt"
	"os"
)

const ColorFlag = "--color="

func FlagChecker(s string) string {
	color := ""

	if len(s) >= 8 && s[:8] == ColorFlag {
		_, ok := Colors[s[8:]]
		if !ok {
			fmt.Println("This color not supported")
			os.Exit(2)
		}
		color = s[8:]
	}
	return color
}
