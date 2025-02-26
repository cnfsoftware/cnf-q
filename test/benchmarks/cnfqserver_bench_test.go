package benchmarks

import (
	"testing"

	"cnf-q/pkg/queueclient"
)

func Benchmark_CNFQServer_Push(b *testing.B) {
	client := queueclient.NewClient("http://localhost:8080")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = client.Push("a", []byte("a"))
	}
}

func Benchmark_CNFQServer_Pop(b *testing.B) {
	client := queueclient.NewClient("http://localhost:8080")
	for i := 0; i < 1000; i++ {
		_ = client.Push("b", []byte("a"))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = client.Pop("b")
	}
}

func Benchmark_CNFQServer_Peek(b *testing.B) {
	client := queueclient.NewClient("http://localhost:8080")
	for i := 0; i < 1000; i++ {
		_ = client.Push("c", []byte("a"))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = client.Peek("c")
	}
}
