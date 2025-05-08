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
**FailReason** | **string** | The reason for withdrawal failure is that there is a value when status &#x3D; CANCEL, and the rest of the state is empty | [optional] 
**Timestamp2** | **string** | The withdrawal end time, i.e.: withdrawal cancel time or withdrawal success time When status &#x3D; CANCEL, the corresponding cancel time When status &#x3D; DONE and block_number &gt; 0, it is the time to withdrawal success | [optional] 
**Memo** | **string** | Additional remarks with regards to the withdrawal | [optional] 
**Status** | **string** | Transaction status  - DONE: Completed (block_number &gt; 0 is considered to be truly completed) - CANCEL: Canceled - REQUEST: Requesting - MANUAL: Pending manual review - BCODE: Recharge code operation - EXTPEND: Sent awaiting confirmation - FAIL: Failure on the chain awaiting confirmation - INVALID: Invalid order - VERIFY: Verifying - PROCES: Processing - PEND: Processing - DMOVE: pending manual review - REVIEW: Under review | [optional] [readonly] 
**Chain** | **string** | Name of the chain used in withdrawals | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


