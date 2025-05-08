# FuturesAutoDeleverage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **int64** | Automatic deleveraging time | [optional] [readonly] 
**User** | **int64** | User ID | [optional] [readonly] 
**OrderId** | **int64** | Order ID. Order IDs before 2023-02-20 are null | [optional] [readonly] 
**Contract** | **string** | Futures contract | [optional] [readonly] 
**Leverage** | **string** | Position leverage | [optional] [readonly] 
**CrossLeverageLimit** | **string** | Cross margin leverage(valid only when &#x60;leverage&#x60; is 0) | [optional] [readonly] 
**EntryPrice** | **string** | Average entry price | [optional] [readonly] 
**FillPrice** | **string** | Average fill price | [optional] [readonly] 
**TradeSize** | **int64** | Trading size | [optional] [readonly] 
**PositionSize** | **int64** | Positions after auto-deleveraging | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


