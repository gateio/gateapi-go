# SubCrossMarginAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int64** | User ID of the cross margin account. 0 means that the subaccount has not yet opened a cross margin account | [optional] 
**Locked** | **bool** | Whether account is locked | [optional] 
**Balances** | [**map[string]CrossMarginBalance**](CrossMarginBalance.md) |  | [optional] 
**Total** | **string** | Total account value in USDT, i.e., the sum of all currencies&#39; &#x60;(available+freeze)*price*discount&#x60; | [optional] 
**Borrowed** | **string** | Total borrowed value in USDT, i.e., the sum of all currencies&#39; &#x60;borrowed*price*discount&#x60; | [optional] 
**BorrowedNet** | **string** | Total borrowed value in USDT * borrowed factor | [optional] 
**Net** | **string** | Total net assets in USDT | [optional] 
**Leverage** | **string** | Position leverage | [optional] 
**Interest** | **string** | Total unpaid interests in USDT, i.e., the sum of all currencies&#39; &#x60;interest*price*discount&#x60; | [optional] 
**Risk** | **string** | Risk rate. When it belows 110%, liquidation will be triggered. Calculation formula: &#x60;total / (borrowed+interest)&#x60; | [optional] 
**TotalInitialMargin** | **string** | Total initial margin | [optional] 
**TotalMarginBalance** | **string** | Total margin balance | [optional] 
**TotalMaintenanceMargin** | **string** | Total maintenance margin | [optional] 
**TotalInitialMarginRate** | **string** | Total initial margin rate | [optional] 
**TotalMaintenanceMarginRate** | **string** | Total maintenance margin rate | [optional] 
**TotalAvailableMargin** | **string** | Total available margin | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


