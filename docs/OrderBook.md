# OrderBook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Order book ID, which is updated whenever the order book is changed. Valid only when &#x60;with_id&#x60; is set to &#x60;true&#x60; | [optional] 
**Current** | **int64** | Response data generation timestamp in milliseconds | [optional] 
**Update** | **int64** | Order book changed timestamp in milliseconds | [optional] 
**Asks** | [**[][]string**](array.md) | Asks order depth | 
**Bids** | [**[][]string**](array.md) | Bids order depth | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


