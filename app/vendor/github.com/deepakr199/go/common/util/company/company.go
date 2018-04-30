package company

import "github.com/deepakr199/go/common/util"

func GetCompany(id string) string {
	companyDB := util.CompanyDB()
	return companyDB[id]
}

func SetCompany(id string, company string) string {
	companyDB := util.CompanyDB()
	companyDB[id] = company
	return company
}

func GetCompanyTotal() int {
	companyDB := util.CompanyDB()
	return len(companyDB)
}
