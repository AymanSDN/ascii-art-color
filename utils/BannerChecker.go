package utils

func BannerChecker(s string) bool {
	arr := []string{"standard", "shadow", "thinkertoy"}
	for _, elemnt := range arr {
		if elemnt == s {
			return true
		}
	}
	return false
}
