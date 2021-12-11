package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo      string `json:"document_no"`
		BusinessPartner string `json:"deliver_to"`
		Quantity        string `json:"quantity"`
		PickedQuantity  string `json:"picked_quantity"`
		Price           string `json:"price"`
		Batch           string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema     string `json:"api_schema"`
	MaterialCode  string `json:"material_code"`
	Plant         string `json:"plant/supplier"`
	Stock         string `json:"stock"`
	DocumentType  string `json:"document_type"`
	DocumentNo    string `json:"document_no"`
	PlannedDate   string `json:"planned_date"`
	ValidatedDate string `json:"validated_date"`
	Deleted       bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey   string `json:"connection_key"`
	Result          bool   `json:"result"`
	RedisKey        string `json:"redis_key"`
	Filepath        string `json:"filepath"`
	BusinessPartner struct {
		BusinessPartner     string `json:"BusinessPartner"`
		BusinessPartnerRole string `json:"BusinessPartnerRole"`
		ValidFrom           string `json:"ValidFrom"`
		ValidTo             string `json:"ValidTo"`
		Address             struct {
			AddressID         string `json:"AddressID"`
			ValidityStartDate string `json:"ValidityStartDate"`
			ValidityEndDate   string `json:"ValidityEndDate"`
			Country           string `json:"Country"`
			Region            string `json:"Region"`
			StreetName        string `json:"StreetName"`
			CityName          string `json:"CityName"`
			PostalCode        string `json:"PostalCode"`
			Language          string `json:"Language"`
		} `json:"AddressID"`
		PurchasingOrganization struct {
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
		} `json:"Purchasing_Organization"`
		Company struct {
			CompanyCode                 string `json:"CompanyCode"`
			PaymentBlockingReason       string `json:"PaymentBlockingReason"`
			PaymentMethodsList          string `json:"PaymentMethodsList"`
			PaymentTerms                string `json:"PaymentTerms"`
			ClearCustomerSupplier       bool   `json:"ClearCustomerSupplier"`
			HouseBank                   string `json:"HouseBank"`
			ReconciliationAccount       string `json:"ReconciliationAccount"`
			SupplierIsBlockedForPosting bool   `json:"SupplierIsBlockedForPosting"`
			DeletionIndicator           bool   `json:"DeletionIndicator"`
		} `json:"Company"`
	} `json:"business_partner"`
	APISchema string   `json:"api_schema"`
	Accepter  []string `json:"accepter"`
	Supplier  string   `json:"business_partner_code"`
	Deleted   bool     `json:"deleted"`
}
