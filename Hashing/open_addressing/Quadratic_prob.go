package openaddressing

import "fmt"

const sizeq = 10

type HashTable struct {
	Items [sizeq]*HashItemsq
}

type HashItemsq struct {
	Val int
	Key string
}

func HashQuadratic(Key string) int {
	sum := 0
	for _, val := range Key {
		sum += int(val)
	}
	return sum % sizeq
}

func (h *HashTable) Insert(key string, val int) {

	index := HashQuadratic(key)

	for i := 0; i < sizeq; i++ {
		newindex := (index + i*i) % size
		if h.Items[newindex] == nil || h.Items[newindex].Key == key {
			h.Items[newindex] = &HashItemsq{Val: val, Key: key}
			return
		}
	}
	fmt.Println("Hash table is filled")
}

func (h *HashTable) Delete(key string, val int) {

	index := HashQuadratic(key)

	for i := 0; i < sizeq; i++ {
		newindex := (index + i*i) % size
		if h.Items[newindex].Key == key {
			h.Items[newindex] = nil
			return
		}
	}
	fmt.Println("key is not exist")
}

func (h *HashTable) Get(key string) (int, bool) {

	index := HashQuadratic(key)

	for i := 0; i < sizeq; i++ {
		newindex := (index + i*i) % size
		if h.Items[newindex].Key == key {
			return h.Items[newindex].Val, true

		}
	}
	return 0, false
}

func (h *HashTable) Display() {
	for _, val := range h.Items {
		if val != nil {
			fmt.Printf("key is %s values is %d \n", val.Key, val.Val)
		}
	}
}
