/*
	Source: 		Column BD in Gradebook
	Author: 		Mohamad Mahdi Ziaee
	Description:	Create a webpage which writes a cookie to the client's machine. Though this is NOT A BEST PRACTICE,
					you will store some session data in the cookie. Make sure you use HMAC to ensure that session data is not changed by a user.
*/
package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Data to be sent to user
var data string = "Some data"

// Encrypted data using HMAC method. Keeping this for verification purposes
var digest string

const (
	PRIVATE_KEY string = "private key"
)

func main() {

	// Registering the URL path and binding it to userNameForm handler
	http.Handle("/favicon.ico", http.NotFoundHandler())

	// The handler
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

		isCookieSet, cookieValue := getCookie(req, "cookie")

		if !isCookieSet {
			// Setting the cookie which contains some data
			fmt.Println("Setting the cookie...")
			cookie := &http.Cookie{
				Name:     "cookie",
				Value:    data,
				HttpOnly: true, // Comment out this line if you wish to change the cookie on client side for testing purposes.
			}
			// Setting the cookie on the response back to the client
			http.SetCookie(res, cookie)

			// Saving the encrypted data
			digest = encrypt(data)

		} else {
			fmt.Println("Client value:[" + cookieValue + "]")
			fmt.Println("Server value:[" + data + "]") // We don't really need this in reality sometimes, like passwords
			fmt.Println("Client digest:[" + digest + "]")
			fmt.Println("Server digest:[" + encrypt(cookieValue) + "]")

			// Verifying the coming data
			if encrypt(cookieValue) == digest {
				fmt.Println("The are the same! You are good")
			} else {
				fmt.Println("The are NOT the same!")
			}
		}
	})

	// Setting the listener on port 8080
	log.Println("Listening to 8080 ...")
	http.ListenAndServe(":8080", nil)
}

// Finds the given cookieName from the request. If it is found, it returns true and its string value, otherwize false and empty string
func getCookie(req *http.Request, cookieName string) (bool, string) {

	// Getting the cookie for the given name
	clientCookie, err := req.Cookie(cookieName)

	// Logging any possible errors
	logError(err)

	if clientCookie != nil {
		return true, clientCookie.Value
	}
	return false, ""
}

func encrypt(data string) string {
	h := hmac.New(sha256.New, []byte(PRIVATE_KEY))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Logs the error given into log
func logError(err error) {
	if err != nil {
		log.Println(err)
	}
}
