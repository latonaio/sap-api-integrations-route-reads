# sap-api-integrations-route-reads  
sap-api-integrations-route-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API ルートデータを取得するマイクロサービスです。  
sap-api-integrations-route-reads には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-route-reads は、オンプレミス版である（＝クラウド版ではない）SAPC4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/route/overview  

## 動作環境
sap-api-integrations-route-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。   
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。   
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須） 

## クラウド環境での利用  
sap-api-integrations-route-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-route-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/route/overview
* APIサービス名(=baseURL): c4codataapi

## 本レポジトリ に 含まれる API名
sap-api-integrations-route-reads には、次の API をコールするためのリソースが含まれています。  

* RouteCollection（SAP ルート - ルート）
* RouteAccountCollection（SAP ルート - ルートアカウント）
* RouteInvolvedPartiesCollection（SAP ルート - 関係者）


## API への 値入力条件 の 初期値
sap-api-integrations-route-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.RouteCollection.ID（ID）
* inoutSDC.RouteAccountCollection.RouteID（ルートID）
* inoutSDC.RouteInvolvedPartiesCollection.PartyID（関係者ID）


## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"RouteCollection" が指定されています。    
  
```
	"api_schema": "RouteCollection",
	"accepter": ["RouteCollection"],
	"route_code": "11",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "RouteCollection",
	"accepter": ["All"],
	"route_code": "11",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetRoute(iD, RouteID string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "RouteCollection":
			func() {
				c.RouteCollection(iD)
				wg.Done()
			}()
		case "RouteAccountCollection":
			func() {
				c.RouteAccountCollection(RouteID)
				wg.Done()
			}()
		case "RouteInvolvedPartiesCollection":
			func() {
				c.RouteInvolvedPartiesCollection(PartyID)
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
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP の ルートデータ が取得された結果の JSON の例です。  
以下の項目のうち、"ObjectID" ～ "WorkingDayCalendarCodeText" は、/SAP_API_Output_Formatter/type.go 内 の Type RouteCollection {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-route-reads/SAP_API_Caller/caller.go#L53",
	"function": "sap-api-integrations-route-reads/SAP_API_Caller.(*SAPAPICaller).RouteCollection",
	"level": "INFO",
	"message": [
		{
			"ObjectID": "00163E05A2E01EE4B2FD0B2BAC14CB9D",
			"ETag": "2015-07-14T05:42:19+09:00",
			"ID": "11",
			"Name": "REWE - Mannheim Tour",
			"RouteTypeCode": "1",
			"RouteTypeCodeText": "Template",
			"StartDate": "",
			"ExcludeMondayIndicator": false,
			"ExcludeTuesdayIndicator": false,
			"ExcludeWednesdayIndicator": false,
			"ExcludeThursdayIndicator": false,
			"ExcludeFridayIndicator": false,
			"ExcludeSaturdayIndicator": false,
			"ExcludeSundayIndicator": false,
			"SalesTerritoryID": "",
			"SalesOrganisationID": "",
			"OrganizerPartyID": "",
			"DefaultStartTime": "PT08H00M00S",
			"DefaultPreparationTime": "PT1H",
			"DefaultDuration": "PT1H",
			"VisitTypeCode": "",
			"VisitTypeCodeText": "",
			"Status": "1",
			"StatusText": "In Preparation",
			"ProcessingStatus": "1",
			"ProcessingStatusText": "Open",
			"OwnerPartyID": "",
			"OwnerPartyUUID": "",
			"PerfectStoreIndicator": false,
			"DistributionChannelCode": "",
			"DistributionChannelCodeText": "",
			"DivisionCode": "",
			"DivisionCodeText": "",
			"AutomaticResequencing": true,
			"EndDate": "",
			"WorkingDayCalendarCode": "",
			"WorkingDayCalendarCodeText": ""
		}
	],
	"time": "2022-07-25T18:07:28+09:00"
}


```