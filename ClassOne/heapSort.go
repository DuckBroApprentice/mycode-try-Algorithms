package classone

func HeapSort(nums []int) []int {
	root := []int{0}
	root = append(root, nums...)
	for i := len(root) - 1; i/2 > 0; i-- { //從1開始整理資料
		if root[i] < root[i/2] {
			root[i], root[i/2] = root[i/2], root[i]
		}
	}
	//root{0,heap shape}
	sorted := make([]int, 0, len(nums))
	for len(root) > 1 {
		sorted = append(sorted, root[1])
		root[1] = root[len(root)-1]
		root = root[:len(root)-1]
		for i := len(root) - 1; i/2 > 0; i-- { //其實只需要檢查1關聯的就行，只是一直寫不對
			if root[i] < root[i/2] {
				root[i], root[i/2] = root[i/2], root[i]
			}
		}
	}
	return sorted
}

/*
heap:
一坨數據，怎麼排的我不管
我需要兩個功能：
1.insert
2.extract min

如果 insert 時間複雜度為 O(1) 的話
extract min 的邏輯就會變的很重

反過來 如果 extract min 時間複雜度階 O(1) 的話
insert 的邏輯就會變重

目的：
insert -> O(logn) <- extract min

想法：
一個binary tree
對於每一個node來說，Left 及 Right 都要比 Val 大
                  1
			5           2
		9      6     3      7

Q：如果insert的數比最後一層都要小  似乎不符合規則？
A：跟上一層交換 直到符合規則
heap property才會對
shape property插到對的位置就會對


*/

/*
先實現array的heap
paretn == k , left = 2k right =2k+1
k => 2k 2k+1
*/
