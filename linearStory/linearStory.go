package main

import (
	"bufio"
	"fmt"
	"os"
)

type storyPage struct {
	text     string
	nextPage *storyPage
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(scanner)
}
