package main

import "fmt"

type CircularQueue struct {
	Data         []int
	Front        int
	Rear         int
	Capacity     int
	CurrrentSize int
}

func NewCircularQueue(cap int) *CircularQueue {
	return &CircularQueue{
		Data:         make([]int, cap),
		Front:        0,
		Rear:         -1,
		Capacity:     cap,
		CurrrentSize: 0,
	}
}

func (c *CircularQueue) Insert(val int) {
	if c.Capacity == c.CurrrentSize {
		fmt.Println("can't add value,que is full")
		return
	}
	c.Rear = (c.Rear + 1) % c.Capacity
	c.Data[c.Rear] = val
	c.CurrrentSize++

}


func (c *CircularQueue) Pop() int {
	if c.Capacity == 0 {
		fmt.Println("can't add value,que is full")
		return 0
	}
	val := c.Data[c.Front]
	c.Front = (c.Front + 1) % c.Capacity
	c.CurrrentSize--
	return val

}

func (c *CircularQueue) FrontVal() int {
	return c.Data[c.Front]
}

func (c *CircularQueue) RearVal() int {
	return c.Data[c.Rear]
}

func (c *CircularQueue) Print() {
	st := c.Front
	for {
		fmt.Printf("%d\n", c.Data[st])
		if st == c.Rear {
			break
		}
		st = (st + 1) % c.Capacity
	}
}

func main() {
	c := NewCircularQueue(5)

	for i := 0; i < 5; i++ {
		c.Insert(i)
	}
	c.Pop()
	c.Pop()
	c.Print()
}
