# sap-api-integrations-business-partner-reads-supplier
sap-api-integrations-business-partner-reads-supplier は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で BP(ビジネスパートナ) - 仕入先 データを取得するマイクロサービスです。    
sap-api-integrations-business-partner-reads-supplier には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-business-partner-reads-supplier は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。   
https://api.sap.com/api/OP_API_BUSINESS_PARTNER_SRV/overview   

## 動作環境  
sap-api-integrations-business-partner-reads-supplier は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）    

## クラウド環境での利用
sap-api-integrations-business-partner-reads-supplier は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-business-partner-reads-supplier が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_BUSINESS_PARTNER_SRV/overview    
* APIサービス名(=baseURL): API_BUSINESS_PARTNER

## 本レポジトリ に 含まれる API名
sap-api-integrations-business-partner-reads-supplier には、次の API をコールするためのリソースが含まれています。  

* A_BusinessPartnerRole（ビジネスパートナ - ロール）
* A_BusinessPartnerAddress（ビジネスパートナ - アドレス）
* A_Supplier('{Supplier}')/to_SupplierPurchasingOrg（ビジネスパートナ - 仕入先購買組織）
* A_Supplier('{Supplier}')/to_SupplierCompany（ビジネスパートナ - 仕入先会社コード）

## API への 値入力条件 の 初期値
sap-api-integrations-business-partner-reads-supplier において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.BusinessPartner.BusinessPartner（ビジネスパートナ）
* inoutSDC.BusinessPartner.BusinessPartnerRole（ビジネスパートナロール）
* inoutSDC.BusinessPartner.PurchasingOrganization.PurchasingOrganization（購買組織）
* inoutSDC.BusinessPartner.Company.CompanyCode（会社コード）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header" が指定されています。    
  
```
  "api_schema": "sap.s4.beh.businesspartner.v1.BusinessPartner.Created.v1",
  "accepter": ["Role"],
  "business_partner_code":  "9980002060",
  "deleted":  false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
  "api_schema": "sap.s4.beh.businesspartner.v1.BusinessPartner.Created.v1",
  "accepter": ["All"],
  "business_partner_code":  "9980002060",
  "deleted":  false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetBPSupplier(businessPartner, businessPartnerRole, addressID, purchasingOrganization, supplier, companyCode string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
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
```

## Output  
本マイクロサービスでは、[golang-logging-library](https://github.com/latonaio/golang-logging-library) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP ビジネスパートナ仕入先  の ロール が取得された結果の JSON の例です。  
以下の項目のうち、"BusinessPartner" ～ "ValidTo" は、/SAP_API_Output_Formatter/type.go 内 の Type Product {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-business-partner-reads-supplier/SAP_API_Caller/caller.go#L69",
	"function": "sap-api-integrations-business-partner-reads-supplier/SAP_API_Caller.(*SAPAPICaller).Role",
	"level": "INFO",
	"message": {
		"BusinessPartner": "9980002060",
		"BusinessPartnerRole": "BUP003",
		"ValidFrom": "/Date(1470355200000+0000)/",
		"ValidTo": "/Date(253402300799000+0000)/"
	},
	"time": "2021-12-18T22:09:07.026509+09:00"
}
```