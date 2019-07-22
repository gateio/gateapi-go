# FuturesOrder

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Futures order ID | [optional] 
**User** | **int32** | User ID | [optional] 
**CreateTime** | **float32** | Order creation time | [optional] 
**FinishTime** | **float32** | Order finished time. Not returned if order is open | [optional] 
**FinishAs** | **string** | How the order is finished.  - filled: all filled - cancelled: manually cancelled - liquidated: cancelled because of liquidation - ioc: time in force is &#x60;IOC&#x60;, finish immediately - auto_deleveraged: finished by ADL - reduce_only: cancelled because of increasing position while &#x60;reduce-only&#x60; set | [optional] 
**Status** | **string** | Order status  - &#x60;open&#x60;: waiting to be traded - &#x60;finished&#x60;: finished | [optional] 
**Contract** | **string** | Futures contract | 
**Size** | **int64** | Order size. Specify positive number to make a bid, and negative number to ask | [optional] 
**Iceberg** | **int64** | Display size for iceberg order. 0 for non-iceberg. Note that you would pay the taker fee for the hidden size | [optional] 
**Price** | **string** | Order price. 0 for market order with &#x60;tif&#x60; set as &#x60;ioc&#x60; | [optional] 
**Close** | **bool** | Set as &#x60;true&#x60; to close the position, with &#x60;size&#x60; set to 0 | [optional] [default to false]
**IsClose** | **bool** | Is the order to close position | [optional] 
**ReduceOnly** | **bool** | Set as &#x60;true&#x60; to be post-only order | [optional] [default to false]
**IsReduceOnly** | **bool** | Is the order reduce-only | [optional] 
**IsLiq** | **bool** | Is the order for liquidation | [optional] 
**Tif** | **string** | Time in force  - gtc: GoodTillCancelled - ioc: ImmediateOrCancelled, taker only - poc: PendingOrCancelled, post-only | [optional] [default to TIF_GTC]
**Left** | **int64** | Size left to be traded | [optional] 
**FillPrice** | **string** | Fill price of the order | [optional] 
**Text** | **string** | User defined information. If not empty, must follow the rules below:  1. prefixed with &#x60;t-&#x60; 2. no longer than 16 bytes without &#x60;t-&#x60; prefix 3. can only include 0-9, A-Z, a-z, underscore(_), hyphen(-) or dot(.) Besides user defined information, reserved contents are listed below, denoting how the order is created:  - web: from web - api: from API - app: from mobile phones - auto_deleveraging: from ADL - liquidation: from liquidation - insurance: from insurance  | [optional] 
**Tkfr** | **string** | Taker fee | [optional] 
**Mkfr** | **string** | Maker fee | [optional] 
**Refu** | **int32** | Reference user ID | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


