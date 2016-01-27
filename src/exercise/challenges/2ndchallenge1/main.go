/*
	Source: 		Column AP in Gradebook
	Author: 		Mohamad Mahdi Ziaee
	Description:	1) Write a function which takes an integer. The function will have two returns.
					The first return should be the argument divided by 2.
					The second return should be a bool that let’s us know whether or not the argument was even.
					For example:
					half(1) returns (0, false)
					half(2) returns (1, true)
*/
package main

import (
	"fmt"
)

func main() {

	// Test 1
	output, even := half(1)
	fmt.Printf("Result: %v, isEven: %v\n", output, even)

	// Test 2
	output, even = half(2)
	fmt.Printf("Result: %v, isEven: %v", output, even)
}

// Divides the input into 2 and returns the result and a boolean to determine if the input is even.
func half(input int) (int, bool) {

	// Checking to see if the input is even
	even := input%2 == 0

	// Dividing the input by 2
	result := input / 2

	return result, even
}
