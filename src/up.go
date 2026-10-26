package reloaded

import (
	"strconv"
	"strings"
)

func Up(text string) string {
	words := strings.Fields(text)
	var sortie []string
	for i := 0; i < len(words); i++ {
		word := words[i]
		var nombre int
		if len(word) >= 4 && (word[:4] == "(up," || word[:4] == "(Up,") { // On precise apres la virgule le nombre de mots qu'on souhaite mettre en majuscule.
			instruction := word
			if (word == "(up," || word == "(Up,") && i+1 < len(words) {
				instruction += words[i+1]
			}
			if instruction[len(instruction)-1] != ')' {
				sortie = append(sortie, word)
				continue
			}
			// On récupère ce qui se trouve entre "(up," et ")".
			valeur, err := strconv.Atoi(instruction[4 : len(instruction)-1])
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
			sortie[j] = strings.ToUpper(sortie[j]) // l'inverse de ToLower
		}
	}
	return strings.Join(sortie, " ")
}
