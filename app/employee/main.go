package main

import (
	"net/http"
	"fmt"
	"github.com/deepakr199/go/app/employee/service"
	"github.com/golang/glog"
	"strconv"
)

func main() {
	http.HandleFunc("/app/employee/total", getEmployeeTotalHandler)
	http.HandleFunc("/app/employee/get", getEmployeeHandler)
	http.HandleFunc("/app/employee/set", setEmployeeHandler)

	glog.Info("Running Employee App at 8003");
	http.ListenAndServe(":8003", nil)
}

func getEmployeeTotalHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, strconv.Itoa(service.GetEmployeeTotal()))
}

func getEmployeeHandler(writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")
	fmt.Fprintf(writer, service.GetEmployee(id))
}

func setEmployeeHandler(writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")
	employee := request.URL.Query().Get("employee")
	fmt.Fprintf(writer, service.SetEmployee(id, employee))
}