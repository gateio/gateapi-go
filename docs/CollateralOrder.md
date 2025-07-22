# CollateralOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **int64** | Order ID. | [optional] 
**CollateralCurrency** | **string** | Collateral. | [optional] 
**CollateralAmount** | **string** | Collateral amount. | [optional] 
**BorrowCurrency** | **string** | Borrowed currency. | [optional] 
**BorrowAmount** | **string** | Borrowing amount. | [optional] 
**RepaidAmount** | **string** | Repaid amount. | [optional] 
**RepaidPrincipal** | **string** | Repaid principal. | [optional] 
**RepaidInterest** | **string** | Repaid interest. | [optional] 
**InitLtv** | **string** | The initial collateralization rate. | [optional] 
**CurrentLtv** | **string** | The current collateralization rate. | [optional] 
**LiquidateLtv** | **string** | The liquidation collateralization rate. | [optional] 
**Status** | **string** | Order status: - initial: Initial state after placing the order - collateral_deducted: Collateral deduction successful - collateral_returning: Loan failed - Collateral return pending - lent: Loan successful - repaying: Repayment in progress - liquidating: Liquidation in progress - finished: Order completed - closed_liquidated: Liquidation and repayment completed | [optional] 
**BorrowTime** | **int64** | Borrowing time, timestamp in seconds. | [optional] 
**LeftRepayTotal** | **string** | Outstanding principal and interest (outstanding principal + outstanding interest) | [optional] 
**LeftRepayPrincipal** | **string** | outstanding principal. | [optional] 
**LeftRepayInterest** | **string** | outstanding interest. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


