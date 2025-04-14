# FuturesInitialOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contract** | **string** | Futures contract | 
**Size** | **int64** | Represents the number of contracts that need to be closed, full closing: size&#x3D;0 Partial closing: plan-close-short-position size&gt;0  Partial closing: plan-close-long-position size&lt;0 | [optional] 
**Price** | **string** | Order price. Set to 0 to use market price | 
**Close** | **bool** | When all positions are closed in a single position mode, it must be set to true to perform the closing operation When partially closed positions in single-store mode/double-store mode, you can not set close, or close&#x3D;false | [optional] [default to false]
**Tif** | **string** | Time in force strategy, default is gtc, market order currently only supports ioc mode Market order currently only supports ioc mode  - gtc: GoodTillCancelled - ioc: ImmediateOrCancelled | [optional] [default to TIF_GTC]
**Text** | **string** | The source of the order, including: - web: web - api: api - app: app | [optional] 
**ReduceOnly** | **bool** | When set to true, perform automatic position reduction operation. Set to true to ensure that the order will not open a new position, and is only used to close or reduce positions | [optional] [default to false]
**AutoSize** | **string** | Do not set auto_size When the dual-position mode is closed all positions (size&#x3D;0), auto_size, close_long, close_short, short When the double-storey mode partially closes the position (size ≠ 0), there is no need to set auto_size | [optional] 
**IsReduceOnly** | **bool** | Is the order reduce-only | [optional] [readonly] 
**IsClose** | **bool** | Is the order to close position | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


