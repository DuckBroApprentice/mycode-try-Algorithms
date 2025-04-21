package classfour

//5-4*3+2
type treeNode struct {
	val         string
	left, right *treeNode
}

func expressionTree(expression []string) *treeNode {
	stack := []*treeNode{}
	curr := &treeNode{}
	for _, ele := range expression {
		if ele == "+" || ele == "-" || ele == "*" || ele == "/" {
			if len(stack) == 0 {
				curr.val = ele
				stack = append(stack, curr)
				//這時候左右兩邊都還無法決定 ... maybe
			}
		}
	}
}
