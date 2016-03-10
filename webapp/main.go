package main

import (
	"github.com/momazia/GoTraining/exercise/favorite.quote"
	"log"
	"net/http"
)

const PORT string = "8080"

/*
Main function to handle all in coming requests.
*/
func main() {

	http.HandleFunc("/", favoriteQuote.FavoriteQuote)

	log.Println("listening to port [" + PORT + "] ...")
	http.ListenAndServe(":"+PORT, nil)
}
