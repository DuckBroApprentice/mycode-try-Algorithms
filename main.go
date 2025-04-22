package main

import (
	"Algorithms/other"
	"fmt"
)

// func main() {
// 	for i := 1; i <= 100; i++ {
// 		fmt.Printf("input i: %d,alive i: %d\n", i, classtwo.Josephus(i))
// 	}
// }

// func nStep(n int) int {
// 	if n == 1 {
// 		return 1
// 	}
// 	return n * nStep(n-1)
// }

//	func main() {
//		a := nStep(9)
//		fmt.Print(a)
//	}
// func main() {
// 	var s1 struct{}
// 	var s2 struct{}
// 	var a1 [0]int
// 	var a2 [0]int
// 	n1 := &classfour.Node{}
// 	n2 := &classfour.Node{}
// 	for i := 0; i < 10; i++ {
// 		n3 := &classfour.Node{}
// 		fmt.Printf("Address of n3: %p\n", &n3)
// 		//每次跑完for迴圈都會被丟掉
// 	}

// 	fmt.Printf("Address of s1: %p\n", &s1)
// 	fmt.Printf("Address of s2: %p\n", &s2)
// 	fmt.Printf("Address of a1: %p\n", &a1)
// 	fmt.Printf("Address of a2: %p\n", &a2)
// 	fmt.Printf("Address of n1: %p\n", &n1)
// 	fmt.Printf("Address of n2: %p\n", &n2)
// }

func main() {
	// fmt.Print(other.Fibonacci(11))
	for i := 1; i <= 11; i++ {
		fmt.Print(other.Fibonacci(i))
	}
	fmt.Println(other.FibonacciArray(11))
}
