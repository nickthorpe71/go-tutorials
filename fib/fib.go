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

func fibDynamic(n int, memo []int) int {
	if memo[n] != 0 {
		fmt.Println("memo!")
		return memo[n]
	}
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		result := fibDynamic(n-1, memo) + fibDynamic(n-2, memo)
		memo[n] = result
		return result
	}
}

func main() {
	fmt.Println(fibIterative(5)) // iterative
	fmt.Println(fibRecursive(5)) // recursive

	memo := make([]int, 5+1)
	fmt.Println(fibDynamic(5, memo)) // memoized

	// write dynamic version
	// bottom up version

	// https://www.youtube.com/watch?v=vYquumk4nWw&ab_channel=CSDojo
}
