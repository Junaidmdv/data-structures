package main

import "fmt"

type Stack struct {
	Item []int
}

func (s *Stack) Push(val int) {
	s.Item = append(s.Item, val)

}

func (s *Stack) Pop() int {
	if len(s.Item) == 0 {
		fmt.Println("empty stack")
		return -1
	}
	l := len(s.Item) - 1
	poped_val := s.Item[l]
	s.Item = s.Item[:l]
	return poped_val
}

func (s *Stack) Peek() int {
	return s.Item[len(s.Item)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.Item) == 0
}

func DeleteMidle(s *Stack, current, size int) {
	if current == size/2 {
		s.Pop()
		return
	}
	top := s.Peek()
	s.Pop()
	DeleteMidle(s, current+1, size)

	s.Push(top)

}

func main() {
	s := &Stack{}

	for i := 1; i < 10; i++ {
		s.Push(i)
	}

	// delte midle element of stack
	DeleteMidle(s, 0, len(s.Item))
	fmt.Println(s.Item)
	//queue using stack

	Queue := &Stack{}
	for i := 1; i < 10; i++ {
		Enqueue(Queue, i)
	}
	fmt.Println(Queue.Item)

	Dequeue(s)
	fmt.Println(s.Item)

	// sort the stack using temp stack
	unsorted := &Stack{}

	unsorted.Push(23)
	unsorted.Push(43)
	unsorted.Push(13)
	unsorted.Push(33)
	unsorted.Push(93)
	unsorted.Push(3)

	sortedStack := SortStack(unsorted)

	fmt.Println("sorted stack", sortedStack.Item)

	// check the give paranthis is balanced or not
	// true condition
	fmt.Println("Is balanced:", IsBalanced("{{}}[]"))
	// false condition
	fmt.Println("Is balanced:", IsBalanced("{{}["))

}
