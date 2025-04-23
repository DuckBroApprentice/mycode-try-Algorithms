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
	//case...two children
	{
		//predecesson
		temp, _ := t.Predecessor(x)
		curr := x.Parent
		if temp.Val > curr.Val {
			//curr.Right
			//處理temp原先的關係
			{
				temp_parent := temp.Parent
				temp_child := temp.Left //一定只有left，如果.right != nil 就找錯predecessor
				temp_parent.Left = temp_child
				temp_child.Parent = temp_parent
			}
			//temp調整位置後的關係
			{
				curr.Right = temp
				temp.Parent = curr
				temp.Right = x.Right
				child := x.Right
				child.Parent = temp
			}
		} else {
			//curr.left
			{
				temp_parent := temp.Parent
				temp_child := temp.Left //一定只有left，如果.right != nil 就找錯predecessor
				temp_parent.Left = temp_child
				temp_child.Parent = temp_parent
			}
			//temp調整位置後的關係
			{
				curr.Left = temp
				temp.Parent = curr
				temp.Left = x.Left
				child := x.Left
				child.Parent = temp
			}
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
