package parent

//多了Parent才能方便做Update
//leetcode沒試過
type treeNode struct {
	Val                 int
	Parent, Left, Right *treeNode
}

func (t *treeNode) Insert(val int) *treeNode {
	if t.Right == nil && t.Val < val {
		newNode := &treeNode{Val: val, Parent: t}
		t.Right = newNode
		return t.Right
	} else if t.Left == nil && t.Val > val {
		newNode := &treeNode{Val: val, Parent: t}
		t.Left = newNode
		return t.Left
	}
	if t.Val < val {
		return t.Right.Insert(val)
	} else if t.Val > val {
		return t.Left.Insert(val)
	}
	return t
}
