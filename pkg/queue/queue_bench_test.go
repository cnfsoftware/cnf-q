package queue

import "testing"

// BenchmarkQueue_Push tests the performance of pushing elements to the queue
func BenchmarkQueue_Push(b *testing.B) {
	q := &Queue{items: make([][]byte, 0)}
	b.ResetTimer()

	// Pushing 1000 elements
	for i := 0; i < b.N; i++ {
		q.Push([]byte(""))
	}
}

// BenchmarkQueue_Pop tests the performance of popping elements from the queue
func BenchmarkQueue_Pop(b *testing.B) {
	q := &Queue{items: make([][]byte, 1000)} // Pre-fill the queue with 1000 items
	for i := 0; i < 1000; i++ {
		q.items[i] = []byte("")
	}
	b.ResetTimer()

	// Popping 1000 elements
	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}

// BenchmarkQueue_Peek tests the performance of peeking elements from the queue
func BenchmarkQueue_Peek(b *testing.B) {
	q := &Queue{items: make([][]byte, 1000)} // Pre-fill the queue with 1000 items
	for i := 0; i < 1000; i++ {
		q.items[i] = []byte("")
	}
	b.ResetTimer()

	// Peeking 1000 elements
	for i := 0; i < b.N; i++ {
		q.Peek()
	}
}
