package strings

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

const (
	// Security limit to prevent DoS attacks through large input strings
	MaxSlugifyInputLength = 1000
)

var nonAlnum = regexp.MustCompile(`[^a-z0-9\-]+`)

// Slugify converts an arbitrary string to a lowercase URL slug.
// NOTE: intentionally limited: accents/diacritics are not handled.
// TODO: Handle accents (e.g., "México" -> "mexico") and non-Latin scripts.
// Includes security measures to prevent ReDoS attacks.
func Slugify(s string) string {
	// Security check: limit input length to prevent DoS attacks
	if len(s) > MaxSlugifyInputLength {
		return ""
	}

	// Security check: validate UTF-8 to prevent potential issues
	if !utf8.ValidString(s) {
		return ""
	}

	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	// Replace various characters with dashes
	replacements := []string{"—", "–", "_", " "}
	for _, char := range replacements {
		s = strings.ReplaceAll(s, char, "-")
	}

	// More efficient approach to collapse multiple dashes (prevents ReDoS)
	s = collapseDashes(s)

	// Remove anything not a-z, 0-9, or dash
	s = nonAlnum.ReplaceAllString(s, "")

	// Trim leading/trailing dashes again, just in case
	s = strings.Trim(s, "-")
	return s
}

// collapseDashes efficiently collapses multiple consecutive dashes into single dashes
func collapseDashes(s string) string {
	var result strings.Builder
	result.Grow(len(s))

	prevDash := false
	for _, r := range s {
		if r == '-' {
			if !prevDash {
				result.WriteRune(r)
				prevDash = true
			}
		} else {
			result.WriteRune(r)
			prevDash = false
		}
	}

	return result.String()
}
