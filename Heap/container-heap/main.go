package main

import (
	"container/heap"
	"fmt"
)

//package implement max heap using this standard go package. While declaring five method, it implement heap interface.

type MaxHeap []int

func (m MaxHeap) Len() int           { return len(m) }
func (m MaxHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m MaxHeap) Less(i, j int) bool { return m[i] > m[j] }

func (m *MaxHeap) Push(val interface{}) {
	*m = append(*m, val.(int))

}

func (m *MaxHeap) Pop() interface{} {
	old := *m
	l := len(old) - 1
	val := old[l]
	*m = old[:l]
	return val

}

func main() {
	h := &MaxHeap{23, 5, 21, 32, 5, 6, 6, 43, 2}
	heap.Init(h)
	fmt.Println(heap.Pop(h))
	fmt.Println(*h)
}
