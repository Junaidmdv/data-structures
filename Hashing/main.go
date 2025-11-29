package main

import (
	openaddressing "GO_BOOKSTORE/dsa-revision/Hashing/open_addressing"
)

func main() {
	hash := openaddressing.HashTable{}
	hash.Insert("junaid", 23)
	hash.Insert("junaid", 21)
	hash.Insert("abhi", 29)
	hash.Insert("shambu", 23)
	hash.Insert("geetha", 99)
	hash.Display()



}
