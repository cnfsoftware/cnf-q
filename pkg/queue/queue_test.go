package queue

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sync"
	"testing"
)

// TestNewQueueManager ensures a new QueueManager is properly initialized
func TestNewQueueManager(t *testing.T) {
	qm := NewQueueManager()
	require.NotNil(t, qm, "QueueManager should not be nil")
}

// TestQueueManager_GetQueue ensures that queues are correctly created and retrieved
func TestQueueManager_GetQueue(t *testing.T) {
	qm := NewQueueManager()

	q1 := qm.GetQueue("testQueue")
	q2 := qm.GetQueue("testQueue")

	assert.Same(t, q1, q2, "Expected to get the same queue instance")

	q3 := qm.GetQueue("anotherQueue")
	assert.NotSame(t, q1, q3, "Expected different queue instances")
}

// TestQueueManager_ListQueues ensures that created queues are correctly listed
func TestQueueManager_ListQueues(t *testing.T) {
	qm := NewQueueManager()

	qm.GetQueue("queue1")
	qm.GetQueue("queue2")

	queues := qm.ListQueues()

	expected := map[string]bool{
		"queue1": true,
		"queue2": true,
	}

	assert.Len(t, queues, len(expected), "The number of queues should match the expected")

	for _, q := range queues {
		assert.Contains(t, expected, q, "Unexpected queue name: "+q)
	}
}

// TestQueue_PushAndPop ensures elements are added and removed in FIFO order
func TestQueue_PushAndPop(t *testing.T) {
	q := &Queue{items: make([][]byte, 0)}

	q.Push([]byte("first"))
	q.Push([]byte("second"))

	item, err := q.Pop()
	require.NoError(t, err, "Pop should not return an error")
	assert.Equal(t, "first", string(item), "First item should be 'first'")

	item, err = q.Pop()
	require.NoError(t, err, "Pop should not return an error")
	assert.Equal(t, "second", string(item), "Second item should be 'second'")
}

// TestQueue_PopEmpty ensures popping an empty queue returns an error
func TestQueue_PopEmpty(t *testing.T) {
	q := &Queue{items: make([][]byte, 0)}

	_, err := q.Pop()
	assert.Error(t, err, "Pop on an empty queue should return an error")
	assert.Equal(t, "queue is empty", err.Error(), "Error message should be 'queue is empty'")
}

// TestQueue_Peek ensures peek retrieves the last item without removing it
func TestQueue_Peek(t *testing.T) {
	q := &Queue{items: make([][]byte, 0)}

	q.Push([]byte("first"))
	q.Push([]byte("second"))

	item, err := q.Peek()
	require.NoError(t, err, "Peek should not return an error")
	assert.Equal(t, "second", string(item), "Peek should return 'second'")

	// Ensure the queue is still intact
	assert.Len(t, q.items, 2, "Queue length should remain 2 after peek")
}

// TestQueue_PeekEmpty ensures peeking an empty queue returns an error
func TestQueue_PeekEmpty(t *testing.T) {
	q := &Queue{items: make([][]byte, 0)}

	_, err := q.Peek()
	assert.Error(t, err, "Peek on an empty queue should return an error")
	assert.Equal(t, "queue is empty", err.Error(), "Error message should be 'queue is empty'")
}

// TestQueue_Concurrency ensures the queue is thread-safe
func TestQueue_Concurrency(t *testing.T) {
	q := &Queue{items: make([][]byte, 0)}

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			q.Push([]byte{byte(i)})
		}(i)
	}
	wg.Wait()

	assert.Len(t, q.items, 100, "Queue should contain 100 items after concurrency test")
}
