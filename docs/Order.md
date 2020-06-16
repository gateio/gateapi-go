# Order

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Order ID | [optional] 
**Text** | **string** | User defined information. If not empty, must follow the rules below:  1. prefixed with &#x60;t-&#x60; 2. no longer than 16 bytes without &#x60;t-&#x60; prefix 3. can only include 0-9, A-Z, a-z, underscore(_), hyphen(-) or dot(.)  | [optional] 
**CreateTime** | **string** | Order creation time | [optional] 
**UpdateTime** | **string** | Order last modification time | [optional] 
**Status** | **string** | Order status  - &#x60;open&#x60;: to be filled - &#x60;closed&#x60;: filled - &#x60;cancelled&#x60;: cancelled | [optional] 
**CurrencyPair** | **string** | Currency pair | 
**Type** | **string** | Order type. limit - limit order | [optional] [default to TYPE_LIMIT]
**Account** | **string** | Account type. spot - use spot account; margin - use margin account | [optional] [default to ACCOUNT_SPOT]
**Side** | **string** | Order side | 
**Amount** | **string** | Trade amount | 
**Price** | **string** | Order price | 
**TimeInForce** | **string** | Time in force  - gtc: GoodTillCancelled - ioc: ImmediateOrCancelled, taker only - poc: PendingOrCancelled, makes a post-only order that always enjoys a maker fee | [optional] [default to TIME_IN_FORCE_GTC]
**AutoBorrow** | **bool** | Used in margin trading(i.e. &#x60;account&#x60; is &#x60;margin&#x60;) to allow automatic loan of insufficient part if balance is not enough. | [optional] 
**Left** | **string** | Amount left to fill | [optional] 
**FillPrice** | **string** | Total filled in quote currency. Deprecated in favor of &#x60;filled_total&#x60; | [optional] 
**FilledTotal** | **string** | Total filled in quote currency | [optional] 
**Fee** | **string** | Fee deducted | [optional] 
**FeeCurrency** | **string** | Fee currency unit | [optional] 
**PointFee** | **string** | Point used to deduct fee | [optional] 
**GtFee** | **string** | GT used to deduct fee | [optional] 
**GtDiscount** | **bool** | Whether GT fee discount is used | [optional] 
**RebatedFee** | **string** | Rebated fee | [optional] 
**RebatedFeeCurrency** | **string** | Rebated fee currency unit | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


