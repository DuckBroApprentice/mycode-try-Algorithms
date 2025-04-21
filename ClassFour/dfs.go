package classfour

func PostorderTraversal(root *treeNode) []string {
	var res []string
	var postorder func(node *treeNode)
	postorder = func(node *treeNode) {
		if node == nil {
			return
		}
		postorder(node.left)
		postorder(node.right)
		res = append(res, node.val)
	}
	postorder(root)
	return res
}
