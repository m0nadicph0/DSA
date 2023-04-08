package main

import (
	"fmt"
	queue "queue/queue"
)

func main() {
	q := queue.NewSliceQueue()

	for i := 0; i < 10; i++ {
		q.Enqueue(i)
	}
	front, _ := q.Front()
	rear, _ := q.Rear()
	fmt.Printf("Front=%d, Rear=%d\n", front, rear)

	for !q.IsEmpty() {
		value, _ := q.Dequeue()
		fmt.Println("Dequeued value=", value)
	}
}
