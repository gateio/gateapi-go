# SubAccountKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int64** | User ID. | [optional] [readonly] 
**Mode** | **int32** | Mode: 1 - classic 2 - portfolio account. | [optional] 
**Name** | **string** | API key name. | [optional] 
**Perms** | [**[]SubAccountKeyPerms**](SubAccountKey_perms.md) |  | [optional] 
**IpWhitelist** | **[]string** | ip white list (list will be removed if no value is passed). | [optional] 
**Key** | **string** | API Key. | [optional] [readonly] 
**State** | **int32** | State 1 - normal 2 - locked 3 - frozen. | [optional] [readonly] 
**CreatedAt** | **int64** | Creation time. | [optional] [readonly] 
**UpdatedAt** | **int64** | Last update time. | [optional] [readonly] 
**LastAccess** | **int64** | Last access time. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


