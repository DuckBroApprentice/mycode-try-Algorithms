package other

//網友聊到費波那契數列，提到資工系居然不知道這個，就試試看

//遞迴
func Fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n > 0 && n < 2 {
		return 1
	}

	sum := 0
	sum = Fibonacci(n-1) + Fibonacci(n-2)
	return sum
}

//Array
func FibonacciArray(n int) []int {
	res := make([]int, 0, n)
	sum := 0
	for i := 0; i < n; i++ {
		if i < 2 {
			res = append(res, 1)
			continue
		}
		sum = res[i-2] + res[i-1]
		res = append(res, sum)
	}
	return res
}
