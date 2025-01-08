# WithdrawalRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Record ID | [optional] [readonly] 
**Txid** | **string** | Hash record of the withdrawal | [optional] [readonly] 
**BlockNumber** | **string** | 区块编号 | [optional] [readonly] 
**WithdrawOrderId** | **string** | Client order id, up to 32 length and can only include 0-9, A-Z, a-z, underscore(_), hyphen(-) or dot(.)  | [optional] 
**Timestamp** | **string** | Operation time | [optional] [readonly] 
**Amount** | **string** | Currency amount | 
**Fee** | **string** | fee | [optional] [readonly] 
**Currency** | **string** | Currency name | 
**Address** | **string** | Withdrawal address. Required for withdrawals | [optional] 
**Memo** | **string** | Additional remarks with regards to the withdrawal | [optional] 
**Status** | **string** | Transaction status  - DONE: Completed (block_number &gt; 0 is considered to be truly completed) - CANCEL: Canceled - REQUEST: Requesting - MANUAL: Pending manual review - BCODE: Recharge code operation - EXTPEND: Sent awaiting confirmation - FAIL: Failure on the chain awaiting confirmation - INVALID: Invalid order - VERIFY: Verifying - PROCES: Processing - PEND: Processing - DMOVE: pending manual review - SPLITPEND: cny withdrawal is greater than 50,000, orders will be split automatically | [optional] [readonly] 
**Chain** | **string** | Name of the chain used in withdrawals | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


