package main

import (
	"net/http"
	"fmt"
	"github.com/deepakr199/go/app/company/service"
	"strconv"
	"github.com/golang/glog"
)

func main() {
	http.HandleFunc("/app/company/total", getCompanyTotalHandler)
	http.HandleFunc("/app/company/get", getCompanyHandler)
	http.HandleFunc("/app/company/set", setCompanyHandler)

	glog.Info("Company App running at 8002");
	http.ListenAndServe(":8002", nil)
}

func getCompanyTotalHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, strconv.Itoa(service.GetCompanyTotal()))
}

func getCompanyHandler(writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")
	fmt.Fprintf(writer, service.GetCompany(id))
}

func setCompanyHandler(writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")
	company := request.URL.Query().Get("company")
	fmt.Fprintf(writer, service.SetCompany(id, company))
}