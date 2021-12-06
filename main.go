package main

import (
	sap_api_caller "sap-api-integrations-business-partner-reads-supplier/SAP_API_Caller"
	"sap-api-integrations-business-partner-reads-supplier/sap_api_input_reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs//SDC_Business_Partner_Supplier_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	caller.AsyncGetBPSupplier(
		inoutSDC.BusinessPartner.BusinessPartner,
        inoutSDC.BusinessPartner.BusinessPartnerRole,
		inoutSDC.BusinessPartner.PurchasingOrganization.PurchasingOrganization,
		inoutSDC.BusinessPartner.Company.CompanyCode,
	)
}
