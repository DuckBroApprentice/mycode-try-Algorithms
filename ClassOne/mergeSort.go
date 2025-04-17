package classone

func mergeSort(nums []int) {
	if len(nums) <= 2 {
		return
	}
	mid := len(nums) / 2
	left := nums[:mid]
	right := nums[mid:]
	mergeSort(left)
	//結束切割後要做的邏輯：
	//兩兩排序
	//同樣的邏輯right也要做一遍
	//這裡是同時存在left及right，是一起處理嗎?
	//right還會進到下一個遞迴，只處理left
	for i := 0; i < len(left); i++ {

	}
	mergeSort(right)
}

/*
{
2,4,9,8,1,3,7,6,
	left{
		2,4,9,8
			left{
			2,4
			}
			right{
			9,8
			}
			len <= 2 return  不再切割
			在這一次遞迴中同時存在left跟right
			退出下一輪遞迴的是func(left)
			sort left
			sort right (8,9)
			merge left ,right (2,4,8,9)
			這個改動要能影響到原數據(nums)
			linked-list?
			{
			type ListNode struct {
				Val int
				Next *ListNode
			}
			要先歷遍長度 for head != nil {lenght++}
			if length <= 2 {
				return
			}
			for i:=0;i<lnght /2 ;i++{left}
			for i:=length /2;i<length;i++{right}
			func(left)

			func(right)
			}
			這樣處理每一次都會有一個O(n)遍歷長度
			再來切割*2  +起來也是一個O(n)
			一直到length <= 2時，分的層數會是O(logn)的時間複雜度
			[]*int可行嗎?
			{
			似乎很麻煩  要給每個地址都賦值 [2fac49] *[2fac49] = [1]
			好像想錯了
			已知array/slice 是一個連續的記憶體空間
			原始資料[]int
			每一個元素都有自己的記憶體位址
			想法不應該是[]*int  這個獨立的陣列操作並不會影響到原始陣列
			而是指向這個陣列中的元素修改值
			結論又回到鏈表了 XD

			另外有個想法  //我大概率是處理不了
			int64的記憶體分配是 size:8 align:8 (內存對齊)
			nums []int
			numspoint := &nums
			就能獲得nums的記憶體位址
			透過規則可以推算每個元素的記憶體位址
			不過其實也不知道能怎麼操作  想想就好
			}
	}
	right{
		1,3,7,6
			left{
			1,3
			}
			right{
			7,6
			}
			邏輯應該會跟處理left一樣
	}
}


*/
