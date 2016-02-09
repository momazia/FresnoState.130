package favoriteQuote

import (
	"io"
	"net/http"
)

const QUOTE string = "A man's enjoyment of all good things is in exact proportion to the pains he has undergone to gain them. \n - Cyrus the Great"

/*
Prints out a famous quote into http response.
*/
func FavoriteQuote(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, QUOTE)
}
