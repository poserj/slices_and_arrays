package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type CircularQueue struct {
	values  []int
	isFull  bool // queue is full
	isEmpty bool // queue is empty
	start   int  // start index
	end     int
}

func NewCircularQueue(size int) CircularQueue {
	cq := CircularQueue{start: 0, end: 0}
	cq.values = make([]int, size)
	cq.isEmpty = true
	if size == 0 {
		cq.isFull = true
	}
	return cq
}

func (q *CircularQueue) Push(value int) bool {
	if q.isFull {
		return false
	}

	q.values[q.end] = value
	q.end = (q.end + 1) % len(q.values)
	q.isFull = q.end == q.start
	q.isEmpty = false
	return true
}

func (q *CircularQueue) Pop() bool {
	if !q.isFull && q.start == q.end {
		//empty queue
		q.isEmpty = true
		return false
	} else {
		q.start = (q.start + 1) % len(q.values)
		q.isFull = false
		q.isEmpty = false
		return true
	}
}

func (q *CircularQueue) Front() int {
	if q.Empty() {
		return -1
	}
	return q.values[q.start]
}

func (q *CircularQueue) Back() int {
	if q.Empty() {
		return -1
	}
	if q.end == 0 {
		return q.values[len(q.values)-1]
	}
	return q.values[q.end-1]
}

func (q *CircularQueue) Empty() bool {
	return q.isEmpty
}

func (q *CircularQueue) Full() bool {
	return q.isFull
}

func TestCircularQueue(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())

	assert.Equal(t, -1, queue.Front())
	assert.Equal(t, -1, queue.Back())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, queue.values))

	assert.False(t, queue.Empty())
	assert.True(t, queue.Full())

	assert.Equal(t, 1, queue.Front())
	assert.Equal(t, 3, queue.Back())

	assert.True(t, queue.Pop())
	assert.False(t, queue.Empty())
	assert.False(t, queue.Full())
	assert.True(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{4, 2, 3}, queue.values))

	assert.Equal(t, 2, queue.Front())
	assert.Equal(t, 4, queue.Back())

	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())
}
