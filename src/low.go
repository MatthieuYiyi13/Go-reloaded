package reloaded

import "strings"

func Low(text string) string {
	words := strings.Fields(text)
	var result []string
	for _, word := range words {
		last := len(result) - 1
		if word == "(Low)" || word == "(low)" {
			if last >= 0 {
				result[last] = strings.ToLower(result[last])
			}
		} else {
			result = append(result, word)
		}
	}
	return strings.Join(result, " ")
}
