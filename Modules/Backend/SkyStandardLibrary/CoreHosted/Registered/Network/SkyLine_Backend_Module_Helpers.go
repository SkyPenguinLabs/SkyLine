package SkyLine_Network

import "strings"

func CleanCharacter(Value string, char rune) string {
	return strings.Map(func(r rune) rune {
		if r == char {
			return -1
		}
		return r
	}, Value)
}
