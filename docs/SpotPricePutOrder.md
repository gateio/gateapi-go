# SpotPricePutOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Order type, default to &#x60;limit&#x60; | [optional] [default to limit]
**Side** | **string** | Order side  - buy: buy side - sell: sell side | 
**Price** | **string** | Order price | 
**Amount** | **string** | Order amount | 
**Account** | **string** | Trading account type.  Portfolio margin account must set to &#x60;cross_margin&#x60;  - normal: spot trading - margin: margin trading - cross_margin: cross_margin trading  | [default to ACCOUNT_NORMAL]
**TimeInForce** | **string** | time_in_force  - gtc: GoodTillCancelled - ioc: ImmediateOrCancelled, taker only  | [optional] [default to TIME_IN_FORCE_GTC]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


