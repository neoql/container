package queue

import (
	"container/list"
	"sync"
	"errors"
)

// Queue is a scalable FIFO queue.
type Queue struct {
	lock	sync.RWMutex
	list	*list.List
	cond	*sync.Cond
	closed	bool
}

// ErrClosed will be return when operate a closed queue.
var ErrClosed  = errors.New("the queue is already closed")

// New instance of Queue
func New() *Queue {
	q := &Queue{closed: false}
	q.list = list.New()
	q.cond = sync.NewCond(&q.lock)
	return q
}

// Put an item to queue back.
func (queue *Queue) Put(item interface{}) error {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.closed {
		return ErrClosed
	}

	queue.list.PushBack(item)
	queue.cond.Signal()
	return nil
}

// Pop front value from queue, Returns false if queue closed.
func (queue *Queue) Pop() (item interface{}, flag bool) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	for queue.list.Len() == 0 {
		if queue.closed {
			return nil, false
		}
		queue.cond.Wait()
	}
	top := queue.list.Front()
	queue.list.Remove(top)

	return top.Value, true
}

// Size returns the size of the queue.
func (queue *Queue) Size() int {
	queue.lock.RLock()
	defer queue.lock.RUnlock()
	return queue.list.Len()
}

// Close queue.
func (queue *Queue) Close() error {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.closed {
		return ErrClosed
	}
	queue.closed = true
	queue.cond.Broadcast()
	return nil
}

// IsClosed returns true when queue is closed else false.
func (queue *Queue) IsClosed() bool {
	queue.lock.RLock()
	defer queue.lock.RUnlock()
	return queue.closed
}
