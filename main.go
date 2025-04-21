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
	var test []string
	test = []string{"5", "-", "4", "*", "3", "+", "2"} //expect 5 4 3 * - 2 +
	test2 := []string{"6", "*", "7", "+", "1"}         //6 7 * 1 +
	tree := classfour.ExpressionTree(test)
	tree2 := classfour.ExpressionTree(test2)
	fmt.Println(classfour.PostorderTraversal(tree))  //correct
	fmt.Println(classfour.PostorderTraversal(tree2)) //correct
}
