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

func (t *treeNode) Del(x *treeNode) *treeNode {
	//case root
	if x == t {
		if x.Left == nil && x.Right == nil {
			return nil
		}
		if x.Left == nil && x.Right != nil {
			t = t.Right
			t.Parent = nil
			return t
		} else if x.Left != nil && x.Right == nil {
			t = t.Left
			t.Parent = nil
			return t
		}
		if x.Left != nil && x.Right != nil {
			pre, _ := t.FindMax(x.Left)
			if pre == x.Left {
				pre.Right = x.Right
				x.Right.Parent = pre
				pre.Parent = nil
				x = nil
				return t
			}
			if pre.Left != nil {
				pre_parent := pre.Parent
				pre_child := pre.Left
				pre_parent.Right = pre_child
				pre_child.Parent = pre_parent
			}
			pre.Left = x.Left
			x.Left.Parent = pre
			pre.Right = x.Right
			x.Right.Parent = pre
			x = nil
			return t
		}
	}
	//case1...no child
	if x.Left == nil && x.Right == nil {
		curr := x.Parent
		//at curr left or right?
		if curr.Val < x.Val {
			//right
			curr.Right = nil
		} else {
			curr.Left = nil
		}
		return t
	}
	//case2...one child only
	if x.Left != nil || x.Right != nil { //上面的if已經排除都nil的情況
		if x.Left == nil {
			//子節點right
			curr := x.Parent
			if curr.Val > x.Val {
				//在parent的左邊
				curr.Left = x.Right
			} else {
				curr.Right = x.Right
			}
		} else {
			curr := x.Parent
			if curr.Val > x.Val {
				//curr.left
				curr.Left = x.Left
			} else {
				curr.Left = x.Left
			}
		}
		return t
	}
	//case3...two children
	//有很多情況要處理 修正版
	{
		//case3-1 delete leaf
		if x.Left == nil && x.Right == nil {
			curr := x.Parent
			if curr.Left == x {
				curr.Left = nil
				x.Parent = nil
				return t
			} else if curr.Right == x {
				curr.Right = nil
				x.Parent = nil //x才會被GC
				return t
			}
		}
		//case3-2 one child only
		if x.Left == nil && x.Right != nil {
			//x == paretn.Right
			curr := x.Parent
			// curr.where = x.here  (right)
			if curr.Left == x {
				curr.Left = x.Right
				x.Right.Parent = curr
			} else if curr.Right == x {
				curr.Right = x.Right
				x.Right.Parent = curr
			}
			return t
		} else if x.Left != nil && x.Right == nil {
			curr := x.Parent
			if curr.Left == x {
				curr.Left = x.Left
				x.Left.Parent = curr
			} else if curr.Right == x {
				curr.Right = x.Left
				x.Left.Parent = curr
			}
			return t
		}
		//case3-2 2 children
		curr := x.Parent
		//先決定which node (pre)取代x
		//pre方案用 FindMax(x.Left)處理
		pre, _ := t.FindMax(x.Left)
		//再維護node的屬性
		//x提供了 parent left right的指針先保留 用完再指向nil 觸發GC
		//pre替代x原先的位置 要繼承x的parent left right
		//pre原先位置的關係也要處理
		//1...pre是個leaf
		if pre.Left == nil && pre.Right == nil {
			pre_parent := pre.Parent
			pre_parent.Right = nil
			pre.Parent = curr
			pre.Left = x.Left
			pre.Right = x.Right
			x.Parent, x.Right, x.Left = nil, nil, nil //GC
			return t
		}

		//2...pre has a child (pre.Left)
		//pre = FindMax(x.Left) ===> pre.Right == nil
		if pre.Left != nil {
			pre_parent := pre.Parent
			pre_child := pre.Left
			pre_parent.Right = pre_child
			pre_child.Parent = pre_parent
			curr := x.Parent
			pre.Parent = curr
			pre.Left = x.Left
			pre.Right = x.Right
			x.Parent, x.Right, x.Left = nil, nil, nil //GC
			return t
		}
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

func (t *treeNode) FindMax(tree *treeNode) (*treeNode, error) {
	if t == nil {
		return nil, fmt.Errorf("tree has no node")
	}
	for tree.Right != nil {
		tree = tree.Right
	}
	return tree, nil
}

//節點左邊的最大值  //Successor則是節點右邊的最小值 (懶的做...
func (t *treeNode) Predecessor(x *treeNode) (*treeNode, error) { //這個val是樹的某個節點
	if x.Left != nil {
		return t.FindMax(x.Left)
	}
	if x.Left == nil {
		for {
			curr := x.Parent
			if curr == nil {
				return nil, fmt.Errorf("not exists")
			}
			if curr.Val < x.Val {
				return curr, nil
			}
			x = curr
		}
	}
	return nil, fmt.Errorf("input is min") //x本身就是最小值
}
