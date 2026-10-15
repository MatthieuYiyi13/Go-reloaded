package reloaded

import (
	"fmt"
	"strconv"
	"strings"
)

func Bin(text string) string {
	words := strings.Fields(text)
	var sortie []string
	for _, word := range words {
		fin := len(sortie) - 1
		if (word == "(bin)" || word == "(Bin)") && fin >= 0 {
			number, err := strconv.ParseInt(sortie[fin], 2, 64)
			if err == nil {
				sortie[fin] = fmt.Sprint(number)
				continue
			}
		}
		sortie = append(sortie, word)
	}
	return strings.Join(sortie, " ")
}
