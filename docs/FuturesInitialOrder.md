# FuturesInitialOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contract** | **string** | Futures contract | 
**Size** | **int64** | Order size. Positive size means to buy, while negative one means to sell. Set to 0 to close the position | [optional] 
**Price** | **string** | Order price. Set to 0 to use market price | 
**Close** | **bool** | Set to true if trying to close the position | [optional] [default to false]
**Tif** | **string** | Time in force. If using market price, only &#x60;ioc&#x60; is supported.  - gtc: GoodTillCancelled - ioc: ImmediateOrCancelled | [optional] [default to TIF_GTC]
**Text** | **string** | How the order is created. Possible values are: web, api and app | [optional] 
**ReduceOnly** | **bool** | Set to true to create a reduce-only order | [optional] [default to false]
**IsReduceOnly** | **bool** | Is the order reduce-only | [optional] [readonly] 
**IsClose** | **bool** | Is the order to close position | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


