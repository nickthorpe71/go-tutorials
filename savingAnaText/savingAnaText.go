package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type item struct {
	name        string
	description string
}

type inventory struct {
	items []*item
}

type choice struct {
	cmd         string
	description string
	nextNode    *storyNode
}

// graph
type storyNode struct {
	text    string
	choices []*choice
	items   []*item
}

func (node *storyNode) addItem(name string, description string) {
	item := &item{name, description}
	node.items = append(node.items, item)
}

func (currentInventory *inventory) addItemToInventory(inputItem *item) {
	newItem := &item{inputItem.name, inputItem.description}
	currentInventory.items = append(currentInventory.items, newItem)
}

func (node *storyNode) addChoice(cmd string, description string, nextNode *storyNode) {
	choice := &choice{cmd, description, nextNode}
	node.choices = append(node.choices, choice)
}

func (node *storyNode) render() {
	fmt.Println(node.text)
	if node.choices != nil {
		for _, choice := range node.choices {
			fmt.Println(choice.cmd, ":", choice.description)
		}
		fmt.Println("i : show inventory")
	}
}

func (currentInventory *inventory) display() {
	if currentInventory.items != nil {
		for _, item := range currentInventory.items {
			fmt.Println(item.name, ":", item.description)
		}
	} else {
		fmt.Println("empty")
	}
}

func (node *storyNode) executeCmd(cmd string, currentInventory inventory) *storyNode {
	if strings.Contains(strings.ToLower(cmd), "use") {
		for _, item := range currentInventory.items {
			if strings.Contains(strings.ToLower(cmd), strings.ToLower(item.name)) {
				fmt.Println("You use ", item.name)
				return node
			}
		}
	} else if strings.Contains(strings.ToLower(cmd), "pick up") || strings.Contains(strings.ToLower(cmd), "grab") {
		for _, item := range node.items {
			if strings.Contains(strings.ToLower(cmd), strings.ToLower(item.name)) {
				fmt.Println("You pick up the", item.name)
				currentInventory.addItemToInventory(item)
				return node
			}
		}
	} else {
		for _, choice := range node.choices {
			if strings.ToLower(choice.cmd) == strings.ToLower(cmd) {
				return choice.nextNode
			}
		}
	}
	fmt.Println("Sorry, I didn't understand that.")
	return node
}

var scanner *bufio.Scanner

func (node *storyNode) play(currentInventory inventory) {
	node.render()
	if node.choices != nil {
		scanner.Scan()
		if strings.ToLower(scanner.Text()) == "i" {
			currentInventory.display()
			node.play(currentInventory)
		} else {
			node.executeCmd(scanner.Text(), currentInventory).play(currentInventory)
		}
	}
}

func main() {
	scanner = bufio.NewScanner(os.Stdin)

	currentInventory := inventory{}

	start := createAllNodesAndChoices()

	start.play(currentInventory)
	fmt.Println("The end.")
}

func createAllNodesAndChoices() storyNode {

	start := storyNode{text: `
	You wake up parched and covered in sand.
	To the left you see a large figure standing in front of what looks to be a cave.
	To the right you see green plants behing a blanket of desert wind.
	Stright ahead you see nothing but gusts of sand.
	`}

	caveEntranceWithCreature := storyNode{text: `
	You see a grotesque beast guarding the entrance to a cave. 
	Do you want to approach the cave or head back?
	`}

	approachCreature := storyNode{text: `
	In a deep gritty female voice the creature calls out in a language you cannot understand. Faster than you can blink or react the creature has one of their pale fingered feet wrapped arount your head and begins removing your limbs with the others... RIP
	`}

	oasis := storyNode{text: `
	Aftwer walking throught a blinding sand storm you arrive at what appears to be an oaisis of lush greenary. 
	There is an axe lodged in a stump.
	Lay down and rest?
	`}

	lostInTheDesert := storyNode{text: `
	You wander into the desert. Looking back you can no longer see the area you started. You wander for days and eventually die due to dehydration.
	`}

	// Start choices
	start.addChoice("Left", "Go left toward the large figure", &caveEntranceWithCreature)
	start.addChoice("Straight", "Go into the unknown desert", &lostInTheDesert)
	start.addChoice("Right", "Go towards the green", &oasis)

	caveEntranceWithCreature.addChoice("Go back", "Head back to where you started", &start)
	caveEntranceWithCreature.addChoice("Approach", "Approach the cave to investigate the creature", &approachCreature)

	oasis.addChoice("Go Back", "Head back to where you started", &start)
	oasis.addChoice("Pick Up Rusty Axe", "Add this item to your inventory", &oasis)
	oasis.addItem("Rusty Axe", "An old rusty ace that Liser found lodged in a stump.")

	return start
}
