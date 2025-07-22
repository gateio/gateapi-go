# FuturesOrderAmendment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | **int64** | New order size, including filled part.  - If new size is less than or equal to filled size, the order will be cancelled. - Order side must be identical to the original one. - Close order size cannot be changed. - For reduce only orders, increasing size may leads to other reduce only orders being cancelled. - If price is not changed, decreasing size will not change its precedence in order book, while increasing will move it to the last at current price. | [optional] 
**Price** | **string** | New order price. | [optional] 
**AmendText** | **string** | Custom info during amending order. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


