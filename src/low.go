package reloaded

import "strings"

func Low(text string) string {
	words := strings.Fields(text)
	var sortie []string
	for _, word := range words {
		fin := len(sortie) - 1
		if word == "(Low)" || word == "(low)" {
			if fin >= 0 {
				sortie[fin] = strings.ToLower(sortie[fin]) // Conversion en minuscules du mot précédent avec ToLower
			}
		} else {
			sortie = append(sortie, word) // append ajoute le mot à la fin de la slice sortie
		}
	}
	return strings.Join(sortie, " ") // Join réunit les mots pour cée une seule et meme string
}
