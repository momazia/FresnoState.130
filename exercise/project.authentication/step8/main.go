/*
	Source: 		Column BP in Gradebook
	Author: 		Mohamad Mahdi Ziaee
	Description:	PROJECT STEP 8 - Allow the user to logout. Show a log-in button when
					the user is not logged-in. Show a log-out button only when the user is logged in.
*/

package main

import (
	//	"crypto/hmac"
	//	"crypto/sha256"
	//	"encoding/base64"
	//	"encoding/json"
	//	"fmt"
	//	"github.com/nu7hatch/gouuid"
	//	"io"
	"log"
	"net/http"
	"auth"
	"text/template"
)

type User struct {
	UserName string
	Password string
}

var validUsers = []User{
	{UserName: "Mahdi", Password: "123"},
	{UserName: "Amin", Password: "321"},
}

func loginHandler(res http.ResponseWriter, req *http.Request) {
	var isInvalidUser bool = false
	if req.Method == "POST" {
		if isValidUser(req) {
			// Settin the user in session and going to userForm
			http.Redirect(res, req, "/userForm", http.StatusFound)
			return
		} else {
			isInvalidUser = true
		}
	}
	temp, err := template.ParseFiles("login.html")
	logFatalError(err)

	temp.Execute(res, isInvalidUser)
	logFatalError(err)
}

func isValidUser(req *http.Request) bool {
	formUserName := req.FormValue("userName")
	formPassword := req.FormValue("password")
}

func userFormHandler(res http.ResponseWriter, req *http.Request) {
	temp, err := template.ParseFiles("userForm.html")
	logFatalError(err)

	temp.Execute(res, nil)
	logFatalError(err)
}

func landingPageHandler(res http.ResponseWriter, req *http.Request) {
	http.Redirect(res, req, "/userForm", http.StatusFound)
	return
}

func main() {
	// Ignoring favicon.ico
	http.Handle("/favicon.ico", http.NotFoundHandler())

	// Setting the landing page
	http.HandleFunc("/", authenticationFilter(landingPageHandler))

	// Setting the handler for login
	http.HandleFunc("/login", loginHandler)

	// Setting the handler for userForm
	http.HandleFunc("/userForm", authenticationFilter(userFormHandler))

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
