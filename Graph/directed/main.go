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

	if _, exist := g.AdjList[dest]; !exist {
		g.AdjList[dest] = []*Edge{}
	}

	g.AdjList[src] = append(g.AdjList[src], edge)

}

func (g *Graph) BFS(start int, visited map[int]bool) {
	queue := []int{start}
	visited[start] = true

	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]
		fmt.Println(current, " ")

		for _, v := range g.AdjList[current] {
			if !visited[v.dest] {
				queue = append(queue, v.dest)
				visited[v.dest] = true
			}
		}
	}

}

// print all the possible path for the src to destination.Here we are using a modified dfs method and visited only used in the undirectional graph where src and destination are connected so it will  make cycle it will be an error so here only undirect node.

func PrintPaths(g *Graph, cur int, tar int, visited map[int]bool, path string) {
	if cur == tar {
		fmt.Println(path)
		return
	}

	visited[cur] = true
	for _, v := range g.AdjList[cur] {
		if !visited[v.dest] {
			newPath := fmt.Sprintf("%s->%d", path, v.dest)
			PrintPaths(g, v.dest, tar, visited, newPath)
		}
	}
	visited[cur] = false
}




func main() {
	g := NewGraph()

	g.Insert(0, 1, 1)
	g.Insert(0, 2, 1)
	g.Insert(1, 0, 1)
	g.Insert(1, 3, 1)
	g.Insert(2, 0, 1)
	g.Insert(2, 4, 1)
	g.Insert(3, 1, 1)
	g.Insert(3, 4, 1)
	g.Insert(3, 5, 1)
	g.Insert(4, 2, 1)
	g.Insert(4, 3, 1)
	g.Insert(4, 5, 1)

	g.Insert(5, 3, 1)
	g.Insert(5, 4, 1)
	g.Insert(5, 6, 1)
	g.Insert(6, 5, 1)
	visited := make(map[int]bool)

	for k := range g.AdjList {
		fmt.Printf("%d: ", k)
		for _, val := range g.AdjList[k] {
			fmt.Print(val.dest)
		}
		fmt.Println()
	}

	// for k := range g.AdjList {
	// 	if !visited[k] {
	// 		g.BFS(k, visited)
	// 	}
	// }
	PrintPaths(g, 0, 5, visited, "0")
}
