package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-business-partner-reads-supplier/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

func ConvertToRole(raw []byte, l *logger.Logger) (*Role, error) {
	pm := &responses.Role{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Role. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &Role{
		BusinessPartner:     data.BusinessPartner,
		BusinessPartnerRole: data.BusinessPartnerRole,
		ValidFrom:           data.ValidFrom,
		ValidTo:             data.ValidTo,
	}, nil
}

func ConvertToAddress(raw []byte, l *logger.Logger) (*Address, error) {
	pm := &responses.Address{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Address. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &Address{
		BusinessPartner:   data.BusinessPartner,
		AddressID:         data.AddressID,
		ValidityStartDate: data.ValidityStartDate,
		ValidityEndDate:   data.ValidityEndDate,
		Country:           data.Country,
		Region:            data.Region,
		StreetName:        data.StreetName,
		CityName:          data.CityName,
		PostalCode:        data.PostalCode,
		Language:          data.Language,
	}, nil
}

func ConvertToPurchasingOrganization(raw []byte, l *logger.Logger) (*PurchasingOrganization, error) {
	pm := &responses.PurchasingOrganization{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to PurchasingOrganization. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &PurchasingOrganization{
		Supplier:                       data.Supplier,
        PurchasingOrganization:         data.PurchasingOrganization,
		IncotermsClassification:        data.IncotermsClassification,
		InvoiceIsGoodsReceiptBased:     data.InvoiceIsGoodsReceiptBased,
		PaymentTerms:                   data.PaymentTerms,
		PurOrdAutoGenerationIsAllowed:  data.PurOrdAutoGenerationIsAllowed,
		PurchaseOrderCurrency:          data.PurchaseOrderCurrency,
		PurchasingGroup:                data.PurchasingGroup,
		ShippingCondition:              data.ShippingCondition,
		SupplierPhoneNumber:            data.SupplierPhoneNumber,
		SupplierRespSalesPersonName:    data.SupplierRespSalesPersonName,
		PurchasingIsBlockedForSupplier: data.PurchasingIsBlockedForSupplier,
		DeletionIndicator:              data.DeletionIndicator,
	}, nil
}

func ConvertToCompany(raw []byte, l *logger.Logger) (*Company, error) {
	pm := &responses.Company{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Company. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &Company{
		Supplier:                       data.Supplier,
		CompanyCode:                    data.CompanyCode,
		PaymentBlockingReason:          data.PaymentBlockingReason,
		PaymentMethodsList:             data.PaymentMethodsList,
		PaymentTerms:                   data.PaymentTerms,
		ClearCustomerSupplier:          data.ClearCustomerSupplier,
		HouseBank:                      data.HouseBank,
		ReconciliationAccount:          data.ReconciliationAccount,
		SupplierIsBlockedForPosting:    data.SupplierIsBlockedForPosting,
		DeletionIndicator:              data.DeletionIndicator,
	}, nil
}
