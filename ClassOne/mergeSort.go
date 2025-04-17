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
2,4,5,8,1,3,7,6,
	left{
		2,4,5,8
			left{
			2,4
			}
			right{
			5,8
			}
			len <= 2 return

	}
	right{
		1,3,7,6
			left{
			1,3
			}
			right{
			7,6
			}
	}
}


*/
