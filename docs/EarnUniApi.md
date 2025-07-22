# EarnUniApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListUniCurrencies**](EarnUniApi.md#ListUniCurrencies) | **Get** /earn/uni/currencies | List currencies for lending.
[**GetUniCurrency**](EarnUniApi.md#GetUniCurrency) | **Get** /earn/uni/currencies/{currency} | Get currency detail for lending.
[**ListUserUniLends**](EarnUniApi.md#ListUserUniLends) | **Get** /earn/uni/lends | List user&#39;s lending orders.
[**CreateUniLend**](EarnUniApi.md#CreateUniLend) | **Post** /earn/uni/lends | Lend or redeem.
[**ChangeUniLend**](EarnUniApi.md#ChangeUniLend) | **Patch** /earn/uni/lends | Amend lending order.
[**ListUniLendRecords**](EarnUniApi.md#ListUniLendRecords) | **Get** /earn/uni/lend_records | List records of lending.
[**GetUniInterest**](EarnUniApi.md#GetUniInterest) | **Get** /earn/uni/interests/{currency} | Get the user&#39;s total interest income of specified currency.
[**ListUniInterestRecords**](EarnUniApi.md#ListUniInterestRecords) | **Get** /earn/uni/interest_records | List interest records.
[**GetUniInterestStatus**](EarnUniApi.md#GetUniInterestStatus) | **Get** /earn/uni/interest_status/{currency} | query currency interest compounding status.
[**ListUniChart**](EarnUniApi.md#ListUniChart) | **Get** /earn/uni/chart | UniLoan currency annualized trend chart.
[**ListUniRate**](EarnUniApi.md#ListUniRate) | **Get** /earn/uni/rate | Currency estimate annualized interest rate.


## ListUniCurrencies

> []UniCurrency ListUniCurrencies(ctx, )

List currencies for lending.

### Required Parameters


### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
)

func main() {
    client := gateapi.NewAPIClient(gateapi.NewConfiguration())
    // uncomment the next line if your are testing against testnet
    // client.ChangeBasePath("https://fx-api-testnet.gateio.ws/api/v4")
    ctx := context.Background()
    
    result, _, err := client.EarnUniApi.ListUniCurrencies(ctx)
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

[**[]UniCurrency**](UniCurrency.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUniCurrency

> UniCurrency GetUniCurrency(ctx, currency)

Get currency detail for lending.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currency** | **string**| Currency. | 

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
)

func main() {
    client := gateapi.NewAPIClient(gateapi.NewConfiguration())
    // uncomment the next line if your are testing against testnet
    // client.ChangeBasePath("https://fx-api-testnet.gateio.ws/api/v4")
    ctx := context.Background()
    currency := "btc" // string - Currency.
    
    result, _, err := client.EarnUniApi.GetUniCurrency(ctx, currency)
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

[**UniCurrency**](UniCurrency.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListUserUniLends

> []UniLend ListUserUniLends(ctx, optional)

List user's lending orders.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListUserUniLendsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUserUniLendsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currency** | **optional.String**| Retrieve data of the specified currency. | 
**page** | **optional.Int32**| Page number. | [default to 1]
**limit** | **optional.Int32**| Maximum response items. Default: 100, minimum: 1, Maximum: 100. | [default to 100]

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
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
    
    result, _, err := client.EarnUniApi.ListUserUniLends(ctx, nil)
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

[**[]UniLend**](UniLend.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateUniLend

> CreateUniLend(ctx, createUniLend)

Lend or redeem.

Lending: When lending, a minimum lending rate must be set. After successful lending is determined on an hourly basis, earnings will be calculated based on the determined rate.  Earnings for each hour will be settled at the top of the hour. If lending fails due to an excessively high interest rate, no interest will be earned for that hour.  If funds are redeemed before the hourly for that hour.  Priority: Under the same interest rate, wealth management products created or modified earlier will be prioritized for lending.  Redemption: For funds that failed to be lent, redemption will be credited immediately. For funds successfully lent, they are entitled to the earnings for that hour, and redemption will be credited in the next hourly interval.  Note: The two minutes before and after the hourly mark are the settlement period, during which lending and redemption are prohibited. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createUniLend** | [**CreateUniLend**](CreateUniLend.md)|  | 

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
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
    createUniLend := gateapi.CreateUniLend{} // CreateUniLend - 
    
    result, _, err := client.EarnUniApi.CreateUniLend(ctx, createUniLend)
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

 (empty response body)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ChangeUniLend

> ChangeUniLend(ctx, patchUniLend)

Amend lending order.

Currently only supports amending the minimum interest rate (hour).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patchUniLend** | [**PatchUniLend**](PatchUniLend.md)|  | 

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
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
    patchUniLend := gateapi.PatchUniLend{} // PatchUniLend - 
    
    result, _, err := client.EarnUniApi.ChangeUniLend(ctx, patchUniLend)
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

 (empty response body)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListUniLendRecords

> []UniLendRecord ListUniLendRecords(ctx, optional)

List records of lending.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListUniLendRecordsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUniLendRecordsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currency** | **optional.String**| Retrieve data of the specified currency. | 
**page** | **optional.Int32**| Page number. | [default to 1]
**limit** | **optional.Int32**| Maximum response items. Default: 100, minimum: 1, Maximum: 100. | [default to 100]
**from** | **optional.Int64**| Start timestamp  Specify start time, time format is Unix timestamp. If not specified, it defaults to (the data start time of the time range actually returned by to and limit) | 
**to** | **optional.Int64**| Termination Timestamp  Specify the end time. If not specified, it defaults to the current time, and the time format is a Unix timestamp | 
**type_** | **optional.String**| type: lend - lend, redeem - redeem. | 

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
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
    
    result, _, err := client.EarnUniApi.ListUniLendRecords(ctx, nil)
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

[**[]UniLendRecord**](UniLendRecord.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUniInterest

> UniLendInterest GetUniInterest(ctx, currency)

Get the user's total interest income of specified currency.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currency** | **string**| Currency. | 

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
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
    currency := "btc" // string - Currency.
    
    result, _, err := client.EarnUniApi.GetUniInterest(ctx, currency)
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

[**UniLendInterest**](UniLendInterest.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListUniInterestRecords

> []UniInterestRecord ListUniInterestRecords(ctx, optional)

List interest records.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListUniInterestRecordsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUniInterestRecordsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currency** | **optional.String**| Retrieve data of the specified currency. | 
**page** | **optional.Int32**| Page number. | [default to 1]
**limit** | **optional.Int32**| Maximum response items. Default: 100, minimum: 1, Maximum: 100. | [default to 100]
**from** | **optional.Int64**| Start timestamp  Specify start time, time format is Unix timestamp. If not specified, it defaults to (the data start time of the time range actually returned by to and limit) | 
**to** | **optional.Int64**| Termination Timestamp  Specify the end time. If not specified, it defaults to the current time, and the time format is a Unix timestamp | 

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
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
    
    result, _, err := client.EarnUniApi.ListUniInterestRecords(ctx, nil)
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

[**[]UniInterestRecord**](UniInterestRecord.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUniInterestStatus

> UniCurrencyInterest GetUniInterestStatus(ctx, currency)

query currency interest compounding status.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currency** | **string**| Currency. | 

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
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
    currency := "btc" // string - Currency.
    
    result, _, err := client.EarnUniApi.GetUniInterestStatus(ctx, currency)
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

[**UniCurrencyInterest**](UniCurrencyInterest.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListUniChart

> []InlineResponse200 ListUniChart(ctx, from, to, asset)

UniLoan currency annualized trend chart.

Project-Id-Version: GateApiTools 1.0.0 Report-Msgid-Bugs-To: EMAIL@ADDRESS POT-Creation-Date: 2025-07-17 21:35+0800 PO-Revision-Date: 2019-01-02 17:30+0800 Last-Translator: FULL NAME <EMAIL@ADDRESS> Language: en Language-Team: en <L@li.org> Plural-Forms: nplurals=2; plural=(n !=1) MIME-Version: 1.0 Content-Type: text/plain; charset=utf-8 Content-Transfer-Encoding: 8bit Generated-By: Babel 2.8.0 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**from** | **int64**| Start timestamp, unit s, maximum span of 30 days. | 
**to** | **int64**| End timestamp, unit s, maximum span of 30 days. | 
**asset** | **string**| Currency name. | 

### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
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
    from := 1719763200 // int64 - Start timestamp, unit s, maximum span of 30 days.
    to := 1722441600 // int64 - End timestamp, unit s, maximum span of 30 days.
    asset := "BTC" // string - Currency name.
    
    result, _, err := client.EarnUniApi.ListUniChart(ctx, from, to, asset)
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

[**[]InlineResponse200**](inline_response_200.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListUniRate

> []InlineResponse2001 ListUniRate(ctx, )

Currency estimate annualized interest rate.

Project-Id-Version: GateApiTools 1.0.0 Report-Msgid-Bugs-To: EMAIL@ADDRESS POT-Creation-Date: 2025-07-17 21:35+0800 PO-Revision-Date: 2019-01-02 17:30+0800 Last-Translator: FULL NAME <EMAIL@ADDRESS> Language: en Language-Team: en <L@li.org> Plural-Forms: nplurals=2; plural=(n !=1) MIME-Version: 1.0 Content-Type: text/plain; charset=utf-8 Content-Transfer-Encoding: 8bit Generated-By: Babel 2.8.0 

### Required Parameters


### Example

```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v6"
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
    
    result, _, err := client.EarnUniApi.ListUniRate(ctx)
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

[**[]InlineResponse2001**](inline_response_200_1.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
