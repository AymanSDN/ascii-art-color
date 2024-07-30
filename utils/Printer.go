package utils

import "fmt"

func PrintAscii(lines []string, table [][]string) {
	for j := 0; j < len(lines); j++ {

		if lines[j] == "" {
			fmt.Println()
			continue
		}

		line := lines[j]

		for a := 0; a < 8; a++ {
			for i := 0; i < len(line); i++ {
				if line[i] >= 32 && line[i] <= 126 { // Print only printable character
					fmt.Print(table[line[i]-32][a])
				} else {
					fmt.Println("Erorr special caracter")
					return
				}
			}
			fmt.Println()
		}

	}
}

func ColorPrinter(lines []string, table [][]string, ascii *Ascii) {
	// khdmo hadok l3bid
}
