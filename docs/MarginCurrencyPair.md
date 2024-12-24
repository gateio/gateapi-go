# MarginCurrencyPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Currency pair | [optional] 
**Base** | **string** | Base currency | [optional] 
**Quote** | **string** | Quote currency | [optional] 
**Leverage** | **int32** | Leverage | [optional] 
**MinBaseAmount** | **string** | Minimum base currency to loan, &#x60;null&#x60; means no limit | [optional] 
**MinQuoteAmount** | **string** | Minimum quote currency to loan, &#x60;null&#x60; means no limit | [optional] 
**MaxQuoteAmount** | **string** | Maximum borrowable amount for quote currency. Base currency limit is calculated by quote maximum and market price. &#x60;null&#x60; means no limit | [optional] 
**Status** | **int32** | Currency pair status   - &#x60;0&#x60;: disabled  - &#x60;1&#x60;: enabled | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


