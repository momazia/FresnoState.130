/*
	Source: 		Column AP in Gradebook
	Author: 		Mohamad Mahdi Ziaee
	Description:	6) Write a program that prints the numbers from 1 to 100.
					But for multiples of three print "Fizz" instead of the number and for the multiples of five print "Buzz".
					For numbers which are multiples of both three and five print "FizzBuzz".
*/
package main

import (
	"fmt"
)

func main() {

	// Looping through 100 digits, starting with 1 until 100 (included)
	for index := 1; index <= 100; index++ {

		// Checking to see if the index is multiple of 3
		var multipleOfThree bool = (index%3 == 0)
		// Checking to see if the index is multiple of 5
		var multipleOfFive bool = (index%5 == 0)

		if multipleOfThree && multipleOfFive {
			fmt.Println("FizzBuzz")
		} else if multipleOfThree {
			fmt.Println("Fizz")
		} else if multipleOfFive {
			fmt.Println("Buzz")
		} else {
			// Printing the numbers which are not multiple of 3, 5 and both
			fmt.Println(index)
		}
	}

}
