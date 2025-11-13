//
// Filename: exercise03.go | Alpha Beta Debug Go Tutorial
// Description: Number guessing game
// Copyright (c) 2025 Ryan Smith
//

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const maxNum = 100 					// The max value we can guess
	secretNum := rand.Intn(maxNum + 1) 	// Random number in the range 0 - 100
	guesses := 0 						// The number of guesses we've made
	var guess int 						// The guess the user just made 

	fmt.Println("I am thinking of a number between 0 and 100. Try to guess it!")

	for guess != secretNum {
		fmt.Print("\nEnter your guess: ")
		fmt.Scanf("%d", &guess)

		// Game logic
		if guess < secretNum {
			fmt.Println("Too low. Guess again!")
		} else if guess > secretNum {
			fmt.Println("Too high. Guess again!")
		} else {
			fmt.Println("Correct!")
		}

		guesses++
	}

	// Handle singular/plural (grammar matters in your code!)
	end := ""
	if guesses != 1 {
		end = "es"
	}

	fmt.Printf("\nYou guessed the number in %d guess%s\n", guesses, end)
}