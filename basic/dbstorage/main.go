package main

import (
	"net/http"
	"fmt"
	"github.com/deepakr199/go/basic/dbstorage/get"
	"github.com/deepakr199/go/basic/dbstorage/set"
)

func main() {
	http.HandleFunc("/basic/storage/get", getHandler)
	http.HandleFunc("/basic/storage/set", setHandler)

	http.ListenAndServe(":8000", nil)
}

func getHandler(writer http.ResponseWriter, request *http.Request) {
	key := request.URL.Query().Get("key")
	fmt.Fprintf(writer, get.Get(key))
}

func setHandler(writer http.ResponseWriter, request *http.Request) {
	key := request.URL.Query().Get("key")
	value := request.URL.Query().Get("value")
	fmt.Fprintf(writer, set.Set(key, value))
}
