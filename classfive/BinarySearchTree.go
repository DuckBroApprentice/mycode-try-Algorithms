package classfive

import "fmt"

type treeNode struct {
	Val         int
	Left, Right *treeNode
}

func (t *treeNode) Insert(target int) *treeNode {
	if t == nil {
		return &treeNode{Val: target}
	}
	if target > t.Val {
		return t.Right.Insert(target)
	} else if target < t.Val {
		return t.Left.Insert(target)
	}
	//不重覆
	return t
}

func (t *treeNode) FindMin(tree *treeNode) (int, error) {
	if t == nil {
		return 0, fmt.Errorf("tree has no node")
	}
	for tree.Left != nil {
		tree = tree.Left
	}
	return tree.Val, nil
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

//DFS
func DFSSearch(tree *treeNode, target int) bool {
	if tree == nil {
		return false
	}
	if tree.Val == target {
		return true
	}
	return DFSSearch(tree.Left, target) || DFSSearch(tree.Right, target)
}

// func Search(tree *treeNode, target int) bool {
// 	if tree == nil {
// 		fmt.Println("target in not exists: ", target)
// 		return false
// 	}
// 	if tree.Val == target {
// 		fmt.Println("target is exists: ", target)
// 		return tree.Val == target
// 	}
// 	if target > tree.Val {
// 		return Search(tree.Right, target)
// 	}
// 	if target < tree.Val {
// 		return Search(tree.Left, target)
// 	}
// 	return tree.Val == target
// }
func (t *treeNode) Search(target int) bool {
	if t == nil {
		fmt.Println("target in not exists: ", target)
		return false
	}
	if t.Val == target {
		fmt.Println("target is exists: ", target)
		return t.Val == target
	}
	if target > t.Val {
		return t.Right.Search(target)
	}
	if target < t.Val {
		return t.Left.Search(target)
	}
	return t.Val == target
}
