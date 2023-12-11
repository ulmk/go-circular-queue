package main

import "fmt"

func main() {
	q := &Queue{l: 0, r: 0, count: 0}

	for i := 1; i <= 5; i++ {
		q.Enqueue(i)
	}

	for i := 1; i <= 5; i++ {
		fmt.Println(q.Dequeue())
	}
}
