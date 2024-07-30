package utils

func CheckEmpty(arr []string) int {
	counter := 0
	for _, v := range arr {
		if v == "" {
			counter++
		}
	}
	return counter
}
