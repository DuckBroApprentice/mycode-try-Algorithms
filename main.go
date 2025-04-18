package main

import (
	"Algorithms/classone"
	"fmt"
)

func main() {
	test := []int{28, -6, -22, 8, 44, 17}
	fmt.Println(test)
	test = classone.HeapSort(test)
	fmt.Println("sorted : ", test)
}
