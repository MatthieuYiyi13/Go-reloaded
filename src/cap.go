package reloaded

import (
	"strings"
)
func Cap(nom string) string {
for _, r := range strings {
			if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')) {
			}
			if len(nom) > 0 {
				strings = strings[:0] + strings.ToUpper(nom[0:])
			}
			if len(nom) > 1 {
				strings = strings[:1] + strings.ToLower(nom[1:])
			}
		}
	}