package service

import "github.com/deepakr199/go/common/util/company"

func GetCompanyTotal() int {
	return company.GetCompanyTotal()
}

func GetCompany(id string) string {
	return company.GetCompany(id)
}

func SetCompany(id string, companyname string) string {
	 return company.SetCompany(id, companyname)
}