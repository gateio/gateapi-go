# FlashSwapApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListFlashSwapCurrencyPair**](FlashSwapApi.md#ListFlashSwapCurrencyPair) | **Get** /flash_swap/currency_pairs | List All Supported Currency Pairs In Flash Swap.
[**ListFlashSwapOrders**](FlashSwapApi.md#ListFlashSwapOrders) | **Get** /flash_swap/orders | List all flash swap orders.
[**CreateFlashSwapOrder**](FlashSwapApi.md#CreateFlashSwapOrder) | **Post** /flash_swap/orders | Create a flash swap order.
[**GetFlashSwapOrder**](FlashSwapApi.md#GetFlashSwapOrder) | **Get** /flash_swap/orders/{order_id} | Get a single flash swap order&#39;s detail.
[**PreviewFlashSwapOrder**](FlashSwapApi.md#PreviewFlashSwapOrder) | **Post** /flash_swap/orders/preview | Initiate a flash swap order preview.


## ListFlashSwapCurrencyPair

> []FlashSwapCurrencyPair ListFlashSwapCurrencyPair(ctx, optional)

List All Supported Currency Pairs In Flash Swap.

`BTC_GT` represents selling BTC and buying GT. The limits for each currency may vary across different currency pairs.  It is not necessary that two currencies that can be swapped instantaneously can be exchanged with each other. For example, it is possible to sell BTC and buy GT, but it does not necessarily mean that GT can be sold to buy BTC.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListFlashSwapCurrencyPairOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFlashSwapCurrencyPairOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currency** | **optional.String**| Retrieve data of the specified currency. | 
**page** | **optional.Int32**| Page number. | [default to 1]
**limit** | **optional.Int32**| Maximum response items. Default: 100, minimum: 1, Maximum: 1000. | [default to 1000]

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
    
    result, _, err := client.FlashSwapApi.ListFlashSwapCurrencyPair(ctx, nil)
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

[**[]FlashSwapCurrencyPair**](FlashSwapCurrencyPair.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListFlashSwapOrders

> []FlashSwapOrder ListFlashSwapOrders(ctx, optional)

List all flash swap orders.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListFlashSwapOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFlashSwapOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**status** | **optional.Int32**| Flash swap order status  &#x60;1&#x60; - success &#x60;2&#x60; - failure | 
**sellCurrency** | **optional.String**| Currency to sell which can be retrieved from supported currency list API &#x60;GET /flash_swap/currencies&#x60; | 
**buyCurrency** | **optional.String**| Currency to buy which can be retrieved from supported currency list API &#x60;GET /flash_swap/currencies&#x60; | 
**reverse** | **optional.Bool**| If results are sorted by id in reverse order. Default to &#x60;true&#x60;  - &#x60;true&#x60;: sort by id in descending order(recent first) - ascending order(oldest first) | 
**limit** | **optional.Int32**| Maximum number of records to be returned in a single list. | [default to 100]
**page** | **optional.Int32**| Page number. | [default to 1]

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
    
    result, _, err := client.FlashSwapApi.ListFlashSwapOrders(ctx, nil)
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

[**[]FlashSwapOrder**](FlashSwapOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateFlashSwapOrder

> FlashSwapOrder CreateFlashSwapOrder(ctx, flashSwapOrderRequest)

Create a flash swap order.

Initiate a flash swap preview in advance because order creation requires a preview result

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flashSwapOrderRequest** | [**FlashSwapOrderRequest**](FlashSwapOrderRequest.md)|  | 

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
    flashSwapOrderRequest := gateapi.FlashSwapOrderRequest{} // FlashSwapOrderRequest - 
    
    result, _, err := client.FlashSwapApi.CreateFlashSwapOrder(ctx, flashSwapOrderRequest)
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

[**FlashSwapOrder**](FlashSwapOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetFlashSwapOrder

> FlashSwapOrder GetFlashSwapOrder(ctx, orderId)

Get a single flash swap order's detail.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int32**| Flash swap order ID. | 

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
    orderId := 1 // int32 - Flash swap order ID.
    
    result, _, err := client.FlashSwapApi.GetFlashSwapOrder(ctx, orderId)
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

[**FlashSwapOrder**](FlashSwapOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PreviewFlashSwapOrder

> FlashSwapOrderPreview PreviewFlashSwapOrder(ctx, flashSwapPreviewRequest)

Initiate a flash swap order preview.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flashSwapPreviewRequest** | [**FlashSwapPreviewRequest**](FlashSwapPreviewRequest.md)|  | 

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
    flashSwapPreviewRequest := gateapi.FlashSwapPreviewRequest{} // FlashSwapPreviewRequest - 
    
    result, _, err := client.FlashSwapApi.PreviewFlashSwapOrder(ctx, flashSwapPreviewRequest)
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

[**FlashSwapOrderPreview**](FlashSwapOrderPreview.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
