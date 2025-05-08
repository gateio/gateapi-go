# UnifiedAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int64** | User ID | [optional] 
**RefreshTime** | **int64** | Time of the most recent refresh | [optional] 
**Locked** | **bool** | Whether the account is locked, valid in cross-currency margin/combined margin mode, false in other modes such as single-currency margin mode | [optional] 
**Balances** | [**map[string]UnifiedBalance**](UnifiedBalance.md) |  | [optional] 
**Total** | **string** | Total account assets converted to USD, i.e. the sum of &#x60;(available + freeze) * price&#x60;  in all currencies (deprecated, to be deprecated, replaced by unified_account_total) | [optional] 
**Borrowed** | **string** | The total borrowed amount of the account converted into USD, i.e. the sum of &#x60;borrowed * price&#x60; of all currencies (excluding Point Cards). It is valid in cross-currency margin/combined margin mode, and is 0 in other modes such as single-currency margin mode. | [optional] 
**TotalInitialMargin** | **string** | Total initial margin, valid in cross-currency margin/combined margin mode, 0 in other modes such as single-currency margin mode | [optional] 
**TotalMarginBalance** | **string** | Total margin balance, valid in cross-currency margin/combined margin mode, 0 in other modes such as single-currency margin mode | [optional] 
**TotalMaintenanceMargin** | **string** | Total maintenance margin is valid in cross-currency margin/combined margin mode, and is 0 in other modes such as single-currency margin mode | [optional] 
**TotalInitialMarginRate** | **string** | Total initial margin rate, valid in cross-currency margin/combined margin mode, 0 in other modes such as single-currency margin mode | [optional] 
**TotalMaintenanceMarginRate** | **string** | Total maintenance margin rate, valid in cross-currency margin/combined margin mode, 0 in other modes such as single-currency margin mode | [optional] 
**TotalAvailableMargin** | **string** | Available margin amount, valid in cross-currency margin/combined margin mode, 0 in other modes such as single-currency margin mode | [optional] 
**UnifiedAccountTotal** | **string** | Unify the total account assets, valid in single currency margin/cross-currency margin/combined margin mode | [optional] 
**UnifiedAccountTotalLiab** | **string** | Unify the total loan of the account, valid in the cross-currency margin/combined margin mode, and 0 in other modes such as single-currency margin mode | [optional] 
**UnifiedAccountTotalEquity** | **string** | Unify the total account equity, valid in single currency margin/cross-currency margin/combined margin mode | [optional] 
**Leverage** | **string** | Actual leverage, valid in cross-currency margin/combined margin mode | [optional] [readonly] 
**SpotOrderLoss** | **string** | Total pending order loss, in USDT, valid in cross-currency margin/combined margin mode, 0 in other modes such as single-currency margin mode | [optional] 
**SpotHedge** | **bool** | Spot hedging status, true - enabled, false - not enabled. | [optional] 
**UseFunding** | **bool** | Whether to use funds as margin | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


