package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-business-partner-reads-supplier/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetBPSupplier(businessPartner, businessPartnerRole, addressID, purchasingOrganization, supplier, companyCode string) {
	wg := &sync.WaitGroup{}

	wg.Add(4)
	func() {
		c.Role(businessPartner, businessPartnerRole)
		wg.Done()
	}()
	func() {
		c.Address(businessPartner, addressID)
		wg.Done()
	}()
	func() {
		c.PurchasingOrganization(supplier, purchasingOrganization)
		wg.Done()
	}()
	func() {
		c.Company(supplier, companyCode)
		wg.Done()
	}()
	wg.Wait()
}

func (c *SAPAPICaller) Role(businessPartner, businessPartnerRole string) {
	data, err := c.callBPSupplierSrvAPIRequirementRole("A_BusinessPartnerRole", businessPartner, businessPartnerRole)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callBPSupplierSrvAPIRequirementRole(api, businessPartner, businessPartnerRole string) (*sap_api_output_formatter.Role, error) {
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithRole(req, businessPartner, businessPartnerRole)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToRole(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Address(businessPartner, addressID string) {
	data, err := c.callBPSupplierSrvAPIRequirementAddress("A_BusinessPartnerAddress", businessPartner, addressID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callBPSupplierSrvAPIRequirementAddress(api, businessPartner, addressID string) (*sap_api_output_formatter.Address, error) {
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithAddress(req, businessPartner, addressID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToAddress(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) PurchasingOrganization(supplier, purchasingOrganization string) {
	data, err := c.callBPSupplierSrvAPIRequirementPurchasingOrganization("A_SupplierPurchasingOrg", supplier, purchasingOrganization)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callBPSupplierSrvAPIRequirementPurchasingOrganization(api, supplier, purchasingOrganization string) (*sap_api_output_formatter.PurchasingOrganization, error) {
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithPurchasingOrganization(req, supplier, purchasingOrganization)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToPurchasingOrganization(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Company(supplier, companyCode string) {
	data, err := c.callBPSupplierSrvAPIRequirementCompany("A_SupplierCompany", supplier, companyCode)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callBPSupplierSrvAPIRequirementCompany(api, supplier, companyCode string) (*sap_api_output_formatter.Company, error) {
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithCompany(req, supplier, companyCode)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToCompany(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithRole(req *http.Request, businessPartner, businessPartnerRole string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("BusinessPartner eq '%s' and BusinessPartnerRole eq '%s'", businessPartner, businessPartnerRole))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithAddress(req *http.Request, businessPartner, addressID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("BusinessPartner eq '%s' and AddressID eq '%s'", businessPartner, addressID))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithPurchasingOrganization(req *http.Request, supplier, purchasingOrganization string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Supplier eq '%s' and PurchasingOrganization eq '%s'", supplier, purchasingOrganization))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithCompany(req *http.Request, supplier, companyCode string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Supplier eq '%s' and CompanyCode eq '%s'", supplier, companyCode))
	req.URL.RawQuery = params.Encode()
}
