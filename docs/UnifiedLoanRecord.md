# UnifiedLoanRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | id | [optional] [readonly] 
**Type** | **string** | type: borrow - borrow, repay - repay | [optional] [readonly] 
**RepaymentType** | **string** | Repayment type, none - No repayment type, manual_repay - Manual repayment, auto_repay - Automatic repayment, cancel_auto_repay - Automatic repayment after withdrawal, different_currencies_repayment - Different currency repayment | [optional] [readonly] 
**BorrowType** | **string** | Loan type, returned when querying loan records. manual_borrow - Manual repayment , auto_borrow - Automatic repayment | [optional] 
**CurrencyPair** | **string** | Currency pair | [optional] [readonly] 
**Currency** | **string** | Currency | [optional] [readonly] 
**Amount** | **string** | The amount of lending or repaying | [optional] [readonly] 
**CreateTime** | **int64** | Created time | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


