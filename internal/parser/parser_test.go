package parser

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseConfig_Sample(t *testing.T) {
	configPath := filepath.Join("..", "..", "testdata", "sample_config.json")
	jsonData, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatalf("read: %v", err)
	}
	config, err := ParseConfig(jsonData)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if config.AppName == "" {
		t.Errorf("AppName should not be empty")
	}
}
