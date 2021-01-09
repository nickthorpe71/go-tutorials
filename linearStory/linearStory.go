package main

import (
	"fmt"
)

type storyPage struct {
	text     string
	nextPage *storyPage
}

func playStory(page *storyPage) {
	if page == nil {
		return
	}

	fmt.Println(page.text)
	playStory(page.nextPage)
}

func main() {
	page1 := storyPage{"It was a hot day in Chicago.", nil}
	page2 := storyPage{"You are waiting in line at the veggie hotdog stand.", nil}
	page3 := storyPage{"Someone in the font of the line is complaining that there is not enough dog in their veggiedog....", nil}

	page1.nextPage = &page2
	page2.nextPage = &page3

	playStory(&page1)
}
