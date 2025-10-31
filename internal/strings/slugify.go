package strings

import (
	"regexp"
	"strings"
)

var nonAlnum = regexp.MustCompile(`[^a-z0-9\-]+`)

// Slugify converts an arbitrary string to a lowercase URL slug.
// NOTE: intentionally limited: accents/diacritics are not handled.
// TODO: Handle accents (e.g., "México" -> "mexico") and non-Latin scripts.
func Slugify(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	// Replace various characters with dashes
	replacements := []string{"—", "–", "_", " "}
	for _, char := range replacements {
		s = strings.ReplaceAll(s, char, "-")
	}

	// Collapse multiple dashes
	for strings.Contains(s, "--") {
		s = strings.ReplaceAll(s, "--", "-")
	}

	// Remove anything not a-z, 0-9, or dash
	s = nonAlnum.ReplaceAllString(s, "")

	// Trim leading/trailing dashes again, just in case
	s = strings.Trim(s, "-")
	return s
}
