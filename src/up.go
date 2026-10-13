package reloaded

import "strings"

func Up(text string) string {
	words := strings.Fields(text)
	var result []string
	for _, word := range words {
		last := len(result) - 1
		if word == "(up)" || word == "(Up)" {
			if last >= 0 {
				result[last] = strings.ToUpper(result[last])
			}
		} else {
			result = append(result, word)
		}
	}
	return strings.Join(result, " ")
}
