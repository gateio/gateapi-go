# AmendOrderResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **string** | Order ID | [optional] [readonly] 
**Amount** | **string** | Trade amount | [optional] [readonly] 
**Price** | **string** | Order price | [optional] [readonly] 
**AmendText** | **string** | The custom data that the user remarked when amending the order | [optional] [readonly] 
**Succeeded** | **bool** | Update success status | [optional] [readonly] 
**Label** | **string** | Error indicator for failed modifications; empty when successful | [optional] [readonly] 
**Message** | **string** | Error description for failed modifications; empty when successful | [optional] [readonly] 
**Account** | **string** | Account types， spot - spot account, margin - margin account, portfolio - portfolio margin account, cross_margin - cross margin account.Portfolio margin accounts can only be set to &#x60;cross_margin&#x60; | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


