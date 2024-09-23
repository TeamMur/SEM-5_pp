package stringutils

func StringReverse(s *string) {

	runes := []rune(*s)
	var reversed_runes []rune
	for i := len(runes) - 1; i >= 0; i-- {
		reversed_runes = append(reversed_runes, runes[i])
	}
	*s = string(reversed_runes)
}
