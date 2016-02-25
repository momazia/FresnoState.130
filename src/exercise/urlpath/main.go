/*
	Source: 		Column AV in Gradebook
	Author: 		Mohamad Mahdi Ziaee
	Description:	Create a webpage that displays the URL path using req.URL.Path
*/
package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	// Registering the URL path and binding it to surferPage handler
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "URL:"+req.URL.Path)
	})

	// Setting the listener on port 8080
	log.Println("Listening to 8080 ...")
	http.ListenAndServe(":8080", nil)
}
