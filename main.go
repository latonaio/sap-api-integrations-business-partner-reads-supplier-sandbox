package main

import (
	sap_api_caller "sap-api-integrations-business-partner-reads-supplier/SAP_API_Caller"
	"sap-api-integrations-business-partner-reads-supplier/sap_api_input_reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs//SDC_Business_Partner_Supplier_Company_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"General", "Role", "Address", "Bank", "BPName", "PurchasingOrganization", "Company",
		}
	}

	caller.AsyncGetBPSupplier(
		inoutSDC.BusinessPartner.BusinessPartner,
		inoutSDC.BusinessPartner.Role.BusinessPartnerRole,
		inoutSDC.BusinessPartner.Address.AddressID,
		inoutSDC.BusinessPartner.Bank.BankCountryKey,
		inoutSDC.BusinessPartner.Bank.BankNumber,
		inoutSDC.BusinessPartner.BusinessPartnerName,
		inoutSDC.BusinessPartner.SupplierData.Supplier,
		inoutSDC.BusinessPartner.SupplierData.PurchasingOrganization.PurchasingOrganization,
		inoutSDC.BusinessPartner.Company.CompanyCode,
		accepter,
	)
}
