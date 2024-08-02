package pyroscopeaction

import "testing"

var result int

func BenchmarkFib10(b *testing.B) {
	n := 0
	for i := 0; i < b.N; i++ {
		n = Fib(10)
	}
	result = n
}

func BenchmarkFib100_000_000(b *testing.B) {
	n := 0
	for i := 0; i < b.N; i++ {
		n = Fib(100_000_000)
	}
	result = n
}
