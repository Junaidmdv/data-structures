package main

func (t *Node) SearchNode(key int) bool {
	if t == nil {
		return false
	}

	if t.Data == key {
		return true
	}

	if key < t.Data {
		return t.Left.SearchNode(key)
	} else {
		return t.Right.SearchNode(key)
	}

}

func DeleteNode(t *Node, key int) *Node {
	if key < t.Data {
		t.Left = DeleteNode(t.Left, key)
	} else if key > t.Data {
		t.Right = DeleteNode(t.Right, key)
	} else {

		if t.Right == nil {
			return t.Left
		} else if t.Left == nil {
			return t.Right
		} else {

			Is := InOrderSuccessor(t.Right)
			t.Data = Is.Data
			t.Right = DeleteNode(t.Right, key)

		}
	}
	return t
}

func InOrderSuccessor(t *Node) *Node {
	if t != nil && t.Left != nil {
		t = t.Left
	}
	return t
}
