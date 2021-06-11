# CrossMarginAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int64** | User ID | [optional] 
**Locked** | **bool** | Whether account is locked | [optional] 
**Balances** | [**map[string]CrossMarginBalance**](CrossMarginBalance.md) |  | [optional] 
**Total** | **string** | Total account value in USDT, i.e., the sum of all currencies&#39; &#x60;(available+freeze)*price*discount&#x60; | [optional] 
**Borrowed** | **string** | Total borrowed value in USDT, i.e., the sum of all currencies&#39; &#x60;borrowed*price*discount&#x60; | [optional] 
**Interest** | **string** | Total unpaid interests in USDT, i.e., the sum of all currencies&#39; &#x60;interest*price*discount&#x60; | [optional] 
**Risk** | **string** | Risk rate. When it belows 110%, liquidation will be triggered. Calculation formula: &#x60;total / (borrowed+interest)&#x60; | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


