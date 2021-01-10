package main

import (
	"fmt"
)

type storyPage struct {
	text     string
	nextPage *storyPage
}

func (page *storyPage) playStory() {
	for page != nil {
		fmt.Println(page.text)
		page = page.nextPage
	}
}

func (page *storyPage) addToEnd(pageToAdd *storyPage) {

}

func deletePage(page *storyPage, prevPage *storyPage) {
	prevPage.nextPage = page.nextPage
}

func main() {
	page1 := storyPage{"It was a hot day in Chicago.", nil}
	page2 := storyPage{"You are waiting in line at the veggie hotdog stand.", nil}
	page3 := storyPage{"Someone in the font of the line is complaining that there is not enough dog in their veggiedog....", nil}

	page1.nextPage = &page2
	page2.nextPage = &page3

	// insertPage(&page2, "The smell of roasted vegetables was in the air.")
	// deletePage(&page2, &page1)

	page1.playStory()
}
