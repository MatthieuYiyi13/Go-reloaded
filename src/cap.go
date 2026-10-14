package reloaded

import "strings"

func Cap(text string) string {
	words := strings.Fields(text)
	var result []string
	for _, word := range words {
		last := len(result) - 1
		if word == "(cap)" || word == "(Cap)" {
			if last >= 0 {
				letters := []rune(result[last])
				result[last] = strings.ToUpper(string(letters[0])) + strings.ToLower(string(letters[1:]))
			}
		} else {
			result = append(result, word)
		}
	}
	return strings.Join(result, " ")
}
