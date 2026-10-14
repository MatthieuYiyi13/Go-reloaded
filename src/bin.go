package reloaded

import (
	"strconv"
	"strings"
)
func Bin(text string) string {
	words := strings.Fields(text)
	var result []string
	for _, word := range words {
		last := len(result) - 1
		if word == "(bin)" || word == "(Bin)" && last >= 0 {
			number, err := strconv.ParseInt(result[last], 2, 64)
			if err == nil {
				result[last] = strconv.FormatInt(number, 10)
				continue
			}
		}
		result = append(result, word)
	}
	return strings.Join(result, " ")
}