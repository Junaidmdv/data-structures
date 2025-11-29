package main

import "fmt"

type StackStr struct {
	Item []rune
}

func (s *StackStr) Push(val rune) {

	s.Item = append(s.Item, val)

}

func (s *StackStr) Pop() rune {
	l := len(s.Item) - 1

	val := s.Item[l]

	s.Item = s.Item[:l]

	return val

}

func (s *StackStr) IsEmpty() bool {
	return len(s.Item) == 0

}

func (s *StackStr) Top() rune {
	l := len(s.Item) - 1
	if len(s.Item) == 0 {
		fmt.Println("empty stack")
		return -1
	}
	return s.Item[l]

}

func IsBalanced(str string) bool {
	s := &StackStr{}

	for _, val := range str {
		if val == '{' || val == '[' || val == '(' {
			s.Push(val)
		} else {
			if s.IsEmpty() {
				return false
			}

			if val == '}' && s.Top() != '{' ||
				val == ']' && s.Top() != '[' ||
				val == ')' && s.Top() != '(' {
				return false
			}
			s.Pop()

		}

	}
	return s.IsEmpty()

}
