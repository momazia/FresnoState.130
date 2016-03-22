package main

import (

)
var sessions = []string{}

func authenticationFilter(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !isUserLoggedIn(res, req) {
			http.Redirect(res, req, "/login", http.StatusFound)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

func isUserLoggedIn(res http.ResponseWriter, req *http.Request) bool {
	sessionIdCookie, err := req.Cookie("SESSIONID")
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