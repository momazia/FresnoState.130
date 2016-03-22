/*
	Source: 		Column BP in Gradebook
	Author: 		Mohamad Mahdi Ziaee
	Description:	PROJECT STEP 8 - Allow the user to logout. Show a log-in button when
					the user is not logged-in. Show a log-out button only when the user is logged in.
					PROJECT STEP 9 - A user should not be able to access the form to upload user data when they are not logged in.
*/

package main

import (
	"encoding/base64"
	"encoding/json"
	"github.com/nu7hatch/gouuid"
	"log"
	"net/http"
	"text/template"
)

const (
	PRIVATE_KEY   string = "Some Private Key"
	SESSIONID     string = "SESSIONID"
	URL_LOGIN     string = "/login"
	URL_USERFORM  string = "/userForm"
	URL_FRONTPAGE string = "/"
)

var sessions []string

type User struct {
	UserName string
	Password string
}

type UserData struct {
	Name string
	Age  string
}

var validUsers = []User{
	{UserName: "Mahdi", Password: "123"},
	{UserName: "Amin", Password: "321"},
}

// Login handler which validates if teh user logged in is authorized. If not, it shows an error message.
func loginHandler(res http.ResponseWriter, req *http.Request) {
	var isInvalidUser bool = false
	if req.Method == "POST" {
		if isValidUser(req) {
			// Settin the user in session and going to userForm
			setSession(res)
			http.Redirect(res, req, URL_USERFORM, http.StatusFound)
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

// Creates a new UUID and sets it on the cookie and the session.
func setSession(res http.ResponseWriter) {
	// Generating a new ID and storing it on the session
	newUuid, err := uuid.NewV4()
	sessionId := newUuid.String()
	logFatalError(err)
	sessions = append(sessions, sessionId)
	createCookie(&res, SESSIONID, sessionId)
}

// Checks and see if the username and password are matching the hardcoded values in validUsers (this represents DB)
func isValidUser(req *http.Request) bool {
	formUserName := req.FormValue("userName")
	formPassword := req.FormValue("password")
	for _, user := range validUsers {
		if user.UserName == formUserName && user.Password == formPassword {
			return true
		}
	}
	return false
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
	// Setting the cookie on the response back to the client
	http.SetCookie(*res, cookie)
}

// User form handler which shows the user's age and name
func userFormHandler(res http.ResponseWriter, req *http.Request) {
	var userDataBytes []byte
	if req.Method == "POST" {
		// Setting the form data into cookie
		userData := UserData{
			Age:  req.FormValue("age"),
			Name: req.FormValue("name"),
		}
		var err error
		userDataBytes, err = json.Marshal(userData)
		logFatalError(err)
		createCookie(&res, "USERDATA", base64.StdEncoding.EncodeToString(userDataBytes))
	}

	temp, err := template.ParseFiles("userForm.html")
	logFatalError(err)

	temp.Execute(res, string(userDataBytes))
	logFatalError(err)
}

// Logs out the user by clearing cookie and the session ID associated with it.
func logoutHandler(res http.ResponseWriter, req *http.Request) {
	// Clearing the cookie.
	cookie, err := req.Cookie(SESSIONID)
	if err == nil {
		// Clearing the cookie
		cookie.MaxAge = -1
		http.SetCookie(res, cookie)
		// Clearing the session
		clearSession(cookie.Value)
	} else {
		log.Println(err)
	}
	// Redirecting the user to front page.
	http.Redirect(res, req, URL_FRONTPAGE, http.StatusFound)
}

// Clears the sessionId given from sessions
func clearSession(sessionId string) {
	var index int = 0
	for i, session := range sessions {
		if session == sessionId {
			index = i
		}
	}
	sessions = append(sessions[:index], sessions[index+1:]...)
}

// The blank landing page which simply redirects the user to userForm
func landingPageHandler(res http.ResponseWriter, req *http.Request) {
	http.Redirect(res, req, URL_USERFORM, http.StatusFound)
	return
}

func main() {
	// Ignoring favicon.ico
	http.Handle("/favicon.ico", http.NotFoundHandler())

	// Setting the landing page
	http.Handle(URL_FRONTPAGE, authenticationFilter(http.HandlerFunc(landingPageHandler)))

	// Setting the handler for login
	http.HandleFunc(URL_LOGIN, loginHandler)

	// Setting the handler for logout
	http.HandleFunc("/logout", logoutHandler)

	// Setting the handler for userForm
	http.Handle(URL_USERFORM, authenticationFilter(http.HandlerFunc(userFormHandler)))

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

// The main authentication filter which is to be used for all the handlers wherever the user must be authenticated.
// If user does not have access, it will redirect him to the front page.
func authenticationFilter(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		if !isUserLoggedIn(responseWriter, request) {
			http.Redirect(responseWriter, request, URL_LOGIN, http.StatusFound)
			return
		}
		handler.ServeHTTP(responseWriter, request)
	})
}

// Checks to see if the user is logged in by looking at the sessionID stored on the request cookie
func isUserLoggedIn(res http.ResponseWriter, req *http.Request) bool {
	sessionIdCookie, err := req.Cookie(SESSIONID)
	if err != nil {
		log.Println("Error reading SESSIONID:" + err.Error())
		return false
	}
	for _, session := range sessions {
		if session == sessionIdCookie.Value {
			return true
		}
	}
	return false
}
