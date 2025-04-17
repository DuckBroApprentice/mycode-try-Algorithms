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
			要先歷遍長度 for head != nil {lenght++}
			for i:=0;i<lnght /2 ;i++{left}
			for i:=length /2;i<length;i++{right}
			}
			[]*int可行嗎?
			{
			似乎很麻煩  要給每個地址都賦值 [2fac49] *[2fac49] = [1]
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
