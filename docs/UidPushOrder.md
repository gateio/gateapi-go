# UidPushOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Order ID. | [optional] 
**PushUid** | **int64** | Initiator User ID. | [optional] 
**ReceiveUid** | **int64** | Recipient User ID. | [optional] 
**Currency** | **string** | Currency name. | [optional] 
**Amount** | **string** | Transfer amount. | [optional] 
**CreateTime** | **int64** | Creation time. | [optional] 
**Status** | **string** | Withdrawal Status  - CREATING: Creating - PENDING: Waiting for receiving(Please contact the other party to accept the transfer on the Gate official website) - CANCELLING: Cancelling - CANCELLED: Revoked - REFUSING: Rejection - REFUSED: Rejected - RECEIVING: Receiving - RECEIVED: Success | [optional] 
**Message** | **string** | PENDING Reason Tips. | [optional] 
**TransactionType** | **string** | Order Type. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


