package main

import "errors"

type StackList struct {
	Head *Node
}

type Node struct {
	Item int
	Next *Node
}



func (list *StackList)Push(val int){
	 NewNode:=&Node{Item: val}
	if list.Head==nil{
		 list.Head=NewNode
	}else{
		NewNode.Next=list.Head
		list.Head=NewNode
	}
} 


func (list *StackList)Pop()(int,error){
	 val:=0
	 if list.Head==nil{
        return val,errors.New("empty stack")
	 }else{
		val=list.Head.Item
		list.Head=list.Head.Next
	 }
	 return val,nil
}


func (list *StackList)Peak()(int,error){
	if list.Head==nil{
		return 0,errors.New("empty stack")
	}

	return list.Head.Item,nil
}

func (list *StackList)IsEmpty()bool{
	return list.Head ==nil
}