# SpotApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListCurrencies**](SpotApi.md#ListCurrencies) | **Get** /spot/currencies | List all currencies&#39; details
[**GetCurrency**](SpotApi.md#GetCurrency) | **Get** /spot/currencies/{currency} | Get details of a specific currency
[**ListCurrencyPairs**](SpotApi.md#ListCurrencyPairs) | **Get** /spot/currency_pairs | List all currency pairs supported
[**GetCurrencyPair**](SpotApi.md#GetCurrencyPair) | **Get** /spot/currency_pairs/{currency_pair} | Get details of a specifc currency pair
[**ListTickers**](SpotApi.md#ListTickers) | **Get** /spot/tickers | Retrieve ticker information
[**ListOrderBook**](SpotApi.md#ListOrderBook) | **Get** /spot/order_book | Retrieve order book
[**ListTrades**](SpotApi.md#ListTrades) | **Get** /spot/trades | Retrieve market trades
[**ListCandlesticks**](SpotApi.md#ListCandlesticks) | **Get** /spot/candlesticks | Market candlesticks
[**GetFee**](SpotApi.md#GetFee) | **Get** /spot/fee | Query user trading fee rates
[**GetBatchSpotFee**](SpotApi.md#GetBatchSpotFee) | **Get** /spot/batch_fee | Query a batch of user trading fee rates
[**ListSpotAccounts**](SpotApi.md#ListSpotAccounts) | **Get** /spot/accounts | List spot accounts
[**ListSpotAccountBook**](SpotApi.md#ListSpotAccountBook) | **Get** /spot/account_book | Query account book
[**CreateBatchOrders**](SpotApi.md#CreateBatchOrders) | **Post** /spot/batch_orders | Create a batch of orders
[**ListAllOpenOrders**](SpotApi.md#ListAllOpenOrders) | **Get** /spot/open_orders | List all open orders
[**CreateCrossLiquidateOrder**](SpotApi.md#CreateCrossLiquidateOrder) | **Post** /spot/cross_liquidate_orders | close position when cross-currency is disabled
[**ListOrders**](SpotApi.md#ListOrders) | **Get** /spot/orders | List orders
[**CreateOrder**](SpotApi.md#CreateOrder) | **Post** /spot/orders | Create an order
[**CancelOrders**](SpotApi.md#CancelOrders) | **Delete** /spot/orders | Cancel all &#x60;open&#x60; orders in specified currency pair
[**CancelBatchOrders**](SpotApi.md#CancelBatchOrders) | **Post** /spot/cancel_batch_orders | Cancel a batch of orders with an ID list
[**GetOrder**](SpotApi.md#GetOrder) | **Get** /spot/orders/{order_id} | Get a single order
[**CancelOrder**](SpotApi.md#CancelOrder) | **Delete** /spot/orders/{order_id} | Cancel a single order
[**AmendOrder**](SpotApi.md#AmendOrder) | **Patch** /spot/orders/{order_id} | Amend an order
[**ListMyTrades**](SpotApi.md#ListMyTrades) | **Get** /spot/my_trades | List personal trading history
[**GetSystemTime**](SpotApi.md#GetSystemTime) | **Get** /spot/time | Get server current time
[**CountdownCancelAllSpot**](SpotApi.md#CountdownCancelAllSpot) | **Post** /spot/countdown_cancel_all | Countdown cancel orders
[**AmendBatchOrders**](SpotApi.md#AmendBatchOrders) | **Post** /spot/amend_batch_orders | Batch modification of orders
[**GetSpotInsuranceHistory**](SpotApi.md#GetSpotInsuranceHistory) | **Get** /spot/insurance_history | Query spot insurance fund historical data
[**ListSpotPriceTriggeredOrders**](SpotApi.md#ListSpotPriceTriggeredOrders) | **Get** /spot/price_orders | Retrieve running auto order list
[**CreateSpotPriceTriggeredOrder**](SpotApi.md#CreateSpotPriceTriggeredOrder) | **Post** /spot/price_orders | Create a price-triggered order
[**CancelSpotPriceTriggeredOrderList**](SpotApi.md#CancelSpotPriceTriggeredOrderList) | **Delete** /spot/price_orders | Cancel all open orders
[**GetSpotPriceTriggeredOrder**](SpotApi.md#GetSpotPriceTriggeredOrder) | **Get** /spot/price_orders/{order_id} | Get a price-triggered order
[**CancelSpotPriceTriggeredOrder**](SpotApi.md#CancelSpotPriceTriggeredOrder) | **Delete** /spot/price_orders/{order_id} | cancel a price-triggered order


## ListCurrencies

> []Currency ListCurrencies(ctx, )

List all currencies' details

Currency has two forms:  1. Only currency name, e.g., BTC, USDT 2. `<currency>_<chain>`, e.g., `HT_ETH`  The latter one occurs when one currency has multiple chains. Currency detail contains a `chain` field whatever the form is. To retrieve all chains of one currency, you can use use all the details which has the name of the currency or name starting with `<currency>_`.

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
    
    result, _, err := client.SpotApi.ListCurrencies(ctx)
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

[**[]Currency**](Currency.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetCurrency

> Currency GetCurrency(ctx, currency)

Get details of a specific currency

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currency** | **string**| Currency name | 

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
    currency := "GT" // string - Currency name
    
    result, _, err := client.SpotApi.GetCurrency(ctx, currency)
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

[**Currency**](Currency.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListCurrencyPairs

> []CurrencyPair ListCurrencyPairs(ctx, )

List all currency pairs supported

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
    
    result, _, err := client.SpotApi.ListCurrencyPairs(ctx)
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

[**[]CurrencyPair**](CurrencyPair.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetCurrencyPair

> CurrencyPair GetCurrencyPair(ctx, currencyPair)

Get details of a specifc currency pair

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currencyPair** | **string**| Currency pair | 

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
    currencyPair := "ETH_BTC" // string - Currency pair
    
    result, _, err := client.SpotApi.GetCurrencyPair(ctx, currencyPair)
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

[**CurrencyPair**](CurrencyPair.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListTickers

> []Ticker ListTickers(ctx, optional)

Retrieve ticker information

Return only related data if `currency_pair` is specified; otherwise return all of them

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListTickersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListTickersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currencyPair** | **optional.String**| Currency pair | 
**timezone** | **optional.String**| Timezone | 

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
    
    result, _, err := client.SpotApi.ListTickers(ctx, nil)
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

[**[]Ticker**](Ticker.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListOrderBook

> OrderBook ListOrderBook(ctx, currencyPair, optional)

Retrieve order book

Order book will be sorted by price from high to low on bids; low to high on asks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currencyPair** | **string**| Currency pair | 
**optional** | **ListOrderBookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListOrderBookOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**interval** | **optional.String**| Order depth. 0 means no aggregation is applied. default to 0 | [default to 0]
**limit** | **optional.Int32**| Maximum number of order depth data in asks or bids | [default to 10]
**withId** | **optional.Bool**| Return order book ID | [default to false]

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
    currencyPair := "BTC_USDT" // string - Currency pair
    
    result, _, err := client.SpotApi.ListOrderBook(ctx, currencyPair, nil)
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

[**OrderBook**](OrderBook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListTrades

> []Trade ListTrades(ctx, currencyPair, optional)

Retrieve market trades

Supports specifying `from` and `to` to query by time range or paging query based on `last_id`. The default query is by time range, and the query range is the last 30 days.  The query method based on `last_id` paging is no longer recommended. If `last_id` is specified, the time range query parameter will be ignored.  When using the limit&page paging function to retrieve data, the maximum number of pages is 100,000, that is, (limit page - 1) <= 100000.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currencyPair** | **string**| Currency pair | 
**optional** | **ListTradesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListTradesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**limit** | **optional.Int32**| Maximum number of records to be returned in a single list.  Default: 100, Minimum: 1, Maximum: 1000 | [default to 100]
**lastId** | **optional.String**| Specify list staring point using the &#x60;id&#x60; of last record in previous list-query results | 
**reverse** | **optional.Bool**| Whether the id of records to be retrieved should be less than the last_id specified. Default to false.  When &#x60;last_id&#x60; is specified. Set &#x60;reverse&#x60; to &#x60;true&#x60; to trace back trading history; &#x60;false&#x60; to retrieve latest tradings.  No effect if &#x60;last_id&#x60; is not specified. | [default to false]
**from** | **optional.Int64**| Start timestamp of the query | 
**to** | **optional.Int64**| Time range ending, default to current time | 
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
    currencyPair := "BTC_USDT" // string - Currency pair
    
    result, _, err := client.SpotApi.ListTrades(ctx, currencyPair, nil)
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

[**[]Trade**](Trade.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListCandlesticks

> [][]string ListCandlesticks(ctx, currencyPair, optional)

Market candlesticks

Maximum of 1000 points can be returned in a query. Be sure not to exceed the limit when specifying from, to and interval

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currencyPair** | **string**| Currency pair | 
**optional** | **ListCandlesticksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCandlesticksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**limit** | **optional.Int32**| Maximum recent data points to return. &#x60;limit&#x60; is conflicted with &#x60;from&#x60; and &#x60;to&#x60;. If either &#x60;from&#x60; or &#x60;to&#x60; is specified, request will be rejected. | [default to 100]
**from** | **optional.Int64**| Start time of candlesticks, formatted in Unix timestamp in seconds. Default to&#x60;to - 100 * interval&#x60; if not specified | 
**to** | **optional.Int64**| End time of candlesticks, formatted in Unix timestamp in seconds. Default to current time | 
**interval** | **optional.String**| Interval time between data points. Note that &#x60;30d&#x60; means 1 natual month, not 30 days | [default to 30m]

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
    currencyPair := "BTC_USDT" // string - Currency pair
    
    result, _, err := client.SpotApi.ListCandlesticks(ctx, currencyPair, nil)
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

[**[][]string**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetFee

> SpotFee GetFee(ctx, optional)

Query user trading fee rates

This API is deprecated in favour of new fee retrieving API `/wallet/fee`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **GetFeeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetFeeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currencyPair** | **optional.String**| Specify a currency pair to retrieve precise fee rate  This field is optional. In most cases, the fee rate is identical among all currency pairs | 

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
    
    result, _, err := client.SpotApi.GetFee(ctx, nil)
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

[**SpotFee**](SpotFee.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetBatchSpotFee

> map[string]SpotFee GetBatchSpotFee(ctx, currencyPairs)

Query a batch of user trading fee rates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currencyPairs** | **string**| A request can only query up to 50 currency pairs | 

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
    currencyPairs := "BTC_USDT,ETH_USDT" // string - A request can only query up to 50 currency pairs
    
    result, _, err := client.SpotApi.GetBatchSpotFee(ctx, currencyPairs)
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

[**map[string]SpotFee**](SpotFee.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListSpotAccounts

> []SpotAccount ListSpotAccounts(ctx, optional)

List spot accounts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListSpotAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSpotAccountsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currency** | **optional.String**| Retrieve data of the specified currency | 

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
    
    result, _, err := client.SpotApi.ListSpotAccounts(ctx, nil)
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

[**[]SpotAccount**](SpotAccount.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListSpotAccountBook

> []SpotAccountBook ListSpotAccountBook(ctx, optional)

Query account book

The record query time range is not allowed to exceed 30 days.  When using the limit&page paging function to retrieve data, the maximum number of pages is 100,000, that is, (limit page - 1) <= 100000.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListSpotAccountBookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSpotAccountBookOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currency** | **optional.String**| Retrieve data of the specified currency | 
**from** | **optional.Int64**| Start timestamp of the query | 
**to** | **optional.Int64**| Time range ending, default to current time | 
**page** | **optional.Int32**| Page number | [default to 1]
**limit** | **optional.Int32**| Maximum number of records to be returned in a single list | [default to 100]
**type_** | **optional.String**| Only retrieve changes of the specified type. All types will be returned if not specified. | 

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
    
    result, _, err := client.SpotApi.ListSpotAccountBook(ctx, nil)
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

[**[]SpotAccountBook**](SpotAccountBook.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateBatchOrders

> []BatchOrder CreateBatchOrders(ctx, order, optional)

Create a batch of orders

Batch orders requirements:  1. custom order field `text` is required 2. At most 4 currency pairs, maximum 10 orders each, are allowed in one request 3. No mixture of spot orders and margin orders, i.e. `account` must be identical for all orders 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**order** | [**[]Order**](Order.md)|  | 
**optional** | **CreateBatchOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateBatchOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xGateExptime** | **optional.Int64**| Specify the expiration time (milliseconds); if the GATE receives the request time greater than the expiration time, the request will be rejected | 

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
    order := []gateapi.Order{gateapi.Order{}} // []Order - 
    
    result, _, err := client.SpotApi.CreateBatchOrders(ctx, order, nil)
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

[**[]BatchOrder**](BatchOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListAllOpenOrders

> []OpenOrders ListAllOpenOrders(ctx, optional)

List all open orders

List open orders in all currency pairs.  Note that pagination parameters affect record number in each currency pair's open order list. No pagination is applied to the number of currency pairs returned. All currency pairs with open orders will be returned.  Spot,portfolio and margin orders are returned by default. To list cross margin orders, `account` must be set to `cross_margin`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListAllOpenOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAllOpenOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**page** | **optional.Int32**| Page number | [default to 1]
**limit** | **optional.Int32**| Maximum number of records returned in one page in each currency pair | [default to 100]
**account** | **optional.String**| Specify operation account. Default to spot ,portfolio and margin account if not specified. Set to &#x60;cross_margin&#x60; to operate against margin account.  Portfolio margin account must set to &#x60;cross_margin&#x60; only | 

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
    
    result, _, err := client.SpotApi.ListAllOpenOrders(ctx, nil)
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

[**[]OpenOrders**](OpenOrders.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateCrossLiquidateOrder

> Order CreateCrossLiquidateOrder(ctx, liquidateOrder)

close position when cross-currency is disabled

Currently, only cross-margin accounts are supported to close position when cross currencies are disabled.  Maximum buy quantity = (unpaid principal and interest - currency balance - the amount of the currency in the order book) / 0.998

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**liquidateOrder** | [**LiquidateOrder**](LiquidateOrder.md)|  | 

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
    liquidateOrder := gateapi.LiquidateOrder{} // LiquidateOrder - 
    
    result, _, err := client.SpotApi.CreateCrossLiquidateOrder(ctx, liquidateOrder)
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

[**Order**](Order.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListOrders

> []Order ListOrders(ctx, currencyPair, status, optional)

List orders

Spot, portfolio and margin orders are returned by default. If cross margin orders are needed, `account` must be set to `cross_margin`  When `status` is `open`, i.e., listing open orders, only pagination parameters `page` and `limit` are supported and `limit` cannot be larger than 100. Query by `side` and time range parameters `from` and `to` are not supported.  When `status` is `finished`, i.e., listing finished orders, pagination parameters, time range parameters `from` and `to`, and `side` parameters are all supported. Time range parameters are handled as order finish time.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currencyPair** | **string**| Retrieve results with specified currency pair. It is required for open orders, but optional for finished ones. | 
**status** | **string**| List orders based on status  &#x60;open&#x60; - order is waiting to be filled &#x60;finished&#x60; - order has been filled or cancelled  | 
**optional** | **ListOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**page** | **optional.Int32**| Page number | [default to 1]
**limit** | **optional.Int32**| Maximum number of records to be returned. If &#x60;status&#x60; is &#x60;open&#x60;, maximum of &#x60;limit&#x60; is 100 | [default to 100]
**account** | **optional.String**| Specify operation account. Default to spot ,portfolio and margin account if not specified. Set to &#x60;cross_margin&#x60; to operate against margin account.  Portfolio margin account must set to &#x60;cross_margin&#x60; only | 
**from** | **optional.Int64**| Start timestamp of the query | 
**to** | **optional.Int64**| Time range ending, default to current time | 
**side** | **optional.String**| All bids or asks. Both included if not specified | 

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
    currencyPair := "BTC_USDT" // string - Retrieve results with specified currency pair. It is required for open orders, but optional for finished ones.
    status := "open" // string - List orders based on status  `open` - order is waiting to be filled `finished` - order has been filled or cancelled 
    
    result, _, err := client.SpotApi.ListOrders(ctx, currencyPair, status, nil)
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

[**[]Order**](Order.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateOrder

> Order CreateOrder(ctx, order, optional)

Create an order

Supports spot, margin, leverage, and cross margin orders. Use different accounts through the `account` field, the default is `spot`, that is, use the spot account to place orders, if the user is `unified` account, the default is to use the unified account to place orders  When using margin account trading, that is, when `account` is set to `margin`, you can set `auto_borrow` to `true`, When the account balance is insufficient, the system automatically executes `POST marginuniloans` to borrow the insufficient amount. Whether the assets obtained after placing a margin order are automatically used to return the borrowing order of the isolated margin account depends on the automatic repayment settings of the user's isolated margin account, The automatic repayment settings of this account can be queried and set through `marginauto_repay`.  When using cross margin account trading, that is, when `account` is set to `cross_margin`, `auto_borrow` To realize the automatic borrowing of the insufficient part, but unlike the isolated margin account, whether the entrustment of the cross margin account is automatically repaid depends on the time when the order is placed `auto_repay` setting, this setting only takes effect for the current order, that is, only the assets obtained after the order is completed will be used to repay the borrowing order of the cross margin account. Placing orders on cross margin accounts currently supports enabling `auto_borrow` and `auto_repay` at the same time.  Automatic repayment will be triggered when the order ends, that is, the `status` is `cancelled` or `closed`.  Delegation status  The order status in the pending order is `open` and remains `open` until all the quantities are filled. If all are filled, the order ends and the status becomes `closed`. If the order is canceled before all transactions are completed, the status will change to `cancelled` regardless of whether there is partial execution or not.  Iceberg Commission  `iceberg` is used to set the number of iceberg orders displayed. If you need to hide it completely, set it to `-1`. Note that when hiding partial transactions, follow the instructions The taker's handling fee is charged.  Restrict users to self-transact  Set `stp_act` to decide to use a strategy that limits users' self-transactions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**order** | [**Order**](Order.md)|  | 
**optional** | **CreateOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateOrderOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xGateExptime** | **optional.Int64**| Specify the expiration time (milliseconds); if the GATE receives the request time greater than the expiration time, the request will be rejected | 

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
    order := gateapi.Order{} // Order - 
    
    result, _, err := client.SpotApi.CreateOrder(ctx, order, nil)
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

[**Order**](Order.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CancelOrders

> []OrderCancel CancelOrders(ctx, optional)

Cancel all `open` orders in specified currency pair

If `account` is not set, all open orders, including spot, portfolio, margin and cross margin ones, will be cancelled. If `currency_pair` is not specified, all pending orders for trading pairs will be cancelled. You can set `account` to cancel only orders within the specified account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **CancelOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CancelOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currencyPair** | **optional.String**| Currency pair | 
**side** | **optional.String**| All bids or asks. Both included if not specified | 
**account** | **optional.String**| Specify account type:  - Classic account: Includes all if not specified - Unified account: Specify &#x60;unified&#x60; - Unified account (legacy): Can only specify &#x60;cross_margin&#x60; | 
**actionMode** | **optional.String**| Processing Mode  When placing an order, different fields are returned based on the action_mode  - ACK: Asynchronous mode, returns only key order fields - RESULT: No clearing information - FULL: Full mode (default) | 
**xGateExptime** | **optional.Int64**| Specify the expiration time (milliseconds); if the GATE receives the request time greater than the expiration time, the request will be rejected | 

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
    
    result, _, err := client.SpotApi.CancelOrders(ctx, nil)
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

[**[]OrderCancel**](OrderCancel.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CancelBatchOrders

> []CancelOrderResult CancelBatchOrders(ctx, cancelBatchOrder, optional)

Cancel a batch of orders with an ID list

Multiple currency pairs can be specified, but maximum 20 orders are allowed per request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cancelBatchOrder** | [**[]CancelBatchOrder**](CancelBatchOrder.md)|  | 
**optional** | **CancelBatchOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CancelBatchOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xGateExptime** | **optional.Int64**| Specify the expiration time (milliseconds); if the GATE receives the request time greater than the expiration time, the request will be rejected | 

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
    cancelBatchOrder := []gateapi.CancelBatchOrder{gateapi.CancelBatchOrder{}} // []CancelBatchOrder - 
    
    result, _, err := client.SpotApi.CancelBatchOrders(ctx, cancelBatchOrder, nil)
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

[**[]CancelOrderResult**](CancelOrderResult.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetOrder

> Order GetOrder(ctx, orderId, currencyPair, optional)

Get a single order

Spot, portfolio and margin orders are queried by default. If cross margin orders are needed or portfolio margin account are used, account must be set to cross_margin.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string**| Order ID returned, or user custom ID(i.e., &#x60;text&#x60; field). Operations based on custom ID can only be checked when the order is in orderbook.  When the order is finished, it can be checked within 1 hour after the end of the order.  After that, only order ID is accepted. | 
**currencyPair** | **string**| Specify the transaction pair to query. If you are querying pending order records, this field is required. If you are querying traded records, this field can be left blank. | 
**optional** | **GetOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetOrderOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**account** | **optional.String**| Specify operation account. Default to spot ,portfolio and margin account if not specified. Set to &#x60;cross_margin&#x60; to operate against margin account.  Portfolio margin account must set to &#x60;cross_margin&#x60; only | 

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
    orderId := "12345" // string - Order ID returned, or user custom ID(i.e., `text` field). Operations based on custom ID can only be checked when the order is in orderbook.  When the order is finished, it can be checked within 1 hour after the end of the order.  After that, only order ID is accepted.
    currencyPair := "BTC_USDT" // string - Specify the transaction pair to query. If you are querying pending order records, this field is required. If you are querying traded records, this field can be left blank.
    
    result, _, err := client.SpotApi.GetOrder(ctx, orderId, currencyPair, nil)
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

[**Order**](Order.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CancelOrder

> Order CancelOrder(ctx, orderId, currencyPair, optional)

Cancel a single order

Spot,portfolio and margin orders are cancelled by default. If trying to cancel cross margin orders or portfolio margin account are used, account must be set to cross_margin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string**| Order ID returned, or user custom ID(i.e., &#x60;text&#x60; field). Operations based on custom ID can only be checked when the order is in orderbook.  When the order is finished, it can be checked within 1 hour after the end of the order.  After that, only order ID is accepted. | 
**currencyPair** | **string**| Currency pair | 
**optional** | **CancelOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CancelOrderOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**account** | **optional.String**| Specify operation account. Default to spot ,portfolio and margin account if not specified. Set to &#x60;cross_margin&#x60; to operate against margin account.  Portfolio margin account must set to &#x60;cross_margin&#x60; only | 
**actionMode** | **optional.String**| Processing Mode  When placing an order, different fields are returned based on the action_mode  - ACK: Asynchronous mode, returns only key order fields - RESULT: No clearing information - FULL: Full mode (default) | 
**xGateExptime** | **optional.Int64**| Specify the expiration time (milliseconds); if the GATE receives the request time greater than the expiration time, the request will be rejected | 

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
    orderId := "12345" // string - Order ID returned, or user custom ID(i.e., `text` field). Operations based on custom ID can only be checked when the order is in orderbook.  When the order is finished, it can be checked within 1 hour after the end of the order.  After that, only order ID is accepted.
    currencyPair := "BTC_USDT" // string - Currency pair
    
    result, _, err := client.SpotApi.CancelOrder(ctx, orderId, currencyPair, nil)
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

[**Order**](Order.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## AmendOrder

> Order AmendOrder(ctx, orderId, orderPatch, optional)

Amend an order

By default, orders for spot, margin and isolated margin accounts are modified. If you need to modify orders for cross margin accounts, you must specify `account` as `cross_margin` and unify the account.`account` can only be specified as `cross_margin`.  Currently both the request body and query support currency_pair and account parameters, but the request body has a higher priority  currency_pair must be filled in either the request body or query  Currently only supports modifying price or quantity (choose one of the two)  About speed limit: Modify orders and create orders to share speed limit rules  About matching priority: only modifying the quantity to make it smaller will not affect the matching priority. If you modify the price or modify the quantity to become larger, the priority will be adjusted to the end of the new price  Note: The modified quantity is smaller than the completed quantity, which will trigger the order cancellation operation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string**| Order ID returned, or user custom ID(i.e., &#x60;text&#x60; field). Operations based on custom ID can only be checked when the order is in orderbook.  When the order is finished, it can be checked within 1 hour after the end of the order.  After that, only order ID is accepted. | 
**orderPatch** | [**OrderPatch**](OrderPatch.md)|  | 
**optional** | **AmendOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AmendOrderOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currencyPair** | **optional.String**| Currency pair | 
**account** | **optional.String**| Specify operation account. Default to spot ,portfolio and margin account if not specified. Set to &#x60;cross_margin&#x60; to operate against margin account.  Portfolio margin account must set to &#x60;cross_margin&#x60; only | 
**xGateExptime** | **optional.Int64**| Specify the expiration time (milliseconds); if the GATE receives the request time greater than the expiration time, the request will be rejected | 

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
    orderId := "12345" // string - Order ID returned, or user custom ID(i.e., `text` field). Operations based on custom ID can only be checked when the order is in orderbook.  When the order is finished, it can be checked within 1 hour after the end of the order.  After that, only order ID is accepted.
    orderPatch := gateapi.OrderPatch{} // OrderPatch - 
    
    result, _, err := client.SpotApi.AmendOrder(ctx, orderId, orderPatch, nil)
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

[**Order**](Order.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListMyTrades

> []Trade ListMyTrades(ctx, optional)

List personal trading history

Spot,portfolio and margin trades are queried by default. If cross margin trades are needed, `account` must be set to `cross_margin`  You can also set `from` and(or) `to` to query by time range. If you don't specify `from` and/or `to` parameters, only the last 7 days of data will be retured. The range of `from` and `to` is not alloed to exceed 30 days.  Time range parameters are handled as order finish time. When using the limit&page paging function to retrieve data, the maximum number of pages is 100,000, that is, (limit * page - 1) <= 100000.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListMyTradesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListMyTradesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currencyPair** | **optional.String**| Retrieve results with specified currency pair | 
**limit** | **optional.Int32**| Maximum number of records to be returned in a single list.  Default: 100, Minimum: 1, Maximum: 1000 | [default to 100]
**page** | **optional.Int32**| Page number | [default to 1]
**orderId** | **optional.String**| Filter trades with specified order ID. &#x60;currency_pair&#x60; is also required if this field is present | 
**account** | **optional.String**| Specify operation account. Default to spot ,portfolio and margin account if not specified. Set to &#x60;cross_margin&#x60; to operate against margin account.  Portfolio margin account must set to &#x60;cross_margin&#x60; only | 
**from** | **optional.Int64**| Start timestamp of the query | 
**to** | **optional.Int64**| Time range ending, default to current time | 

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
    
    result, _, err := client.SpotApi.ListMyTrades(ctx, nil)
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

[**[]Trade**](Trade.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetSystemTime

> SystemTime GetSystemTime(ctx, )

Get server current time

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
    
    result, _, err := client.SpotApi.GetSystemTime(ctx)
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

[**SystemTime**](SystemTime.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CountdownCancelAllSpot

> TriggerTime CountdownCancelAllSpot(ctx, countdownCancelAllSpotTask)

Countdown cancel orders

When the timeout set by the user is reached, if there is no cancel or set a new countdown, the related pending orders will be automatically cancelled.  This endpoint can be called repeatedly to set a new countdown or cancel the countdown. For example, call this endpoint at 30s intervals, each countdown`timeout` is set to 30s. If this endpoint is not called again within 30 seconds, all pending orders on the specified `market` will be automatically cancelled, if no `market` is specified, all market pending orders will be cancelled. If the `timeout` is set to 0 within 30 seconds, the countdown timer will expire and the cacnel function will be cancelled.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countdownCancelAllSpotTask** | [**CountdownCancelAllSpotTask**](CountdownCancelAllSpotTask.md)|  | 

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
    countdownCancelAllSpotTask := gateapi.CountdownCancelAllSpotTask{} // CountdownCancelAllSpotTask - 
    
    result, _, err := client.SpotApi.CountdownCancelAllSpot(ctx, countdownCancelAllSpotTask)
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

[**TriggerTime**](triggerTime.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## AmendBatchOrders

> []BatchOrder AmendBatchOrders(ctx, batchAmendItem, optional)

Batch modification of orders

Default modification of orders for spot, portfolio, and margin accounts. To modify orders for a cross margin account, the `account` parameter must be specified as `cross_margin`.  For portfolio margin accounts, the `account` parameter can only be specified as `cross_margin`. Currently, only modifications to price or quantity (choose one) are supported. When modifying unfinished orders, a maximum of 5 orders can be batch-modified in one request. The request parameters should be passed in an array format. During batch modification, if one order modification fails, the modification process will continue with the next order. After execution, the response will include corresponding failure information for the failed orders. The sequence of calling for batch order modification should be consistent with the order in the order list. The response content order for batch order modification will also be consistent with the order in the order list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchAmendItem** | [**[]BatchAmendItem**](BatchAmendItem.md)|  | 
**optional** | **AmendBatchOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AmendBatchOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xGateExptime** | **optional.Int64**| Specify the expiration time (milliseconds); if the GATE receives the request time greater than the expiration time, the request will be rejected | 

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
    batchAmendItem := []gateapi.BatchAmendItem{gateapi.BatchAmendItem{}} // []BatchAmendItem - 
    
    result, _, err := client.SpotApi.AmendBatchOrders(ctx, batchAmendItem, nil)
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

[**[]BatchOrder**](BatchOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetSpotInsuranceHistory

> []SpotInsuranceHistory GetSpotInsuranceHistory(ctx, business, currency, from, to, optional)

Query spot insurance fund historical data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**business** | **string**| Leverage business, margin - position by position; unified - unified account | 
**currency** | **string**| Currency | 
**from** | **int64**| Start timestamp, seconds | 
**to** | **int64**| End timestamp, in seconds | 
**optional** | **GetSpotInsuranceHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSpotInsuranceHistoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**page** | **optional.Int32**| Page number | [default to 1]
**limit** | **optional.Int32**| The maximum number of items returned in the list, the default value is 30 | [default to 30]

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
    business := "margin" // string - Leverage business, margin - position by position; unified - unified account
    currency := "BTC" // string - Currency
    from := 1547706332 // int64 - Start timestamp, seconds
    to := 1547706332 // int64 - End timestamp, in seconds
    
    result, _, err := client.SpotApi.GetSpotInsuranceHistory(ctx, business, currency, from, to, nil)
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

[**[]SpotInsuranceHistory**](SpotInsuranceHistory.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListSpotPriceTriggeredOrders

> []SpotPriceTriggeredOrder ListSpotPriceTriggeredOrders(ctx, status, optional)

Retrieve running auto order list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**status** | **string**| Only list the orders with this status | 
**optional** | **ListSpotPriceTriggeredOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSpotPriceTriggeredOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**market** | **optional.String**| Currency pair | 
**account** | **optional.String**| Trading account type.  Portfolio margin account must set to &#x60;cross_margin&#x60; | 
**limit** | **optional.Int32**| Maximum number of records to be returned in a single list | [default to 100]
**offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]

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
    status := "status_example" // string - Only list the orders with this status
    
    result, _, err := client.SpotApi.ListSpotPriceTriggeredOrders(ctx, status, nil)
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

[**[]SpotPriceTriggeredOrder**](SpotPriceTriggeredOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateSpotPriceTriggeredOrder

> TriggerOrderResponse CreateSpotPriceTriggeredOrder(ctx, spotPriceTriggeredOrder)

Create a price-triggered order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spotPriceTriggeredOrder** | [**SpotPriceTriggeredOrder**](SpotPriceTriggeredOrder.md)|  | 

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
    spotPriceTriggeredOrder := gateapi.SpotPriceTriggeredOrder{} // SpotPriceTriggeredOrder - 
    
    result, _, err := client.SpotApi.CreateSpotPriceTriggeredOrder(ctx, spotPriceTriggeredOrder)
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

[**TriggerOrderResponse**](TriggerOrderResponse.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CancelSpotPriceTriggeredOrderList

> []SpotPriceTriggeredOrder CancelSpotPriceTriggeredOrderList(ctx, optional)

Cancel all open orders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **CancelSpotPriceTriggeredOrderListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CancelSpotPriceTriggeredOrderListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**market** | **optional.String**| Currency pair | 
**account** | **optional.String**| Trading account type.  Portfolio margin account must set to &#x60;cross_margin&#x60; | 

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
    
    result, _, err := client.SpotApi.CancelSpotPriceTriggeredOrderList(ctx, nil)
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

[**[]SpotPriceTriggeredOrder**](SpotPriceTriggeredOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetSpotPriceTriggeredOrder

> SpotPriceTriggeredOrder GetSpotPriceTriggeredOrder(ctx, orderId)

Get a price-triggered order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string**| Retrieve the data of the order with the specified ID | 

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
    orderId := "orderId_example" // string - Retrieve the data of the order with the specified ID
    
    result, _, err := client.SpotApi.GetSpotPriceTriggeredOrder(ctx, orderId)
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

[**SpotPriceTriggeredOrder**](SpotPriceTriggeredOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CancelSpotPriceTriggeredOrder

> SpotPriceTriggeredOrder CancelSpotPriceTriggeredOrder(ctx, orderId)

cancel a price-triggered order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string**| Retrieve the data of the order with the specified ID | 

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
    orderId := "orderId_example" // string - Retrieve the data of the order with the specified ID
    
    result, _, err := client.SpotApi.CancelSpotPriceTriggeredOrder(ctx, orderId)
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

[**SpotPriceTriggeredOrder**](SpotPriceTriggeredOrder.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
