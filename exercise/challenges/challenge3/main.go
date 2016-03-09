/*
	Source: 		Column AP in Gradebook
	Author: 		Mohamad Mahdi Ziaee
	Description:	3) Create a program that prints to the terminal asking for a user to enter their name.
					Use fmt.Scan to read a user’s name entered at the terminal.
					Print “Hello <NAME>” with <NAME> replaced with what the user entered at the terminal.
*/
package main

import (
	"fmt"
)

func main() {

	// Getting the user's name
	fmt.Printf("Name: ")
	var name string
	fmt.Scan(&name)

	// Printing the name
	fmt.Printf("Hello " + name)
}
