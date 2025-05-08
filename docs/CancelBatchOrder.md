# CancelBatchOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyPair** | **string** | Order currency pair | 
**Id** | **string** | Order ID or user custom ID. Custom ID are accepted only within 30 minutes after order creation | 
**Account** | **string** | If the canceled order is a unified account apikey, this field must be specified and set to &#x60;unified&#x60; | [optional] 
**ActionMode** | **string** | Processing Mode: When placing an order, different fields are returned based on action_mode. This field is only valid during the request and is not included in the response result ACK: Asynchronous mode, only returns key order fields RESULT: No clearing information FULL: Full mode (default) | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


