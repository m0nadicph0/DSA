package main

import (
	"fmt"
	queue "queue/queue"
)

func main() {
	q := queue.NewLinkedListQueue()

	for i := 0; i < 10; i++ {
		q.Enqueue(i + 1)
	}
	front, _ := q.Front()
	rear, _ := q.Rear()
	fmt.Printf("Front=%d, Rear=%d\n", front, rear)

	for !q.IsEmpty() {
		value, _ := q.Dequeue()
		fmt.Println("Dequeued value=", value)
	}
	q.Enqueue(11)
	q.Enqueue(12)
	fmt.Printf("Size=%d\n", q.Size())
	value, _ := q.Dequeue()
	fmt.Println(value)
	value, _ = q.Dequeue()
	fmt.Println(value)
}
