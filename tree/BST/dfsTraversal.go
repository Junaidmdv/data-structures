package main

import "fmt"




func (t *Node) InOrder() {
	if t == nil {
		return
	}

	t.Left.InOrder()
	fmt.Println(t.Data)
	t.Right.InOrder()

}

func(t *Node) NonDecreasingOrder(){

	if t==nil{
		return 
	}

	t.Right.NonDecreasingOrder()
	 fmt.Println(t.Data)
	t.Left.NonDecreasingOrder()

}


func (t *Node) PostOrder() {
  if t==nil{
	return 
  }

    t.Left.PostOrder()
	t.Right.PostOrder()
	fmt.Println(t.Data)
	
}

func (t *Node) PreOrder() {

	if t==nil{
		return 
	}
	
	fmt.Println(t.Data)
	t.Left.PreOrder()
	t.Right.PreOrder()

}

