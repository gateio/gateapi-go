# CancelOrderResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyPair** | **string** | Order currency pair | [optional] 
**Id** | **string** | Order ID | [optional] 
**Succeeded** | **bool** | Whether cancellation succeeded | [optional] 
**Label** | **string** | Error label when failed to cancel the order; emtpy if succeeded | [optional] 
**Message** | **string** | Error message when failed to cancel the order; empty if succeeded | [optional] 
**Account** | **string** | Empty by default. If cancelled order is cross margin order, this field is set to &#x60;cross_margin&#x60; | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


