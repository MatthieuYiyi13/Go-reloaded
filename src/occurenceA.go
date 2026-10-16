package reloaded

import "strings"

func OccurenceA(text string) string {
	words := strings.Fields(text)

	for index := 0; index < len(words)-1; index++ {
		next := strings.ToLower(words[index+1])
// Ajout de l'occurence de "a" devant les voyelles et le "h"
		if next[0] == 'a' || next[0] == 'e' || next[0] == 'i' || next[0] == 'o' || next[0] == 'u' || next[0] == 'h' {
			if words[len(words)-1] == "a" {
				words[index] = "an"
			} else if words[len(words)-1] == "A" {
				words[index] = "An"
			}
		}
	}

	return strings.Join(words, " ")
}
