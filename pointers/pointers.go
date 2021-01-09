package main

import "fmt"

func addOne(x int) {
	x = x + 1
}

func mutatePlusOne(x *int) { // take pointer as a parameter
	*x = *x + 1
}

func main() {

	x := 5
	fmt.Println(x)

	// xPtr := &x
	var xPtr *int = &x
	fmt.Println(xPtr)

	addOne(x)      // sends a copy of x
	fmt.Println(x) // so x remains 5

	mutatePlusOne(xPtr)
	fmt.Println(x)
}
