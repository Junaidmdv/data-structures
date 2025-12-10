package main

import (
	"fmt"
)

type Graph struct {
	AdjList map[int][]*Edge
}

type Edge struct {
	src    int
	dest   int
	weight int
}

func NewGraph() *Graph {
	return &Graph{
		AdjList: make(map[int][]*Edge),
	}
}

func (g *Graph) Insert(src, dest, weight int) {
	edge := &Edge{src: src, dest: dest, weight: weight}
	revedge := &Edge{src: dest, dest: src, weight: weight}
	g.AdjList[dest] = append(g.AdjList[dest], revedge)
	g.AdjList[src] = append(g.AdjList[src], edge)
}

func (g *Graph) DFS(cur int, visited map[int]bool) {

	visited[cur] = true
	fmt.Println(cur)
	for _, v := range g.AdjList[cur] {
		if !visited[v.dest] {
			visited[v.dest] = true
			g.DFS(v.dest, visited)
		}
	}
}

func (g *Graph) ShortestPath(start, tar int) {
	visited := make(map[int]bool)
	parent := make(map[int]int)

	if _, exist := g.AdjList[start]; !exist {
		fmt.Println("starting is missing")
		return
	}
	if _, exist := g.AdjList[tar]; !exist {
		fmt.Println("tart is missing")
		return
	}

	queue := []int{start}
	visited[start] = true
	parent[start] = -1

	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]
		if current == tar {
			break
		}

		for _, v := range g.AdjList[current] {
			if !visited[v.dest] {
				visited[v.dest] = true
				parent[v.dest] = current
				queue = append(queue, v.dest)
			}
		}
	}

	if !visited[tar] {
		fmt.Println("there is not path to reach tar")
		return
	}

	path := []int{}
	v := tar

	for v != -1 {
		path = append(path, v)
		v = parent[v]
	}

	i := 0
	j := len(path) - 1

	for i < j {
		path[i], path[j] = path[j], path[i]
		i++
		j--
	}
	fmt.Println(path)
}

// detect cycle in the directed graph

func (g *Graph) IsCycle(cur int, visited map[int]bool, par int) bool {
	visited[cur] = true

	for _, v := range g.AdjList[cur] {
		if visited[v.dest] && cur != par {
			return true
		}
		if !visited[v.dest] {
			return g.IsCycle(v.dest, visited, cur)
		}
	}
	return false
}

func main() {
	g := NewGraph()
	g.Insert(1, 2, 5)
	g.Insert(1, 3, 5)
	g.Insert(1, 4, 5)
	g.Insert(2, 3, 5)
	g.Insert(4, 5, 5)
	g.Insert(1, 5, 5)

	visited := make(map[int]bool)

	for k := range g.AdjList {
		if !visited[k] {
			g.DFS(k, visited)
		}
	}
	g.ShortestPath(1, 5)

}
