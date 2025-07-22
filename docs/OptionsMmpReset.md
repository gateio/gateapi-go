# OptionsMmpReset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Underlying** | **string** | Underlying. | 
**Window** | **int32** | Time window (milliseconds), between 1-5000, 0 means disabling MMP. | [optional] [readonly] 
**FrozenPeriod** | **int32** | Freeze duration (milliseconds), 0 means always frozen, need to call reset API to unfreeze | [optional] [readonly] 
**QtyLimit** | **string** | Trading volume upper limit (positive number, up to 2 decimal places). | [optional] [readonly] 
**DeltaLimit** | **string** | Upper limit of net delta value (positive number, up to 2 decimal places). | [optional] [readonly] 
**TriggerTimeMs** | **int64** | Trigger freeze time (milliseconds), 0 means no freeze is triggered. | [optional] [readonly] 
**FrozenUntilMs** | **int64** | Unfreeze time (milliseconds). If the freeze duration is not configured, there will be no unfreeze time after the freeze is triggered. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


