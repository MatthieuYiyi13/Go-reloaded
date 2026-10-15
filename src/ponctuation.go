package reloaded

import "strings"

func Ponc(text string) string {
	words := strings.Fields(text)
	sortie := ""
	for _, word := range words {
		if word[0] == '.' || word[0] == ',' || word[0] == '!' || word[0] == '?' || word[0] == ':' || word[0] == ';' {
			sortie += word
		} else {
			if sortie != "" {
				sortie += " "
			}
			sortie += word
		}
	}
	return sortie
}

