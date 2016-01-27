/*
	Source: 		Column AP in Gradebook
	Author: 		Mohamad Mahdi Ziaee
	Description:	4) Create a program that prints to the terminal asking for a user to enter
					a small number and a larger number. Print the remainder of the larger number
					divided by the smaller number.
*/
package main

import (
	"fmt"
)

func main() {

	// Handling the errors in a generic way
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Something went wrong: ", err)
		}
	}()

	// Getting the numbers
	firstNumber := getNumber(1)
	secondNumber := getNumber(2)

	// Printing the remainder of the larger number divided by the smaller number
	if firstNumber > secondNumber {
		fmt.Printf("Remainder: %v", firstNumber%secondNumber)
	} else {
		fmt.Printf("Remainder: %v", secondNumber%firstNumber)
	}

}

// getNumber takes a number as a sequencial digit and returns user's input.
func getNumber(number int) int {
	fmt.Printf("Enter number %v: ", number)
	var input int
	fmt.Scan(&input)
	return input
}
