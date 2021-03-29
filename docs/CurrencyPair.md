# CurrencyPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Currency pair | [optional] 
**Base** | **string** | Base currency | [optional] 
**Quote** | **string** | Quote currency | [optional] 
**Fee** | **string** | Trading fee | [optional] 
**MinBaseAmount** | **string** | Minimum amount of base currency to trade, &#x60;null&#x60; means no limit | [optional] 
**MinQuoteAmount** | **string** | Minimum amount of quote currency to trade, &#x60;null&#x60; means no limit | [optional] 
**AmountPrecision** | **int32** | Amount scale | [optional] 
**Precision** | **int32** | Price scale | [optional] 
**TradeStatus** | **string** | How currency pair can be traded  - untradable: cannot be bought or sold - buyable: can be bought - sellable: can be sold - tradable: can be bought or sold | [optional] 
**SellStart** | **int64** | Sell start unix timestamp in seconds | [optional] 
**BuyStart** | **int64** | Buy start unix timestamp in seconds | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


