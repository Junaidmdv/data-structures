package main

import "fmt"

type MaxHeap struct {
	Item []int
}

func (m *MaxHeap) Insert(val int) {
	m.Item = append(m.Item, val)
	m.HeapifyUp(len(m.Item) - 1)
}
func (m *MaxHeap) Remove() int {
	val := m.Item[0]
	l := len(m.Item)-1
	m.Item[0] = m.Item[l]
	m.Item = m.Item[:l]
	m.HeapifyDown(0)
	return val

}

func (m *MaxHeap) HeapifyDown(idx int) {
	max := idx
	left := idx*2 + 1
	right := idx*2 + 2

	if left < len(m.Item) && m.Item[left] > m.Item[max] {
		max = left
	}

	if right < len(m.Item) && m.Item[right] > m.Item[max] {
		max = right
	}

	if idx != max {
		m.Swap(max, idx)
		m.HeapifyDown(max)
	}

}

func (m *MaxHeap) Swap(idx1, idx2 int) {
	m.Item[idx1], m.Item[idx2] = m.Item[idx2], m.Item[idx1]

}

func (m *MaxHeap) HeapifyUp(idx int) {

	parent := (idx - 1) / 2

	if parent >= 0 && m.Item[parent] < m.Item[idx] {
		m.Swap(idx, parent)
		m.HeapifyUp(parent)
	}

}

func main() {
	m := &MaxHeap{}
	m.Insert(13)
	m.Insert(23)
	m.Insert(12)
	m.Insert(52)
	fmt.Println("value removed:", m.Remove())
	fmt.Println(m.Item)
}
