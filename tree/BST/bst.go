package main

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func (b *Node) Insert(data int) {

	if data < b.Data {

		if b.Left == nil {
			b.Left = &Node{Data: data}
			return
		} else {
			b.Left.Insert(data)
		}

	} else if data > b.Data {

		if b.Right == nil {
			b.Right = &Node{Data: data}
			return
		} else {
			b.Right.Insert(data)
		}

	}
}
