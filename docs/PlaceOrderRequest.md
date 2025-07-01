# PlaceOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | Trading Symbol | 
**Side** | **string** | 买单或者卖单 - buy - sell | 
**Amount** | **string** | Trade Quantity - &#x60;side&#x60; : &#x60;buy&#x60; refers to the quote currency, i.e., &#x60;USDT&#x60; - &#x60;side&#x60; : &#x60;sell&#x60; refers to the base currency | 
**GasMode** | **string** | Trading mode affects slippage selection - &#x60;speed&#x60; : Smart mode - &#x60;custom&#x60; : Custom mode, uses &#x60;slippage&#x60; parameter | 
**Slippage** | **string** | Slippage value of 10 represents a 10% tolerance | [optional] 
**QuoteId** | **string** | The quote_id returned by the quotation API | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


