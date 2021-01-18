package main

import (
	"bufio"
	"fmt"
	"os"

	// "spaceFlight"
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
	fmt.Println("\033[2J")
	for _, choice := range node.choices {
		if strings.ToLower(choice.cmd) == strings.ToLower(cmd) {
			if choice.hasDialog {
				fmt.Println(choice.dialog)
				fmt.Println("Press Enter to return")
				scanner.Scan()
			}
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

	// spaceFlight.Fly()

	fmt.Println("\033[2J")
	start.play()
	fmt.Println("The end.")
}

func createAllNodesAndChoices() storyNode {

	// Misc
	start := storyNode{text: `
	You wake up to see a man standing by a light post. It appears you are on some kind wooden paltform. "Hey Kid what the hell just happened??" You sit on the ground still stunned. "We were talking about your next target and all of a sudden you collapsed like an invisible refridgerator unit fell on your head." Your body feels different, you are in a place you have never seen before and you have never met the man in front of you...
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
	You jet toward Port 402 feeling that now is not the time to chase Jer Banta. Instead, you are focused on Zella. You get to Pot City Diner and see her sitting in her usual booth. You approach and say "Is now a good time to take you up on that breakfast?"
	`}

	confrontBounty := storyNode{text: `
	You confront Jer Banta and he tries to tell you why he has a bounty to stop you from taking him in.
	`}

	// P401
	lobby := storyNode{text: `
	This looks like a boardwalk floating in space mixed with an airport. There is someone in a cloak leaning against a lamp post straight ahead of you smoking a cigarette. To the left you see what looks like a run down bodega and to the right is a large platform with strange aircrafts.
	`}

	port401 := storyNode{text: `
	There are many ships on this platform but one stands out to you. Seeing it is like seeing an old friend. 
	`}

	spaceBodega := storyNode{text: `
	You enter a dingy run down space bodega. Behind the counter sits a crusty old dude who is shirtless and singing to himself under his breath.
	`}

	// P402
	port402 := storyNode{text: `
	This platform is filled with a warm light and the smell of delicious food. This port is well designed and looks very modern and clean. You see two signs, "Pot City Diner" and "Restroom"
	`}

	diner := storyNode{text: `
	As you enter the diner through the decorated metal door, you are welcomed by the sent of delicious food and the gaze of a beautiful woman sitting in a booth. 
	`}

	bathroom := storyNode{text: `
	You admire the luxury of this bathroom. It is very clean and well kept with marble counters and shiny tile floor. The only thing that seems out of place is that the mirror above the gold plated sink has a large crack running diagonallly from the top right to the bottom left.
	`}

	// P403
	port403Locked := storyNode{text: `
	This port is very dark and the smell of alcohol and cigarette smoke fills the air. In the corner you notice a suspicious man who seems to be trying to avoid being seen. There are two signs, "Smokin' Gunz Bar" and "The Lookout"
	`}

	port403Unlocked := storyNode{text: `
	You see Jer Banata! He is no longer in the corner hiding but is walking toward his ship!
	`}

	theLookout := storyNode{text: `
	This room is mostly empty with a few lounges and one giant window looking out into space. 
	`}

	bar := storyNode{text: `
	Upon entering the bar you hear the sound of lofi jazz and people chattering. The bartender is working, but still maanger to welcome you with a friendly nod.
	`}

	// Choices
	// Misc
	start.addChoice("Stand up", "<- Type the command on the left and hit enter to take that action", false, "", &lobby)

	inShip401.addChoice("402", "Fly to port 402", false, "", &port402)
	inShip401.addChoice("403", "Fly to port 403", false, "", &port403Locked)

	inShip402.addChoice("401", "Fly to port 401", false, "", &port401)
	inShip402.addChoice("403", "Fly to port 403", false, "", &port403Locked)

	inShip403.addChoice("401", "Fly to port 401", false, "", &port401)
	inShip403.addChoice("402", "Fly to port 402", false, "", &port402)

	confrontBounty.addChoice("Jump", "in your ship and chase Jer Banta!", false, "", &endChase)
	confrontBounty.addChoice("Abandon", "this bounty and head to the lookout", false, "", &endView)
	confrontBounty.addChoice("Leave", "the port and head to Smokin' Gunz", false, "", &endBar)
	confrontBounty.addChoice("Get in", "your ship and fly back to Port401", false, "", &endPort401)
	confrontBounty.addChoice("Decide", "to board your ship and head to Port402", false, "", &endPort402)

	// -------------------

	// Port 401
	// Lobby
	lobby.addChoice("Walk", "to the dock of Port 401", false, "", &port401)
	lobby.addChoice("Saunter", "to the space bodega", false, "", &spaceBodega)
	lobby.addChoice("Talk", "to man standing by the lightpost", true, "Topps: \"Whelp looks like the life of a bounty hunter is just wandering around talking to people, must be nice! This Jer Banata fella you are after is a small fry compared to some fo the fish you've fried. Capturing him shouldn't take you more than 15 to 20 minutes.\"", &lobby)
	lobby.addChoice("Inspect", "man standing by the lightpost", true, "--Scanner--\nName: Glenn (Uno) Topps\nOccupation: n/a\nHe is wearing a ragged brown cloak which covers his face. All that portrudes from the hood is a lit cigarette and a red glow from what looks to be where one of his eyes might be.", &lobby)

	// P401
	port401.addChoice("Amble", "to the space bodega", false, "", &spaceBodega)
	port401.addChoice("Stroll", "to the lobby", false, "", &lobby)
	port401.addChoice("Board", "your ship and fly into space", false, "", &inShip401)

	// Bodega
	spaceBodega.addChoice("Ramble", "to the lobby", false, "", &lobby)
	spaceBodega.addChoice("Hike", "to the dock of Port 401", false, "", &port401)
	spaceBodega.addChoice("Smoke", "a cigarette", true, "You have a strange desire to smoke. The man behind the counter gives you a confused look and hurls a pack of cigarettes at you, then watches you closely as if waiting for you to do something. You take a cigarette out of the pack and place it in your mouth. Reaching in your pocket you find a lighter. You spark up the coffin nail and inhale deeply. The earthy sweet taste gives your body relief, relaxing your muscles and mind.", &spaceBodega)
	spaceBodega.addChoice("Eat", "a candy bar", true, "Below the counter where the shop keep is muttering to himself you see a package labeled \"Rocket Bar\". You place it on the counter. The man turns and looks at you saying \"Just take it! You aren't Kid anyway so I couldn't take the money from you even if you do give it to me using her hands!\" As disturbing as this is you are very hungry. You remove the wrapping and take a bite. The flavor is creamy and sweet but the high sugar content hurts your molars.", &spaceBodega)
	spaceBodega.addChoice("Drink", "an energy drink", true, "\"The good stuff!!\" yells the man behind the counter as you grab a can that says \"Sparks Energy\" from the fridge. You crack the can and guzzle the contents. Your body buzzes with vitality!", &spaceBodega)

	// -------------------

	// Port 402
	// P402
	port402.addChoice("March", "into the bathroom", false, "", &bathroom)
	port402.addChoice("Plod", "over to Pot City Diner", false, "", &diner)
	port402.addChoice("Get in", "your sweet galactic ride and cruise", false, "", &inShip402)

	// Pot City Diner
	diner.addChoice("Trek", "to the dock of Port 402", false, "", &port402)
	diner.addChoice("Roam", "to the restroom", false, "", &spaceBodega)
	diner.addChoice("Order", "a meal from the diner", true, "describe a tasty meal", &diner)
	diner.addChoice("Drink", "coffee", true, "describe coffee", &diner)
	diner.addChoice("Speak", "to the woman in booth 13", true, "she tells you about the bounty, says you are cute, and asks if you want to join her for breakfast some time", &diner)
	diner.addChoice("Check out", "the woman in booth 13", true, "Scanner: Name, Occupation - description, clothes, appearance, demeaner, etc. Attractive", &diner)

	// Bathroom
	bathroom.addChoice("Prowl", "to the dock of Port 402", false, "", &port402)
	bathroom.addChoice("Traipse", "to to Pot City Diner", false, "", &diner)
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
	theLookout.addChoice("Advance ", "to the dock of Port 403", false, "", &port403Locked)
	theLookout.addChoice("Gaze", "out into space from what is said to be the best view in the galaxy", true, "You see the most beautiful sight you've ever seen. Words don't describe how amazing this view is. One would have to see it to truly understand, and once they see it, their view of life will never be the same.", &theLookout)

	// Smokin' Gunz Bar
	bar.addChoice("Stumble Right", "to the dock of Port 403", false, "", &port403Locked)
	bar.addChoice("Float Left", "to the lookout", false, "", &theLookout)
	bar.addChoice("Light", "a cigarette", true, "describe smoking", &bar)
	bar.addChoice("Order Whisky", "because it's been one of those days", true, "describe whisky drinking here", &bar)
	bar.addChoice("Order Beer", "casually", true, "describe beer drinking here and maybe conversation with the bartender", &bar)
	bar.addChoice("Watch", "the bartender", true, "Scanner: Name, Occupation - describe Anigun while they work", &bar)
	bar.addChoice("Chat", "with the bartender", true, "conversation with the bartender", &bar)

	return start
}
