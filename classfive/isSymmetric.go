//101. Symmetric Tree

package classfive

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func isSymmetric(root *TreeNode) bool {
// 	var stack []int
// 	preAndPostOrder(root, stack)
// 	for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
// 		if stack[i] != stack[j] {
// 			return false
// 		}
// 	}
// 	return true
// }

// func preAndPostOrder(root *TreeNode, stack []int) []int {
// 	if root == nil {
// 		stack = append(stack, 101)
// 		return stack
// 	}
// 	stack = append(stack, root.Val)
// 	stack = preAndPostOrder(root.Left, stack)
// 	stack = preAndPostOrder(root.Right, stack)
// 	stack = append(stack, root.Val)
// 	return stack
// }
// Runtime 89ms Beats 2.39%
// Memory 8.84MB Beats 4.38%
func isSymmetric(root *TreeNode) bool {
	if root.Left.Val != root.Right.Val {
		return false
	}
	return startLeft(root.Left) == startRight(root.Right)
}
func startLeft(root *TreeNode) int {
	if root == nil {
		return 101
	}
	startLeft(root.Left)
	startLeft(root.Right)
	return root.Val
}
func startRight(root *TreeNode) int {
	if root == nil {
		return 101
	}
	startRight(root.Right)
	startRight(root.Left)
	return root.Val
}
