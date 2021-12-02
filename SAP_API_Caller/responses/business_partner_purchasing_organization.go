package responses

type PurchasingOrganization struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			BusinessPartner                string `json:"BusinessPartner"`
			SupplierDesc                   string `json:"Supplier_desc"`
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
		} `json:"results"`
	} `json:"d"`
}
