# PlaceDualInvestmentOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanId** | **string** | Plan ID | 
**Copies** | **string** | The number of copies is mutually exclusive with the amount field and will be deprecated soon. It is recommended to use the amount parameter. | [optional] 
**IsMax** | **int32** | Whether to purchase at the maximum. Mutually exclusive with the amount field. Will be deprecated soon. It is recommended to use the amount parameter. | [optional] 
**Amount** | **string** | Subscription amount, mutually exclusive with the copies field | 
**Text** | **string** | User defined information. If not empty, must follow the rules below:  1. prefixed with &#x60;t-&#x60; 2. no longer than 28 bytes without &#x60;t-&#x60; prefix 3. can only include 0-9, A-Z, a-z, underscore(_), hyphen(-) or dot(.)  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


