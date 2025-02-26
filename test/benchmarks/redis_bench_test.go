package benchmarks

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
)

func Benchmark_Redis_Push(b *testing.B) {
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client.RPush(ctx, "a", "a")
	}
}

func Benchmark_Redis_Pop(b *testing.B) {
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379"})

	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		client.RPush(ctx, "b", "a").Err()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = client.LPop(ctx, "b").Result()
	}
}
