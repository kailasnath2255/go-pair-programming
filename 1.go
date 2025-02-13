package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Println("Enter the 1 st number")
	fmt.Scanln(&a)
	fmt.Println("Enter the 2 nd number")
	fmt.Scanln(&b)
	fmt.Println("Enter the 3 rd number")
	fmt.Scanln(&c)

	if (a > b) && (a > c) {
		fmt.Println("the first nymber is greater:", a)
	} else if (b > a) && (b > c) {
		fmt.Println("the second digit is greater", b)
	} else {
		fmt.Println("the third digit is greater", c)
	}
}
