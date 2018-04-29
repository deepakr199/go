package main

import (
	"fmt"
	"net/http"
	"github.com/deepakr199/go/basic/dbstorage/get"
	"github.com/deepakr199/go/basic/dbstorage/set"
)

func main() {
	http.HandleFunc("/basic/proxy/get", handleGet)
	http.HandleFunc("/basic/proxy/set", handleSet)

	http.ListenAndServe(":8001", nil)
}

func handleGet(writer http.ResponseWriter, request *http.Request) {
	key := request.URL.Query().Get("key")
	fmt.Fprintf(writer, get.Get(key))
}

func handleSet(writer http.ResponseWriter, request *http.Request) {
	key := request.URL.Query().Get("key")
	value := request.URL.Query().Get("value")
	fmt.Fprintf(writer, set.Set(key, value))
}
