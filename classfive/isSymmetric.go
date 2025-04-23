//101. Symmetric Tree

package classfive

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
思路：
用一個陣列存左右半邊的pre+post order
陣列前->後：左半邊的pre  (會混到post)
陣列後->前：右半邊的post (同樣會混到pre)
雖然都有混到不需要的數據，但是 左 post 同樣也會是 右 pre  (陣列夾著看的話)
核心在比較左的pre跟右的post，且為了使陣列能真實反映樹，nil都補上101(-100<=val<=100)
*/

func isSymmetric(root *TreeNode) bool {
	var stack []int
	preAndPostOrder(root, stack)
	for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
		if stack[i] != stack[j] {
			return false
		}
	}
	return true
}

func preAndPostOrder(root *TreeNode, stack []int) []int {
	if root == nil {
		stack = append(stack, 101)
		return stack
	}
	stack = append(stack, root.Val)
	stack = preAndPostOrder(root.Left, stack)
	stack = preAndPostOrder(root.Right, stack)
	stack = append(stack, root.Val)
	return stack
}

// Runtime 89ms Beats 2.39%
// Memory 8.84MB Beats 4.38%
