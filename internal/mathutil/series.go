package mathutil

const (
	// Fibonacci sequence constants
	FibonacciBaseCase0 = 0 // F(0) = 0
	FibonacciBaseCase1 = 1 // F(1) = 1
	FibonacciZeroValue = 0 // Return value for negative inputs
)

// Fibonacci returns the nth Fibonacci number.
// This implementation is intentionally naive (recursive) to invite refactoring.
// TODO: Consider an iterative approach or memoization for better performance.
func Fibonacci(n int) int {
	if n < 0 {
		return FibonacciZeroValue // questionable behavior: negative inputs collapse to zero
	}
	if n == 0 {
		return FibonacciBaseCase0
	}
	if n == 1 {
		return FibonacciBaseCase1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
