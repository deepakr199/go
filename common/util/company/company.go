package company

import (
	"github.com/deepakr199/go/common/util"
	"github.com/golang/glog"
)

func GetCompany(id string) string {
	glog.Info("Common/Util/Company!!! Request to get Company");
	companyDB := util.CompanyDB()
	return companyDB[id]
}

func SetCompany(id string, company string) string {
	glog.Info("Common/Util/Company!!! Request to set Company");
	companyDB := util.CompanyDB()
	companyDB[id] = company
	return company
}

func GetCompanyTotal() int {
	glog.Info("Common/Util/Company!!! Request to get Company total");
	companyDB := util.CompanyDB()
	return len(companyDB)
}
