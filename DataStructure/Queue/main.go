package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	queue *list.List
}

func (q *Queue) Enqueue(value string) {
	q.queue.PushBack(value)
}

func (q *Queue) Dequeue() error {
	if q.queue.Len() > 0 {
		ele := q.queue.Front()
		q.queue.Remove(ele)
	}
	return fmt.Errorf("Pop Error: Queue is empty")
}

func (q *Queue) Front() (string, error) {
	if q.queue.Len() > 0 {
		if val, ok := q.queue.Front().Value.(string); ok {
			return val, nil
		}
		return "", fmt.Errorf("Peep Error: Queue Datatype is incorrect")
	}
	return "", fmt.Errorf("Peep Error: Queue is empty")
}

func (q *Queue) Size() int {
	return q.queue.Len()
}

func (q *Queue) Empty() bool {
	return q.queue.Len() == 0
}

func main() {
	Queue := &Queue{
		queue: list.New(),
	}
	fmt.Printf("Enqueue: A\n")
	Queue.Enqueue("Thor")
	fmt.Printf("Enqueue: B\n")
	Queue.Enqueue("Hulk")
	fmt.Printf("Size: %d\n", Queue.Size())
	for Queue.Size() > 0 {
		frontVal, _ := Queue.Front()
		fmt.Printf("Front: %s\n", frontVal)
		fmt.Printf("Dequeue: %s\n", frontVal)
		Queue.Dequeue()
	}
	fmt.Printf("Size: %d\n", Queue.Size())
}
