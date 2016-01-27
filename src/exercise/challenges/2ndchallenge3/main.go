/*
	Source: 		Column AP in Gradebook
	Author: 		Mohamad Mahdi Ziaee
	Description:	3) Write a function with one variadic parameter that finds the greatest number in a list of numbers.
*/
package main

import (
	"fmt"
)

func main() {

	// Test 1
	biggestNumber := findBiggestNumber(-60, 20, 30)
	fmt.Println("Biggest number: ", biggestNumber)

	// Test 2
	biggestNumber = findBiggestNumber()
	fmt.Println("Biggest number: ", biggestNumber)

}

// Finds the biggest number in the input
func findBiggestNumber(inputs ...int) int {

	biggestNumber := 0

	for index := range inputs {
		// If we find a bigger number, we want to keep that
		if biggestNumber < inputs[index] {
			biggestNumber = inputs[index]
		}
	}

	return biggestNumber
}
