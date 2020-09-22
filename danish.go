package slugify

import (
	"regexp"
	"strings"
)

// Danish returns slugified text in Danish
func Danish(s string) string {
	var re = regexp.MustCompile("[^a-z0-9]+")

	s = strings.Trim(strings.ToLower(s), "-")

	s = strings.ReplaceAll(s, "æ", "ae")
	s = strings.ReplaceAll(s, "ø", "oe")
	s = strings.ReplaceAll(s, "å", "aa")

	return re.ReplaceAllString(s, "-")
}
