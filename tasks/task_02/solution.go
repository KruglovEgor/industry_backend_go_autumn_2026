package main

func rotateRunes(s string, shift int) string {
	runes := []rune(s)
	if len(runes) == 0 {
		return ""
	}
	shift %= len(runes)
	if shift < 0 {
		shift += len(runes)
	}
	return string(runes[shift:]) + string(runes[:shift])
}
