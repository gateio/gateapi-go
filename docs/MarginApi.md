# MarginApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListMarginCurrencyPairs**](MarginApi.md#ListMarginCurrencyPairs) | **Get** /margin/currency_pairs | List all supported currency pairs supported in margin trading
[**ListFundingBook**](MarginApi.md#ListFundingBook) | **Get** /margin/funding_book | Order book of lending loans
[**ListMarginAccounts**](MarginApi.md#ListMarginAccounts) | **Get** /margin/accounts | Margin account list
[**ListMarginAccountBook**](MarginApi.md#ListMarginAccountBook) | **Get** /margin/account_book | List margin account balance change history
[**ListFundingAccounts**](MarginApi.md#ListFundingAccounts) | **Get** /margin/funding_accounts | Funding account list
[**ListLoans**](MarginApi.md#ListLoans) | **Get** /margin/loans | List all loans
[**CreateLoan**](MarginApi.md#CreateLoan) | **Post** /margin/loans | Lend or borrow
[**MergeLoans**](MarginApi.md#MergeLoans) | **Post** /margin/merged_loans | Merge multiple lending loans
[**GetLoan**](MarginApi.md#GetLoan) | **Get** /margin/loans/{loan_id} | Retrieve one single loan detail
[**CancelLoan**](MarginApi.md#CancelLoan) | **Delete** /margin/loans/{loan_id} | Cancel lending loan
[**UpdateLoan**](MarginApi.md#UpdateLoan) | **Patch** /margin/loans/{loan_id} | Modify a loan
[**ListLoanRepayments**](MarginApi.md#ListLoanRepayments) | **Get** /margin/loans/{loan_id}/repayment | List loan repayment records
[**RepayLoan**](MarginApi.md#RepayLoan) | **Post** /margin/loans/{loan_id}/repayment | Repay a loan
[**ListLoanRecords**](MarginApi.md#ListLoanRecords) | **Get** /margin/loan_records | List repayment records of specified loan
[**GetLoanRecord**](MarginApi.md#GetLoanRecord) | **Get** /margin/loan_records/{loan_record_id} | Get one single loan record
[**UpdateLoanRecord**](MarginApi.md#UpdateLoanRecord) | **Patch** /margin/loan_records/{loan_record_id} | Modify a loan record


## ListMarginCurrencyPairs

> []MarginCurrencyPair ListMarginCurrencyPairs(ctx, )

List all supported currency pairs supported in margin trading

### Required Parameters


### Example

```golang
package main

import (
    "context"
    "fmt"
    "gateapi"
)

func main() {
    client := gateapi.NewAPIClient(gateapi.NewConfiguration())
    // uncomment the next line if your are testing against testnet
    // client.ChangeBasePath("https://fx-api-testnet.gateio.ws/api/v4")
    ctx := context.Background()
    
    result, _, err := client.MarginApi.ListMarginCurrencyPairs(ctx)
    if err != nil {
        if e, ok := err.(gateapi.GateAPIError); ok {
            fmt.Printf("gate api error: %s\n", e.Error())
        } else {
            fmt.Printf("generic error: %s\n", err.Error())
        }
    } else {
        fmt.Println(result)
    }
}
```


### Return type

[**[]MarginCurrencyPair**](MarginCurrencyPair.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListFundingBook

> []FundingBookItem ListFundingBook(ctx, currency)

Order book of lending loans

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currency** | **string**| Retrieved specified currency related data | 

### Example

```golang
package main

import (
    "context"
    "fmt"
    "gateapi"
)

func main() {
    client := gateapi.NewAPIClient(gateapi.NewConfiguration())
    // uncomment the next line if your are testing against testnet
    // client.ChangeBasePath("https://fx-api-testnet.gateio.ws/api/v4")
    ctx := context.Background()
    currency := "BTC" // string - Retrieved specified currency related data
    
    result, _, err := client.MarginApi.ListFundingBook(ctx, currency)
    if err != nil {
        if e, ok := err.(gateapi.GateAPIError); ok {
            fmt.Printf("gate api error: %s\n", e.Error())
        } else {
            fmt.Printf("generic error: %s\n", err.Error())
        }
    } else {
        fmt.Println(result)
    }
}
```


### Return type

[**[]FundingBookItem**](FundingBookItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListMarginAccounts

> []MarginAccount ListMarginAccounts(ctx, optional)

Margin account list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListMarginAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListMarginAccountsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currencyPair** | **optional.String**| Currency pair | 

### Example

```golang
package main

import (
    "context"
    "fmt"
    "gateapi"
)

func main() {
    client := gateapi.NewAPIClient(gateapi.NewConfiguration())
    // uncomment the next line if your are testing against testnet
    // client.ChangeBasePath("https://fx-api-testnet.gateio.ws/api/v4")
    ctx := context.WithValue(context.Background(),
                             gateapi.ContextGateAPIV4,
                             gateapi.GateAPIV4{
                                 Key:    "YOUR_API_KEY",
                                 Secret: "YOUR_API_SECRET",
                             }
                            )
    
    result, _, err := client.MarginApi.ListMarginAccounts(ctx, nil)
    if err != nil {
        if e, ok := err.(gateapi.GateAPIError); ok {
            fmt.Printf("gate api error: %s\n", e.Error())
        } else {
            fmt.Printf("generic error: %s\n", err.Error())
        }
    } else {
        fmt.Println(result)
    }
}
```


### Return type

[**[]MarginAccount**](MarginAccount.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListMarginAccountBook

> []MarginAccountBook ListMarginAccountBook(ctx, optional)

List margin account balance change history

Only transferring from or to margin account are provided for now. Time range allows 30 days at most

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListMarginAccountBookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListMarginAccountBookOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currency** | **optional.String**| List records related to specified currency only. If specified, &#x60;currency_pair&#x60; is also required. | 
**currencyPair** | **optional.String**| List records related to specified currency pair. Used in combination with &#x60;currency&#x60;. Ignored if &#x60;currency&#x60; is not provided | 
**from** | **optional.Int64**| Time range beginning, default to 7 days before current time | 
**to** | **optional.Int64**| Time range ending, default to current time | 
**page** | **optional.Int32**| Page number | [default to 1]
**limit** | **optional.Int32**| Maximum number of records returned in one list | [default to 100]

### Example

```golang
package main

import (
    "context"
    "fmt"
    "gateapi"
)

func main() {
    client := gateapi.NewAPIClient(gateapi.NewConfiguration())
    // uncomment the next line if your are testing against testnet
    // client.ChangeBasePath("https://fx-api-testnet.gateio.ws/api/v4")
    ctx := context.WithValue(context.Background(),
                             gateapi.ContextGateAPIV4,
                             gateapi.GateAPIV4{
                                 Key:    "YOUR_API_KEY",
                                 Secret: "YOUR_API_SECRET",
                             }
                            )
    
    result, _, err := client.MarginApi.ListMarginAccountBook(ctx, nil)
    if err != nil {
        if e, ok := err.(gateapi.GateAPIError); ok {
            fmt.Printf("gate api error: %s\n", e.Error())
        } else {
            fmt.Printf("generic error: %s\n", err.Error())
        }
    } else {
        fmt.Println(result)
    }
}
```


### Return type

[**[]MarginAccountBook**](MarginAccountBook.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListFundingAccounts

> []FundingAccount ListFundingAccounts(ctx, optional)

Funding account list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListFundingAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFundingAccountsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currency** | **optional.String**| Retrieved specified currency related data | 

### Example

```golang
package main

import (
    "context"
    "fmt"
    "gateapi"
)

func main() {
    client := gateapi.NewAPIClient(gateapi.NewConfiguration())
    // uncomment the next line if your are testing against testnet
    // client.ChangeBasePath("https://fx-api-testnet.gateio.ws/api/v4")
    ctx := context.WithValue(context.Background(),
                             gateapi.ContextGateAPIV4,
                             gateapi.GateAPIV4{
                                 Key:    "YOUR_API_KEY",
                                 Secret: "YOUR_API_SECRET",
                             }
                            )
    
    result, _, err := client.MarginApi.ListFundingAccounts(ctx, nil)
    if err != nil {
        if e, ok := err.(gateapi.GateAPIError); ok {
            fmt.Printf("gate api error: %s\n", e.Error())
        } else {
            fmt.Printf("generic error: %s\n", err.Error())
        }
    } else {
        fmt.Println(result)
    }
}
```


### Return type

[**[]FundingAccount**](FundingAccount.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListLoans

> []Loan ListLoans(ctx, status, side, optional)

List all loans

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**status** | **string**| Loan status | 
**side** | **string**| Lend or borrow | 
**optional** | **ListLoansOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListLoansOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currency** | **optional.String**| Retrieved specified currency related data | 
**currencyPair** | **optional.String**| Currency pair | 
**sortBy** | **optional.String**| Specify which field is used to sort. &#x60;create_time&#x60; or &#x60;rate&#x60; is supported. Default to &#x60;create_time&#x60; | 
**reverseSort** | **optional.Bool**| Whether to sort in descending order. Default to &#x60;true&#x60; | 
**page** | **optional.Int32**| Page number | [default to 1]
**limit** | **optional.Int32**| Maximum number of records returned in one list | [default to 100]

### Example

```golang
package main

import (
    "context"
    "fmt"
    "gateapi"
)

func main() {
    client := gateapi.NewAPIClient(gateapi.NewConfiguration())
    // uncomment the next line if your are testing against testnet
    // client.ChangeBasePath("https://fx-api-testnet.gateio.ws/api/v4")
    ctx := context.WithValue(context.Background(),
                             gateapi.ContextGateAPIV4,
                             gateapi.GateAPIV4{
                                 Key:    "YOUR_API_KEY",
                                 Secret: "YOUR_API_SECRET",
                             }
                            )
    status := "open" // string - Loan status
    side := "lend" // string - Lend or borrow
    
    result, _, err := client.MarginApi.ListLoans(ctx, status, side, nil)
    if err != nil {
        if e, ok := err.(gateapi.GateAPIError); ok {
            fmt.Printf("gate api error: %s\n", e.Error())
        } else {
            fmt.Printf("generic error: %s\n", err.Error())
        }
    } else {
        fmt.Println(result)
    }
}
```


### Return type

[**[]Loan**](Loan.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateLoan

> Loan CreateLoan(ctx, loan)

Lend or borrow

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loan** | [**Loan**](Loan.md)|  | 

### Example

```golang
package main

import (
    "context"
    "fmt"
    "gateapi"
)

func main() {
    client := gateapi.NewAPIClient(gateapi.NewConfiguration())
    // uncomment the next line if your are testing against testnet
    // client.ChangeBasePath("https://fx-api-testnet.gateio.ws/api/v4")
    ctx := context.WithValue(context.Background(),
                             gateapi.ContextGateAPIV4,
                             gateapi.GateAPIV4{
                                 Key:    "YOUR_API_KEY",
                                 Secret: "YOUR_API_SECRET",
                             }
                            )
    loan := gateapi.Loan{} // Loan - 
    
    result, _, err := client.MarginApi.CreateLoan(ctx, loan)
    if err != nil {
        if e, ok := err.(gateapi.GateAPIError); ok {
            fmt.Printf("gate api error: %s\n", e.Error())
        } else {
            fmt.Printf("generic error: %s\n", err.Error())
        }
    } else {
        fmt.Println(result)
    }
}
```


### Return type

[**Loan**](Loan.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## MergeLoans

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
package main

import (
    "context"
    "fmt"
    "gateapi"
)

func main() {
    client := gateapi.NewAPIClient(gateapi.NewConfiguration())
    // uncomment the next line if your are testing against testnet
    // client.ChangeBasePath("https://fx-api-testnet.gateio.ws/api/v4")
    ctx := context.WithValue(context.Background(),
                             gateapi.ContextGateAPIV4,
                             gateapi.GateAPIV4{
                                 Key:    "YOUR_API_KEY",
                                 Secret: "YOUR_API_SECRET",
                             }
                            )
    currency := "BTC" // string - Retrieved specified currency related data
    ids := "123,234,345" // string - Lending loan ID list separated by `,`. Maximum of 20 IDs are allowed in one request
    
    result, _, err := client.MarginApi.MergeLoans(ctx, currency, ids)
    if err != nil {
        if e, ok := err.(gateapi.GateAPIError); ok {
            fmt.Printf("gate api error: %s\n", e.Error())
        } else {
            fmt.Printf("generic error: %s\n", err.Error())
        }
    } else {
        fmt.Println(result)
    }
}
```


### Return type

[**Loan**](Loan.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetLoan

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
package main

import (
    "context"
    "fmt"
    "gateapi"
)

func main() {
    client := gateapi.NewAPIClient(gateapi.NewConfiguration())
    // uncomment the next line if your are testing against testnet
    // client.ChangeBasePath("https://fx-api-testnet.gateio.ws/api/v4")
    ctx := context.WithValue(context.Background(),
                             gateapi.ContextGateAPIV4,
                             gateapi.GateAPIV4{
                                 Key:    "YOUR_API_KEY",
                                 Secret: "YOUR_API_SECRET",
                             }
                            )
    loanId := "12345" // string - Loan ID
    side := "lend" // string - Lend or borrow
    
    result, _, err := client.MarginApi.GetLoan(ctx, loanId, side)
    if err != nil {
        if e, ok := err.(gateapi.GateAPIError); ok {
            fmt.Printf("gate api error: %s\n", e.Error())
        } else {
            fmt.Printf("generic error: %s\n", err.Error())
        }
    } else {
        fmt.Println(result)
    }
}
```


### Return type

[**Loan**](Loan.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CancelLoan

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
package main

import (
    "context"
    "fmt"
    "gateapi"
)

func main() {
    client := gateapi.NewAPIClient(gateapi.NewConfiguration())
    // uncomment the next line if your are testing against testnet
    // client.ChangeBasePath("https://fx-api-testnet.gateio.ws/api/v4")
    ctx := context.WithValue(context.Background(),
                             gateapi.ContextGateAPIV4,
                             gateapi.GateAPIV4{
                                 Key:    "YOUR_API_KEY",
                                 Secret: "YOUR_API_SECRET",
                             }
                            )
    loanId := "12345" // string - Loan ID
    currency := "BTC" // string - Retrieved specified currency related data
    
    result, _, err := client.MarginApi.CancelLoan(ctx, loanId, currency)
    if err != nil {
        if e, ok := err.(gateapi.GateAPIError); ok {
            fmt.Printf("gate api error: %s\n", e.Error())
        } else {
            fmt.Printf("generic error: %s\n", err.Error())
        }
    } else {
        fmt.Println(result)
    }
}
```


### Return type

[**Loan**](Loan.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateLoan

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
package main

import (
    "context"
    "fmt"
    "gateapi"
)

func main() {
    client := gateapi.NewAPIClient(gateapi.NewConfiguration())
    // uncomment the next line if your are testing against testnet
    // client.ChangeBasePath("https://fx-api-testnet.gateio.ws/api/v4")
    ctx := context.WithValue(context.Background(),
                             gateapi.ContextGateAPIV4,
                             gateapi.GateAPIV4{
                                 Key:    "YOUR_API_KEY",
                                 Secret: "YOUR_API_SECRET",
                             }
                            )
    loanId := "12345" // string - Loan ID
    loanPatch := gateapi.LoanPatch{} // LoanPatch - 
    
    result, _, err := client.MarginApi.UpdateLoan(ctx, loanId, loanPatch)
    if err != nil {
        if e, ok := err.(gateapi.GateAPIError); ok {
            fmt.Printf("gate api error: %s\n", e.Error())
        } else {
            fmt.Printf("generic error: %s\n", err.Error())
        }
    } else {
        fmt.Println(result)
    }
}
```


### Return type

[**Loan**](Loan.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListLoanRepayments

> []Repayment ListLoanRepayments(ctx, loanId)

List loan repayment records

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loanId** | **string**| Loan ID | 

### Example

```golang
package main

import (
    "context"
    "fmt"
    "gateapi"
)

func main() {
    client := gateapi.NewAPIClient(gateapi.NewConfiguration())
    // uncomment the next line if your are testing against testnet
    // client.ChangeBasePath("https://fx-api-testnet.gateio.ws/api/v4")
    ctx := context.WithValue(context.Background(),
                             gateapi.ContextGateAPIV4,
                             gateapi.GateAPIV4{
                                 Key:    "YOUR_API_KEY",
                                 Secret: "YOUR_API_SECRET",
                             }
                            )
    loanId := "12345" // string - Loan ID
    
    result, _, err := client.MarginApi.ListLoanRepayments(ctx, loanId)
    if err != nil {
        if e, ok := err.(gateapi.GateAPIError); ok {
            fmt.Printf("gate api error: %s\n", e.Error())
        } else {
            fmt.Printf("generic error: %s\n", err.Error())
        }
    } else {
        fmt.Println(result)
    }
}
```


### Return type

[**[]Repayment**](Repayment.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## RepayLoan

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
package main

import (
    "context"
    "fmt"
    "gateapi"
)

func main() {
    client := gateapi.NewAPIClient(gateapi.NewConfiguration())
    // uncomment the next line if your are testing against testnet
    // client.ChangeBasePath("https://fx-api-testnet.gateio.ws/api/v4")
    ctx := context.WithValue(context.Background(),
                             gateapi.ContextGateAPIV4,
                             gateapi.GateAPIV4{
                                 Key:    "YOUR_API_KEY",
                                 Secret: "YOUR_API_SECRET",
                             }
                            )
    loanId := "12345" // string - Loan ID
    repayRequest := gateapi.RepayRequest{} // RepayRequest - 
    
    result, _, err := client.MarginApi.RepayLoan(ctx, loanId, repayRequest)
    if err != nil {
        if e, ok := err.(gateapi.GateAPIError); ok {
            fmt.Printf("gate api error: %s\n", e.Error())
        } else {
            fmt.Printf("generic error: %s\n", err.Error())
        }
    } else {
        fmt.Println(result)
    }
}
```


### Return type

[**Loan**](Loan.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListLoanRecords

> []LoanRecord ListLoanRecords(ctx, loanId, optional)

List repayment records of specified loan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loanId** | **string**| Loan ID | 
**optional** | **ListLoanRecordsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListLoanRecordsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**status** | **optional.String**| Loan record status | 
**page** | **optional.Int32**| Page number | [default to 1]
**limit** | **optional.Int32**| Maximum number of records returned in one list | [default to 100]

### Example

```golang
package main

import (
    "context"
    "fmt"
    "gateapi"
)

func main() {
    client := gateapi.NewAPIClient(gateapi.NewConfiguration())
    // uncomment the next line if your are testing against testnet
    // client.ChangeBasePath("https://fx-api-testnet.gateio.ws/api/v4")
    ctx := context.WithValue(context.Background(),
                             gateapi.ContextGateAPIV4,
                             gateapi.GateAPIV4{
                                 Key:    "YOUR_API_KEY",
                                 Secret: "YOUR_API_SECRET",
                             }
                            )
    loanId := "12345" // string - Loan ID
    
    result, _, err := client.MarginApi.ListLoanRecords(ctx, loanId, nil)
    if err != nil {
        if e, ok := err.(gateapi.GateAPIError); ok {
            fmt.Printf("gate api error: %s\n", e.Error())
        } else {
            fmt.Printf("generic error: %s\n", err.Error())
        }
    } else {
        fmt.Println(result)
    }
}
```


### Return type

[**[]LoanRecord**](LoanRecord.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetLoanRecord

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
package main

import (
    "context"
    "fmt"
    "gateapi"
)

func main() {
    client := gateapi.NewAPIClient(gateapi.NewConfiguration())
    // uncomment the next line if your are testing against testnet
    // client.ChangeBasePath("https://fx-api-testnet.gateio.ws/api/v4")
    ctx := context.WithValue(context.Background(),
                             gateapi.ContextGateAPIV4,
                             gateapi.GateAPIV4{
                                 Key:    "YOUR_API_KEY",
                                 Secret: "YOUR_API_SECRET",
                             }
                            )
    loanRecordId := "12345" // string - Loan record ID
    loanId := "12345" // string - Loan ID
    
    result, _, err := client.MarginApi.GetLoanRecord(ctx, loanRecordId, loanId)
    if err != nil {
        if e, ok := err.(gateapi.GateAPIError); ok {
            fmt.Printf("gate api error: %s\n", e.Error())
        } else {
            fmt.Printf("generic error: %s\n", err.Error())
        }
    } else {
        fmt.Println(result)
    }
}
```


### Return type

[**LoanRecord**](LoanRecord.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateLoanRecord

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
package main

import (
    "context"
    "fmt"
    "gateapi"
)

func main() {
    client := gateapi.NewAPIClient(gateapi.NewConfiguration())
    // uncomment the next line if your are testing against testnet
    // client.ChangeBasePath("https://fx-api-testnet.gateio.ws/api/v4")
    ctx := context.WithValue(context.Background(),
                             gateapi.ContextGateAPIV4,
                             gateapi.GateAPIV4{
                                 Key:    "YOUR_API_KEY",
                                 Secret: "YOUR_API_SECRET",
                             }
                            )
    loanRecordId := "12345" // string - Loan record ID
    loanPatch := gateapi.LoanPatch{} // LoanPatch - 
    
    result, _, err := client.MarginApi.UpdateLoanRecord(ctx, loanRecordId, loanPatch)
    if err != nil {
        if e, ok := err.(gateapi.GateAPIError); ok {
            fmt.Printf("gate api error: %s\n", e.Error())
        } else {
            fmt.Printf("generic error: %s\n", err.Error())
        }
    } else {
        fmt.Println(result)
    }
}
```


### Return type

[**LoanRecord**](LoanRecord.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
