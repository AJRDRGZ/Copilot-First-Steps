package parser

import (
	"strings"
	"testing"
)

func Test_ParseConfig_SecurityValidation(t *testing.T) {
	tests := []struct {
		name        string
		jsonData    []byte
		expectError bool
		errorMsg    string
	}{
		{
			name:        "input too large",
			jsonData:    []byte(strings.Repeat("a", MaxJSONSize+1)),
			expectError: true,
			errorMsg:    "input too large",
		},
		{
			name:        "invalid UTF-8",
			jsonData:    []byte("{\xff\xfe\"invalid\": \"utf8\"}"),
			expectError: true,
			errorMsg:    "invalid UTF-8 encoding",
		},
		{
			name:        "appName too long",
			jsonData:    []byte(`{"appName":"` + strings.Repeat("a", MaxAppNameLen+1) + `","maxUsers":10,"features":[],"threshold":0.5}`),
			expectError: true,
			errorMsg:    "appName too long",
		},
		{
			name:        "maxUsers negative",
			jsonData:    []byte(`{"appName":"Test","maxUsers":-1,"features":[],"threshold":0.5}`),
			expectError: true,
			errorMsg:    "maxUsers cannot be negative",
		},
		{
			name:        "maxUsers too large",
			jsonData:    []byte(`{"appName":"Test","maxUsers":2000000,"features":[],"threshold":0.5}`),
			expectError: true,
			errorMsg:    "maxUsers too large",
		},
		{
			name:        "threshold out of range - negative",
			jsonData:    []byte(`{"appName":"Test","maxUsers":10,"features":[],"threshold":-0.1}`),
			expectError: true,
			errorMsg:    "threshold must be between 0 and 1",
		},
		{
			name:        "threshold out of range - too large",
			jsonData:    []byte(`{"appName":"Test","maxUsers":10,"features":[],"threshold":1.1}`),
			expectError: true,
			errorMsg:    "threshold must be between 0 and 1",
		},
		{
			name:        "too many features",
			jsonData:    createJSONWithManyFeatures(),
			expectError: true,
			errorMsg:    "too many features",
		},
		{
			name:        "feature name too long",
			jsonData:    []byte(`{"appName":"Test","maxUsers":10,"features":["` + strings.Repeat("a", MaxFeatureNameLen+1) + `"],"threshold":0.5}`),
			expectError: true,
			errorMsg:    "feature name too long",
		},
		{
			name:        "valid config at limits",
			jsonData:    []byte(`{"appName":"` + strings.Repeat("a", MaxAppNameLen) + `","maxUsers":1000000,"features":["test"],"threshold":1.0}`),
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config, err := ParseConfig(tt.jsonData)
			
			if tt.expectError {
				if err == nil {
					t.Errorf("ParseConfig() expected error %q but got nil", tt.errorMsg)
				} else if err.Error() != tt.errorMsg {
					t.Errorf("ParseConfig() error = %q, want %q", err.Error(), tt.errorMsg)
				}
			} else {
				if err != nil {
					t.Errorf("ParseConfig() unexpected error: %v", err)
				}
				if config.AppName == "" {
					t.Error("ParseConfig() returned empty appName for valid input")
				}
			}
		})
	}
}

// createJSONWithManyFeatures creates a JSON config with too many features
func createJSONWithManyFeatures() []byte {
	features := make([]string, MaxFeaturesLen+1)
	for i := range features {
		features[i] = "feature" + string(rune('0'+i%10))
	}
	
	json := `{"appName":"Test","maxUsers":10,"features":[`
	for i, feature := range features {
		if i > 0 {
			json += ","
		}
		json += `"` + feature + `"`
	}
	json += `],"threshold":0.5}`
	
	return []byte(json)
}