# FuturesPriceTriggeredOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Initial** | [**FuturesInitialOrder**](FuturesInitialOrder.md) |  | 
**Trigger** | [**FuturesPriceTrigger**](FuturesPriceTrigger.md) |  | 
**Id** | **int64** | Auto order ID | [optional] [readonly] 
**User** | **int32** | User ID | [optional] [readonly] 
**CreateTime** | **float64** | Creation time | [optional] [readonly] 
**FinishTime** | **float64** | Finished time | [optional] [readonly] 
**TradeId** | **int64** | ID of the newly created order on condition triggered | [optional] [readonly] 
**Status** | **string** | Auto order status  - &#x60;open&#x60;: order is active - &#x60;finished&#x60;: order is finished - &#x60;inactive&#x60;: order is not active, only for close-long-order or close-short-order - &#x60;invalid&#x60;: order is invalid, only for close-long-order or close-short-order | [optional] [readonly] 
**FinishAs** | **string** | How order is finished | [optional] [readonly] 
**Reason** | **string** | Additional remarks on how the order was finished | [optional] [readonly] 
**OrderType** | **string** | Take-profit/stop-loss types, which include:  - &#x60;close-long-order&#x60;: order take-profit/stop-loss, close long position - &#x60;close-short-order&#x60;: order take-profit/stop-loss, close short position - &#x60;close-long-position&#x60;: position take-profit/stop-loss, close long position - &#x60;close-short-position&#x60;: position take-profit/stop-loss, close short position - &#x60;plan-close-long-position&#x60;: position planned take-profit/stop-loss, close long position - &#x60;plan-close-short-position&#x60;: position planned take-profit/stop-loss, close short position  The order take-profit/stop-loss can not be passed by request. These two types are read only. | [optional] 
**MeOrderId** | **string** | Corresponding order ID of order take-profit/stop-loss. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


