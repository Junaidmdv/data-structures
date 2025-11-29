package openaddressing

import "testing"

func TestInsertLinear(t *testing.T){
    hash:=&HashTableLinear{}
	hash.Insert("junaid",20)
	hash.Insert("abhi",23)
	hash.Insert("babu",21)
	
}