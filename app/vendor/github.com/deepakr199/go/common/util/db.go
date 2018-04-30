package util

var companyMap = make(map[string]string)
var employeeMap = make(map[string]string)

func CompanyDB() map[string]string {
	return companyMap
}

func EmployeeDB() map[string]string {
	return employeeMap
}
