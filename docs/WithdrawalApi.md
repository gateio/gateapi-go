# \WithdrawalApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Withdraw**](WithdrawalApi.md#Withdraw) | **Post** /withdrawals | Withdraw


# **Withdraw**
> LedgerRecord Withdraw(ctx, ledgerRecord)
Withdraw

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ledgerRecord** | [**LedgerRecord**](LedgerRecord.md)|  | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.WithdrawalApi
ledgerRecord := new (gateapi.LedgerRecord); // LedgerRecord - 

result, _, err := api.Withdraw(nil, ledgerRecord)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**LedgerRecord**](LedgerRecord.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

