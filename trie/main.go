package main

import "fmt"

type TrieNode struct {
	Child map[rune]*TrieNode
	IsEnd bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		Child: make(map[rune]*TrieNode),
		IsEnd: false,
	}
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

func (t *Trie) Insert(word string) {
	cur := t.root
	for _, ch := range word {
		if _, exist := cur.Child[ch]; !exist {
			cur.Child[ch] = NewTrieNode()
		}
		cur = cur.Child[ch]
	}
	cur.IsEnd = true
}
func (t *Trie) Search(word string) bool {
	cur := t.root
	for _, ch := range word {
		if _, exist := cur.Child[ch]; !exist {
			return false
		}
		cur = cur.Child[ch]
	}

	return cur.IsEnd
}

func main() {
	t := NewTrie()
	t.Insert("Hello")
	t.Insert("Hel")
	fmt.Println(t.Search("Hel"))
}
