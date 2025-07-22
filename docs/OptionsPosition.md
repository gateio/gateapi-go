# OptionsPosition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | **int32** | User ID. | [optional] [readonly] 
**Underlying** | **string** | Underlying. | [optional] [readonly] 
**UnderlyingPrice** | **string** | Underlying price (quote currency). | [optional] [readonly] 
**Contract** | **string** | Options contract name. | [optional] [readonly] 
**Size** | **int64** | Position size (contract size). | [optional] [readonly] 
**EntryPrice** | **string** | Entry size (quote currency). | [optional] [readonly] 
**MarkPrice** | **string** | Current mark price (quote currency). | [optional] [readonly] 
**MarkIv** | **string** | Implied volatility. | [optional] [readonly] 
**RealisedPnl** | **string** | Realized PNL. | [optional] [readonly] 
**UnrealisedPnl** | **string** | Unrealized PNL. | [optional] [readonly] 
**PendingOrders** | **int32** | Current open orders. | [optional] [readonly] 
**CloseOrder** | Pointer to [**OptionsPositionCloseOrder**](OptionsPosition_close_order.md) |  | [optional] 
**Delta** | **string** | Delta. | [optional] [readonly] 
**Gamma** | **string** | Gamma. | [optional] [readonly] 
**Vega** | **string** | Vega. | [optional] [readonly] 
**Theta** | **string** | Theta. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


