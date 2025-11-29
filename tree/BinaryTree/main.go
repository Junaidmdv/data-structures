package main

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func (t *Node) LevelOrderInsert(data int) {
	NewNode := &Node{Data: data}
	queue := []*Node{t}
	if t == nil {
		fmt.Println("root value is empty")
		return
	}

	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Left == nil {
			current.Left = NewNode
			return
		} else {
			queue = append(queue, current.Left)
		}

		if current.Right == nil {
			current.Right = NewNode
			return
		} else {
			queue = append(queue, current.Right)
		}
	}

}

func LevelOrderTraversal(root *Node) [][]int {
	var result [][]int
	queue := []*Node{root}

	for len(queue) != 0 {
		var level []int
		len := len(queue)
		for i := 0; i < len; i++ {
			current := queue[0]
			queue = queue[1:]
			level = append(level, current.Data)

			if current.Left != nil {
				queue = append(queue, current.Left)
			}

			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
		result = append(result, level)
	}

	return result
}

//print all the leaf node in the tree

func (t *Node) PrintLeafNode() {

	if t == nil {
		return
	}

	if t.Left == nil && t.Right == nil {
		fmt.Println(t.Data)
		return
	}

	if t.Left != nil {
		t.Left.PrintLeafNode()
	}

	if t.Right != nil {
		t.Right.PrintLeafNode()
	}
}

func (t *Node) HeightofTree() int {
	if t == nil {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	return 1 + max(t.Left.HeightofTree(), t.Right.HeightofTree())
}

func (t *Node) DepthOfNode(tar int) int {

	if t == nil {
		return -1
	}

	if tar == t.Data {
		return 0
	}

	left := t.Left.DepthOfNode(tar)
	if left != -1 {
		return left+1
	}

	right := t.Right.DepthOfNode(tar)
	if right != -1 {
		return right+1
	}

	return -1

}

func main() {
	t := &Node{Data: 50}
	t.LevelOrderInsert(23)
	t.LevelOrderInsert(23)
	t.LevelOrderInsert(98)
	t.LevelOrderInsert(23)
	t.LevelOrderInsert(93)
	t.LevelOrderInsert(65)
	t.LevelOrderInsert(23)
	t.LevelOrderInsert(20)
	t.LevelOrderInsert(85)
	t.LevelOrderInsert(43)
	t.LevelOrderInsert(26)
	result := LevelOrderTraversal(t)

	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			fmt.Printf("%d  ", result[i][j])
		}
		fmt.Println()
	}

	fmt.Printf("Height of tree:%d \n", t.HeightofTree())

	t.PrintLeafNode()

	fmt.Printf("depth of node:%d \n", t.DepthOfNode(23))
}
