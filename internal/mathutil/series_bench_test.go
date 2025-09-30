package mathutil

import "testing"

func BenchmarkFibonacci10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Fibonacci(10)
	}
}

func BenchmarkFibonacci25(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Fibonacci(25)
	}
}
