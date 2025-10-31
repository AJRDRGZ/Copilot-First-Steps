package mathutil

import "testing"

func TestFibonacciImplementations(t *testing.T) {
	// Test cases with expected Fibonacci values
	testCases := []struct {
		input    int
		expected int
	}{
		{-1, 0},  // negative input
		{0, 0},   // F(0) = 0
		{1, 1},   // F(1) = 1
		{2, 1},   // F(2) = 1
		{3, 2},   // F(3) = 2
		{4, 3},   // F(4) = 3
		{5, 5},   // F(5) = 5
		{6, 8},   // F(6) = 8
		{7, 13},  // F(7) = 13
		{8, 21},  // F(8) = 21
		{9, 34},  // F(9) = 34
		{10, 55}, // F(10) = 55
		{15, 610}, // F(15) = 610
		{20, 6765}, // F(20) = 6765
	}

	// Test all implementations
	implementations := map[string]func(int) int{
		"Recursive":    Fibonacci,
		"Iterative":    FibonacciIterative,
		"Memoized":     FibonacciMemoized,
		"DP":           FibonacciDP,
		"DP Optimized": FibonacciDPOptimized,
	}

	for name, fn := range implementations {
		t.Run(name, func(t *testing.T) {
			for _, tc := range testCases {
				// Skip recursive for large inputs to avoid long test times
				if name == "Recursive" && tc.input > 25 {
					continue
				}
				
				result := fn(tc.input)
				if result != tc.expected {
					t.Errorf("%s(%d) = %d, expected %d", name, tc.input, result, tc.expected)
				}
			}
		})
	}
}

// Test consistency between all implementations for a range of values
func TestFibonacciConsistency(t *testing.T) {
	for n := 0; n <= 30; n++ {
		// Skip recursive for larger values
		if n <= 25 {
			recursive := Fibonacci(n)
			iterative := FibonacciIterative(n)
			if recursive != iterative {
				t.Errorf("Inconsistency at F(%d): recursive=%d, iterative=%d", n, recursive, iterative)
			}
		}

		// Test all O(n) implementations
		iterative := FibonacciIterative(n)
		memoized := FibonacciMemoized(n)
		dp := FibonacciDP(n)
		dpOptimized := FibonacciDPOptimized(n)

		if iterative != memoized {
			t.Errorf("Inconsistency at F(%d): iterative=%d, memoized=%d", n, iterative, memoized)
		}
		if iterative != dp {
			t.Errorf("Inconsistency at F(%d): iterative=%d, dp=%d", n, iterative, dp)
		}
		if iterative != dpOptimized {
			t.Errorf("Inconsistency at F(%d): iterative=%d, dpOptimized=%d", n, iterative, dpOptimized)
		}
	}
}