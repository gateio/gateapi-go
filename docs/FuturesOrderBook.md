# FuturesOrderBook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Order Book ID. Increases by 1 on every order book change. Set &#x60;with_id&#x3D;true&#x60; to include this field in response | [optional] 
**Current** | **float64** | Response data generation timestamp | [optional] 
**Update** | **float64** | Order book changed timestamp | [optional] 
**Asks** | [**[]FuturesOrderBookItem**](futures_order_book_item.md) | Asks order depth | 
**Bids** | [**[]FuturesOrderBookItem**](futures_order_book_item.md) | Bids order depth | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


