# MultiCollateralOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **string** | Order ID | [optional] 
**OrderType** | **string** | current - current, fixed - fixed | [optional] 
**FixedType** | **string** | Fixed interest rate loan periods: 7d - 7 days, 30d - 30 days. | [optional] 
**FixedRate** | **string** | Fixed interest rate | [optional] 
**ExpireTime** | **int64** | Expiration time, timestamp, unit in seconds. | [optional] 
**AutoRenew** | **bool** | Fixed interest rate, automatic renewal | [optional] 
**AutoRepay** | **bool** | Fixed interest rate, automatic repayment | [optional] 
**CurrentLtv** | **string** | The current collateralization rate | [optional] 
**Status** | **string** | Order status: - initial: Initial state after placing the order - collateral_deducted: Collateral deduction successful - collateral_returning: Loan failed - Collateral return pending - lent: Loan successful - repaying: Repayment in progress - liquidating: Liquidation in progress - finished: Order completed - closed_liquidated: Liquidation and repayment completed | [optional] 
**BorrowTime** | **int64** | Borrowing time, timestamp in seconds | [optional] 
**TotalLeftRepayUsdt** | **string** | Value of Left repay amount converted in USDT | [optional] 
**TotalLeftCollateralUsdt** | **string** | Value of Collateral amount in USDT | [optional] 
**BorrowCurrencies** | [**[]BorrowCurrencyInfo**](BorrowCurrencyInfo.md) | Borrowing Currency List | [optional] 
**CollateralCurrencies** | [**[]CollateralCurrencyInfo**](CollateralCurrencyInfo.md) | Collateral Currency List | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


