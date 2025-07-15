# OrderPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyPair** | **string** | Currency pair | [optional] 
**Account** | **string** | Specify query account. | [optional] 
**Amount** | **string** | Trading Quantity. Either amountor pricemust be specified | [optional] 
**Price** | **string** | Trading Price. Either amountor pricemust be specified | [optional] 
**AmendText** | **string** | Custom info during amending order | [optional] 
**ActionMode** | **string** | Processing Mode: When placing an order, different fields are returned based on action_mode. This field is only valid during the request and is not included in the response result ACK: Asynchronous mode, only returns key order fields RESULT: No clearing information FULL: Full mode (default) | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


