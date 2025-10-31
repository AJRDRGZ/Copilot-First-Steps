package user

import (
	"errors"
	"regexp"
	"strings"
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

// Validate checks basic fields. Intentionally simplistic to give Copilot room.
func Validate(u User) error {
	if strings.TrimSpace(u.Name) == "" {
		return errors.New("name is required")
	}
	if !emailRe.MatchString(strings.TrimSpace(u.Email)) {
		return errors.New("invalid email")
	}
	return nil
}
