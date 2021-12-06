package sap_api_output_formatter

type BusinessPartner struct {
	 ConnectionKey       string `json:"connection_key"`
	 Result              bool   `json:"result"`
	 RedisKey            string `json:"redis_key"`
	 Filepath            string `json:"filepath"`
	 APISchema           string `json:"api_schema"`
	 Deleted             bool   `json:"deleted"`    
}

type Role struct {
     BusinessPartner     string   `json:"BusinessPartner"`
     BusinessPartnerRole string   `json:"BusinessPartnerRole"`
     ValidFrom           string   `json:"ValidFrom"`
     ValidTo             string   `json:"ValidTo"`
}

type Address struct {
     BusinessPartner     string `json:"BusinessPartner"`
     AddressID           string `json:"AddressID"`
     ValidityStartDate   string `json:"ValidityStartDate"`
     ValidityEndDate     string `json:"ValidityEndDate"`
     Country             string `json:"Country"`
     Region              string `json:"Region"`
     StreetName          string `json:"StreetName"`
     CityName            string `json:"CityName"`
     PostalCode          string `json:"PostalCode"`
     Language            string `json:"Language"`
}

type PurchaseOrganization struct {
     Supplier                       string `json:"BusinessPartner"`
     PurchasingOrganization         string `json:"PurchasingOrganization"`
     IncotermsClassification        string `json:"IncotermsClassification"`
     InvoiceIsGoodsReceiptBased     bool   `json:"InvoiceIsGoodsReceiptBased"`
     PaymentTerms                   string `json:"PaymentTerms"`
     PurOrdAutoGenerationIsAllowed  bool   `json:"PurOrdAutoGenerationIsAllowed"`
     PurchaseOrderCurrency          string `json:"PurchaseOrderCurrency"`
     PurchasingGroup                string `json:"PurchasingGroup"`
     ShippingCondition              string `json:"ShippingCondition"`
     SupplierPhoneNumber            string `json:"SupplierPhoneNumber"`
     SupplierRespSalesPersonName    string `json:"SupplierRespSalesPersonName"`
     PurchasingIsBlockedForSupplier bool   `json:"PurchasingIsBlockedForSupplier"`
     DeletionIndicator              bool   `json:"DeletionIndicator"`
}

type Company struct {
     Supplier                    string `json:"BusinessPartner"`
     CompanyCode                 string `json:"CompanyCode"`
     PaymentBlockingReason       string `json:"PaymentBlockingReason"`
     PaymentMethodsList          string `json:"PaymentMethodsList"`
     PaymentTerms                string `json:"PaymentTerms"`
     ClearCustomerSupplier       bool   `json:"ClearCustomerSupplier"`
     HouseBank                   string `json:"HouseBank"`
     ReconciliationAccount       string `json:"ReconciliationAccount"`
     SupplierIsBlockedForPosting bool   `json:"SupplierIsBlockedForPosting"`
     DeletionIndicator           bool   `json:"DeletionIndicator"`
}
