package seperatechaining

import "fmt"

const size = 10

type HashTable struct {
	Item [size]*Bucket
}

type Bucket struct {
	Head *BucketNode
}

type BucketNode struct {
	Key  string
	Val  int
	Next *BucketNode
}

func Hash(key string) int {
	sum := 0

	for _, val := range key {
		sum += int(val)
	}
	return sum % size
}

func Init() *HashTable {
	hash := &HashTable{}
	for i := range hash.Item {
		hash.Item[i] = &Bucket{}
	}
	return hash
}

func (h *HashTable) Insert(key string, val int) {
	index := Hash(key)
	h.Item[index].insert(key, val)
}

func (b *Bucket) insert(key string, val int) {
	NewNode := &BucketNode{Val: val, Key: key}

	if b.Head == nil {
		b.Head = NewNode
		return
	}
	current := b.Head

	for current.Next != nil {
		current = current.Next
	}
	current.Next = NewNode

}

func (h *HashTable) Search(key string, val int) (int, bool) {
	index := Hash(key)
	return h.Item[index].search(key)
}

func (b *Bucket) search(key string) (int, bool) {
	current := b.Head

	for current != nil {
		if current.Key == key {
			return current.Val, true
		}
		current = current.Next
	}

	fmt.Println("key is not exist")
	return 0, false
}

func (h *HashTable) Delete(key string) {
	index := Hash(key)
	h.Item[index].delete(key)
}

func (b *Bucket) delete(key string) {
	if b.Head.Key == key {
		b.Head.Next = b.Head
	}
	prev := b.Head
	current := b.Head

	for current != nil {
		if current.Key == key {
			prev.Next = current.Next
			fmt.Println("value deleted")
			return
		}
		prev = current

		current = current.Next
	}

}
