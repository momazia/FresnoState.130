/*
	Source: 		Column BM in Gradebook
	Author: 		Mohamad Mahdi Ziaee
	Description:	PROJECT STEP 5 - continuing to build our application,
					integrate HMAC into our application to ensure that nobody tampers with the cookie.
	Comment:		Change the cookie on the client by running the following JS: document.cookie="userData=new data"
*/

package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"io"
	"log"
	"net/http"
	"text/template"
)

type User struct {
	Age  string
	Name string
}

var cookiesDigests = make(map[string]string)

const (
	PRIVATE_KEY string = "Some Private Key"
)

func loginHandler(res http.ResponseWriter, req *http.Request) {

	// validating the cookies
	validateCookie(req)

	setCookies(res, req)

	temp, err := template.ParseFiles("login.html")
	// Logging possible errors
	logFatalError(err)

	temp.Execute(res, nil)
	// Logging possible errors
	logFatalError(err)
}

// Validates to make sure the cookies have been modified on client side.
func validateCookie(req *http.Request) {
	for key, value := range cookiesDigests {
		log.Println("Key:", key, "Value:", value)
		cookie, err := req.Cookie(key)
		if err == nil {
			log.Println("Cookie:" + cookie.String())
			// Logging the possible errors
			if value != generateDigest(cookie.Value) {
				log.Println("The following cookie has been changed: [" + key + "]")
				return
			}
		} else {
			log.Println(err)
		}
	}
}

// Sets the cookies on the response
func setCookies(res http.ResponseWriter, req *http.Request) {

	// Generating a new ID
	id, err := uuid.NewV4()
	logFatalError(err)
	createCookie(&res, "session-fino", id.String())

	// Setting user's data on cookie
	user := User{
		Age:  req.FormValue("age"),
		Name: req.FormValue("name"),
	}

	// Avoiding setting the cookie for the first visit.
	if req.Method == "POST" {
		userBytes, err := json.Marshal(user)
		logFatalError(err)
		createCookie(&res, "userData", base64.StdEncoding.EncodeToString(userBytes))
	}
}

// Creates a cookie for the given name and value and sets it on the response
func createCookie(res *http.ResponseWriter, cookieName, cookieValue string) {

	// Setting the cookie
	cookie := &http.Cookie{
		Name:  cookieName,
		Value: cookieValue,
		//		Secure: true,
		HttpOnly: false, // Setting this to false so we can modify it on the client side using JS
	}

	// Storing the digest in the map for verification purporses
	cookiesDigests[cookieName] = generateDigest(cookieValue)

	// Setting the cookie on the response back to the client
	http.SetCookie(*res, cookie)
}

// Generates a digest out of the message given.
func generateDigest(str string) string {
	h := hmac.New(sha256.New, []byte(PRIVATE_KEY))
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {

	// Ignoring favicon.ico
	http.Handle("/favicon.ico", http.NotFoundHandler())

	// Setting the handler for login
	http.HandleFunc("/", loginHandler)

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
