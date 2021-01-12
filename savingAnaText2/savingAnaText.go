package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type choice struct {
	cmd         string
	description string
	hasDialog   bool
	dialog      string
	nextNode    *storyNode
}

type storyNode struct {
	text    string
	choices []*choice
}

func (node *storyNode) addChoice(cmd string, description string, hasDialog bool, dialog string, nextNode *storyNode) {
	choice := &choice{cmd, description, hasDialog, dialog, nextNode}
	node.choices = append(node.choices, choice)
}

func (node *storyNode) render() {
	fmt.Println(node.text)
	if node.choices != nil {
		for _, choice := range node.choices {
			fmt.Println(choice.cmd, ":", choice.description)
		}
	}
}

func (node *storyNode) executeCmd(cmd string) *storyNode {
	for _, choice := range node.choices {
		if strings.ToLower(choice.cmd) == strings.ToLower(cmd) {
			return choice.nextNode
		}
	}
	fmt.Println("Sorry, I didn't understand that.")
	return node
}

var scanner *bufio.Scanner

func (node *storyNode) play() {
	node.render()
	if node.choices != nil {
		scanner.Scan()
		node.executeCmd(scanner.Text()).play()
	}
}

func main() {
	scanner = bufio.NewScanner(os.Stdin)

	start := createAllNodesAndChoices()

	start.play()
	fmt.Println("The end.")
}

