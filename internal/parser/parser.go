package parser

import (
	"encoding/json"
	"errors"
	"strings"
	"unicode/utf8"
)

const (
	// Security limits to prevent DoS attacks
	MaxJSONSize       = 1024 * 1024 // 1MB limit for JSON input
	MaxAppNameLen     = 255
	MaxFeaturesLen    = 100
	MaxFeatureNameLen = 50
)

type Config struct {
	AppName   string   `json:"appName"`
	MaxUsers  int      `json:"maxUsers"`
	Features  []string `json:"features"`
	Threshold float64  `json:"threshold"`
}

// ParseConfig loads a JSON-encoded config with comprehensive security validation.
func ParseConfig(jsonData []byte) (Config, error) {
	var config Config

	// Security check: limit input size to prevent DoS attacks
	if len(jsonData) == 0 {
		return config, errors.New("empty input")
	}
	if len(jsonData) > MaxJSONSize {
		return config, errors.New("input too large")
	}

	// Security check: validate UTF-8 encoding
	if !utf8.Valid(jsonData) {
		return config, errors.New("invalid UTF-8 encoding")
	}

	if err := json.Unmarshal(jsonData, &config); err != nil {
		return config, err
	}

	// Enhanced validation with security checks
	if err := validateConfig(config); err != nil {
		return config, err
	}

	return config, nil
}

// validateConfig performs comprehensive validation of config data
func validateConfig(config Config) error {
	// Validate AppName
	appName := strings.TrimSpace(config.AppName)
	if appName == "" {
		return errors.New("missing appName")
	}
	if len(config.AppName) > MaxAppNameLen {
		return errors.New("appName too long")
	}
	if !utf8.ValidString(config.AppName) {
		return errors.New("appName contains invalid UTF-8 characters")
	}

	// Validate MaxUsers (reasonable bounds)
	if config.MaxUsers < 0 {
		return errors.New("maxUsers cannot be negative")
	}
	if config.MaxUsers > 1000000 { // Reasonable upper limit
		return errors.New("maxUsers too large")
	}

	// Validate Features array
	if len(config.Features) > MaxFeaturesLen {
		return errors.New("too many features")
	}
	for _, feature := range config.Features {
		if len(feature) > MaxFeatureNameLen {
			return errors.New("feature name too long")
		}
		if !utf8.ValidString(feature) {
			return errors.New("feature name contains invalid UTF-8 characters")
		}
	}

	// Validate Threshold (must be between 0 and 1)
	if config.Threshold < 0 || config.Threshold > 1 {
		return errors.New("threshold must be between 0 and 1")
	}

	return nil
}
