package classfour

type Node struct {
	Val       int
	Neighbors []*Node
}

var res []*Node

func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}
	//要指到新的記憶體位址
	//用key:val記錄(map)
	data := visite(node)
	cur := &Node{}
	result := cur
	for i, nodeArray := range data {
		if i == 0 {
			continue
		}
		cur.Val = nodeArray.Val
		cur.Neighbors = nodeArray.Neighbors
	}
	return result
}

// 遍歷node
func visite(node *Node) []*Node {
	newNode := &Node{}
	newNode.Val = node.Val
	newNode.Neighbors = node.Neighbors
	// res := make([]*Node, 0, 101)
	if node.Val == 0 {
		return res
	}
	res = append(res, newNode) //順序會一致
	node.Val = 0
	for i := 0; i < len(node.Neighbors); i++ {
		node = node.Neighbors[i]
		visite(node)
	}
	return res
}

//用map處理的clone會亂序
//用array處理要怎麼刷新記憶體位址
// node is : &{1 [0xc0000b4020 0xc0000b4080]}
// [0xc0000b4000 0xc0000b4020 0xc0000b4080 0xc0000b4040] 存到同樣的位址
