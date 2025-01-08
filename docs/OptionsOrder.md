# OptionsOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Options order ID | [optional] [readonly] 
**User** | **int32** | User ID | [optional] [readonly] 
**CreateTime** | **float64** | Creation time of order | [optional] [readonly] 
**FinishTime** | **float64** | Order finished time. Not returned if order is open | [optional] [readonly] 
**FinishAs** | **string** | Ending method, including:  - filled: fully completed - canceled: user canceled - liquidated: forced liquidation cancellation - ioc: Not fully filled immediately because tif is set to ioc - auto_deleveraged: automatic deleveraging cancel - reduce_only: Increased position is cancelled, because reduce_only is set or the position is closed - position_closed: Because the position was closed, the pending order was canceled - reduce_out: Only reduce the excluded pending orders that are not easy to be filled - mmp_cancelled: MMP canceled | [optional] [readonly] 
**Status** | **string** | Order status  - &#x60;open&#x60;: waiting to be traded - &#x60;finished&#x60;: finished | [optional] [readonly] 
**Contract** | **string** | Contract name | 
**Size** | **int64** | Order size. Specify positive number to make a bid, and negative number to ask | 
**Iceberg** | **int64** | Display size for iceberg order. 0 for non-iceberg. Note that you will have to pay the taker fee for the hidden size | [optional] 
**Price** | **string** | Order price. 0 for market order with &#x60;tif&#x60; set as &#x60;ioc&#x60; (USDT) | [optional] 
**Close** | **bool** | Set as &#x60;true&#x60; to close the position, with &#x60;size&#x60; set to 0 | [optional] [default to false]
**IsClose** | **bool** | Is the order to close position | [optional] [readonly] 
**ReduceOnly** | **bool** | Set as &#x60;true&#x60; to be reduce-only order | [optional] [default to false]
**IsReduceOnly** | **bool** | Is the order reduce-only | [optional] [readonly] 
**IsLiq** | **bool** | Is the order for liquidation | [optional] [readonly] 
**Mmp** | **bool** | When set to true, delegate to MMP | [optional] [default to false]
**IsMmp** | **bool** | Whether it is MMP delegation. Corresponds to &#x60;mmp&#x60; in the request. | [optional] [readonly] 
**Tif** | **string** | Time in force  - gtc: GoodTillCancelled - ioc: ImmediateOrCancelled, taker only - poc: PendingOrCancelled, makes a post-only order that always enjoys a maker fee | [optional] [default to TIF_GTC]
**Left** | **int64** | Size left to be traded | [optional] [readonly] 
**FillPrice** | **string** | Fill price of the order | [optional] [readonly] 
**Text** | **string** | User defined information. If not empty, must follow the rules below:  1. prefixed with &#x60;t-&#x60; 2. no longer than 28 bytes without &#x60;t-&#x60; prefix 3. can only include 0-9, A-Z, a-z, underscore(_), hyphen(-) or dot(.) Besides user defined information, reserved contents are listed below, denoting how the order is created:  - web: from web - api: from API - app: from mobile phones - auto_deleveraging: from ADL - liquidation: from liquidation - insurance: from insurance  | [optional] 
**Tkfr** | **string** | Taker fee | [optional] [readonly] 
**Mkfr** | **string** | Maker fee | [optional] [readonly] 
**Refu** | **int32** | Reference user ID | [optional] [readonly] 
**Refr** | **string** | Referrer rebate | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


