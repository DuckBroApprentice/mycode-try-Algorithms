package parent

import "fmt"

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

func (t *treeNode) Search(x *treeNode) *treeNode {
	if t == nil {
		fmt.Println("node in not exists: ", x)
		return nil
	}
	if t.Val == x.Val {
		fmt.Println("node is exists: ", x)
		return x
	}
	if x.Val > t.Val {
		return t.Right.Search(x)
	}
	if x.Val < t.Val {
		return t.Left.Search(x)
	}
	return x
}

func (t *treeNode) FindMax(tree *treeNode) (int, error) {
	if t == nil {
		return 0, fmt.Errorf("tree has no node")
	}
	for tree.Right != nil {
		tree = tree.Right
	}
	return tree.Val, nil
}

//節點左邊的最大值  //Successor則是節點右邊的最小值 (懶的做...
func (t *treeNode) Predecessor(x *treeNode) *treeNode { //這個val是樹的某個節點
	if x.Left != nil {
		return t.FindMax(x.Left)
	}
	if x.Left == nil {
		for {
			curr := x.Parent
			if curr == nil {
				return nil
			}
			if curr.Val < x.Val {
				return curr
			}
			x = curr
		}
	}
	return nil //x本身就是最小值
}
