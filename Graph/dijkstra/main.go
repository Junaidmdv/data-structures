package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Node struct {
	Index int
	wgt   int
}

type PriorityQ []*Node

func (pq PriorityQ) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq PriorityQ) Less(i, j int) bool { return pq[i].wgt < pq[j].wgt }

func (pq PriorityQ) Len() int { return len(pq) }

func (pq *PriorityQ) Push(val interface{}) {
	*pq = append(*pq, val.(*Node))
}

func (pq *PriorityQ) Pop() interface{} {
	old := *pq
	n := len(old)
	val := old[n-1]
	old = old[:n-1]
	*pq = old
	return val

}

func Dijkstra(graph [][]Node, start, end int) []int {
	visited := make(map[int]bool)
	n := len(graph)
	dest := make([]int, n)

	for i := 0; i < n; i++ {
		dest[i] = math.MaxInt
		visited[i] = false
	}
	dest[start] = 0
	pq := &PriorityQ{}
	heap.Init(pq)
	heap.Push(pq, &Node{start, 0})

	for pq.Len() > 0 {
		u := heap.Pop(pq).(*Node)

		if visited[u.Index] {
			continue
		}

		visited[u.Index] = true

		for _, v := range graph[u.Index] {
			if !visited[v.Index] && dest[u.Index]+v.wgt < dest[v.Index] {
				dest[v.Index] = dest[u.Index] + v.wgt
				heap.Push(pq, &Node{v.Index, dest[u.Index] + v.wgt})

			}
		}

	}
	return dest
}

func main() {
	graph := [][]Node{
		{{1, 5}, {3, 9}},
		{{0, 5}, {2, 2}},
		{{1, 2}, {3, 3}, {4, 7}},
		{{0, 9}, {2, 3}},
		{{2, 7}},
	}
	fmt.Println("The given nodes are:", graph)
	start := 0
	end := 4
	dist := Dijkstra(graph, start, end)
	fmt.Printf("Shortest path from node %d to %d: %d\n", start, end, dist[end])

}
