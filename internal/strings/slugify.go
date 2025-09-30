package strings

import (
\t"regexp"
\t"strings"
)

var nonAlnum = regexp.MustCompile(`[^a-z0-9\-]+`)

// Slugify converts an arbitrary string to a lowercase URL slug.
// NOTE: intentionally limited: accents/diacritics are not handled.
// TODO: Handle accents (e.g., "México" -> "mexico") and non-Latin scripts.
func Slugify(s string) string {
\ts = strings.TrimSpace(s)
\ts = strings.ToLower(s)
\ts = strings.ReplaceAll(s, "—", "-")
\ts = strings.ReplaceAll(s, "–", "-")
\ts = strings.ReplaceAll(s, "_", "-")
\ts = strings.ReplaceAll(s, " ", "-")

\t// Collapse multiple dashes
\tfor strings.Contains(s, "--") {
\t\ts = strings.ReplaceAll(s, "--", "-")
\t}

\t// Remove anything not a-z, 0-9, or dash
\ts = nonAlnum.ReplaceAllString(s, "")

\t// Trim leading/trailing dashes again, just in case
\ts = strings.Trim(s, "-")
\treturn s
}
