package main

import (
	"exercise/favorite.quote"
	"io"
	"log"
	"net/http"
)

const PORT string = "8080"

/*
Main function to handle all in coming requests.
*/
func main() {

	http.HandleFunc("/", favoriteQuote)

	log.Println("listening to port [" + PORT + "] ...")
	http.ListenAndServe(":"+PORT, nil)
}
