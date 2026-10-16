package reloaded

import (
	"strconv"
	"strings"
	"fmt"
)

func Hex(text string) string {
	words := strings.Fields(text)
	var sortie []string
	for _, word := range words {
		 fin := len(sortie) - 1
		 if (word == "(hex)" || word == "(Hex)") && fin >= 0 { // meme chose que bin mais avec la base 16 (hexadecimal)
			number, err := strconv.ParseInt(sortie[fin], 16, 64) //ParsInt convertit le mot en question sur une base 0, 2, 8, 16, 32, 64
			if err == nil {
				sortie[fin] = fmt.Sprint(number) // Sprint remet la valeur a son premier format (string) pour l'afficher
				continue
			}
		}
		sortie = append(sortie, word)
	}
	return strings.Join(sortie, " ")
}