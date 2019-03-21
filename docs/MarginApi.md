# \MarginApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelLoan**](MarginApi.md#CancelLoan) | **Delete** /margin/loans/{loan_id} | Cancel lending loan
[**CreateLoan**](MarginApi.md#CreateLoan) | **Post** /margin/loans | Lend or borrow
[**GetLoan**](MarginApi.md#GetLoan) | **Get** /margin/loans/{loan_id} | Retrieve one single loan detail
[**GetLoanRecord**](MarginApi.md#GetLoanRecord) | **Get** /margin/loan_records/{loan_record_id} | Get one single loan record
[**ListFundingAccounts**](MarginApi.md#ListFundingAccounts) | **Get** /margin/funding_accounts | Funding account list
[**ListFundingBook**](MarginApi.md#ListFundingBook) | **Get** /margin/funding_book | Order book of lending loans
[**ListLoanRecords**](MarginApi.md#ListLoanRecords) | **Get** /margin/loan_records | List repayment records of specified loan
[**ListLoanRepayments**](MarginApi.md#ListLoanRepayments) | **Get** /margin/loans/{loan_id}/repayment | List loan repayment records
[**ListLoans**](MarginApi.md#ListLoans) | **Get** /margin/loans | List all loans
[**ListMarginAccounts**](MarginApi.md#ListMarginAccounts) | **Get** /margin/accounts | Margin account list
[**ListMarginCurrencyPairs**](MarginApi.md#ListMarginCurrencyPairs) | **Get** /margin/currency_pairs | List all supported currency pairs supported in margin trading
[**MergeLoans**](MarginApi.md#MergeLoans) | **Post** /margin/merged_loans | Merge multiple lending loans
[**RepayLoan**](MarginApi.md#RepayLoan) | **Post** /margin/loans/{loan_id}/repayment | Repay a loan
[**UpdateLoan**](MarginApi.md#UpdateLoan) | **Patch** /margin/loans/{loan_id} | Modify a loan
[**UpdateLoanRecord**](MarginApi.md#UpdateLoanRecord) | **Patch** /margin/loan_records/{loan_record_id} | Modify a loan record


# **CancelLoan**
> Loan CancelLoan(ctx, loanId, currency)
Cancel lending loan

Only lending loans can be cancelled

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loanId** | **string**| Loan ID | 
  **currency** | **string**| Retrieved specified currency related data | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.MarginApi
loanId := "12345"; // string - Loan ID
currency := "BTC"; // string - Retrieved specified currency related data

result, _, err := api.CancelLoan(nil, loanId, currency)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**Loan**](Loan.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLoan**
> Loan CreateLoan(ctx, loan)
Lend or borrow

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loan** | [**Loan**](Loan.md)|  | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.MarginApi
loan := new (gateapi.Loan); // Loan - 

result, _, err := api.CreateLoan(nil, loan)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**Loan**](Loan.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLoan**
> Loan GetLoan(ctx, loanId, side)
Retrieve one single loan detail

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loanId** | **string**| Loan ID | 
  **side** | **string**| Lend or borrow | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.MarginApi
loanId := "12345"; // string - Loan ID
side := "lend"; // string - Lend or borrow

result, _, err := api.GetLoan(nil, loanId, side)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**Loan**](Loan.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLoanRecord**
> LoanRecord GetLoanRecord(ctx, loanRecordId, loanId)
Get one single loan record

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loanRecordId** | **string**| Loan record ID | 
  **loanId** | **string**| Loan ID | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.MarginApi
loanRecordId := "12345"; // string - Loan record ID
loanId := "12345"; // string - Loan ID

result, _, err := api.GetLoanRecord(nil, loanRecordId, loanId)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**LoanRecord**](LoanRecord.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFundingAccounts**
> []FundingAccount ListFundingAccounts(ctx, optional)
Funding account list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListFundingAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListFundingAccountsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | **optional.String**| Retrieved specified currency related data | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.MarginApi

result, _, err := api.ListFundingAccounts(nil, nil)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**[]FundingAccount**](FundingAccount.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFundingBook**
> []FundingBookItem ListFundingBook(ctx, currency)
Order book of lending loans

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **currency** | **string**| Retrieved specified currency related data | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
api := client.MarginApi
currency := "BTC"; // string - Retrieved specified currency related data

result, _, err := api.ListFundingBook(nil, currency)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**[]FundingBookItem**](FundingBookItem.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLoanRecords**
> []LoanRecord ListLoanRecords(ctx, loanId, optional)
List repayment records of specified loan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loanId** | **string**| Loan ID | 
 **optional** | ***ListLoanRecordsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListLoanRecordsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **optional.String**| Loan record status | 
 **page** | **optional.Int32**| Page number | [default to 1]
 **limit** | **optional.Int32**| Maximum number of record returned in one list | [default to 100]

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.MarginApi
loanId := "12345"; // string - Loan ID

result, _, err := api.ListLoanRecords(nil, loanId, nil)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**[]LoanRecord**](LoanRecord.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLoanRepayments**
> []Repayment ListLoanRepayments(ctx, loanId)
List loan repayment records

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loanId** | **string**| Loan ID | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.MarginApi
loanId := "12345"; // string - Loan ID

result, _, err := api.ListLoanRepayments(nil, loanId)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**[]Repayment**](Repayment.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLoans**
> []Loan ListLoans(ctx, status, side, optional)
List all loans

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **status** | **string**| Loan status | 
  **side** | **string**| Lend or borrow | 
 **optional** | ***ListLoansOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListLoansOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **currency** | **optional.String**| Retrieved specified currency related data | 
 **page** | **optional.Int32**| Page number | [default to 1]
 **limit** | **optional.Int32**| Maximum number of record returned in one list | [default to 100]

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.MarginApi
status := "open"; // string - Loan status
side := "lend"; // string - Lend or borrow

result, _, err := api.ListLoans(nil, status, side, nil)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**[]Loan**](Loan.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMarginAccounts**
> []MarginAccount ListMarginAccounts(ctx, optional)
Margin account list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListMarginAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListMarginAccountsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currencyPair** | **optional.String**| Currency pair | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.MarginApi

result, _, err := api.ListMarginAccounts(nil, nil)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**[]MarginAccount**](MarginAccount.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMarginCurrencyPairs**
> []MarginCurrencyPair ListMarginCurrencyPairs(ctx, )
List all supported currency pairs supported in margin trading

### Required Parameters
This endpoint does not need any parameter.

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
api := client.MarginApi
result, _, err := api.ListMarginCurrencyPairs(nil)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**[]MarginCurrencyPair**](MarginCurrencyPair.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MergeLoans**
> Loan MergeLoans(ctx, currency, ids)
Merge multiple lending loans

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **currency** | **string**| Retrieved specified currency related data | 
  **ids** | **string**| Lending loan ID list separated by &#x60;,&#x60;. Maximum of 20 IDs are allowed in one request | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.MarginApi
currency := "BTC"; // string - Retrieved specified currency related data
ids := "123,234,345"; // string - Lending loan ID list separated by `,`. Maximum of 20 IDs are allowed in one request

result, _, err := api.MergeLoans(nil, currency, ids)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**Loan**](Loan.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepayLoan**
> Loan RepayLoan(ctx, loanId, repayRequest)
Repay a loan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loanId** | **string**| Loan ID | 
  **repayRequest** | [**RepayRequest**](RepayRequest.md)|  | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.MarginApi
loanId := "12345"; // string - Loan ID
repayRequest := new (gateapi.RepayRequest); // RepayRequest - 

result, _, err := api.RepayLoan(nil, loanId, repayRequest)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**Loan**](Loan.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLoan**
> Loan UpdateLoan(ctx, loanId, loanPatch)
Modify a loan

Only `auto_renew` modification is supported currently

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loanId** | **string**| Loan ID | 
  **loanPatch** | [**LoanPatch**](LoanPatch.md)|  | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.MarginApi
loanId := "12345"; // string - Loan ID
loanPatch := new (gateapi.LoanPatch); // LoanPatch - 

result, _, err := api.UpdateLoan(nil, loanId, loanPatch)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**Loan**](Loan.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLoanRecord**
> LoanRecord UpdateLoanRecord(ctx, loanRecordId, loanPatch)
Modify a loan record

Only `auto_renew` modification is supported currently

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loanRecordId** | **string**| Loan record ID | 
  **loanPatch** | [**LoanPatch**](LoanPatch.md)|  | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.MarginApi
loanRecordId := "12345"; // string - Loan record ID
loanPatch := new (gateapi.LoanPatch); // LoanPatch - 

result, _, err := api.UpdateLoanRecord(nil, loanRecordId, loanPatch)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**LoanRecord**](LoanRecord.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

