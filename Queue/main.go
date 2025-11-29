package main

import "fmt"

type Queue struct {
	Item []int
}

func (q *Queue) EnQueue(val int) {
	q.Item = append(q.Item, val)
}

func (q *Queue) DeQueue() int {
	val := q.Item[0]
	q.Item = q.Item[1:]
	return val
}

func (q *Queue) Top() int {
	return q.Item[0]
}

func (q *Queue) IsEmpty() bool {
	return len(q.Item) == 0
}

func (q *Queue) PrintQueue() {
	for _, val := range q.Item {
		fmt.Printf("%d \n", val)
	}
}
