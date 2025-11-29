package main



// here comparision with parent the node.left side always  <= max and right side always >= min, 

func ValidateBST(t *Node, min *Node, max *Node) bool {
	if t == nil {
		return true
	}
	if min != nil && min.Data >= t.Data {
		return false
	}
	if max != nil && max.Data <= t.Data {
		return false
	}

	return ValidateBST(t.Left, min, t) && ValidateBST(t.Right, t, max)
}
