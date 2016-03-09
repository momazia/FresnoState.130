/*
	Source: 		Column BJ in Gradebook
	Author: 		Mohamad Mahdi Ziaee
	Description:	PROJECT STEP 2 - have the application write a cookie called "session-fino" with a UUID. The cookie should serve HttpOnly and
					you should have the "Secure" flag set also though comment the "Secure" flag out as we're not using https.
*/

package main

import (
	"github.com/nu7hatch/gouuid"
	"log"
	"net/http"
	"text/template"
)

func login(res http.ResponseWriter, req *http.Request) {

	createCookie(&res)

	temp, err := template.ParseFiles("login.html")
	// Logging possible errors
	logFatalError(err)

	temp.Execute(res, nil)
	// Logging possible errors
	logFatalError(err)
}

func createCookie(res *http.ResponseWriter) {
	// Generating a new ID
	id, err := uuid.NewV4()
	// Logging possible errors
	logFatalError(err)

	// Setting the cookie
	cookie := &http.Cookie{
		Name:     "session-fino",
		Value:    id.String(),
//		Secure: true,
		HttpOnly: true,
	}

	// Setting the cookie on the response back to the client
	http.SetCookie(*res, cookie)
}

func main() {

	// Ignoring favicon.ico
	http.Handle("/favicon.ico", http.NotFoundHandler())

	// Setting the handler for login
	http.HandleFunc("/", login)

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
