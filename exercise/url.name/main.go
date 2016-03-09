/*
	Source: 		Column AW in Gradebook
	Author: 		Mohamad Mahdi Ziaee
	Description:	Create a webpage that serves at localhost:8080 and will display the name in the url when the url is localhost:8080/name - use req.URL.Path to do this
*/
package main

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {

	// Registering the URL path and binding it to surferPage handler
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		name := strings.Split(req.URL.Path, "/")
		log.Println(name)
		io.WriteString(res, "Name: "+name[1])
	})

	// Setting the listener on port 8080
	log.Println("Listening to 8080 ...")
	http.ListenAndServe(":8080", nil)
}
