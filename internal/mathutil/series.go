package mathutil

const (
	// Fibonacci sequence constants
	FibonacciBaseCase0 = 0 // F(0) = 0
	FibonacciBaseCase1 = 1 // F(1) = 1
	FibonacciZeroValue = 0 // Return value for negative inputs
)

// Fibonacci returns the nth Fibonacci number using naive recursion.
// This implementation is intentionally naive (recursive) to invite refactoring.
// Time complexity: O(2^n), Space complexity: O(n)
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

// FibonacciIterative returns the nth Fibonacci number using iteration.
// Time complexity: O(n), Space complexity: O(1)
func FibonacciIterative(n int) int {
	if n < 0 {
		return FibonacciZeroValue
	}
	if n == 0 {
		return FibonacciBaseCase0
	}
	if n == 1 {
		return FibonacciBaseCase1
	}

	prev2, prev1 := 0, 1
	for i := 2; i <= n; i++ {
		current := prev1 + prev2
		prev2, prev1 = prev1, current
	}
	return prev1
}

// FibonacciMemoized returns the nth Fibonacci number using memoization.
// Time complexity: O(n), Space complexity: O(n)
func FibonacciMemoized(n int) int {
	memo := make(map[int]int)
	return fibonacciMemoHelper(n, memo)
}

func fibonacciMemoHelper(n int, memo map[int]int) int {
	if n < 0 {
		return FibonacciZeroValue
	}
	if n == 0 {
		return FibonacciBaseCase0
	}
	if n == 1 {
		return FibonacciBaseCase1
	}

	if val, exists := memo[n]; exists {
		return val
	}

	memo[n] = fibonacciMemoHelper(n-1, memo) + fibonacciMemoHelper(n-2, memo)
	return memo[n]
}

// FibonacciDP returns the nth Fibonacci number using dynamic programming (bottom-up).
// Time complexity: O(n), Space complexity: O(n)
func FibonacciDP(n int) int {
	if n < 0 {
		return FibonacciZeroValue
	}
	if n == 0 {
		return FibonacciBaseCase0
	}
	if n == 1 {
		return FibonacciBaseCase1
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// FibonacciDPOptimized returns the nth Fibonacci number using space-optimized DP.
// Time complexity: O(n), Space complexity: O(1)
// This is essentially the same as FibonacciIterative but shows the DP thinking process.
func FibonacciDPOptimized(n int) int {
	if n < 0 {
		return FibonacciZeroValue
	}
	if n == 0 {
		return FibonacciBaseCase0
	}
	if n == 1 {
		return FibonacciBaseCase1
	}

	dp0, dp1 := 0, 1
	for i := 2; i <= n; i++ {
		dp0, dp1 = dp1, dp0+dp1
	}
	return dp1
}
