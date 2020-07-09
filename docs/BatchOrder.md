# BatchOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | User defined information. If not empty, must follow the rules below:  1. prefixed with &#x60;t-&#x60; 2. no longer than 28 bytes without &#x60;t-&#x60; prefix 3. can only include 0-9, A-Z, a-z, underscore(_), hyphen(-) or dot(.)  | [optional] 
**Succeeded** | **bool** | Whether order succeeds | [optional] 
**Label** | **string** | Error label, empty string if order succeeds | [optional] 
**Message** | **string** | Detailed error message, empty string if order succeeds | [optional] 
**Id** | **string** | Order ID | [optional] [readonly] 
**CreateTime** | **string** | Order creation time | [optional] [readonly] 
**UpdateTime** | **string** | Order last modification time | [optional] [readonly] 
**Status** | **string** | Order status  - &#x60;open&#x60;: to be filled - &#x60;closed&#x60;: filled - &#x60;cancelled&#x60;: cancelled | [optional] [readonly] 
**CurrencyPair** | **string** | Currency pair | [optional] 
**Type** | **string** | Order type. limit - limit order | [optional] [default to TYPE_LIMIT]
**Account** | **string** | Account type. spot - use spot account; margin - use margin account | [optional] [default to ACCOUNT_SPOT]
**Side** | **string** | Order side | [optional] 
**Amount** | **string** | Trade amount | [optional] 
**Price** | **string** | Order price | [optional] 
**TimeInForce** | **string** | Time in force  - gtc: GoodTillCancelled - ioc: ImmediateOrCancelled, taker only - poc: PendingOrCancelled, makes a post-only order that always enjoys a maker fee | [optional] [default to TIME_IN_FORCE_GTC]
**AutoBorrow** | **bool** | Used in margin trading(i.e. &#x60;account&#x60; is &#x60;margin&#x60;) to allow automatic loan of insufficient part if balance is not enough. | [optional] 
**Left** | **string** | Amount left to fill | [optional] [readonly] 
**FillPrice** | **string** | Total filled in quote currency. Deprecated in favor of &#x60;filled_total&#x60; | [optional] [readonly] 
**FilledTotal** | **string** | Total filled in quote currency | [optional] [readonly] 
**Fee** | **string** | Fee deducted | [optional] [readonly] 
**FeeCurrency** | **string** | Fee currency unit | [optional] [readonly] 
**PointFee** | **string** | Point used to deduct fee | [optional] [readonly] 
**GtFee** | **string** | GT used to deduct fee | [optional] [readonly] 
**GtDiscount** | **bool** | Whether GT fee discount is used | [optional] [readonly] 
**RebatedFee** | **string** | Rebated fee | [optional] [readonly] 
**RebatedFeeCurrency** | **string** | Rebated fee currency unit | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


