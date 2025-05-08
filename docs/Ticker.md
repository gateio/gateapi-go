# Ticker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyPair** | **string** | Currency pair | [optional] 
**Last** | **string** | Last trading price | [optional] 
**LowestAsk** | **string** | Recent lowest ask | [optional] 
**LowestSize** | **string** | The latest seller&#39;s lowest price quantity; does not exist for batch query; exists for single query, and is empty if there is no data | [optional] 
**HighestBid** | **string** | Recent highest bid | [optional] 
**HighestSize** | **string** | The latest buyer&#39;s highest price quantity; does not exist for batch query; exists for single query, and is empty if there is no data | [optional] 
**ChangePercentage** | **string** | Change percentage in the last 24h | [optional] 
**ChangeUtc0** | **string** | utc0 timezone, the percentage change in the last 24 hours | [optional] 
**ChangeUtc8** | **string** | utc8 timezone, the percentage change in the last 24 hours | [optional] 
**BaseVolume** | **string** | Base currency trade volume in the last 24h | [optional] 
**QuoteVolume** | **string** | Quote currency trade volume in the last 24h | [optional] 
**High24h** | **string** | Highest price in 24h | [optional] 
**Low24h** | **string** | Lowest price in 24h | [optional] 
**EtfNetValue** | **string** | ETF net value | [optional] 
**EtfPreNetValue** | Pointer to **string** | ETF previous net value at re-balancing time | [optional] 
**EtfPreTimestamp** | Pointer to **int64** | ETF previous re-balancing time | [optional] 
**EtfLeverage** | Pointer to **string** | ETF current leverage | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


