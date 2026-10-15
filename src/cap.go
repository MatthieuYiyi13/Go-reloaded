package reloaded

import "strings"

func Cap(text string) string {
	words := strings.Fields(text)
	var sortie []string
	for _, word := range words {
		fin := len(sortie) - 1
		if word == "(cap)" || word == "(Cap)" {
			if fin >= 0 {
				letters := []rune(sortie[fin])
				sortie[fin] = strings.ToUpper(string(letters[0])) + strings.ToLower(string(letters[1:]))
			}
		} else {
			sortie = append(sortie, word)
		}
	}
	return strings.Join(sortie, " ")
}
