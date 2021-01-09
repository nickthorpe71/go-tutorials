package main

import "fmt"

type Position struct {
	x float32
	y float32
}

func main() {
	p := Position{3, 4}

	fmt.Println(p.x)
}
