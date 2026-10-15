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
		if (word == "(hex)" || word == "(Hex)") && fin >= 0 {
			number, err := strconv.ParseInt(sortie[fin], 16, 64)
			if err == nil {
				sortie[fin] = fmt.Sprint(number)
				continue
			}
		}
		sortie = append(sortie, word)
	}
	return strings.Join(sortie, " ")
}