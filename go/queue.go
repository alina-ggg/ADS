package main

import (
	"fmt"
)

type QueueNode struct {
	data string
	next *QueueNode
}

type Queue struct {
	front *QueueNode
	rear  *QueueNode
}

func (q *Queue) SPUSH(value string) {
	newNode := &QueueNode{data: value, next: nil}
	if q.rear != nil {
		q.rear.next = newNode
	}
	q.rear = newNode
	if q.front == nil {
		q.front = q.rear
	}
}

func (q *Queue) PRINT() {
	curr := q.front
	for curr != nil {
		fmt.Printf("%s ", curr.data)
		curr = curr.next
	}
	fmt.Println()
}
