# FuturesTrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Trade ID. | [optional] 
**CreateTime** | **float64** | Trading time. | [optional] 
**CreateTimeMs** | **float64** | Trading time, with milliseconds set to 3 decimal places. | [optional] 
**Contract** | **string** | Futures contract. | [optional] 
**Size** | **int64** | Trading size. | [optional] 
**Price** | **string** | Trading price (quote currency). | [optional] 
**IsInternal** | **bool** | Whether internal trade. Internal trade refers to the takeover of liquidation orders by the insurance fund and ADL users. Since it is not a normal matching on the market depth, the transaction price may deviate, and it will not be recorded in the K-line. an internal trade, this field will not be returned. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


