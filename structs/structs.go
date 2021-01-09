package main

import "fmt"

type position struct {
	x float32
	y float32
}

type badGuy struct {
	name   string
	health int
	pos    position
}

func whereIsBadGuy(b badGuy) {
	x := b.pos.x
	y := b.pos.y
	fmt.Println(x, y)
}

func main() {
	p := position{3, 4}
	fmt.Println(p.x)

	b := badGuy{"mean person", 100, p}
	fmt.Println(b.name)
	whereIsBadGuy(b)
}
