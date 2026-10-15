package reloaded

import "strings"

func Up(text string) string {
	words := strings.Fields(text)
	var sortie []string
	for _, word := range words {
		fin := len(sortie) - 1
		if word == "(up)" || word == "(Up)" {
			if fin >= 0 {
				sortie[fin] = strings.ToUpper(sortie[fin])
			}
		} else {
			sortie = append(sortie, word)
		}
	}
	return strings.Join(sortie, " ")
}