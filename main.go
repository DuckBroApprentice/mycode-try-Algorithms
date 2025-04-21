package main

import (
	classfour "Algorithms/ClassFour"
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
func main() {
	test := classfour.NewGraph(7)
	test.AddEdge(2, 4)
	test.AddEdge(7, 0)
	test.AddEdge(7, 3)
	test.AddEdge(6, 2)

	fmt.Println("4,6 :", test.HasEdge(4, 6))
	fmt.Println("7,0 :", test.HasEdge(7, 0))
	fmt.Println("6,2 :", test.HasEdge(6, 2))
	fmt.Println(test)
}
