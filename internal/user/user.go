package user

import (
	"errors"
	"net/mail"
	"regexp"
	"strings"
	"unicode/utf8"
)

const (
	// Security limits to prevent DoS attacks
	MaxNameLength  = 255
	MaxEmailLength = 320 // RFC 5321 limit for email addresses
)

// User represents a user in the system with an ID, name, and email address.
type User struct {
	ID    int    // unique identifier for the user
	Name  string // full name of the user
	Email string // email address of the user
}

// emailRe is a regular expression for basic email validation.
// Note: This is a simplified regex and may not cover all valid email formats.
var emailRe = regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)

// Validate performs comprehensive validation of user data including security checks.
func Validate(u User) error {
	// Validate name
	trimmedName := strings.TrimSpace(u.Name)
	if trimmedName == "" {
		return errors.New("name is required")
	}

	// Check name length to prevent DoS attacks
	if !utf8.ValidString(u.Name) {
		return errors.New("name contains invalid UTF-8 characters")
	}
	if len(u.Name) > MaxNameLength {
		return errors.New("name too long")
	}

	// Validate email
	trimmedEmail := strings.TrimSpace(u.Email)
	if trimmedEmail == "" {
		return errors.New("invalid email")
	}

	// Check email length to prevent DoS attacks
	if len(u.Email) > MaxEmailLength {
		return errors.New("email too long")
	}

	// Use Go's built-in email validation for better security
	if _, err := mail.ParseAddress(trimmedEmail); err != nil {
		return errors.New("invalid email")
	}

	// Additional regex check for stricter validation
	if !emailRe.MatchString(trimmedEmail) {
		return errors.New("invalid email")
	}

	return nil
}
