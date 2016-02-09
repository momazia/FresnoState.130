package main

import (
	"io"
	"log"
	"net/http"
)

const QUOTE string = "A man's enjoyment of all good things is in exact proportion to the pains he has undergone to gain them. \n - Cyrus the Great"
const PORT string = "8080"

/*
Prints out a famous quote into http response.
*/
func favoriteQuote(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, QUOTE)
}

/*
Main function to handle all in coming requests.
*/
func main() {
	http.HandleFunc("/", favoriteQuote)
	log.Println("listening to port [" + PORT + "] ...")
	http.ListenAndServe(":"+PORT, nil)
}
