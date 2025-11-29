package main

func Enqueue(s *Stack, val int) {
	s.Push(val)
}

func Dequeue(s *Stack) {
	temp := &Stack{}

	for !s.IsEmpty() {
		temp.Push(s.Peek())
		s.Pop()
	}

	temp.Pop()

	for !temp.IsEmpty() {
		s.Push(temp.Peek())
		temp.Pop()
	}

}
