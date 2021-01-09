package main

import "fmt"

func main() {

	x := 5
	fmt.Println(x)

	// xPtr := &x
	var xPtr *int = &x
	fmt.Println(xPtr)
}
