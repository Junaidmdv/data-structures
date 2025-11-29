package main

import "fmt"

func LevelOrderAsLevel(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	var result [][]int
	queue := []*Node{root}

	for len(queue) != 0 {
		length := len(queue)
		var level []int

		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Data)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}
	return result
}

func LevelOrderTraversal(root *Node) {
	queue := []*Node{root}

	for len(queue) != 0 {
		current := queue[0]
		fmt.Println(current.Data)
        queue=queue[1:]

		if current.Left != nil{
			queue=append(queue, current.Left)
		}

		if current.Right != nil{
			queue=append(queue, current.Right)
		}

	}

}
