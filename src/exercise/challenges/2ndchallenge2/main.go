/*
	Source: 		Column AP in Gradebook
	Author: 		Mohamad Mahdi Ziaee
	Description:	2) Modify the previous program to use a func expression.
*/
package main

import (
	"fmt"
)

func main() {

	// Divides the input into 2 and returns the result and a boolean to determine if the input is even.
	half := func(input int) (int, bool) {

		// Checking to see if the input is even
		even := input%2 == 0

		// Dividing the input by 2
		result := input / 2

		return result, even
	}

	// Test 1
	output, even := half(1)
	fmt.Printf("Result: %v, isEven: %v\n", output, even)

	// Test 2
	output, even = half(2)
	fmt.Printf("Result: %v, isEven: %v", output, even)
}
