# MockMarginResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Position combination type &#x60;original_position&#x60; - Original position &#x60;long_delta_original_position&#x60; - Positive delta + Original position &#x60;short_delta_original_position&#x60; - Negative delta + Original position | [optional] 
**ProfitLossRanges** | [**[]ProfitLossRange**](ProfitLossRange.md) | The results of 33 pressure scenarios for MR1 | [optional] 
**MaxLoss** | [**ProfitLossRange**](.md) | 最大损失 | [optional] 
**Mr1** | **string** | Stress testing | [optional] 
**Mr2** | **string** | Basis spread risk | [optional] 
**Mr3** | **string** | Volatility spread risk | [optional] 
**Mr4** | **string** | Option short risk | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


