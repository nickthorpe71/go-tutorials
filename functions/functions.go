package main

import "fmt"

func sayHello(name string) {
	fmt.Println("hi,", name)
}

func saySomething() {
	fmt.Println("something... ")
}

func beSocial(name string) {
	sayHello(name)
	saySomething()
}

func addOne(x int) int {
	return x + 1
}

func main() {
	beSocial("human1")
	beSocial("human2")

	x := 5
	x = addOne(x)
	fmt.Println(x)
}
