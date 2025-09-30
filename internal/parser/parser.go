package parser

import (
	"encoding/json"
	"errors"
)

type Config struct {
	AppName  string   `json:"appName"`
	MaxUsers int      `json:"maxUsers"`
	Features []string `json:"features"`
	Threshold float64 `json:"threshold"`
}

// ParseConfig loads a JSON-encoded config. The code below is intentionally
// a bit verbose and opinionated to give you refactoring targets.
func ParseConfig(b []byte) (Config, error) {
	var c Config
	if len(b) == 0 {
		return c, errors.New("empty input")
	}
	if err := json.Unmarshal(b, &c); err != nil {
		return c, err
	}
	// Minimal validation (could be improved)
	if c.AppName == "" {
		return c, errors.New("missing appName")
	}
	if c.MaxUsers < 0 {
		// questionable policy: negative users coerced to zero elsewhere
		c.MaxUsers = 0
	}
	return c, nil
}
