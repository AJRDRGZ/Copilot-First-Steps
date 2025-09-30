package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"copilotkata/internal/mathutil"
	"copilotkata/internal/parser"
	"copilotkata/internal/strings"
)

func main() {
	cfgPath := "testdata/sample_config.json"
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
