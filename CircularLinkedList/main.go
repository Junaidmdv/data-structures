package main

import "fmt"

type CircularList struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Val  int
	Next *Node
}

func (cl *CircularList) InsertAtHead(val int) {
	NewNode := &Node{Val: val}

	if cl.Head == nil {
		cl.Head = NewNode
		cl.Tail = NewNode
		cl.Tail.Next = cl.Head
	} else {
		NewNode.Next = cl.Head
		cl.Head = NewNode
		cl.Tail.Next = cl.Head

	}
}

func (cl *CircularList) Print() {
	current := cl.Head
	for {
		fmt.Printf(" %d >", current.Val)
		current = current.Next
		if current == cl.Head {
			break
		}
	}
	fmt.Println()
}

func (cl *CircularList) InsertAtTail(val int) {
	NewNode := &Node{Val: val}
	if cl.Head == nil {
		cl.Head = NewNode
		cl.Tail = NewNode
		cl.Tail.Next = cl.Head
	} else {
		cl.Tail.Next = NewNode
		cl.Tail = NewNode
		cl.Tail.Next = cl.Head
	}
}

func (cl *CircularList) DelteAtTail() {
	if cl.Head == nil {
		return
	} else if cl.Head == cl.Tail {
		cl.Head = nil
		cl.Tail = nil
	} else {
		prev := cl.Head

		for prev.Next != cl.Tail {
			prev = prev.Next
		}
		cl.Tail = prev
		cl.Tail.Next = cl.Head

	}
}

func (cl *CircularList) DeleteAtHead() {

	if cl.Head == nil {
		return
	} else if cl.Head == cl.Tail {
		cl.Head = nil
		cl.Tail = nil
		return
	} else {
		cl.Head = cl.Head.Next
		cl.Tail.Next = cl.Head
	}
}

func main() {
	cl := CircularList{}
	cl.InsertAtTail(1)
	cl.Print()
	cl.InsertAtHead(2)
	cl.InsertAtHead(4)
	cl.InsertAtHead(6)
	cl.InsertAtHead(8)
	cl.InsertAtHead(10)
	cl.InsertAtTail(12)
	cl.InsertAtTail(14)
	cl.InsertAtTail(16)
	cl.Print()
	cl.DeleteAtHead()
	cl.Print()

}
