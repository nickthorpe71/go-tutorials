package main

import "fmt"

func fibIterative(n int) int {
	a, b, c := 1, 1, 0

	for n > 2 {

		c = b
		b = a
		a = a + c
		n--
	}

	return a
}

func fibRecursive(n uint) uint {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibRecursive(n-1) + fibRecursive(n-2)
	}
}

func main() {
	fmt.Println(fibIterative(5)) // iterative
	fmt.Println(fibRecursive(5)) // recursive
	// write dynamic version
}
