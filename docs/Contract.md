# Contract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Futures contract. | [optional] 
**Type** | **string** | Futures contract type. | [optional] 
**QuantoMultiplier** | **string** | Multiplier used in converting from invoicing to settlement currency. | [optional] 
**LeverageMin** | **string** | Minimum leverage. | [optional] 
**LeverageMax** | **string** | Maximum leverage. | [optional] 
**MaintenanceRate** | **string** | Maintenance rate of margin. | [optional] 
**MarkType** | **string** | Mark price type, internal - based on internal trading, external index price | [optional] 
**MarkPrice** | **string** | Current mark price. | [optional] 
**IndexPrice** | **string** | Current index price. | [optional] 
**LastPrice** | **string** | Last trading price. | [optional] 
**MakerFeeRate** | **string** | Maker fee rate, where negative means rebate. | [optional] 
**TakerFeeRate** | **string** | Taker fee rate. | [optional] 
**OrderPriceRound** | **string** | Minimum order price increment. | [optional] 
**MarkPriceRound** | **string** | Minimum mark price increment. | [optional] 
**FundingRate** | **string** | Current funding rate. | [optional] 
**FundingInterval** | **int32** | Funding application interval, unit in seconds. | [optional] 
**FundingNextApply** | **float64** | Next funding time. | [optional] 
**RiskLimitBase** | **string** | Risk limit base,deprecated. | [optional] 
**RiskLimitStep** | **string** | Step of adjusting risk limit,deprecated. | [optional] 
**RiskLimitMax** | **string** | Maximum risk limit the contract allowed,deprecated,It is recommended to use /futures/{settle}/risk_limit_tiers to query risk limits. | [optional] 
**OrderSizeMin** | **int64** | Minimum order size the contract allowed. | [optional] 
**OrderSizeMax** | **int64** | Maximum order size the contract allowed. | [optional] 
**OrderPriceDeviate** | **string** | deviation between order price and current index price. If price of an order is denoted as order_price, it must meet the following condition:   abs(order_price - mark_price) &lt;&#x3D; mark_price * order_price_deviate | [optional] 
**RefDiscountRate** | **string** | Referral fee rate discount. | [optional] 
**RefRebateRate** | **string** | Referrer commission rate. | [optional] 
**OrderbookId** | **int64** | Current orderbook ID. | [optional] 
**TradeId** | **int64** | Current trade ID. | [optional] 
**TradeSize** | **int64** | Historical accumulated trade size. | [optional] 
**PositionSize** | **int64** | Current total long position size. | [optional] 
**ConfigChangeTime** | **float64** | Last changed time of configuration. | [optional] 
**InDelisting** | **bool** | &#x60;in_delisting&#x3D;true&#x60; And when position_size&gt;0, it means the contract is in the offline transition period &#x60;in_delisting&#x3D;true&#x60; contract is offline | [optional] 
**OrdersLimit** | **int32** | Maximum number of open orders. | [optional] 
**EnableBonus** | **bool** | Whether bouns is enabled. | [optional] 
**EnableCredit** | **bool** | Whether portfolio margin account is enabled. | [optional] 
**CreateTime** | **float64** | Created time of the contract. | [optional] 
**FundingCapRatio** | **string** | The factor for the maximum of the funding rate. Maximum of funding rate &#x3D; (1/market maximum leverage - maintenance margin rate) * funding_cap_ratio | [optional] 
**Status** | **string** | Contract Status Types include: prelaunch, trading, delisting, delisted. | [optional] 
**LaunchTime** | **int64** | Contract expiry timestamp. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


