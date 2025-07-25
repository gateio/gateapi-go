# SpotPricePutOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Order type，default to &#x60;limit&#x60;  - limit : Limit Order - market : Market Order | [optional] [default to TYPE_LIMIT]
**Side** | **string** | Order side  - buy: buy side - sell: sell side | 
**Price** | **string** | Order price. | 
**Amount** | **string** | When &#x60;type&#x60; is limit, it refers to base currency. For instance, &#x60;BTC_USDT&#x60; means &#x60;BTC&#x60;  When different currency according to &#x60;side&#x60;  - &#x60;side&#x60; : &#x60;buy&#x60; means quote currency, &#x60;BTC_USDT&#x60; means &#x60;USDT&#x60; - &#x60;side&#x60; : &#x60;sell&#x60; means base currency，&#x60;BTC_USDT&#x60; means &#x60;BTC&#x60;  | 
**Account** | **string** | Trading account type. Portfolio margin account must set to &#x60;unified&#x60;  -normal: spot trading - margin: margin trading - unified: unified trading  | [default to ACCOUNT_NORMAL]
**TimeInForce** | **string** | time_in_force  - gtc: GoodTillCancelled - ioc: ImmediateOrCancelled, taker only  | [optional] [default to TIME_IN_FORCE_GTC]
**AutoBorrow** | **bool** | Whether to borrow coins automatically. | [optional] [default to false]
**AutoRepay** | **bool** | Whether to repay the loan automatically. | [optional] [default to false]
**Text** | **string** | The source of the order, including: - web: web - api: api - app: app | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


