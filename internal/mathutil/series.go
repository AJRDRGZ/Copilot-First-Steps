package mathutil

// Fibonacci returns the nth Fibonacci number.
// This implementation is intentionally naive (recursive) to invite refactoring.
// TODO: Consider an iterative approach or memoization for better performance.
func Fibonacci(n int) int {
	if n < 0 {
		return 0 // questionable behavior: negative inputs collapse to zero
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
