/*
	Source: 		Column BA in Gradebook
	Author: 		Mohamad Mahdi Ziaee
	Description:	Create a webpage which uses a cookie to track the number of visits of a user. Display the number of visits.
					Make sure that the favicon.ico requests are not also incrementing the number of visits.
*/
package main

import (
	"log"
	"net/http"
	"strconv"
)

func cookieFunc(res http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("Counter")
	logError(err)

	if cookie == nil {
		http.SetCookie(res, &http.Cookie{
			Name:  "Counter",
			Value: "1"})
	} else {
		counter, err := strconv.Atoi(cookie.Value)
		logError(err)
		counter++
		
		http.SetCookie(res, &http.Cookie{
			Name:  "Counter",
			Value: strconv.Itoa(counter)})
	}

}

func main() {

	// Registering the URL path and binding it to userNameForm handler
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", cookieFunc)

	// Setting the listener on port 8080
	log.Println("Listening to 8080 ...")
	http.ListenAndServe(":8080", nil)
}

// Logs the error given into log
func logError(err error) {
	if err != nil {
		log.Println(err)
	}
}
