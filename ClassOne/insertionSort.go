package classone

func InsertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		for j, k := i-1, i; j >= 0; j, k = j-1, k-1 {
			if nums[k] > nums[j] {
				break
			}
			nums[j], nums[k] = nums[k], nums[j]
		}
	}
}
