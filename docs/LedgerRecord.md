# LedgerRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Record ID | [optional] [readonly] 
**Txid** | **string** | Hash record of the withdrawal | [optional] [readonly] 
**Timestamp** | **string** | Operation time | [optional] [readonly] 
**Amount** | **string** | Currency amount | 
**Currency** | **string** | Currency name | 
**Address** | **string** | Withdrawal address. Required for withdrawals | [optional] 
**Memo** | **string** | Extra withdrawal memo | [optional] 
**Status** | **string** | Record status.  - DONE: done - CANCEL: cancelled - REQUEST: requesting - MANUAL: waiting for manual approval - BCODE: GateCode operation | [optional] [readonly] 
**Chain** | **string** | Name of the chain used in withdrawals | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


