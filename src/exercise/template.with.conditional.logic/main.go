/*
	Source: 		Column AN in Gradebook
	Author: 		Mohamad Mahdi Ziaee
	Description:	Create a template that uses conditional logic
*/

package main

import (
	"log"
	"os"
	"text/template"
)

const condition bool = false // Constant used to show difference results using the template.

func main() {

	//Parsing the template
	temp, err := template.ParseFiles("src/exercise/template.with.conditional.logic/conditional.logic.temp.html")

	// Logging possible errors
	logError(err)

	// Executing the template using the constant
	err = temp.Execute(os.Stdout, condition)

	// Logging possible errors
	logError(err)
}

// Logs error
func logError(err error) {

	if err != nil {
		log.Fatal(err)
	}
}
