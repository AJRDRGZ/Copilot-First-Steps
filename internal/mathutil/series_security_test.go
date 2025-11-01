package mathutil

import (
	"testing"
)

func Test_Fibonacci_SecurityLimits_PreventDoS(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "input at safe limit",
			input:    50,          // Safe input within limits
			expected: 12586269025, // F(50)
		},
		{
			name:     "input exceeds limit",
			input:    MaxFibonacciInput + 1,
			expected: FibonacciZeroValue, // Should return 0 for security
		},
		{
			name:     "very large input (potential DoS)",
			input:    1000,
			expected: FibonacciZeroValue, // Should return 0 for security
		},
		{
			name:     "negative input",
			input:    -1,
			expected: FibonacciZeroValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test iterative function (most commonly used)
			result := FibonacciIterative(tt.input)
			if result != tt.expected {
				t.Errorf("FibonacciIterative(%d) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func Test_validateFibonacciInput_SecurityChecks(t *testing.T) {
	tests := []struct {
		name        string
		input       int
		expectError bool
	}{
		{
			name:        "valid input within limits",
			input:       50,
			expectError: false,
		},
		{
			name:        "input at maximum limit",
			input:       MaxFibonacciInput,
			expectError: false,
		},
		{
			name:        "input exceeds limit",
			input:       MaxFibonacciInput + 1,
			expectError: true,
		},
		{
			name:        "very large input",
			input:       1000000,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateFibonacciInput(tt.input)

			if tt.expectError && err == nil {
				t.Error("validateFibonacciInput() expected error but got nil")
			}
			if !tt.expectError && err != nil {
				t.Errorf("validateFibonacciInput() unexpected error: %v", err)
			}
		})
	}
}
