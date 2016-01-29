/*
	Source: 		Column AQ in Gradebook
	Author: 		Mohamad Mahdi Ziaee
	Description:	Create the "surfer page" and serve it using Go
*/
package main

import (
	"html/template"
	"log"
	"net/http"
)

// surferPage renders surfer.html page by putting the page content into the response back to client.
func surferPage(res http.ResponseWriter, req *http.Request) {
	temp, err := template.ParseFiles("./surfer.html")
	// Logging possible errors
	logError(err)
	temp.Execute(res, nil)

}

func main() {
	// Registering the URL path and binding it to surferPage handler
	http.HandleFunc("/", surferPage)

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

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
