# AlphaApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAlphaAccounts**](AlphaApi.md#ListAlphaAccounts) | **Get** /alpha/accounts | API for Alpha Accounts
[**ListAlphaAccountBook**](AlphaApi.md#ListAlphaAccountBook) | **Get** /alpha/account_book | Alpha Asset Transaction API
[**QuoteAlphaOrder**](AlphaApi.md#QuoteAlphaOrder) | **Post** /alpha/quote | Alpha Quotation API
[**ListAlphaOrder**](AlphaApi.md#ListAlphaOrder) | **Get** /alpha/orders | Alpha 查询订单列表接口
[**PlaceAlphaOrder**](AlphaApi.md#PlaceAlphaOrder) | **Post** /alpha/orders | Alpha Order Placement API
[**GetAlphaOrder**](AlphaApi.md#GetAlphaOrder) | **Get** /alpha/order | Alpha 查询单个订单接口
[**ListAlphaCurrencies**](AlphaApi.md#ListAlphaCurrencies) | **Get** /alpha/currencies | 查询币种信息
[**ListAlphaTickers**](AlphaApi.md#ListAlphaTickers) | **Get** /alpha/tickers | 查询币种ticker


## ListAlphaAccounts

> []AccountsResponse ListAlphaAccounts(ctx, )

API for Alpha Accounts

Query Position Assets

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
    
    result, _, err := client.AlphaApi.ListAlphaAccounts(ctx)
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

[**[]AccountsResponse**](AccountsResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListAlphaAccountBook

> []AccountBookResponse ListAlphaAccountBook(ctx, from, optional)

Alpha Asset Transaction API

Query Asset Transactions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**from** | **int64**| Start timestamp of the query | 
**optional** | **ListAlphaAccountBookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAlphaAccountBookOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**to** | **optional.Int64**| Time range ending, default to current time | 
**page** | **optional.Int32**| Page number | 
**limit** | **optional.Int32**| The maximum number of items per page is 100 | 

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
    from := 56 // int64 - Start timestamp of the query
    
    result, _, err := client.AlphaApi.ListAlphaAccountBook(ctx, from, nil)
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

[**[]AccountBookResponse**](AccountBookResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## QuoteAlphaOrder

> QuoteResponse QuoteAlphaOrder(ctx, quoteRequest)

Alpha Quotation API

The quote_id returned by the quotation API is valid for one minute.You must place the order within this time window;otherwise, the quote will expire and a new quotation request is required

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteRequest** | [**QuoteRequest**](QuoteRequest.md)|  | 

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
    quoteRequest := gateapi.QuoteRequest{} // QuoteRequest - 
    
    result, _, err := client.AlphaApi.QuoteAlphaOrder(ctx, quoteRequest)
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

[**QuoteResponse**](QuoteResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListAlphaOrder

> []OrderResponse ListAlphaOrder(ctx, currency, side, status, optional)

Alpha 查询订单列表接口

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currency** | **string**| Trading Symbol | 
**side** | **string**| 买单或者卖单 - buy - sell | 
**status** | **int32**| Order Status - &#x60;0&#x60; : All - &#x60;1&#x60; : Processing - &#x60;2&#x60; : Successful - &#x60;3&#x60; : Failed - &#x60;4&#x60; : Canceled - &#x60;5&#x60; : Buy order placed but transfer not completed - &#x60;6&#x60; : Cancelled order with transfer not complete | 
**optional** | **ListAlphaOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAlphaOrderOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**from** | **optional.Int64**| 查询订单的起始时间 | 
**to** | **optional.Int64**| 查询订单的结束时间，不指定则默认为当前时间 | 
**limit** | **optional.Int32**| Maximum response items.  Default: 100, minimum: 1, Maximum: 100 | [default to 100]
**page** | **optional.Int32**| Page number | [default to 1]

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
    currency := "memeboxsst" // string - Trading Symbol
    side := "buy" // string - 买单或者卖单 - buy - sell
    status := 2 // int32 - Order Status - `0` : All - `1` : Processing - `2` : Successful - `3` : Failed - `4` : Canceled - `5` : Buy order placed but transfer not completed - `6` : Cancelled order with transfer not complete
    
    result, _, err := client.AlphaApi.ListAlphaOrder(ctx, currency, side, status, nil)
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

[**[]OrderResponse**](OrderResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PlaceAlphaOrder

> PlaceOrderResponse PlaceAlphaOrder(ctx, placeOrderRequest)

Alpha Order Placement API

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**placeOrderRequest** | [**PlaceOrderRequest**](PlaceOrderRequest.md)|  | 

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
    placeOrderRequest := gateapi.PlaceOrderRequest{} // PlaceOrderRequest - 
    
    result, _, err := client.AlphaApi.PlaceAlphaOrder(ctx, placeOrderRequest)
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

[**PlaceOrderResponse**](PlaceOrderResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetAlphaOrder

> OrderResponse GetAlphaOrder(ctx, orderId)

Alpha 查询单个订单接口

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string**| Order ID | 

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
    orderId := "fdaf12321" // string - Order ID
    
    result, _, err := client.AlphaApi.GetAlphaOrder(ctx, orderId)
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

[**OrderResponse**](OrderResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListAlphaCurrencies

> []Currency2 ListAlphaCurrencies(ctx, optional)

查询币种信息

When the currency parameter is provided, query and return information for the specified currency. When the currency parameter is not provided, return a paginated list of currencies.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListAlphaCurrenciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAlphaCurrenciesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currency** | **optional.String**| 根据币种符号查询币种信息 | 
**limit** | **optional.Int32**| Maximum number of records to be returned in a single list | [default to 100]
**page** | **optional.Int32**| Page number | [default to 1]

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
    
    result, _, err := client.AlphaApi.ListAlphaCurrencies(ctx, nil)
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

[**[]Currency2**](Currency_2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListAlphaTickers

> []Ticker2 ListAlphaTickers(ctx, optional)

查询币种ticker

When the currency parameter is provided, query and return information for the specified ticker, When the currency parameter is not provided, return a paginated list of tickers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListAlphaTickersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAlphaTickersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currency** | **optional.String**| Retrieve data of the specified currency | 
**limit** | **optional.Int32**| Maximum number of records to be returned in a single list | [default to 100]
**page** | **optional.Int32**| Page number | [default to 1]

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
    
    result, _, err := client.AlphaApi.ListAlphaTickers(ctx, nil)
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

[**[]Ticker2**](Ticker_2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
