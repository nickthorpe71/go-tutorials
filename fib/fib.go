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

	// bottom up version in python (need to convert)
	// def fib_bottom_up(n):
	// if n == 1 or n == 2:
	//     return 1
	// bottom_up = [None] * (n+1)
	// bottom_up[1] = 1
	// bottom_up[2] = 1
	// for i in range(3, n+1):
	//     bottom_up[i] = bottom_up[i-1] + bottom_up[i-2]
	// return bottom_up[n]

	// https://www.youtube.com/watch?v=vYquumk4nWw&ab_channel=CSDojo
}
