# PortfolioAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int64** | User ID | [optional] 
**RefreshTime** | **int64** | Time of the most recent refresh | [optional] 
**Locked** | **bool** | Whether account is locked | [optional] 
**Balances** | [**map[string]PortfolioMarginBalance**](PortfolioMarginBalance.md) |  | [optional] 
**Total** | **string** | Total account value in USDT, i.e., the sum of all currencies&#39; | [optional] 
**Borrowed** | **string** | Total borrowed value in USDT, i.e., the sum of all currencies | [optional] 
**Interest** | **string** | Total unpaid interests in USDT, i.e., the sum of all currencies | [optional] 
**TotalInitialMargin** | **string** | Total initial margin | [optional] 
**TotalMarginBalance** | **string** | Total margin balance | [optional] 
**TotalMaintenanceMargin** | **string** | Total maintenance margin | [optional] 
**TotalInitialMarginRate** | **string** | Total initial margin rate | [optional] 
**TotalMaintenanceMarginRate** | **string** | Total maintenance margin rate | [optional] 
**TotalAvailableMargin** | **string** | Total available margin | [optional] 
**PortfolioMarginTotal** | **string** | Total amount of the portfolio margin account | [optional] 
**PortfolioMarginTotalLiab** | **string** | Total liabilities of the portfolio margin account | [optional] 
**PortfolioMarginTotalEquity** | **string** | Total equity of the portfolio margin account | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


