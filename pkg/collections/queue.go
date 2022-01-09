package collections

import (
	"container/list"
	"fmt"
)

// A simple way to implement a queue.
// ref: [go-queue](https://github.com/phf/go-queue/blob/master/queue/queue.go)

// Using slice
func queueBySlice() {
	var queue []int

	// Enqueue
	queue = append(queue, 1)
	queue = append(queue, 0)

	head := queue[0]
	queue[0] = 0
	// Dequeue
	queue = queue[1:]
	fmt.Println("Head of queue ", head)
}

// Using container/list
func queueByList() {
	// Create
	queue := list.New()

	// Enqueue
	queue.PushBack("Hello ")
	queue.PushBack("world!")

	for queue.Len() > 0 {
		// Print first
		e := queue.Front()
		fmt.Print(e.Value)

		// Dequeue
		queue.Remove(e)
	}
}
