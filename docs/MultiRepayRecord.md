# MultiRepayRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **int64** | Order ID. | [optional] 
**RecordId** | **int64** | Repayment record ID. | [optional] 
**InitLtv** | **string** | The initial collateralization rate. | [optional] 
**BeforeLtv** | **string** | Ltv before the operation. | [optional] 
**AfterLtv** | **string** | Ltv after the operation. | [optional] 
**BorrowTime** | **int64** | Borrowing time, timestamp in seconds. | [optional] 
**RepayTime** | **int64** | Repayment time, timestamp in seconds. | [optional] 
**BorrowCurrencies** | [**[]RepayRecordCurrency**](RepayRecordCurrency.md) | List of borrowing information. | [optional] 
**CollateralCurrencies** | [**[]RepayRecordCurrency**](RepayRecordCurrency.md) | List of collateral information. | [optional] 
**RepaidCurrencies** | [**[]RepayRecordRepaidCurrency**](RepayRecordRepaidCurrency.md) | Repay Currency List. | [optional] 
**TotalInterestList** | [**[]RepayRecordTotalInterest**](RepayRecordTotalInterest.md) | Total Interest List. | [optional] 
**LeftRepayInterestList** | [**[]RepayRecordLeftInterest**](RepayRecordLeftInterest.md) | List of left repay interest. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


