package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"copilotkata/internal/mathutil"
	"copilotkata/internal/parser"
	"copilotkata/internal/strings"
)

// validateConfigPath ensures the config path is safe and within expected bounds
func validateConfigPath(path string) error {
	cleanPath := filepath.Clean(path)
	// Only allow files within testdata directory
	if !filepath.HasPrefix(cleanPath, "testdata/") && cleanPath != "testdata/sample_config.json" {
		return fmt.Errorf("invalid config path: %s", path)
	}
	return nil
}

func main() {
	cfgPath := "testdata/sample_config.json"

	// Validate config path to prevent path traversal attacks
	if err := validateConfigPath(cfgPath); err != nil {
		log.Fatalf("config path validation failed: %v", err)
	}

	b, err := os.ReadFile(cfgPath)
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	cfg, err := parser.ParseConfig(b)
	if err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}

	fmt.Println("App:", cfg.AppName)
	fmt.Println("MaxUsers:", cfg.MaxUsers)
	fmt.Println("Features:", cfg.Features)

	// Demo work: compute a Fibonacci term from config.
	fibN := cfg.MaxUsers % 30
	fmt.Println("Fibonacci(", fibN, ") =", mathutil.Fibonacci(fibN))

	// Slugify app name
	slug := strings.Slugify(cfg.AppName)
	fmt.Println("Slug:", slug)

	// Show the entire config as pretty JSON
	var out map[string]any
	_ = json.Unmarshal(b, &out)
	p, _ := json.MarshalIndent(out, "", "  ")
	fmt.Println(string(p))
}
