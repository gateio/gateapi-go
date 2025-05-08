# CreateMultiCollateralOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **string** | Order ID | [optional] 
**OrderType** | **string** | current - current, fixed - fixed, if not specified, default to current | [optional] 
**FixedType** | **string** | Fixed interest rate loan period: 7d - 7 days, 30d - 30 days. Must be provided for fixed | [optional] 
**FixedRate** | **string** | Fixed interest rate, must be specified for fixed | [optional] 
**AutoRenew** | **bool** | Fixed interest rate, automatic renewal | [optional] 
**AutoRepay** | **bool** | Fixed interest rate, automatic repayment | [optional] 
**BorrowCurrency** | **string** | Borrowed currency | 
**BorrowAmount** | **string** | Borrowing amount | 
**CollateralCurrencies** | [**[]CollateralCurrency**](CollateralCurrency.md) | Collateral currency and amount | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


