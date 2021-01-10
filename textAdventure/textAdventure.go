package main

// linked list
type choices struct {
	cmd         string
	description string
	node        *storyNode
	next        *choices
}

// graph
type storyNode struct {
	text    string
	choices *choices
}

func main() {

}
