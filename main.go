package main

import (
	classtwo "Algorithms/ClassTwo"
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("input i: %d,alive i: %d\n", i, classtwo.Josephus(i))
	}
}
