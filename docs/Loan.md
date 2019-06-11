# Loan

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Loan ID | [optional] 
**CreateTime** | **string** | Creation time | [optional] 
**ExpireTime** | **string** | Repay time of the loan. No value will be returned for lending loan | [optional] 
**Status** | **string** | Loan status  open - not fully loaned loaned - all loaned out for lending loan; loaned in for borrowing side finished - loan is finished, either being all repaid or cancelled by the lender auto_repaid - automatically repaid by the system | [optional] 
**Side** | **string** | Loan side | 
**Currency** | **string** | Loan currency | 
**Rate** | **string** | Loan rate. Only rates in [0.0002, 0.002] are supported.  Not required in lending. Market rate calculated from recent rates will be used if not set | [optional] 
**Amount** | **string** | Loan amount | 
**Days** | **int32** | Loan days | 
**AutoRenew** | **bool** | Auto renew the loan on expiration | [optional] [default to false]
**CurrencyPair** | **string** | Currency pair. Required for borrowing side | [optional] 
**Left** | **string** | Amount not lending out | [optional] 
**Repaid** | **string** | Repaid amount | [optional] 
**PaidInterest** | **string** | Repaid interest | [optional] 
**UnpaidInterest** | **string** | Interest not repaid | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


