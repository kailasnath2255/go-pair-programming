package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1
	var guess int

	fmt.Println("Guess a number between 1 and 100!")

	for {
		fmt.Print("Enter your guess: ")
		fmt.Scan(&guess)

		if guess < target {
			fmt.Println("Too low! Try again.")
		} else if guess > target {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Println("Congratulations! You guessed it right.")
			break
		}
	}
}
