# \SpotApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelOrder**](SpotApi.md#CancelOrder) | **Delete** /spot/orders/{order_id} | Cancel a single order
[**CancelOrders**](SpotApi.md#CancelOrders) | **Delete** /spot/orders | Cancel all &#x60;open&#x60; orders in specified currency pair
[**CreateOrder**](SpotApi.md#CreateOrder) | **Post** /spot/orders | Create an order
[**GetCurrencyPair**](SpotApi.md#GetCurrencyPair) | **Get** /spot/currency_pairs/{currency_pair} | Get detail of one single order
[**GetOrder**](SpotApi.md#GetOrder) | **Get** /spot/orders/{order_id} | Get a single order
[**ListCandlesticks**](SpotApi.md#ListCandlesticks) | **Get** /spot/candlesticks | Market candlesticks
[**ListCurrencyPairs**](SpotApi.md#ListCurrencyPairs) | **Get** /spot/currency_pairs | List all currency pairs supported
[**ListMyTrades**](SpotApi.md#ListMyTrades) | **Get** /spot/my_trades | List personal trading history
[**ListOrderBook**](SpotApi.md#ListOrderBook) | **Get** /spot/order_book | Retrieve order book
[**ListOrders**](SpotApi.md#ListOrders) | **Get** /spot/orders | List futures orders
[**ListSpotAccounts**](SpotApi.md#ListSpotAccounts) | **Get** /spot/accounts | List spot accounts
[**ListTickers**](SpotApi.md#ListTickers) | **Get** /spot/tickers | Retrieve ticker information
[**ListTrades**](SpotApi.md#ListTrades) | **Get** /spot/trades | Retrieve market trades


# **CancelOrder**
> Order CancelOrder(ctx, orderId, currencyPair)
Cancel a single order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **string**| ID returned on order successfully being created | 
  **currencyPair** | **string**| Currency pair | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.SpotApi
orderId := "12345"; // string - ID returned on order successfully being created
currencyPair := "BTC_USDT"; // string - Currency pair

result, _, err = api.CancelOrder(nil, orderId, currencyPair)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**Order**](Order.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelOrders**
> []Order CancelOrders(ctx, currencyPair, optional)
Cancel all `open` orders in specified currency pair

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **currencyPair** | **string**| Currency pair | 
 **optional** | ***CancelOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CancelOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **side** | **optional.String**| All bids or asks. Both included in not specified | 
 **account** | **optional.String**| Specify account type. Default to all account types being included | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.SpotApi
currencyPair := "BTC_USDT"; // string - Currency pair

result, _, err = api.CancelOrders(nil, currencyPair, nil)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**[]Order**](Order.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrder**
> Order CreateOrder(ctx, order)
Create an order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **order** | [**Order**](Order.md)|  | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.SpotApi
order := new (gateapi.Order); // Order - 

result, _, err = api.CreateOrder(nil, order)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**Order**](Order.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrencyPair**
> CurrencyPair GetCurrencyPair(ctx, currencyPair)
Get detail of one single order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **currencyPair** | **string**| Currency pair | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
api := client.SpotApi
currencyPair := "ETH_BTC"; // string - Currency pair

result, _, err = api.GetCurrencyPair(nil, currencyPair)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**CurrencyPair**](CurrencyPair.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrder**
> Order GetOrder(ctx, orderId, currencyPair)
Get a single order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **string**| ID returned on order successfully being created | 
  **currencyPair** | **string**| Currency pair | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.SpotApi
orderId := "12345"; // string - ID returned on order successfully being created
currencyPair := "BTC_USDT"; // string - Currency pair

result, _, err = api.GetOrder(nil, orderId, currencyPair)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**Order**](Order.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCandlesticks**
> [][]string ListCandlesticks(ctx, currencyPair, optional)
Market candlesticks

Candlestick data will start from (current time - limit * interval), end at current time

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **currencyPair** | **string**| Currency pair | 
 **optional** | ***ListCandlesticksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListCandlesticksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Maximum number of record returned in one list | [default to 100]
 **interval** | **optional.String**| Interval time between data points | [default to 30m]

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
api := client.SpotApi
currencyPair := "BTC_USDT"; // string - Currency pair

result, _, err = api.ListCandlesticks(nil, currencyPair, nil)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**[][]string**](array.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCurrencyPairs**
> []CurrencyPair ListCurrencyPairs(ctx, )
List all currency pairs supported

### Required Parameters
This endpoint does not need any parameter.

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
api := client.SpotApi
result, _, err = api.ListCurrencyPairs(nil)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**[]CurrencyPair**](CurrencyPair.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMyTrades**
> []Trade ListMyTrades(ctx, currencyPair, optional)
List personal trading history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **currencyPair** | **string**| Currency pair | 
 **optional** | ***ListMyTradesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListMyTradesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Maximum number of record returned in one list | [default to 100]
 **page** | **optional.Int32**| Page number | [default to 1]
 **orderId** | **optional.String**| List all trades of specified order | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.SpotApi
currencyPair := "BTC_USDT"; // string - Currency pair

result, _, err = api.ListMyTrades(nil, currencyPair, nil)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**[]Trade**](Trade.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrderBook**
> OrderBook ListOrderBook(ctx, currencyPair, optional)
Retrieve order book

Order book will be sorted by price from high to low on bids; reversed on asks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **currencyPair** | **string**| Currency pair | 
 **optional** | ***ListOrderBookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListOrderBookOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **interval** | **optional.String**| Price precision of order book. 0 means no aggregation is applied | [default to 0]
 **limit** | **optional.Int32**| Maximum number of order depth data in asks or bids | [default to 10]

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
api := client.SpotApi
currencyPair := "BTC_USDT"; // string - Currency pair

result, _, err = api.ListOrderBook(nil, currencyPair, nil)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**OrderBook**](OrderBook.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrders**
> []Order ListOrders(ctx, currencyPair, status, optional)
List futures orders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **currencyPair** | **string**| Currency pair | 
  **status** | **string**| List orders based on status  &#x60;open&#x60; - order is waiting to be filled &#x60;finished&#x60; - order has been filled or cancelled  | 
 **optional** | ***ListOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Page number | [default to 1]
 **limit** | **optional.Int32**| Maximum number of record returned in one list | [default to 100]

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.SpotApi
currencyPair := "BTC_USDT"; // string - Currency pair
status := "open"; // string - List orders based on status  `open` - order is waiting to be filled `finished` - order has been filled or cancelled 

result, _, err = api.ListOrders(nil, currencyPair, status, nil)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**[]Order**](Order.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSpotAccounts**
> []SpotAccount ListSpotAccounts(ctx, optional)
List spot accounts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListSpotAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListSpotAccountsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | **optional.String**| Retrieved specified currency related data | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.SpotApi

result, _, err = api.ListSpotAccounts(nil, nil)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**[]SpotAccount**](SpotAccount.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTickers**
> []Ticker ListTickers(ctx, optional)
Retrieve ticker information

Return only related data if `currency_pair` is specified; otherwise return all of them

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListTickersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListTickersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currencyPair** | **optional.String**| Currency pair | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
api := client.SpotApi

result, _, err = api.ListTickers(nil, nil)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**[]Ticker**](Ticker.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTrades**
> []Trade ListTrades(ctx, currencyPair, optional)
Retrieve market trades

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **currencyPair** | **string**| Currency pair | 
 **optional** | ***ListTradesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListTradesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Maximum number of record returned in one list | [default to 100]
 **lastId** | **optional.String**| Specify list staring point using the last record of &#x60;id&#x60; in previous list-query results | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
api := client.SpotApi
currencyPair := "BTC_USDT"; // string - Currency pair

result, _, err = api.ListTrades(nil, currencyPair, nil)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**[]Trade**](Trade.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

