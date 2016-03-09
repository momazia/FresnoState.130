/*
	Source: 		Column BI in Gradebook
	Author: 		Mohamad Mahdi Ziaee
	Description:	PROJECT STEP 1 - create a web application that serves an HTML template.
*/

package main

import (
	"log"
	"net/http"
	"text/template"
)

func login(res http.ResponseWriter, req *http.Request) {

	temp, err := template.ParseFiles("login.html")
	// Logging possible errors
	logFatalError(err)

	temp.Execute(res, nil)
	// Logging possible errors
	logFatalError(err)
}
func main() {

	// Ignoring favicon.ico
	http.Handle("/favicon.ico", http.NotFoundHandler())

	// Setting the handler for login
	http.HandleFunc("/", login)

	// Setting the listener on port 8080
	log.Println("Listening to 8080 ...")
	http.ListenAndServe(":8080", nil)
}

// Logs error at Fatal level given the error is not nil.
func logFatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
