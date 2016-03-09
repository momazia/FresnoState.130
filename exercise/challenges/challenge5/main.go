/*
	Source: 		Column AP in Gradebook
	Author: 		Mohamad Mahdi Ziaee
	Description:	5) Print all of the even numbers between 0 and 100.
*/
package main

import (
	"fmt"
)

func main() {

	// Looping through 100 digits, starting with 0 until 100 (included)
	for index := 0; index <= 100; index++ {
		// Printing only those numbers which their remainders are 0 when divided by 2.
		if index%2 == 0 {
			fmt.Println(index)
		}
	}

}
