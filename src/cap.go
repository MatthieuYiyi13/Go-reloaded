package reloaded

import (
	"strings"
	"unicode"
)

func Cap(text string) string {
	words := strings.Fields(text)
	var result []string
	for _, word := range words {
		last := len(result) - 1
		if word == "(cap)" || word == "(Cap)" {
			if last >= 0 {
				letters := []rune(strings.ToLower(result[last]))
				letters[0] = unicode.ToUpper(letters[0])
				result[last] = string(letters)
			}
		} else {
			result = append(result, word)
		}
	}
	return strings.Join(result, " ")
}
