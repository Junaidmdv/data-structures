package main

import (
	"fmt"
)

type QueueList struct {
	Head *Node
}

type Node struct {
	Data int
	Next *Node
}

func (q *QueueList) EnQueue(val int) {
	NewNode := &Node{Data: val}
	if q.Head == nil {
		q.Head = NewNode
		return
	}

	current := q.Head

	for current.Next != nil {
		current = current.Next
	}

	current.Next = NewNode

}

func (q *QueueList) DeQueue() int {
	if q.Head == nil {
		fmt.Println("empty linked list")
		return 0
	}

	val := q.Head.Data
	q.Head = q.Head.Next
	return val
}

func (q *QueueList) Top() int {
	if q.Head == nil {
		fmt.Println("empty linked list")
		return 0
	}
	return q.Head.Data
}

func (q *QueueList) Print() {
	current := q.Head

	for current != nil {
		fmt.Printf("%d \n", current.Data)
		current = current.Next
	}
}

func (q *QueueList) IsEmpty() bool {
	return q.Head == nil
}
