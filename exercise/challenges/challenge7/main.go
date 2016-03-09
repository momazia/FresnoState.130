/*
	Source: 		Column AP in Gradebook
	Author: 		Mohamad Mahdi Ziaee
	Description:	7) If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
					The sum of these multiples is 23.
					Find the sum of all the multiples of 3 or 5 below 1000.
*/
package main

import (
	"fmt"
)

func main() {

	// results slice holds the values which can be multiples of 3 or 5.
	var results []int

	// Looping through 1000 digits, starting with 1 until 1000 (excluded)
	for index := 1; index < 1000; index++ {

		// Checking to see if the index is multiple of 3
		var multipleOfThree bool = (index%3 == 0)
		// Checking to see if the index is multiple of 5
		var multipleOfFive bool = (index%5 == 0)

		// If the index multiple of three or five, we want to keep them in the result slice.
		if multipleOfThree || multipleOfFive {
			results = append(results, index)
		}
	}

	// Looping through results slice to sum them up
	var sum int = 0
	for index := range results {
		sum = sum + results[index]
	}

	// Printing the sum
	fmt.Println("Sum: ", sum)
}
