package queue

import (
	"errors"
	"sync"
)

// QueueManager manages multiple queues using sync.Map
type QueueManager struct {
	queues sync.Map
}

// Queue is single queue FIFO
type Queue struct {
	items [][]byte
	mu    sync.RWMutex
}

// NewQueueManager creates new QueueManager
func NewQueueManager() *QueueManager {
	return &QueueManager{}
}

// GetQueue retrieve existing queue
func (qm *QueueManager) GetQueue(name string) *Queue {
	queue, _ := qm.queues.LoadOrStore(name, &Queue{items: make([][]byte, 0)})
	return queue.(*Queue)
}

// ListQueues returns list of existing queues
func (qm *QueueManager) ListQueues() []string {
	keys := []string{}
	qm.queues.Range(func(key, _ interface{}) bool {
		keys = append(keys, key.(string))
		return true
	})
	return keys
}

// Push add element to queue
func (q *Queue) Push(item []byte) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, item)
}

// Pop retrieve an element from queue and delete it
func (q *Queue) Pop() ([]byte, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.items) == 0 {
		return nil, errors.New("queue is empty")
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

// Peek retrieves the last element from queue without deleting
func (q *Queue) Peek() ([]byte, error) {
	q.mu.RLock()
	defer q.mu.RUnlock()

	if len(q.items) == 0 {
		return nil, errors.New("queue is empty")
	}

	return q.items[len(q.items)-1], nil
}
