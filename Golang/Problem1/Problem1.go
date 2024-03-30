package main

import (
	"fmt"
	"math/rand"
)

// Create a program that generates a random number between 1 and 100
// (or a customizable range) and prompts the user to guess it.
func main() {
	actual_number := rand.Intn(100)
	guess := 0
	for guess != actual_number {
		print("Guess a number between 1 and 100: ")
		_, err := fmt.Scan(&guess)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}
		if guess < actual_number {
			fmt.Println("Too low!")
		} else if guess > actual_number {
			fmt.Println("Too high!")
		} else {
			fmt.Println("You guessed it!")
		}
	}
}
