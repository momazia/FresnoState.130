/*
	Source: 		Column CB in Gradebook
	Author: 		Mohamad Mahdi Ziaee
	Description:	Create a web app that uses a google cloud storage query (storage.Query) and demonstrates the functionality of the query's delimeter field
*/

package storage

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	storageLog "google.golang.org/appengine/log"
	"google.golang.org/cloud/storage"
	"log"
	"net/http"
	"os"
	"text/template"
)

const BUCKET_NAME = "gotraining-1271.appspot.com"

func init() {
	http.HandleFunc("/", handler)
}

func handler(res http.ResponseWriter, req *http.Request) {

	// Creating new context and client.
	ctx := appengine.NewContext(req)
	client, err := storage.NewClient(ctx)
	logStorageError(ctx, "Could not create a new client", err)
	defer client.Close()

	//Parsing the template
	temp, err := template.ParseFiles("index.html")

	// Logging possible errors
	logError(err)

	// Executing the template using the constant
	err = temp.Execute(os.Stdout, getPhotoNames(ctx, client))
}

func getPhotoNames(ctx context.Context, client *storage.Client) []string {

	query := &storage.Query{
		Delimiter: "/",
		Prefix:    "photos/",
	}
	objs, err := client.Bucket(BUCKET_NAME).List(ctx, query)
	logError(err)

	var names []string
	for _, result := range objs.Results {
		names = append(names, result.Name)
	}
	return names
}

// Logs the error given into log
func logError(err error) {
	if err != nil {
		log.Println(err)
	}
}

// Logs the error given into storage log
func logStorageError(ctx context.Context, errMessage string, err error) {
	if err != nil {
		storageLog.Errorf(ctx, errMessage, err)
	}
}
