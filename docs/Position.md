# Position

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | **int64** | User ID | [optional] 
**Contract** | **string** | Futures contract | [optional] 
**Size** | **int64** | Position size | [optional] 
**Leverage** | **string** | Position leverage. 0 means cross margin; positive number means isolated margin | [optional] 
**RiskLimit** | **string** | Position risk limit | [optional] 
**LeverageMax** | **string** | Maximum leverage under current risk limit | [optional] 
**MaintenanceRate** | **string** | Maintenance rate under current risk limit | [optional] 
**Value** | **string** | Position value calculated in settlement currency | [optional] 
**Margin** | **string** | Position margin | [optional] 
**EntryPrice** | **string** | Entry price | [optional] 
**LiqPrice** | **string** | Liquidation price | [optional] 
**MarkPrice** | **string** | Current mark price | [optional] 
**UnrealisedPnl** | **string** | Unrealized PNL | [optional] 
**RealisedPnl** | **string** | Realized PNL | [optional] 
**HistoryPnl** | **string** | History realized PNL | [optional] 
**LastClosePnl** | **string** | PNL of last position close | [optional] 
**RealisedPoint** | **string** | Realized POINT PNL | [optional] 
**HistoryPoint** | **string** | History realized POINT PNL | [optional] 
**AdlRanking** | **int32** | ADL ranking, range from 1 to 5 | [optional] 
**PendingOrders** | **int32** | Current open orders | [optional] 
**CloseOrder** | [**PositionCloseOrder**](Position_close_order.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


