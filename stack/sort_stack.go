package main

func SortStack(s *Stack) *Stack {
	tempStack := &Stack{}

	for !s.IsEmpty() {
		temp := s.Peek()
		s.Pop()

		for !tempStack.IsEmpty() && tempStack.Peek() > temp {
			s.Push(tempStack.Peek())
			tempStack.Pop()
		}

		tempStack.Push(temp)
	}

	return tempStack
}
