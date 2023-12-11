package main

import (
	"testing"
)

func TestEnqueueDequeue(t *testing.T) {
	q := &Queue{l: 0, r: 0, count: 0}

	values := []int{1, 2, 3, 4, 5}
	// for i := 1; i <= 5; i++ {
	// 	q.Enqueue(i)
	// }
	for _, v := range values {
		q.Enqueue(v)
	}

	// for i := 1; i <= 5; i++ {
	// 	fmt.Println(q.Dequeue())
	// }
	for i := 0; i < len(values); i++ {
		dequeued := q.Dequeue()
		expected := values[i]

		if dequeued != expected {
			t.Errorf("Expected %d, but got %d", expected, dequeued)
		}
	}

}
