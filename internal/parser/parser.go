package parser

import (
	"encoding/json"
	"errors"
)

type Config struct {
	AppName   string   `json:"appName"`
	MaxUsers  int      `json:"maxUsers"`
	Features  []string `json:"features"`
	Threshold float64  `json:"threshold"`
}

// ParseConfig loads a JSON-encoded config. The code below is intentionally
// a bit verbose and opinionated to give you refactoring targets.
func ParseConfig(jsonData []byte) (Config, error) {
	var config Config
	if len(jsonData) == 0 {
		return config, errors.New("empty input")
	}
	if err := json.Unmarshal(jsonData, &config); err != nil {
		return config, err
	}
	// Minimal validation (could be improved)
	if config.AppName == "" {
		return config, errors.New("missing appName")
	}
	if config.MaxUsers < 0 {
		// questionable policy: negative users coerced to zero elsewhere
		config.MaxUsers = 0
	}
	return config, nil
}
