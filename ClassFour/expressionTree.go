package classfour

import "fmt"

//5-4*3+2
type treeNode struct {
	val         string
	left, right *treeNode
}

func ExpressionTree(expression []string) *treeNode {
	stackNum := []*treeNode{}
	stackSym := []*treeNode{}
	root := &treeNode{}
	//build tree first
	//expression 本身可以看成一個preorder
	//課題：將preorde改成postorder
	/*思路：
	5-4*3+2
	stack num {5 , 4 ,3 }
	stack symbol{- * + }  優先級處理+num有2個value
	5
	-
	5 4
	- *
	5 4 3
	※- * len(stackNum)>2 stackSym-pop(*) stackNum-insert(*) stackSym-pop(-) stackNum-insert(-) stackSym-insert(+)
	※=> 5 * ==> -
	※- * ==> - ==> +
	- 2
	finish range clean stacks
	*/
	for i, ele := range expression {

		if ele != "+" && ele != "-" && ele != "*" && ele != "/" {
			//curr.val = ele
			curr := &treeNode{}
			curr.val = ele
			curr.left = nil
			curr.right = nil
			stackNum = append(stackNum, curr)
			//數字一定會是葉子
			if i == len(expression)-1 {
				//一定是數字前面已經append到stackNum
				//並且處理的只有加減乘除 不會有超過1個符號保留在stackSym內
				stackSym[len(stackSym)-1].right = curr
				stackNum = stackNum[:len(stackNum)-1]
				stackSym[len(stackSym)-1].left = stackNum[len(stackNum)-1]
				stackNum = stackNum[:len(stackNum)-1]
				root = stackSym[len(stackSym)-1]
				stackSym = stackSym[:len(stackSym)-1]
			}
			fmt.Println("stackNum now is : ", stackNum, ele)
		}
		if ele == "+" || ele == "-" || ele == "*" || ele == "/" {
			//curr.val = ele
			curr := &treeNode{}
			curr.val = ele
			if len(stackSym) == 0 {
				stackSym = append(stackSym, curr)
				fmt.Println("stackSym now is :", stackSym, ele)
				continue
			}
			if ele == "*" || ele == "/" {
				if stackSym[len(stackSym)-1].val == "*" || stackSym[len(stackSym)-1].val == "/" {
					stackSym[len(stackSym)-1].right = stackNum[len(stackNum)-1]
					stackNum = stackNum[:len(stackNum)-1]
					stackSym[len(stackSym)-1].left = stackNum[len(stackNum)-1]
					stackNum = stackNum[:len(stackNum)-1]
					stackNum = append(stackNum, stackSym[len(stackSym)-1])
					stackSym = stackSym[:len(stackSym)-1]
					stackSym = append(stackSym, curr)
					fmt.Println("stackSym now is :", stackSym, ele)
				} else {
					stackSym = append(stackSym, curr)
					fmt.Println("stackSym now is :", stackSym, ele)
				}
			}
			if ele == "+" || ele == "-" {
				for len(stackSym) > 0 {
					stackSym[len(stackSym)-1].right = stackNum[len(stackNum)-1]
					stackNum = stackNum[:len(stackNum)-1]
					stackSym[len(stackSym)-1].left = stackNum[len(stackNum)-1]
					stackNum = stackNum[:len(stackNum)-1]
					stackNum = append(stackNum, stackSym[len(stackSym)-1])
					stackSym = stackSym[:len(stackSym)-1]
					fmt.Println("stackSym now is :", stackSym, ele)
				}
				stackSym = append(stackSym, curr)
			}
		}

	}
	return root
}

/*if len(stackNum) > 2 { // len(stackNum) >= len(stackSym) is wrong
			if ele == "+" || ele == "-" { //優先度最低 不管前面存在什麼都是後算
				if len(stackSym) > 0 {
					stackSym[len(stackSym)-1].right = stackNum[len(stackNum)-1]
					stackNum = stackNum[:len(stackNum)-1]
					stackSym[len(stackSym)-1].left = stackNum[len(stackNum)-1]
					stackNum = stackNum[:len(stackNum)-1]
					stackNum = append(stackNum, stackSym[len(stackSym)-1])
					stackSym = stackSym[:len(stackSym)-1]
					stackSym = append(stackSym, curr)
				}
			} else if ele == "*" || ele == "/" {
				if stackSym[len(stackSym)-1].val == "*" || stackSym[len(stackSym)-1].val == "/" {
					stackSym[len(stackSym)-1].right = stackNum[len(stackNum)-1]
					stackNum = stackNum[:len(stackNum)-1]
					stackSym[len(stackSym)-1].left = stackNum[len(stackNum)-1]
					stackNum = stackNum[:len(stackNum)-1]
					stackNum = append(stackNum, stackSym[len(stackSym)-1])
					stackSym = stackSym[:len(stackSym)-1]
					stackSym = append(stackSym, curr)
				} else {
					curr.right = stackNum[len(stackNum)-1]
					stackNum = stackNum[:len(stackNum)-1]
					curr.left = stackNum[len(stackNum)-1]
					stackNum = stackNum[:len(stackNum)-1]
					stackNum = append(stackNum, curr)
				}
			}
		} else {
			stackSym = append(stackSym, curr)
		}
	}
}
for len(stackSym) > 0 {
	stackSym[len(stackSym)-1].right = stackNum[len(stackNum)-1]
	stackNum = stackNum[:len(stackNum)-1]
	stackSym[len(stackSym)-1].left = stackNum[len(stackNum)-1]
	stackNum = stackNum[:len(stackNum)-1]
	stackSym = stackSym[:len(stackSym)-1]
	if len(stackSym) == 1 {
		root = stackSym[len(stackSym)-1]
	}
}
return root
*/
