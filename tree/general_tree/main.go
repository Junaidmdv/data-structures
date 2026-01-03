package main

import "fmt"

var cap int

type Node struct {
	val   int
	child []*Node
}

func NewNode(val int) *Node {
	return &Node{
		val:   val,
		child: []*Node{},
	}
}

func (t *Node) Insert(val int) {
	queue := []*Node{t}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if len(cur.child) < cap {
			cur.child = append(cur.child, NewNode(val))
			return
		}
		for _, child := range cur.child {
			queue = append(queue, child)
		}
	}
}

func (t *Node) Print() {
	if t == nil {
		return
	}

	queue := []*Node{t}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		fmt.Print(cur.val, " ")

		for _, c := range cur.child {
			queue = append(queue, c)
		}
	}

	fmt.Println()
}

func main() {
	t := NewNode(5)
	cap = 5

	for i := range 20 {
		t.Insert(i)
	}

	t.Print()
}
