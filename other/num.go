package other

//嘗試實現俄式乘法
//先找到n*2 或 n/2 在二進制的表現
//剛好是<<>>的位移
//&1運算後就能判斷奇、偶

import "fmt"

func ToBinary(n int) string {
	if n/2 == 0 {
		return fmt.Sprint(n % 2)
	}
	return ToBinary(n/2) + fmt.Sprint(n%2)
}

func Multi(n1, n2 int) int {
	if n1 == 0 || n2 == 0 {
		return 0
	}
	//n1 * n2
	var n1Array []int
	var n2Array []int
	for n1 > 0 {
		n1Array = append(n1Array, n1)
		n1 = n1 / 2
	}
	length := len(n1Array)
	for i := 0; i < length; i++ {
		n2Array = append(n2Array, n2)
		n2 = n2 * 2
	}
	// for i, num := range n1Array {
	// 	if num%2 == 0 {
	// 		n2Array[i] = 0
	// 	}
	// }
	sum := 0
	// for i := 0; i < len(n2Array); i++ {
	// 	sum += n2Array[i]
	// }
	for i := 0; i < length; i++ {
		if n1Array[i]%2 == 0 {
			continue
		}
		sum += n2Array[i]
	}
	return sum
}
