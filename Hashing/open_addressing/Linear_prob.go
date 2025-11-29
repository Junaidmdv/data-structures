package openaddressing

import "fmt"

const size = 10

type HashTableLinear struct {
	Items [size]*HashItems
}

type HashItems struct {
	Key string
	Val int
}

func Hash(key string) int {
	sum := 0
	for _, val := range key {
		sum += int(val)
	}
	return sum % size
}

func (h *HashTableLinear) Insert(key string, val int) {
	index := Hash(key)
	start := index

	for {
		if h.Items[index] == nil || h.Items[index].Key == key {
			h.Items[index] = &HashItems{Key: key, Val: val}
			return
		}
		index = (index + 1) % size
		if start == index {
			break
		}
	}
	fmt.Println("hash is filled")
}
func (h *HashTableLinear) Delete(key string, val int) {
	index := Hash(key)
	start := index

	for {
		if h.Items[index].Key == key {
			h.Items[index] = nil
			return
		}
		index = (index + 1) % size
		if start == index {
			break
		}
	}
	fmt.Println("key is not exist")
}
func (h *HashTableLinear) Get(key string, val int) (int, bool) {
	index := Hash(key)
	start := index

	for {
		if h.Items[index].Key == key {
			return h.Items[index].Val, true
		}
		index = (index + 1) % size
		if start == index {
			break
		}
	}
	return 0, false
}

func (h *HashTableLinear) Display() {
	for _, items := range h.Items {
		if items != nil {
			fmt.Printf("key is %s and value %v \n", items.Key, items.Val)
		}
	}
	fmt.Println()
}
