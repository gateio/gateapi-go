# \WalletApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDepositAddress**](WalletApi.md#GetDepositAddress) | **Get** /wallet/deposit_address | Generate currency deposit address
[**ListDeposits**](WalletApi.md#ListDeposits) | **Get** /wallet/deposits | Retrieve deposit records. Time range cannot exceed 30 days
[**ListWithdrawals**](WalletApi.md#ListWithdrawals) | **Get** /wallet/withdrawals | Retrieve withdrawal records. Time range cannot exceed 30 days
[**Transfer**](WalletApi.md#Transfer) | **Post** /wallet/transfers | Transfer between accounts
[**TransferWithSubAccount**](WalletApi.md#TransferWithSubAccount) | **Post** /wallet/sub_account_transfers | Transfer between main and sub accounts


# **GetDepositAddress**
> DepositAddress GetDepositAddress(ctx, currency)
Generate currency deposit address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **currency** | **string**| Currency name | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.WalletApi
currency := "currency_example"; // string - Currency name

result, _, err := api.GetDepositAddress(nil, currency)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**DepositAddress**](DepositAddress.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDeposits**
> []LedgerRecord ListDeposits(ctx, optional)
Retrieve deposit records. Time range cannot exceed 30 days

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListDepositsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListDepositsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | **optional.String**| Filter by currency. Return all currency records if not specified | 
 **from** | **optional.Int64**| Time range beginning, default to 7 days before current time | 
 **to** | **optional.Int64**| Time range ending, default to current time | 
 **limit** | **optional.Int32**| Maximum number of record returned in one list | [default to 100]
 **offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.WalletApi

result, _, err := api.ListDeposits(nil, nil)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**[]LedgerRecord**](LedgerRecord.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListWithdrawals**
> []LedgerRecord ListWithdrawals(ctx, optional)
Retrieve withdrawal records. Time range cannot exceed 30 days

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListWithdrawalsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListWithdrawalsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | **optional.String**| Filter by currency. Return all currency records if not specified | 
 **from** | **optional.Int64**| Time range beginning, default to 7 days before current time | 
 **to** | **optional.Int64**| Time range ending, default to current time | 
 **limit** | **optional.Int32**| Maximum number of record returned in one list | [default to 100]
 **offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.WalletApi

result, _, err := api.ListWithdrawals(nil, nil)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**[]LedgerRecord**](LedgerRecord.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Transfer**
> Transfer(ctx, transfer)
Transfer between accounts

Transfer between different accounts. Currently support transfers between the following:  1. spot - margin 2. spot - futures

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

# **TransferWithSubAccount**
> TransferWithSubAccount(ctx, subAccountTransfer)
Transfer between main and sub accounts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subAccountTransfer** | [**SubAccountTransfer**](SubAccountTransfer.md)|  | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.WalletApi
subAccountTransfer := new (gateapi.SubAccountTransfer); // SubAccountTransfer - 

result, _, err := api.TransferWithSubAccount(nil, subAccountTransfer)
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

