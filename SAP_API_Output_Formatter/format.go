package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-business-partner-reads-supplier/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
)

func ConvertToBPSupplier(raw []byte, l *logger.Logger) *BPSupplier {
	pm := &responses.BPSupplier{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		l.Error(err)
		return nil
	}
	if len(pm.D.Results) == 0 {
		l.Error("Result data is not exist.")
		return nil
	}
	if len(pm.D.Results) > 1 {
		l.Error("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &BPSupplier{
		BusinessPartner                data.BusinessPartner,
		BusinessPartnerRole            data.BusinessPartnerRole,
		ValidFrom                      data.ValidFrom,
		ValidTo                        data.ValidTo,
		AddressID                      data.AddressID,
		ValidityStartDate              data.ValidityStartDate,
		ValidityEndDate                data.ValidityEndDate,
		Country                        data.Country,
		Region                         data.Region,
		StreetName                     data.StreetName,
		CityName                       data.CityName,
		PostalCode                     data.PostalCode,
		Language                       data.Language,
		ToEmailAddress                 data.to_EmailAddress,
		ToFaxNumber                    data.to_FaxNumber,
		ToMobilePhoneNumber            data.to_MobilePhoneNumber,
		ToPhoneNumber                  data.to_PhoneNumber,
		ToURLAddress                   data.to_URLAddress,
		SupplierDesc                   data.Supplier_desc,
		PurchasingOrganization         data.PurchasingOrganization,
		IncotermsClassification        data.IncotermsClassification,
		InvoiceIsGoodsReceiptBased     data.InvoiceIsGoodsReceiptBased,
		PaymentTerms                   data.PaymentTerms,
		PurOrdAutoGenerationIsAllowed  data.PurOrdAutoGenerationIsAllowed,
		PurchaseOrderCurrency          data.PurchaseOrderCurrency,
		PurchasingGroup                data.PurchasingGroup,
		ShippingCondition              data.ShippingCondition,
		SupplierPhoneNumber            data.SupplierPhoneNumber,
		SupplierRespSalesPersonName    data.SupplierRespSalesPersonName,
		PurchasingIsBlockedForSupplier data.PurchasingIsBlockedForSupplier,
		DeletionIndicator              data.DeletionIndicator,
		SupplierDesc                   data.Supplier_desc,
		CompanyCode                    data.CompanyCode,
		PaymentBlockingReason          data.PaymentBlockingReason,
		PaymentMethodsList             data.PaymentMethodsList,
		PaymentTerms                   data.PaymentTerms,
		ClearCustomerSupplier          data.ClearCustomerSupplier,
		HouseBank                      data.HouseBank,
		ReconciliationAccount          data.ReconciliationAccount,
		SupplierIsBlockedForPosting    data.SupplierIsBlockedForPosting,
		DeletionIndicator              data.DeletionIndicator,
	}
}
