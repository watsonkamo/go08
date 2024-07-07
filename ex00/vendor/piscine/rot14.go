package piscine

func Rot14(s string) string {
	runes := []rune(s)
	for i := range runes {
		if (runes[i] >= 'A' && runes[i] <= 'M') || (runes[i] >= 'a' && runes[i] <= 'm') {
			runes[i] += 14
		} else if (runes[i] >= 'N' && runes[i] <= 'Z') || (runes[i] >= 'n' && runes[i] <= 'z') {
			runes[i] -= 12
		}
	}
	return string(runes)
}
