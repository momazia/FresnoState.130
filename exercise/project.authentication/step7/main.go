/*
	Source: 		Column BO in Gradebook
	Author: 		Mohamad Mahdi Ziaee
	Description:	PROJECT STEP 7 - Allow the user to login. Store the information about whether or not a
					user is logged in in both the "user" data type you created and in the cookie. Show a
					"logout" button when the user is logged in
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
	Age      string
	Name     string
	LoggedIn bool
}

type TemplateData struct {
	Error    bool
	LoggedIn bool
}

var cookiesDigests = make(map[string]string)

const (
	PRIVATE_KEY string = "Some Private Key"
)

func loginHandler(res http.ResponseWriter, req *http.Request) {

	// validating the cookies
	cookieChanged := validateCookie(req)

	setCookies(res, req)

	temp, err := template.ParseFiles("login.html")
	// Logging possible errors
	logFatalError(err)

	temp.Execute(res, createTemplateData(req, cookieChanged))
	// Logging possible errors
	logFatalError(err)
}

func createTemplateData(req *http.Request, cookieChanged bool) TemplateData {
	userDataCookie, err := req.Cookie("userData")
	if err == nil {
		var dataChanged bool = false
		bytes, err := base64.StdEncoding.DecodeString(userDataCookie.Value)
		if err != nil {
			dataChanged = true
		}
		var userData User
		err = json.Unmarshal(bytes, &userData)
		if err != nil {
			dataChanged = true
		}
		return TemplateData{
			Error:    cookieChanged || dataChanged,
			LoggedIn: userData.LoggedIn,
		}
	}
	return TemplateData{
		Error:    cookieChanged,
		LoggedIn: false,
	}
}

// Validates to make sure the cookies have been modified on client side.
func validateCookie(req *http.Request) bool {
	for key, value := range cookiesDigests {
		log.Println("Key:", key, "Value:", value)
		cookie, err := req.Cookie(key)
		if err == nil {
			log.Println("Cookie:" + cookie.String())
			// Logging the possible errors
			if value != generateDigest(cookie.Value) {
				log.Println("The following cookie has been changed: [" + key + "]")
				return true
			}
		} else {
			log.Println(err)
		}
	}
	return false
}

// Sets the cookies on the response
func setCookies(res http.ResponseWriter, req *http.Request) {

	// Generating a new ID
	id, err := uuid.NewV4()
	logFatalError(err)
	createCookie(&res, "session-fino", id.String())

	// Setting user's data on cookie
	user := User{
		Age:      req.FormValue("age"),
		Name:     req.FormValue("name"),
		LoggedIn: true,
	}

	userBytes, err := json.Marshal(user)
	logFatalError(err)
	createCookie(&res, "userData", base64.StdEncoding.EncodeToString(userBytes))
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
