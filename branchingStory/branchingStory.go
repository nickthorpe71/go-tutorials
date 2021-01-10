package main

import (
	"bufio"
	"fmt"
	"os"
)

type storyNode struct {
	text    string
	yesPath *storyNode
	noPath  *storyNode
}

func (node *storyNode) printStory(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Print("  ")
	}

	fmt.Println(node.text)
	if node.yesPath != nil {
		node.yesPath.printStory(depth + 1)
	}
	if node.noPath != nil {
		node.noPath.printStory(depth + 1)
	}
}

func (node *storyNode) play() {
	fmt.Println(node.text)

	if node.yesPath != nil && node.noPath != nil {

		scanner := bufio.NewScanner(os.Stdin)

		for {
			scanner.Scan()
			answer := scanner.Text()

			if answer == "yes" {
				node.yesPath.play()
				break
			} else if answer == "no" {
				node.noPath.play()
				break
			} else {
				fmt.Println("must answer 'yes' or 'no'")
			}
		}
	}

}

func main() {
	root := storyNode{"You wake up in a desert oasis and you see a grotesque beast guarding the entrance to a cave. Do you want to approach the cave?", nil, nil}
	winning := storyNode{"You move along.", nil, nil}
	losing := storyNode{"In a deep gritty female voice the creature calls out in a language you cannot understand. Faster than you can blink or react the creature has one of their pale fingered feet wrapped arount your head and begins removing your limbs with the others... RIP", nil, nil}

	root.yesPath = &losing
	root.noPath = &winning

	root.printStory(0)
}
