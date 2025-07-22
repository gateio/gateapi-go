# OptionsContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Options contract name. | [optional] 
**Tag** | **string** | tag. | [optional] 
**CreateTime** | **float64** | Creation time. | [optional] 
**ExpirationTime** | **float64** | Expiration time. | [optional] 
**IsCall** | **bool** | &#x60;true&#x60; means call options, while &#x60;false&#x60; is put options. | [optional] 
**Multiplier** | **string** | Multiplier used in converting from invoicing to settlement currency. | [optional] 
**Underlying** | **string** | Underlying. | [optional] 
**UnderlyingPrice** | **string** | Underlying price (quote currency). | [optional] 
**LastPrice** | **string** | Last trading price. | [optional] 
**MarkPrice** | **string** | Current mark price (quote currency). | [optional] 
**IndexPrice** | **string** | Current index price (quote currency). | [optional] 
**MakerFeeRate** | **string** | Maker fee rate, where negative means rebate. | [optional] 
**TakerFeeRate** | **string** | Taker fee rate. | [optional] 
**OrderPriceRound** | **string** | Minimum order price increment. | [optional] 
**MarkPriceRound** | **string** | Minimum mark price increment. | [optional] 
**OrderSizeMin** | **int64** | Minimum order size the contract allowed. | [optional] 
**OrderSizeMax** | **int64** | Maximum order size the contract allowed. | [optional] 
**OrderPriceDeviate** | **string** | The positive and negative offset allowed between the order price and the current mark price, that &#x60;order_price&#x60; must meet the following conditions:   order_price is within the range of mark_price +/- order_price_deviate * underlying_price and does not distinguish between buy and sell orders | [optional] 
**RefDiscountRate** | **string** | Referral fee rate discount. | [optional] 
**RefRebateRate** | **string** | Referrer commission rate. | [optional] 
**OrderbookId** | **int64** | Current orderbook ID. | [optional] 
**TradeId** | **int64** | Current trade ID. | [optional] 
**TradeSize** | **int64** | Historical accumulated trade size. | [optional] 
**PositionSize** | **int64** | Current total long position size. | [optional] 
**OrdersLimit** | **int32** | Maximum number of open orders. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