func createAllNodesAndChoices() storyNode {

	// Misc
	start := storyNode{text: `
	You wake up to see a man standing by a light post. It appears you are on some kind of boardwalk in outerspace. "Hey Kid what the hell just happened??" You sit on the ground still stunned. "We were talking about your next target and then you just collapsed like an invisible refridgerator unit fell on your head." Your body feels different, you are in a place you have never seen before and you have never met the man in front of you...
	`}

	inShip401 := storyNode{text: `
	You descend into the seat of your cockpit and feel the comfort of your ship surround you. The hiss of the cabin pressurizing, bleebs of the internal systems, and the roar of the engines starting up fill you with determination.
	`}

	inShip402 := storyNode{text: `
	While your engines head up you gaze out of the front of your ship and watch small chunks of space debris flaot past.
	`}

	inShip403 := storyNode{text: `
	You realize that your cockpit smells mildly of cigarette smoke and alcohol. You grab a pizza shaped air freshener form the console and place hang it from a monitor arm. "Much better!"
	`}

	endChase := storyNode{text: `
	After a long chase you catch up to Jer Banta and disable his ship. On your way back to Port 401 you debate why you took this mission so seriously since this isn't even your body. You begin to think of how you might be able to contact Noyaba and how you will find your way back to Ana.
	`}

	endView := storyNode{text: `
	After speaking with Jer Banta you decide that there is no point in destroying anohter persons life in exchange for money. You don't even know for sure what Jer had done to have a bounty placed on his head. You remember that you have much more pressing concerns, such as finding your real body and getting back to Ana.
	`}

	endBar := storyNode{text: `
	Eternity is a long time, but not long enough to be chasing someone who may or may not dererve the bounty that's been placed on them. You head over to Smokin' Gunz and drink with Anigun for days. After all, there's no need to worry about your health when the body you are currently in isn't even your own.
	`}

	endPort401 := storyNode{text: `
	You fly back to Port 401 to tell Topps that you have given up on bounty hunting. Upon arrival you see the lamp post he is usually leaning against but he is nowhere to be found. Just then you notice that on the ground where you originally fell is a small shiny seed. This jolts your memory and you remember everything about Noyaba and Ana...
	`}

	endPort402 := storyNode{text: `
	You jet toward Port 402 feeling that now is not the time to chase Jer Banta. Instead, you are focused on Zella. You get to the diner and see her sitting in her usual booth. You approach and say "Is now a good time to take you up on that breakfast?"
	`}

	confrontBounty := storyNode{text: `
	You confront Jer Banta and he tries to tell you why he has a bounty to stop you from taking him in.
	`}

	// P401
	lobby := storyNode{text: `
	Des
	`}

	port401 := storyNode{text: `
	Des
	`}

	spaceBodega := storyNode{text: `
	Des
	`}

	// P402
	port402 := storyNode{text: `
	Des
	`}

	diner := storyNode{text: `
	Des
	`}

	bathroom := storyNode{text: `
	Des
	`}

	// P403
	port403Locked := storyNode{text: `
	Des
	`}

	port403Unlocked := storyNode{text: `
	Des
	`}

	theLookout := storyNode{text: `
	Des
	`}

	bar := storyNode{text: `
	Des
	`}

	// Choices
	// Misc
	start.addChoice("Stand", "Stand up and try to get your bearings", false, "", &lobby)

	inShip401.addChoice("402", "Fly to port 402", false, "", &port402)
	inShip401.addChoice("403", "Fly to port 403", false, "", &port403Locked)

	inShip402.addChoice("401", "Fly to port 401", false, "", &port401)
	inShip402.addChoice("403", "Fly to port 403", false, "", &port403Locked)

	inShip403.addChoice("401", "Fly to port 401", false, "", &port401)
	inShip403.addChoice("402", "Fly to port 402", false, "", &port402)

	confrontBounty.addChoice("Jump", "in your shop and chase Jer Banta!", false, "", &endChase)
	confrontBounty.addChoice("Abandon", "this bounty and head to the lookout", false, "", &endView)
	confrontBounty.addChoice("Leave", "the port and head to Smokin' Gunz", false, "", &endBar)
	confrontBounty.addChoice("Get in", "your ship and fly back to Port401", false, "", &endPort401)
	confrontBounty.addChoice("Decide", "to board your ship and head to Port402", false, "", &endPort402)

	// -------------------

	// Port 401
	// Lobby
	lobby.addChoice("Walk", "to the dock of Port401", false, "", &port401)
	lobby.addChoice("Saunter", "to the space bodega", false, "", &spaceBodega)
	lobby.addChoice("Talk", "to man standing by the lightpost", true, "He says it must be nice to be a bounty hunter. Tells you this mission shouldn't take long, 15 - 20 mins", &lobby)
	lobby.addChoice("Inspect", "man standing by the lightpost", true, "Scanner: Name: Glenn (Uno) Topps Occupation: n/a -  description, clothes, appearance, demeaner, etc. Hood, red eye smoking", &lobby)

	// P401
	port401.addChoice("Amble", "to the space bodega", false, "", &spaceBodega)
	port401.addChoice("Stroll", "to the lobby", false, "", &lobby)
	port401.addChoice("Board your ship", "and fly into space", false, "", &inShip401)

	// Bodega
	spaceBodega.addChoice("Ramble", "to the lobby", false, "", &lobby)
	spaceBodega.addChoice("Hike", "to the dock of Port401", false, "", &port401)
	spaceBodega.addChoice("Smoke", "a cigarette", true, "describe buying a pack and smoking", &spaceBodega)
	spaceBodega.addChoice("Eat", "a candy bar", true, "describe buying and eating tasty deli food", &spaceBodega)
	spaceBodega.addChoice("Drink", "an energy drink", true, "describe buying and dirnking energizing beverage", &spaceBodega)

	// -------------------

	// Port 402
	// P402
	port402.addChoice("March", "into the bathroom", false, "", &bathroom)
	port402.addChoice("Plod", "to the diner", false, "", &diner)
	port402.addChoice("Get in", "your sweet galactic ride and cruise", false, "", &inShip402)

	// Diner
	diner.addChoice("Trek", "to the dock of Port402", false, "", &port402)
	diner.addChoice("Roam", "to the restroom", false, "", &spaceBodega)
	diner.addChoice("Order", "a meal from the diner", true, "describe a tasty meal", &diner)
	diner.addChoice("Drink", "coffee", true, "describe coffee", &diner)
	diner.addChoice("Speak", "to the woman in booth 13", true, "she tells you about the bounty, says you are cute, and asks if you want to join her for breakfast some time", &diner)
	diner.addChoice("Check out", "the woman in booth 13", true, "Scanner: Name, Occupation - description, clothes, appearance, demeaner, etc. Attractive", &diner)

	// Bathroom
	bathroom.addChoice("Prowl", "to the dock of Port402", false, "", &port402)
	bathroom.addChoice("Traipse", "to the diner", false, "", &diner)
	bathroom.addChoice("Look", "at yourself in the cracked mirror", true, "Scanner: Name: Kid Vessla, Occupation: Bounty Hunter - describe yourself", &bathroom)
	bathroom.addChoice("Sneak", "a cigarette", true, "describe smoking", &bathroom)

	// -------------------

	// Port 403
	// P403 Locked
	port403Locked.addChoice("Proceed Right", "to the lookout", false, "", &theLookout)
	port403Locked.addChoice("Wander Left", "to Smokin' Gunz Bar", false, "", &bar)
	port403Locked.addChoice("Scan", "the suspicious person standing in the corner", true, "Scanner: No Data on this person describe the person in the corner", &port403Locked)
	port403Locked.addChoice("Ease into your space steed", "hoping there are no space police in the area", false, "", &inShip403)

	// P403 Unlocked
	port403Unlocked.addChoice("Mosey Right", "to the space bodega", false, "", &theLookout)
	port403Unlocked.addChoice("Sprint Left", "to Smokin' Gunz Bar", false, "", &bar)
	port403Unlocked.addChoice("Glare at", "Jer Banta, the person who is your bounty", true, "Scanner: Name, Occupation - describe Jer banta", &port403Unlocked)
	port403Unlocked.addChoice("Confront", "the bounty!", false, "", &confrontBounty)
	port403Unlocked.addChoice("Leap into your steel falcon", "and prepare yourself for flight", false, "", &inShip403)

	// The Lookout
	theLookout.addChoice("Stride", "to Smokin' Gunz Bar", false, "", &bar)
	theLookout.addChoice("Advance ", "to the dock of Port403", false, "", &port403Locked)
	theLookout.addChoice("Gaze", "out into space from what is said to be the best view in the galaxy", true, "You see the most beautiful sight you've ever seen. Words don't describe how amazing this view is. One would have to see it to truly understand, and once they see it, their view of life will never be the same.", &theLookout)

	// Smokin' Gunz Bar
	bar.addChoice("Stumble Right", "to the dock of Port403", false, "", &port403Locked)
	bar.addChoice("Float Left", "to the lookout", false, "", &theLookout)
	bar.addChoice("Light", "a cigarette", true, "describe smoking", &bar)
	bar.addChoice("Order Whisky", "because it's been one of those days", true, "describe whisky drinking here", &bar)
	bar.addChoice("Order Beer", "casually", true, "describe beer drinking here and maybe conversation with the bartender", &bar)
	bar.addChoice("Watch", "the bartender", true, "Scanner: Name, Occupation - describe Anigun while they work", &bar)
	bar.addChoice("Chat", "with the bartender", true, "conversation with the bartender", &bar)

	return start
}
