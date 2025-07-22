# UnifiedPortfolioOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaintainMarginTotal** | **string** | Total maintenance margin, including only the portfolio margin calculation results for positions in the risk unit,  excluding borrowed margin. If borrowing exists, conventional borrowing margin requirements will still apply. | [optional] 
**InitialMarginTotal** | **string** | Total initial margin, calculated as the maximum of the following three combinations: position,  position + positive delta orders, orders. | [optional] 
**CalculateTime** | **int64** | Calculate time. | [optional] 
**RiskUnit** | [**[]MockRiskUnit**](MockRiskUnit.md) | Risk unit. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


