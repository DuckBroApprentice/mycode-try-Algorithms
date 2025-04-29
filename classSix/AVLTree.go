package classsix

//簡單處理，val不重覆
type AVLTreeNode struct {
	Val                 int
	Height              int
	Parent, Left, Right *AVLTreeNode
}

//nodeHeight Insert Delete(findMax/findMin) Rotation Search

func NewAVLTreeNode() *AVLTreeNode {
	return &AVLTreeNode{}
}

func (at *AVLTreeNode) GetHeight() int {
	if at == nil {
		return 0
	}
	return at.Height
}

// 插入的node並不會事先知道他會被插入哪層
func (at *AVLTreeNode) SetHeight() {
	at.Height = 1 + max(at.Left.GetHeight(), at.Right.GetHeight())
}

func (at *AVLTreeNode) rebalanceTree() *AVLTreeNode {
	if at == nil {
		return nil
	}
	at.SetHeight()
	//check balance factor(bf)
	bf := at.Left.GetHeight() - at.Right.GetHeight()
	if bf == -2 {
		//右邊重
		//RL
		if at.Right.Left.GetHeight() > at.Right.Right.GetHeight() {
			at.Right = at.Right.rotateRight()
		}
		return at.rotateLeft()
	} else if bf == 2 {
		//左邊重
		//LR
		if at.Left.Right.GetHeight() > at.Left.Left.GetHeight() {
			at.Left = at.Left.rotateLeft()
		}
		return at.rotateRight()
	}
	//bf 1 0 -1
	return at
}

func (at *AVLTreeNode) rotateLeft() *AVLTreeNode {
	temp := at.Right
	at.Right = temp.Left
	temp.Left.Parent = at
	temp.Left = at
	at.Parent = temp

	at.SetHeight()
	temp.SetHeight()
	return temp
}
func (at *AVLTreeNode) rotateRight() *AVLTreeNode {
	temp := at.Left
	at.Left = temp.Right
	temp.Right.Parent = at
	temp.Right = at
	at.Parent = temp

	at.SetHeight()
	temp.SetHeight()
	return temp
}

func (at *AVLTreeNode) Search(target *AVLTreeNode) *AVLTreeNode {
	return at.search(target)
}

func (at *AVLTreeNode) search(target *AVLTreeNode) *AVLTreeNode {
	if at == nil {
		return nil
	}
	if target.Val < at.Val {
		return at.Left.search(target)
	} else if target.Val > at.Val {
		return at.Right.search(target)
	} else {
		return at
	}
}

// 方便起見，插入失敗的不回傳error，直接回傳原樹就好(ex:target不存在)
//how to handle BF(Balance Factory)??? ==>  左右高度、遍歷、rotation
//怎麼遍歷並找到最底層失衡的node ???? ==> 從Insert的node.parent調用BF
func (at *AVLTreeNode) Insert(newNode *AVLTreeNode) *AVLTreeNode {
	if at == nil {
		newNode.Height = newNode.GetHeight() //leaf height = 1
		return newNode
	}
	if newNode.Val < at.Val {
		at.Left = at.Left.Insert(newNode)
		newNode.Parent = at
	} else if newNode.Val > at.Val {
		at.Right = at.Right.Insert(newNode)
		newNode.Parent = at
	} else {
		at = newNode
	}
	return at.rebalanceTree()
}

func (at *AVLTreeNode) Remove(target *AVLTreeNode) *AVLTreeNode {
	if at == nil {
		return nil
	}
	if target.Val < at.Val {
		at.Left = at.Left.Remove(target)
	} else if target.Val > at.Val {
		at.Right = at.Right.Remove(target)
	} else {
		//遍歷後找到目標node
		//跟BST一樣處理3種情況：leaf , one chile, two children
		if at.Left != nil && at.Right != nil {
			//左樹最大值* 或 右樹最小值 邏輯同BST
			//finish findMax before
			leftMax := at.findMax()
			//目前有at(現在是target) leftMax 兩個節點
			at.Val = leftMax.Val
			//at.Parent no chage
			//at.Left & at.Right return rebalance (need rotate)
			at.Left = at.Left.Remove(at.Left)
		} else if at.Left != nil {
			at.Left.Parent = at.Parent
			at = at.Left
		} else if at.Right != nil {
			at.Right.Parent = at.Parent
			at = at.Right
		} else {
			at = nil
			return at //no height-change happened
		}
	}
	return at.rebalanceTree()
}

//findMax
func (at *AVLTreeNode) findMax() *AVLTreeNode {
	if at.Left != nil {
		return at.Left.findMax()
	} else {
		return at
	}
}
