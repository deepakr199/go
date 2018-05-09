package service

import (
	"github.com/deepakr199/go/common/util/company"
	"github.com/golang/glog"
)

func GetCompanyTotal() int {
	glog.Info("App/Company/Service!!! Request to get Company total");
	return company.GetCompanyTotal()
}

func GetCompany(id string) string {
	glog.Info("App/Company/Service!!! Request to get Company");
	return company.GetCompany(id)
}

func SetCompany(id string, companyname string) string {
	glog.Info("App/Company/Service!!! Request to set Company");
	return company.SetCompany(id, companyname)
}