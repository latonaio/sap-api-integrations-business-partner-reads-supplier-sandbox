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

func (c *SAPAPICaller) AsyncGetBPSupplier(businessPartner, businessPartnerRole, addressID, bankCountryKey, bankNumber, supplier, purchasingOrganization, companyCode string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "General":
			func() {
				c.General(businessPartner)
				wg.Done()
			}()
		case "Role":
			func() {
				c.Role(businessPartner, businessPartnerRole)
				wg.Done()
			}()
		case "Address":
			func() {
				c.Address(businessPartner, addressID)
				wg.Done()
			}()
		case "Bank":
			func() {
				c.Bank(businessPartner, bankCountryKey, bankNumber)
				wg.Done()
			}()
		case "Supplier":
			func() {
				c.Supplier(supplier)
				wg.Done()
			}()
		case "PurchasingOrganization":
			func() {
				c.PurchasingOrganization(supplier, purchasingOrganization)
				wg.Done()
			}()
		case "Company":
			func() {
				c.Company(supplier, companyCode)
				wg.Done()
			}()

		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) General(businessPartner string) {
	generalData, err := c.callBPSrvAPIRequirementGeneral("A_BusinessPartner", businessPartner)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(generalData)

	roleData, err := c.callToRole(generalData[0].ToRole)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(roleData)

	addressData, err := c.callToAddress(generalData[0].ToAddress)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(addressData)
	
	bankData, err := c.callToBank(generalData[0].ToBank)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(bankData)

	supplierData, err := c.callToSupplier(generalData[0].ToSupplier)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(supplierData)

	purchasingOrganizationData, err := c.callToPurchasingOrganization(supplierData.ToPurchasingOrganization)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(purchasingOrganizationData)
	
	partnerFunctionData, err := c.callToPartnerFunction(purchasingOrganizationData[0].ToPartnerFunction)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(partnerFunctionData)

	companyData, err := c.callToCompany(supplierData.ToCompany)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(companyData)

}

func (c *SAPAPICaller) callBPSrvAPIRequirementGeneral(api, businessPartner string) ([]sap_api_output_formatter.General, error) {
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithGeneral(req, businessPartner)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToGeneral(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToRole(url string) ([]sap_api_output_formatter.ToRole, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToRole(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToAddress(url string) ([]sap_api_output_formatter.ToAddress, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToAddress(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToBank(url string) ([]sap_api_output_formatter.ToBank, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToBank(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToSupplier(url string) (*sap_api_output_formatter.ToSupplier, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToSupplier(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToPurchasingOrganization(url string) ([]sap_api_output_formatter.ToPurchasingOrganization, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToPurchasingOrganization(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToPartnerFunction(url string) ([]sap_api_output_formatter.ToPartnerFunction, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToPartnerFunction(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToCompany(url string) ([]sap_api_output_formatter.ToCompany, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToCompany(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Role(businessPartner, businessPartnerRole string) {
	data, err := c.callBPSupplierSrvAPIRequirementRole("A_BusinessPartnerRole", businessPartner, businessPartnerRole)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callBPSupplierSrvAPIRequirementRole(api, businessPartner, businessPartnerRole string) ([]sap_api_output_formatter.Role, error) {
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

func (c *SAPAPICaller) callBPSupplierSrvAPIRequirementAddress(api, businessPartner, addressID string) ([]sap_api_output_formatter.Address, error) {
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

func (c *SAPAPICaller) Bank(businessPartner, bankCountryKey, bankNumber string) {
	data, err := c.callBPCustomerSrvAPIRequirementBank("A_BusinessPartnerBank", businessPartner, bankCountryKey, bankNumber)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callBPCustomerSrvAPIRequirementBank(api, businessPartner, bankCountryKey, bankNumber string) ([]sap_api_output_formatter.Bank, error) {
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithBank(req, businessPartner, bankCountryKey, bankNumber)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToBank(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Supplier(supplier string) {
	supplierData, err := c.callBPSupplierSrvAPIRequirementSupplier("A_Supplier", supplier)
	
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(supplierData)

	purchasingOrganizationData, err := c.callToPurchasingOrganization(supplierData[0].ToPurchasingOrganization)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(purchasingOrganizationData)
	
	partnerFunctionData, err := c.callToPartnerFunction(purchasingOrganizationData[0].ToPartnerFunction)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(partnerFunctionData)

	companyData, err := c.callToCompany(supplierData[0].ToCompany)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(companyData)

}

func (c *SAPAPICaller) callBPSupplierSrvAPIRequirementSupplier(api, supplier string) ([]sap_api_output_formatter.Supplier, error) {
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithSupplier(req, supplier)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToSupplier(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToPurchasingOrganization2(url string) ([]sap_api_output_formatter.ToPurchasingOrganization, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToPurchasingOrganization(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToPartnerFunction2(url string) ([]sap_api_output_formatter.ToPartnerFunction, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToPartnerFunction(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToCompany2(url string) ([]sap_api_output_formatter.ToCompany, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToCompany(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) PurchasingOrganization(supplier, purchasingOrganization string) {
	purchasingOrganizationData, err := c.callBPSupplierSrvAPIRequirementPurchasingOrganization("A_SupplierPurchasingOrg", supplier, purchasingOrganization)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(purchasingOrganizationData)

	partnerFunctionData, err := c.callToPartnerFunction(purchasingOrganizationData[0].ToPartnerFunction)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(partnerFunctionData)
}

func (c *SAPAPICaller) callBPSupplierSrvAPIRequirementPurchasingOrganization(api, supplier, purchasingOrganization string) ([]sap_api_output_formatter.PurchasingOrganization, error) {
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

func (c *SAPAPICaller) callToPartnerFunction3(url string) ([]sap_api_output_formatter.ToPartnerFunction, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToPartnerFunction(byteArray, c.log)
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

func (c *SAPAPICaller) callBPSupplierSrvAPIRequirementCompany(api, supplier, companyCode string) ([]sap_api_output_formatter.Company, error) {
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

func (c *SAPAPICaller) getQueryWithGeneral(req *http.Request, businessPartner string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("BusinessPartner eq '%s'", businessPartner))
	req.URL.RawQuery = params.Encode()
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

func (c *SAPAPICaller) getQueryWithBank(req *http.Request, businessPartner, bankCountryKey, bankNumber string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("BusinessPartner eq '%s' and BankCountryKey eq '%s' and BankNumber eq '%s", businessPartner, bankCountryKey, bankNumber))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithSupplier(req *http.Request, supplier string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Supplier eq '%s'", supplier))
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
