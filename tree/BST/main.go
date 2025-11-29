package main

import "fmt"

func main() {
	t := &Node{Data: 50}

	t.Insert(55)
	t.Insert(45)
	t.Insert(32)
	t.Insert(48)
	t.Insert(80)
	//t.InOrder()
	//t.NonDecreasingOrder()
	//t.PostOrder()

	result := LevelOrderAsLevel(t)

	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			fmt.Printf(" %d ", result[i][j])
		}
		fmt.Println()
	}

	fmt.Println("Normal level order Traversal")

	LevelOrderTraversal(t)

	key := 32
	fmt.Printf("Is node %d exist:%v \n", key, t.SearchNode(key))

}
