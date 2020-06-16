# Contract

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Futures contract name | [optional] 
**Type** | **string** | Futures contract type | [optional] 
**QuantoMultiplier** | **string** | Multiplier used in converting from invoicing to settlement currency in quanto futures | [optional] 
**LeverageMin** | **string** | Minimum leverage | [optional] 
**LeverageMax** | **string** | Maximum leverage | [optional] 
**MaintenanceRate** | **string** | Maintenance rate of margin | [optional] 
**MarkType** | **string** | Mark price type, internal - based on internal trading, index - based on external index price | [optional] 
**MarkPrice** | **string** | Current mark price | [optional] 
**IndexPrice** | **string** | Current index price | [optional] 
**LastPrice** | **string** | Last trading price | [optional] 
**MakerFeeRate** | **string** | Maker fee rate, where negative means rebate | [optional] 
**TakerFeeRate** | **string** | Taker fee rate | [optional] 
**OrderPriceRound** | **string** | Minimum order price increment | [optional] 
**MarkPriceRound** | **string** | Minimum mark price increment | [optional] 
**FundingRate** | **string** | Current funding rate | [optional] 
**FundingInterval** | **int32** | Funding application interval, unit in seconds | [optional] 
**FundingNextApply** | **float32** | Next funding time | [optional] 
**RiskLimitBase** | **string** | Risk limit base | [optional] 
**RiskLimitStep** | **string** | Step of adjusting risk limit | [optional] 
**RiskLimitMax** | **string** | Maximum risk limit the contract allowed | [optional] 
**OrderSizeMin** | **int64** | Minimum order size the contract allowed | [optional] 
**OrderSizeMax** | **int64** | Maximum order size the contract allowed | [optional] 
**OrderPriceDeviate** | **string** | deviation between order price and current index price. If price of an order is denoted as order_price, it must meet the following condition:      abs(order_price - mark_price) &lt;&#x3D; mark_price * order_price_deviate | [optional] 
**OrderbookId** | **int64** | Current orderbook ID | [optional] 
**TradeId** | **int64** | Current trade ID | [optional] 
**TradeSize** | **int64** | Historical accumulation trade size | [optional] 
**PositionSize** | **int64** | Current total long position size | [optional] 
**ConfigChangeTime** | **float32** | Configuration&#39;s last changed time | [optional] 
**InDelisting** | **bool** | Contract is delisting | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


