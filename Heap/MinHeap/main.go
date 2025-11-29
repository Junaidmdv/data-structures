package main

import "fmt"

type MinHeap struct {
	Item []int
}

func (m *MinHeap) Insert(val int) {
	m.Item = append(m.Item, val)
	m.HeapifyUp(len(m.Item) - 1)

}

func (m *MinHeap) HeapifyUp(idx int) {
	parent := (idx - 1) / 2

	if parent >= 0 && m.Item[idx] < m.Item[parent] {
		m.Swap(idx, parent)
		m.HeapifyUp(parent)
	}
}

func (m *MinHeap) Remove() int {
	val := m.Item[0]
	l := len(m.Item) - 1
	m.Item[0] = m.Item[l]
	m.Item = m.Item[:l]
	m.HeapifyDown(0)
	return val
}

func (m *MinHeap) HeapifyDown(idx int) {
	min := idx
	left := idx*2 + 1
	right := idx*2 + 2

	if left < len(m.Item) && m.Item[min] > m.Item[left] {
		min = left
	}
	if right < len(m.Item) && m.Item[min] > m.Item[right] {
		min = right
	}

	if min != idx {
		m.Swap(min, idx)
		m.HeapifyDown(min)
	}

}

func (m *MinHeap) Swap(idx1, idx2 int) {
	m.Item[idx1], m.Item[idx2] = m.Item[idx2], m.Item[idx1]
}

func main() {
	m := &MinHeap{}
	m.Insert(23)
	m.Insert(64)
	m.Insert(62)
	m.Insert(12)
	fmt.Println(m.Remove())
	fmt.Println(m.Item)
}
