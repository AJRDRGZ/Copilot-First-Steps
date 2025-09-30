package parser

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseConfig_Sample(t *testing.T) {
	p := filepath.Join("..", "..", "testdata", "sample_config.json")
	b, err := os.ReadFile(p)
	if err != nil {
		t.Fatalf("read: %v", err)
	}
	cfg, err := ParseConfig(b)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if cfg.AppName == "" {
		t.Errorf("AppName should not be empty")
	}
}
