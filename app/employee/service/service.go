package service

import "github.com/deepakr199/go/common/util/employee"

func GetEmployeeTotal() int {
	return employee.GetEmployeeTotal()
}

func GetEmployee(id string) string {
	return employee.GetEmployee(id)
}

func SetEmployee(id string, employeename string) string {
	return employee.SetEmployee(id, employeename)
}