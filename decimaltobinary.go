package main

import (
	"fmt"
)

func main() {
	var n int

	// Get user input
	fmt.Print("Enter a decimal number: ")
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println("Binary: 0")
		return
	}

	var binary []int

	// Convert decimal to binary
	for n > 0 {
		binary = append([]int{n % 2}, binary...)
		n /= 2
	}

	// Print binary representation
	fmt.Print("Binary: ")
	for _, bit := range binary {
		fmt.Print(bit)
	}
	fmt.Println()
}
