package main

import "fmt"

type LinkedList struct {
	Head *Node
}

type Node struct {
	Val  int
	Next *Node
}

func (list *LinkedList) PushFront(val int) {
	NewNode := &Node{Val: val}

	if list.Head == nil {
		list.Head = NewNode
	} else {
		NewNode.Next = list.Head
		list.Head = NewNode
	}

}

// Add values to the end of the linked list

func (list *LinkedList) PushBack(val int) {
	NewNode := &Node{Val: val}
	if list.Head == nil {
		list.Head = NewNode
		return
	}

	current := list.Head

	for current.Next != nil {
		current = current.Next
	}
	current.Next = NewNode

}

// Delete the last occurrence of a given item in a singly linked list

func (list *LinkedList) PopBack() {
	if list.Head == nil || list.Head.Next == nil {
		list.Head = nil
		return
	}

	current := list.Head

	for current.Next.Next != nil {
		current = current.Next
	}

	current.Next = nil
}

// Find and Print the middle of a given singly linked list.
func (list *LinkedList) PrintMidle() {

	if list.Head == nil || list.Head.Next == nil {
		return
	}

	slow := list.Head
	fast := list.Head

	for fast != nil && fast.Next != nil {

		slow = slow.Next
		fast = fast.Next.Next
	}

	fmt.Println(slow.Val)

}

// Find and Delete the middle of a given singly linked list.
func (list *LinkedList) DeleteMidle() {

	if list.Head == nil || list.Head.Next == nil {
		return
	}

	slow := list.Head
	fast := list.Head
	var prev *Node
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	prev.Next = slow.Next

}

// delete all occuerence of the value
func (list *LinkedList) DltOccurence(val int) {
	current := list.Head
	var prev *Node

	for current != nil {
		if current.Val == val {
			if prev == nil {
				list.Head = current.Next
			} else {
				prev.Next = current.Next
			}
			current = current.Next
		} else {
			prev = current
			current = current.Next
		}
	}
}

// 2.Print the linked list
func (list *LinkedList) Print() {
	current := list.Head
	for current != nil {
		fmt.Printf("%d ->", current.Val)
		current = current.Next
	}
	fmt.Println("nil")

}

// 3.Reverse the linked list
func (list *LinkedList) ReverseList() {

	if list.Head == nil || list.Head.Next == nil {
		return
	}

	current := list.Head
	var prev *Node

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	list.Head = prev

}
    
 // find the nth node from the end of a singly linked list

func (list *LinkedList) FindNodeEnd(n int) int {
	refptr := list.Head
	mainptr := list.Head

	for i := 0; i < n; i++ {
		refptr = refptr.Next
		if refptr == nil {
			return -1
		}
	}

	for refptr != nil {
		refptr = refptr.Next
		mainptr = mainptr.Next
	}

	return mainptr.Val

}

func (list *LinkedList) DeleteDublicates() {
	current := list.Head

	for current != nil && current.Next != nil {
		if current.Next.Val == current.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
}

// Detect if a loop exists in a given singly linked list.

func (list *LinkedList) IsCyclic() bool {
	slow := list.Head
	fast := list.Head

	for fast != nil && fast.Next != nil {
		fast = fast.Next
		slow = slow.Next

		if slow == fast {
			return true
		}
	}
	return false
}

// Delete N nodes after skipping M nodes of a singly linked list.

func (list *LinkedList) DeleteNafterM(n, m int) {

	current := list.Head

	for current != nil {
		for i := 1; i < n; i++ {
			if current.Next == nil {
				return
			}
			current = current.Next
		}

		temp := current

		for j := 0; j < m; j++ {
			temp = temp.Next
		}

	}

}

func IsSorted(list1, list2 *Node) *Node {
	dummy := &Node{}
	current := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	for list1 != nil {
		current.Next = list1
		current = current.Next
		list1 = list1.Next
	}

	for list2 != nil {
		current.Next = list2
		current = current.Next
		list2 = list2.Next
	}
	return dummy.Next

}

func singlylis() {
	linkedlist := LinkedList{}
	linkedlist.PushBack(2)
	linkedlist.PushBack(5)
	linkedlist.PushBack(2)
	linkedlist.PushBack(5)
	linkedlist.PushBack(10)
	linkedlist.PushBack(2)
	linkedlist.PushBack(2)
	linkedlist.PushFront(100)
	// linkedlist.PopBack()
	// linkedlist.ReverseList()
	// linkedlist.Print()
	// linkedlist.DeleteMidle()
	// linkedlist.Print()
	// fmt.Println(linkedlist.FindNodeEnd(2))
	// linkedlist.DltOccurence(5)
	linkedlist.Print()

}

func main() {

	singlylis()

}
