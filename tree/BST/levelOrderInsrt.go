package main 



func (t *Node)LevelOrderInsertion(val int){
	NewNode:=&Node{Data: val}

 
	 queue:=[]*Node{t}

	 for len(queue) !=0{
		current:=queue[0]
        queue=queue[1:]

		if current.Left == nil{
			current.Left=NewNode
			return 
		}else{
            queue = append(queue, current.Left)
		}

		if current.Right == nil{
			current.Right=NewNode
			return 
		}else{
			 queue = append(queue, current.Right)
		}
		 
	 }
}