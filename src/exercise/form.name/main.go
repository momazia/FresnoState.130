/*
	Source: 		Column AY in Gradebook
	Author: 		Mohamad Mahdi Ziaee
	Description:	Create a webpage that serves a form and allows the user to enter their name. Once a user has entered their name, show their name on the webpage. Use req.FormValue to do this
*/
package main

import (
	"html/template"
	"log"
	"net/http"
)

type Person struct {
	FirstName string
	LastName  string
}

// Handles userNameForm page.
func userNameForm(res http.ResponseWriter, req *http.Request) {

	temp, err := template.ParseFiles("./userNameForm.html")
	// Logging possible errors
	logError(err)

	firstName := req.FormValue("firstName")
	lastName := req.FormValue("lastName")

	if firstName != "" || lastName != "" {
		err = temp.Execute(res, Person{firstName, lastName})
	}else {
		err = temp.Execute(res, false)
	}
	// Logging possible errors
	logError(err)
}

func main() {

	// Registering the URL path and binding it to userNameForm handler
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", userNameForm)

	// Setting the listener on port 8080
	log.Println("Listening to 8080 ...")
	http.ListenAndServe(":8080", nil)
}

// Logs the error given into log
func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
