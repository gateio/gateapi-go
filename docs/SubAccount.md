# SubAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Remark** | **string** | custom text. | [optional] 
**LoginName** | **string** | Sub-account login name: Only letters, numbers and underscores are supported, and cannot contain other illegal characters | 
**Password** | **string** | The sub-account&#39;s password. (Default: the same as main account&#39;s password). | [optional] 
**Email** | **string** | The sub-account&#39;s email address. (Default: the same as main account&#39;s email address) | [optional] 
**State** | **int32** | State: 1-normal, 2-locked\&quot;. | [optional] [readonly] 
**Type** | **int32** | \&quot;Sub-account type: 1 - sub-account, 3 - cross margin account. | [optional] [readonly] 
**UserId** | **int64** | The user id of the sub-account. | [optional] [readonly] 
**CreateTime** | **int64** | Created time. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


