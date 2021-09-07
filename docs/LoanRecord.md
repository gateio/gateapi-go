# LoanRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Loan record ID | [optional] 
**LoanId** | **string** | Loan ID the record belongs to | [optional] 
**CreateTime** | **string** | Loan time | [optional] 
**ExpireTime** | **string** | Expiration time | [optional] 
**Status** | **string** | Loan record status | [optional] 
**BorrowUserId** | **string** | Garbled user ID | [optional] 
**Currency** | **string** | Loan currency | [optional] 
**Rate** | **string** | Loan rate | [optional] 
**Amount** | **string** | Loan amount | [optional] 
**Days** | **int32** | Loan days | [optional] 
**AutoRenew** | **bool** | Whether the record will auto renew on expiration | [optional] [default to false]
**Repaid** | **string** | Repaid amount | [optional] 
**PaidInterest** | **string** | Repaid interest | [optional] [readonly] 
**UnpaidInterest** | **string** | Outstanding interest yet to be paid | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


