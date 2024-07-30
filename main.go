package main

import (
	"fmt"
	"os"
	"strings"

	"utils/utils"
)

const usage = "Usage: go run . [STRING] [BANNER] \n\nEX: EX: go run . something standard "

// write to standard error
func f_perror(str string) {
	fmt.Fprintln(os.Stdout, str)
	os.Exit(0)
}

func ArgsCheker(arg *[]string, NewAscii *utils.Ascii) {
	if len(*arg) == 0 {
		f_perror(usage)
	}

	color := utils.FlagChecker((*arg)[0])
	if color == "" {
		NewAscii.Input = (*arg)[0]
	} else {
		NewAscii.Color = color
	}

	*arg = (*arg)[1:]

	if len(*arg) == 0 {
		return
	}

	if NewAscii.Input != "" && utils.BannerChecker((*arg)[0]) {

		if len((*arg)) > 1 {
			f_perror(usage)
		}

		NewAscii.Banner = (*arg)[0] + ".txt"
		*arg = (*arg)[1:]
		return
	} else if NewAscii.Color != "" {

		NewAscii.Input = (*arg)[0]
		*arg = (*arg)[1:]

		if len(*arg) >= 1 {
			if utils.BannerChecker((*arg)[len((*arg))-1]) {
				NewAscii.Banner = (*arg)[len((*arg))-1] + ".txt"
				*arg = (*arg)[:len(*arg)-1]
			}
		}

		if len(*arg) == 1 {
			NewAscii.SubString = NewAscii.Input
			NewAscii.Input = (*arg)[0]
			*arg = (*arg)[1:]
		}

	}
}

// check args function :
func main() {
	NewAscii := utils.NewAscii()
	arg := os.Args[1:]

	ArgsCheker(&arg, NewAscii)

	if len(arg) != 0 || (NewAscii.Color != "" && NewAscii.Input == "") {
		f_perror(usage)
	}

	text := utils.Readfile("ascii/" + NewAscii.Banner) // Reading the input and returning the lines as "text []string"

	table := utils.MakeTable(text, NewAscii) // Organizing the "table[][]" with 9 lines for each element

	lines := strings.Split(NewAscii.Input, "\\n") // Spliting the "input" with "\n"

	coutner := utils.CheckEmpty(lines) // Counting empty elements in "input"

	if coutner == len(lines) { // decreasing first element for succesive "\n"
		lines = lines[1:]
	}

	if NewAscii.SubString != "" {
		utils.ColorPrinter(lines, table, &utils.Ascii{}) // Printing table by lines (input)
		return
	}

	utils.PrintAscii(lines, table) // Printing table by lines (input)
}
