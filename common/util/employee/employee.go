package employee

import "github.com/deepakr199/go/common/util"

func GetEmployee(id string) string {
	employeeDB := util.EmployeeDB()
	return employeeDB[id]
}

func SetEmployee(id string, employee string) string {
	employeeDB := util.EmployeeDB()
	employeeDB[id] = employee
	return "ok"
}

func GetEmployeeTotal() int {
	employeeDB := util.EmployeeDB()
	return len(employeeDB)
}