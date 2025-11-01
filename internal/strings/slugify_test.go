package strings

import (
	stdstrings "strings"
	"testing"
)

func Test_Slugify_SecurityLimits_PreventDoS(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "input too long",
			input:    stdstrings.Repeat("a", MaxSlugifyInputLength+1),
			expected: "",
		},
		{
			name:     "max allowed length",
			input:    stdstrings.Repeat("a", MaxSlugifyInputLength),
			expected: stdstrings.Repeat("a", MaxSlugifyInputLength),
		},
		{
			name:     "invalid UTF-8",
			input:    "hello\xff\xfeinvalid",
			expected: "",
		},
		{
			name:     "valid UTF-8 characters",
			input:    "héllo wörld",
			expected: "hllo-wrld", // accents removed by regex
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Slugify(tt.input)
			if result != tt.expected {
				t.Errorf("Slugify(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func Test_collapseDashes_EfficiencyAndSecurity(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "multiple consecutive dashes",
			input:    "hello----world",
			expected: "hello-world",
		},
		{
			name:     "many consecutive dashes (DoS prevention)",
			input:    stdstrings.Repeat("-", 1000),
			expected: "-",
		},
		{
			name:     "mixed content with many dashes",
			input:    "a" + stdstrings.Repeat("-", 100) + "b",
			expected: "a-b",
		},
		{
			name:     "no dashes",
			input:    "hello world",
			expected: "hello world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := collapseDashes(tt.input)
			if result != tt.expected {
				t.Errorf("collapseDashes(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
