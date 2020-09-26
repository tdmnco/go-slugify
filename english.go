package slugify

import (
	"regexp"
	"strings"
)

// English returns slugified text in English
func English(s string) string {
	var re = regexp.MustCompile("[^a-z0-9]+")

	s = strings.Trim(strings.ToLower(s), "-")

	return re.ReplaceAllString(s, "-")
}
