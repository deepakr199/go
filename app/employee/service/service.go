package service

import (
	"github.com/deepakr199/go/common/util/employee"
	"github.com/golang/glog"
)

func GetEmployeeTotal() int {
	glog.Info("App/Employee/Service!!! Request to get Employee total");
	return employee.GetEmployeeTotal()
}

func GetEmployee(id string) string {
	glog.Info("App/Employee/Service!!! Request to get Employee");
	return employee.GetEmployee(id)
}

func SetEmployee(id string, employeename string) string {
	glog.Info("App/Employee/Service!!! Request to set Employee");
	return employee.SetEmployee(id, employeename)
}