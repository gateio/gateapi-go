# Loan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Loan ID | [optional] [readonly] 
**CreateTime** | **string** | Creation time | [optional] [readonly] 
**ExpireTime** | **string** | Repay time of the loan. No value will be returned for lending loan | [optional] [readonly] 
**Status** | **string** | Loan status  open - not fully loaned loaned - all loaned out for lending loan; loaned in for borrowing side finished - loan is finished, either being all repaid or cancelled by the lender auto_repaid - automatically repaid by the system | [optional] [readonly] 
**Side** | **string** | Loan side | 
**Currency** | **string** | Loan currency | 
**Rate** | **string** | Loan rate. Only rates in [0.0002, 0.002] are supported.  Not required in lending. Market rate calculated from recent rates will be used if not set | [optional] 
**Amount** | **string** | Loan amount | 
**Days** | **int32** | Loan days. Only 10 is supported for now | [optional] 
**AutoRenew** | **bool** | Whether to auto renew the loan upon expiration | [optional] [default to false]
**CurrencyPair** | **string** | Currency pair. Required if borrowing | [optional] 
**Left** | **string** | Amount not lent out yet | [optional] [readonly] 
**Repaid** | **string** | Repaid amount | [optional] [readonly] 
**PaidInterest** | **string** | Repaid interest | [optional] [readonly] 
**UnpaidInterest** | **string** | Outstanding interest yet to be paid | [optional] [readonly] 
**FeeRate** | **string** | Loan fee rate | [optional] 
**OrigId** | **string** | Original loan ID of the loan if auto-renewed, otherwise equals to id | [optional] 
**Text** | **string** | User defined custom ID | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


