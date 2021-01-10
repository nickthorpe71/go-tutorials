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

func (page *storyPage) addToEnd(text string) {
	for page.nextPage != nil {
		page = page.nextPage
	}
	page.nextPage = &storyPage{text, nil}
}

func (page *storyPage) addAfter(text string) {
	newPage := &storyPage{text, page.nextPage}
	page.nextPage = newPage
}

func deletePage(page *storyPage, prevPage *storyPage) {
	prevPage.nextPage = page.nextPage
}

func main() {
	page1 := storyPage{"It was a hot day in Chicago.", nil}
	page1.addToEnd("You are waiting in line at the veggie hotdog stand.")
	page1.addToEnd("Someone in the font of the line is complaining that there is not enough dog in their veggiedog....")

	page1.addAfter("Test after add")
	page1.playStory()
}
