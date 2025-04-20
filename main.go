package main

import "fmt"

// func main() {
// 	for i := 1; i <= 100; i++ {
// 		fmt.Printf("input i: %d,alive i: %d\n", i, classtwo.Josephus(i))
// 	}
// }

func nStep(n int) int {
	if n == 1 {
		return 1
	}
	return n * nStep(n-1)
}

func main() {
	a := nStep(9)
	fmt.Print(a)
}
