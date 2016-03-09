/*
	Source: 		Column AZ in Gradebook
	Author: 		Mohamad Mahdi Ziaee
	Description:	Create a webpage that serves a form and allows the user to upload a txt file. You do not need to check if the file is a txt;
					bad programming but just trust the user to follow the instructions.
					Once a user has uploaded a txt file, copy the text from the file and display it on the webpage.
					Use req.FormFile and io.Copy to do this
*/
package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

// Handles uploadTextFile page using a template.
func uploadTextFile(res http.ResponseWriter, req *http.Request) {

	temp, err := template.ParseFiles("./uploadTextFile.html")
	// Logging possible errors
	logError(err)

	if req.Method == "POST" {

		inFile, _, err := req.FormFile("file")
		defer inFile.Close()
		// Logging possible errors
		logError(err)

		contents, err := ioutil.ReadAll(inFile)
		// Logging possible errors
		logError(err)

		err = temp.Execute(res, string(contents))
	} else {
		err = temp.Execute(res, false)
	}
	// Logging possible errors
	logError(err)

}

func main() {

	// Registering the URL path and binding it to userNameForm handler
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", uploadTextFile)

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
