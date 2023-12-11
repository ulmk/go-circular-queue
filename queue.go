package main

import "fmt"

const queueSize = 5

//  [l...r] <- enqueue
type Queue struct {
	entries [queueSize]int // [0]el1 [1]el2
	l       int
	r       int
	count   int
}

func (q *Queue) Enqueue(item int) {
	if q.count == queueSize {
		fmt.Println("Queue is full")
		return

	}
	q.entries[q.r] = item
	q.r = (q.r + 1) % queueSize
	q.count++
}

func (q *Queue) Dequeue() int {
	if q.count == 0 {
		fmt.Println("Queue is empty")
		return -1
	}
	item := q.entries[q.l]
	q.l = (q.l + 1) % queueSize
	q.count--
	return item
}
