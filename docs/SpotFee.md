# SpotFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int64** | User ID | [optional] 
**TakerFee** | **string** | taker fee rate | [optional] 
**MakerFee** | **string** | maker fee rate | [optional] 
**GtDiscount** | **bool** | If GT deduction is enabled | [optional] 
**GtTakerFee** | **string** | Taker fee rate if using GT deduction. It will be 0 if GT deduction is disabled | [optional] 
**GtMakerFee** | **string** | Maker fee rate if using GT deduction. It will be 0 if GT deduction is disabled | [optional] 
**LoanFee** | **string** | Loan fee rate of margin lending | [optional] 
**PointType** | **string** | Point type. 0 - Initial version. 1 - new version since 202009 | [optional] 
**CurrencyPair** | **string** | Currency pair | [optional] 
**DebitFee** | **int32** | Deduction types for rates, 1 - GT deduction, 2 - Point card deduction, 3 - VIP rates | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


