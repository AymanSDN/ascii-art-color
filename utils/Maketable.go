package utils

func MakeTable(text []string, ascii *Ascii) [][]string {
	table := [][]string{}
	compteur := 0
	x := []string{}

	for i := 0; i < len(text); i++ {

		if compteur == 9 {

			table = append(table, x)
			x = []string{}
			compteur = 0
		}
		if compteur != 0 {
			if ascii.Color != "" && ascii.SubString == "" {
				x = append(x, Colors[ascii.Color]+text[i]+"\033[0m")
			} else {
				x = append(x, text[i])
			}
		}

		compteur++
	}
	return table
}
