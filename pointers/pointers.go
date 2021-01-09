package main

import "fmt"

func addOne(x int) {
	x = x + 1
}

func mutatePlusOne(x *int) { // take pointer as a parameter
	*x = *x + 1
}

type position struct {
	x float32
	y float32
}

type badGuy struct {
	name   string
	health int
	pos    position
}

func whereIsBadGuy(b *badGuy) {
	x := b.pos.x
	y := b.pos.y
	fmt.Println(x, y)
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

	p := position{3, 4}
	b := badGuy{"mean person", 100, p}
	whereIsBadGuy(&b)

}
