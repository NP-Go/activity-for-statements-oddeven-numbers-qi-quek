package main

import "fmt"

var userInput int

func main() {
	fmt.Println("Please Input a number")
	fmt.Scan(&userInput)

	if userInput == 0 {
		fmt.Println(userInput, "is an even number")
	} else if userInput%2 == 0 {
		fmt.Println(userInput, "is an even number")
	} else {
		fmt.Println(userInput, "is an odd number")
	}
}
