# \WalletApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Transfer**](WalletApi.md#Transfer) | **Post** /wallet/transfers | Transfer between accounts


# **Transfer**
> Transfer(ctx, transfer)
Transfer between accounts

Transfer between different accounts. Currently support transfers between the following:  1. spot - margin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transfer** | [**Transfer**](Transfer.md)|  | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.WalletApi
transfer := new (gateapi.Transfer); // Transfer - 

result, _, err := api.Transfer(nil, transfer)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

 (empty response body)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

