# Position

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | **int64** | User ID | [optional] [readonly] 
**Contract** | **string** | Futures contract | [optional] [readonly] 
**Size** | **int64** | Position size | [optional] [readonly] 
**Leverage** | **string** | Position leverage. 0 means cross margin; positive number means isolated margin | [optional] 
**RiskLimit** | **string** | Position risk limit | [optional] 
**LeverageMax** | **string** | Maximum leverage under current risk limit | [optional] [readonly] 
**MaintenanceRate** | **string** | Maintenance rate under current risk limit | [optional] [readonly] 
**Value** | **string** | Position value calculated in settlement currency | [optional] [readonly] 
**Margin** | **string** | Position margin | [optional] 
**EntryPrice** | **string** | Entry price | [optional] [readonly] 
**LiqPrice** | **string** | Liquidation price | [optional] [readonly] 
**MarkPrice** | **string** | Current mark price | [optional] [readonly] 
**UnrealisedPnl** | **string** | Unrealized PNL | [optional] [readonly] 
**RealisedPnl** | **string** | Realized PNL | [optional] [readonly] 
**HistoryPnl** | **string** | History realized PNL | [optional] [readonly] 
**LastClosePnl** | **string** | PNL of last position close | [optional] [readonly] 
**RealisedPoint** | **string** | Realized POINT PNL | [optional] [readonly] 
**HistoryPoint** | **string** | History realized POINT PNL | [optional] [readonly] 
**AdlRanking** | **int32** | ADL ranking, ranging from 1 to 5 | [optional] [readonly] 
**PendingOrders** | **int32** | Current open orders | [optional] [readonly] 
**CloseOrder** | Pointer to [**PositionCloseOrder**](Position_close_order.md) |  | [optional] 
**Mode** | **string** | Position mode, including:  - &#x60;single&#x60;: dual mode is not enabled- &#x60;dual_long&#x60;: long position in dual mode- &#x60;dual_short&#x60;: short position in dual mode | [optional] 
**CrossLeverageLimit** | **string** | Cross margin leverage(valid only when &#x60;leverage&#x60; is 0) | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


