package user

import (
	"errors"
	"regexp"
	"strings"
)

type User struct {
	ID    int
	Name  string
	Email string
}

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
