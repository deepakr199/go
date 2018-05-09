package employee

import (
	"github.com/deepakr199/go/common/util"
	"github.com/golang/glog"
)

func GetEmployee(id string) string {
	glog.Info("Common/Util/Employee!!! Request to get Employee");
	employeeDB := util.EmployeeDB()
	return employeeDB[id]
}

func SetEmployee(id string, employee string) string {
	glog.Info("Common/Util/Employee!!! Request to set Employee");
	employeeDB := util.EmployeeDB()
	employeeDB[id] = employee
	return "ok"
}

func GetEmployeeTotal() int {
	glog.Info("Common/Util/Employee!!! Request to get Employee total");
	employeeDB := util.EmployeeDB()
	return len(employeeDB)
}