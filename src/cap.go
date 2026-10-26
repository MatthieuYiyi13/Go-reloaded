package reloaded

import (
	"strconv"
	"strings"
)

func Cap(text string) string {
	words := strings.Fields(text)
	var sortie []string
	for i := 0; i < len(words); i++ {
		word := words[i]
		var nombre int
		if len(word) >= 5 && (word[:5] == "(cap," || word[:5] == "(Cap,") { // On precise apres la virgule le nombre de mots qu'on souhaite mettre avec une premiere lettre en majuscule.
			instruction := word
			if (word == "(cap," || word == "(Cap,") && i+1 < len(words) {
				instruction += words[i+1]
			}
			if instruction[len(instruction)-1] != ')' {
				sortie = append(sortie, word)
				continue
			}
			valeur, err := strconv.Atoi(instruction[5 : len(instruction)-1])
			if err != nil || valeur <= 0 {
				sortie = append(sortie, word)
				continue
			}
			nombre = valeur
			if instruction != word {
				i++
			}
		} else {
			sortie = append(sortie, word)
			continue
		}
		for j := len(sortie) - nombre; j < len(sortie); j++ {
			letters := []rune(sortie[j])
			sortie[j] = strings.ToUpper(string(letters[0])) + strings.ToLower(string(letters[1:]))
		}
	}
	return strings.Join(sortie, " ")
}
