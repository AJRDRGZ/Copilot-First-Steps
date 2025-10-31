package mathutil

import "testing"

// Benchmark small Fibonacci numbers (n=10)
func BenchmarkFibonacci10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Fibonacci(10)
	}
}

func BenchmarkFibonacciIterative10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FibonacciIterative(10)
	}
}

func BenchmarkFibonacciMemoized10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FibonacciMemoized(10)
	}
}

func BenchmarkFibonacciDP10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FibonacciDP(10)
	}
}

func BenchmarkFibonacciDPOptimized10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FibonacciDPOptimized(10)
	}
}

// Benchmark medium Fibonacci numbers (n=25)
func BenchmarkFibonacci25(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Fibonacci(25)
	}
}

func BenchmarkFibonacciIterative25(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FibonacciIterative(25)
	}
}

func BenchmarkFibonacciMemoized25(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FibonacciMemoized(25)
	}
}

func BenchmarkFibonacciDP25(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FibonacciDP(25)
	}
}

func BenchmarkFibonacciDPOptimized25(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FibonacciDPOptimized(25)
	}
}

// Benchmark large Fibonacci numbers (n=40) - only for efficient algorithms
// Note: Recursive approach would be too slow for n=40

func BenchmarkFibonacciIterative40(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FibonacciIterative(40)
	}
}

func BenchmarkFibonacciMemoized40(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FibonacciMemoized(40)
	}
}

func BenchmarkFibonacciDP40(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FibonacciDP(40)
	}
}

func BenchmarkFibonacciDPOptimized40(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FibonacciDPOptimized(40)
	}
}

// Benchmark very large Fibonacci numbers (n=100)
func BenchmarkFibonacciIterative100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FibonacciIterative(100)
	}
}

func BenchmarkFibonacciMemoized100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FibonacciMemoized(100)
	}
}

func BenchmarkFibonacciDP100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FibonacciDP(100)
	}
}

func BenchmarkFibonacciDPOptimized100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FibonacciDPOptimized(100)
	}
}
