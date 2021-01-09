package main

// HOMEWORK
// 1. Make the program print how many tries to win
// 2. Throw an error if the user is guessing outside the range

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	low := 1
	high := 100
	tries := 0
	lastGuess := 0

	fmt.Println("Think of a number from", low, "to", high)
	fmt.Println("Hit Enter when you're ready")

	scanner.Scan()

	for {
		guess := (low + high) / 2
		fmt.Println("Is it", guess, "?")
		fmt.Println("(a) too high?")
		fmt.Println("(b) too low?")
		fmt.Println("(c) correct!")
		scanner.Scan()
		response := scanner.Text()

		if response == "a" {
			high = guess - 1
		} else if response == "b" {
			low = guess + 1
		} else if response == "c" {
			fmt.Println("Yesssss, I win!")
			break
		} else {
			fmt.Println("That isn't one of the options...")
		}

		tries++
		if tries == 1 {
			fmt.Println(tries, "guess")
		} else {
			fmt.Println(tries, "guesses")
		}

		if guess == lastGuess {
			fmt.Println("I don't play with cheaters...")
			break
		}

		lastGuess = guess

	}

}
