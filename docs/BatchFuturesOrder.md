# BatchFuturesOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Succeeded** | **bool** | Whether the batch of orders succeeded. | [optional] 
**Label** | **string** | Error label, only exists if execution fails. | [optional] 
**Detail** | **string** | Error detail, only present if execution failed and details need to be given | [optional] 
**Id** | **int64** | Futures order ID. | [optional] [readonly] 
**User** | **int32** | User ID. | [optional] [readonly] 
**CreateTime** | **float64** | Creation time of order. | [optional] [readonly] 
**FinishTime** | **float64** | Order finished time. Not returned if order is open. | [optional] [readonly] 
**FinishAs** | **string** | How the order was finished.  - filled: all filled - cancelled: manually cancelled - liquidated: cancelled because of liquidation - ioc: time in force is &#x60;IOC&#x60;, finish immediately - auto_deleveraged: finished by ADL - increasing position while &#x60;reduce-only&#x60; set- position_closed: cancelled because of position close - position_closed: canceled because the position was closed - reduce_out: only reduce positions by excluding hard-to-fill orders - stp: cancelled because self trade prevention  | [optional] [readonly] 
**Status** | **string** | Order status  - &#x60;open&#x60;: waiting to be traded - &#x60;finished&#x60;: finished | [optional] [readonly] 
**Contract** | **string** | Futures contract. | [optional] 
**Size** | **int64** | Order size. Specify positive number to make a bid, and negative number to ask | [optional] 
**Iceberg** | **int64** | Display size for iceberg order. 0 for non-iceberg. Note that you will have to pay the taker fee for the hidden size | [optional] 
**Price** | **string** | Order price. 0 for market order with &#x60;tif&#x60; set as &#x60;ioc&#x60;. | [optional] 
**Close** | **bool** | Set as &#x60;true&#x60; to close the position, with &#x60;size&#x60; set to 0. | [optional] [default to false]
**IsClose** | **bool** | Is the order to close position. | [optional] [readonly] 
**ReduceOnly** | **bool** | Set as &#x60;true&#x60; to be reduce-only order. | [optional] [default to false]
**IsReduceOnly** | **bool** | Is the order reduce-only. | [optional] [readonly] 
**IsLiq** | **bool** | Is the order for liquidation. | [optional] [readonly] 
**Tif** | **string** | Time in force  - gtc: GoodTillCancelled - ioc: ImmediateOrCancelled, taker only - poc: PendingOrCancelled, makes a post-only order that always enjoys a maker fee - fok: FillOrKill, fill either completely or none | [optional] [default to TIF_GTC]
**Left** | **int64** | Size left to be traded. | [optional] [readonly] 
**FillPrice** | **string** | Fill price of the order. | [optional] [readonly] 
**Text** | **string** | User defined information. If not empty, must follow the rules below:  1. prefixed with &#x60;t-&#x60; 2. no longer than 28 bytes without &#x60;t-&#x60; prefix 3. can only include 0-9, A-Z, a-z, underscore(_), hyphen(-) or dot(.) Besides user defined information, reserved contents are listed below, denoting how the order is created:  - web: from web - api: from API - app: from mobile phones - auto_deleveraging: from ADL - liquidation: from liquidation - insurance: from insurance  | [optional] 
**Tkfr** | **string** | Taker fee. | [optional] [readonly] 
**Mkfr** | **string** | Maker fee. | [optional] [readonly] 
**Refu** | **int32** | Reference user ID. | [optional] [readonly] 
**AutoSize** | **string** | Set side to close dual-mode position. &#x60;close_long&#x60; closes the long side; while &#x60;close_short&#x60; the short one. Note &#x60;size&#x60; also needs to be set to 0 | [optional] 
**StpAct** | **string** | Self-Trading Prevention Action. Users can use this field to set self-trade prevetion strategies  1. After users join the &#x60;STP Group&#x60;, he can pass &#x60;stp_act&#x60; to limit the user&#39;s self-trade prevetion strategy. If &#x60;stp_act&#x60; is not passed, the default is &#x60;cn&#x60; strategy。 2. When the user does not join the &#x60;STP group&#x60;, an error will be returned when passing the &#x60;stp_act&#x60; parameter。 3. If the user did not use &#39;stp_act&#39; when placing the order, &#39;stp_act&#39; will return &#39;-&#39;  - cn: Cancel newest, Cancel new orders and keep old ones - co: Cancel oldest, new ones - cb: Cancel both, Both old and new orders will be cancelled | [optional] 
**StpId** | **int32** | Orders between users in the same &#x60;stp_id&#x60; group are not allowed to be self-traded  1. If the &#x60;stp_id&#x60; of two orders being matched is non-zero and equal, they will not be executed. Instead, the corresponding strategy will be executed based on the &#x60;stp_act&#x60; of the taker. 2. &#x60;stp_id&#x60; returns &#x60;0&#x60; by default for orders that have not been set for &#x60;STP group&#x60; | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


