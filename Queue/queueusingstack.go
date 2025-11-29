package main

func Push(st1 *QueueList, Val int) {
	temp := QueueList{}

	temp.EnQueue(Val)

	for !st1.IsEmpty() {
		temp.EnQueue(st1.DeQueue())
	}

	*st1 = temp
}

func Pop(st1 *Queue) int {
	val := st1.DeQueue()
	return val
}
