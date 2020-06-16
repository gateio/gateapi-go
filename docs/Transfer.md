# Transfer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | Transfer currency. For futures account, &#x60;currency&#x60; can be set to &#x60;POINT&#x60; or settle currency | 
**From** | **string** | Account transferred from. &#x60;spot&#x60; - spot account. &#x60;margin&#x60; - margin account, &#x60;futures&#x60; - futures account | 
**To** | **string** | Account transferred to. &#x60;spot&#x60; - spot account. &#x60;margin&#x60; - margin account, &#x60;futures&#x60; - futures account | 
**Amount** | **string** | Transfer amount | 
**CurrencyPair** | **string** | Margin currency pair. Required if transfer from or to margin account | [optional] 
**Settle** | **string** | Futures settle currency. Required if &#x60;currency&#x60; is &#x60;POINT&#x60; | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


