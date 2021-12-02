package main

import (
	sap_api_caller "sap-api-integrations-business-partner-reads-supplier/SAP_API_Caller"
	"sap-api-integrations-business-partner-reads-supplier/file_reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := file_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs//SDC_Business_Partner_Supplier_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	caller.AsyncGetBusinessPartnerSupplier(
		inoutSDC.BusinessPartner.BusinessPartnerRole,
		inoutSDC.BusinessPartner.VaridityEndDate,
		inoutSDC.BusinessPartner.PurchasingOrganization,
		inoutSDC.BusinessPartner.CompanyCode,
	)
}