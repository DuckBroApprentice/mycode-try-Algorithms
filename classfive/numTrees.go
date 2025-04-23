//leetcode  Q96. Unique Binary Search Trees

package classfive

//Catalan number C(n) =  C(n-1) * (4*n-2) /(n+1)
func numTrees(n int) int {
	if n == 0 {
		return 1
	}
	return numTrees(n-1) * (4*n - 2) / (n + 1)
	//偷看公式
	//Runtime 0ms
	//Memory 3.89MB
}

//難產
func numTrees2(n int) int {
	/*
		想像一串亂序的[1~n]的陣列
		這種陣列有所有n個元素的排列
		透過一個方法遍歷這個陣列會得到一個正序的結果 -> 合法 return ++
		如果得到的仍然不是正序 -> 不合法 return 0
		//可能可以return 1 || 0 把所有return加總
		//也可能可以return的值在遞迴中處理
		舉例來說，樹最差的情況：
		全走左：在i = 2^(層數-1) 時會發現數字
		全走右：在i = 2^層數 -1 時會發現數字
		array[1] 存在一個 array[2] 大於 array[1] 時 -> 不合法
		//似乎形成一個heap array
	*/
	/*
		1.產生所有陣列  猜測可以用回溯法完成 [n,n-1,n-2....1] n!數量的陣列
		2.提供一個處理陣列的func
	*/
}

func allArray(n int) []int {
	/*
		補充：這個array還要包含nil (用-1表示)
		[1,nil,2,nil,nil,nil,3......]
		[7,6,nil,5,nil,nil,nil,4,nil,nil,nil,....]
			Base Case：元素只剩下一個時，這個元素本身就是一個完整的排列
			剩餘元素取一個
			Append
			需要兩個陣列
			一個存放未使用的
			//[1,2,3,4] 拿走3 -> [1,2,4] 如何實現? [1,2,4]要遞迴 所以要能正確描述
			//[1,2,3,4] 拿走4 -> [1,2,3] [4] 拿走3 -> [1,2] [4,3] [1,2,(4] 拿走2 -> [1] [4,3,2] [1,(4,3] 拿走1 -> [] [4,3,2,1] [(4,3,2]
			//能正確傳入  但是判斷式會很複雜  而且太多切片
			一個往後增加
	*/

}
