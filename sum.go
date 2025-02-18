package main

import (
	"fmt"
)

func main() {
	var n, sum int

	// Get user input
	fmt.Print("Enter a number (N): ")
	fmt.Scan(&n)

	// Loop through numbers from 1 to N
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			sum += i
		}
	}

	// Print the sum of even numbers
	fmt.Println("Sum of even numbers from 1 to", n, "is:", sum)
}
