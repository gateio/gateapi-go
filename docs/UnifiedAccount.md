# UnifiedAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int64** | User ID | [optional] 
**RefreshTime** | **int64** | Time of the most recent refresh | [optional] 
**Locked** | **bool** | Whether account is locked | [optional] 
**Balances** | [**map[string]UnifiedBalance**](UnifiedBalance.md) |  | [optional] 
**Total** | **string** | The total asset value in USDT. Sum of &#x60;(available + freeze) * price&#x60; | [optional] 
**Borrowed** | **string** | The total borrowed amount in USDT equivalent. Sum of &#x60;borrowed * price&#x60;  | [optional] 
**TotalInitialMargin** | **string** | Total initial margin | [optional] 
**TotalMarginBalance** | **string** | Total margin balance | [optional] 
**TotalMaintenanceMargin** | **string** | Total maintenance margin | [optional] 
**TotalInitialMarginRate** | **string** | Total initial margin rate | [optional] 
**TotalMaintenanceMarginRate** | **string** | Total maintenance margin rate | [optional] 
**TotalAvailableMargin** | **string** | Total available margin | [optional] 
**UnifiedAccountTotal** | **string** | Total amount of the portfolio margin account | [optional] 
**UnifiedAccountTotalLiab** | **string** | Total liabilities of the portfolio margin account | [optional] 
**UnifiedAccountTotalEquity** | **string** | Total equity of the portfolio margin account | [optional] 
**Leverage** | **string** | Leverage | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


