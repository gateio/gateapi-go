# \FuturesApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelFuturesOrder**](FuturesApi.md#CancelFuturesOrder) | **Delete** /futures/{settle}/orders/{order_id} | Cancel a single order
[**CancelFuturesOrders**](FuturesApi.md#CancelFuturesOrders) | **Delete** /futures/{settle}/orders | Cancel all &#x60;open&#x60; orders matched
[**CancelPriceTriggeredOrder**](FuturesApi.md#CancelPriceTriggeredOrder) | **Delete** /futures/{settle}/price_orders/{order_id} | Cancel a single order
[**CancelPriceTriggeredOrderList**](FuturesApi.md#CancelPriceTriggeredOrderList) | **Delete** /futures/{settle}/price_orders | Cancel all open orders
[**CreateFuturesOrder**](FuturesApi.md#CreateFuturesOrder) | **Post** /futures/{settle}/orders | Create a futures order
[**CreatePriceTriggeredOrder**](FuturesApi.md#CreatePriceTriggeredOrder) | **Post** /futures/{settle}/price_orders | Create a price-triggered order
[**GetFuturesContract**](FuturesApi.md#GetFuturesContract) | **Get** /futures/{settle}/contracts/{contract} | Get a single contract
[**GetFuturesOrder**](FuturesApi.md#GetFuturesOrder) | **Get** /futures/{settle}/orders/{order_id} | Get a single order
[**GetMyTrades**](FuturesApi.md#GetMyTrades) | **Get** /futures/{settle}/my_trades | List personal trading history
[**GetPosition**](FuturesApi.md#GetPosition) | **Get** /futures/{settle}/positions/{contract} | Get single position
[**GetPriceTriggeredOrder**](FuturesApi.md#GetPriceTriggeredOrder) | **Get** /futures/{settle}/price_orders/{order_id} | Get a single order
[**ListFuturesAccountBook**](FuturesApi.md#ListFuturesAccountBook) | **Get** /futures/{settle}/account_book | Query account book
[**ListFuturesAccounts**](FuturesApi.md#ListFuturesAccounts) | **Get** /futures/{settle}/accounts | Query futures account
[**ListFuturesCandlesticks**](FuturesApi.md#ListFuturesCandlesticks) | **Get** /futures/{settle}/candlesticks | Get futures candlesticks
[**ListFuturesContracts**](FuturesApi.md#ListFuturesContracts) | **Get** /futures/{settle}/contracts | List all futures contracts
[**ListFuturesFundingRateHistory**](FuturesApi.md#ListFuturesFundingRateHistory) | **Get** /futures/{settle}/funding_rate | Funding rate history
[**ListFuturesInsuranceLedger**](FuturesApi.md#ListFuturesInsuranceLedger) | **Get** /futures/{settle}/insurance | Futures insurance balance history
[**ListFuturesOrderBook**](FuturesApi.md#ListFuturesOrderBook) | **Get** /futures/{settle}/order_book | Futures order book
[**ListFuturesOrders**](FuturesApi.md#ListFuturesOrders) | **Get** /futures/{settle}/orders | List futures orders
[**ListFuturesTickers**](FuturesApi.md#ListFuturesTickers) | **Get** /futures/{settle}/tickers | List futures tickers
[**ListFuturesTrades**](FuturesApi.md#ListFuturesTrades) | **Get** /futures/{settle}/trades | Futures trading history
[**ListLiquidates**](FuturesApi.md#ListLiquidates) | **Get** /futures/{settle}/liquidates | List liquidation history
[**ListPositionClose**](FuturesApi.md#ListPositionClose) | **Get** /futures/{settle}/position_close | List position close history
[**ListPositions**](FuturesApi.md#ListPositions) | **Get** /futures/{settle}/positions | List all positions of a user
[**ListPriceTriggeredOrders**](FuturesApi.md#ListPriceTriggeredOrders) | **Get** /futures/{settle}/price_orders | List all auto orders
[**UpdatePositionLeverage**](FuturesApi.md#UpdatePositionLeverage) | **Post** /futures/{settle}/positions/{contract}/leverage | Update position leverage
[**UpdatePositionMargin**](FuturesApi.md#UpdatePositionMargin) | **Post** /futures/{settle}/positions/{contract}/margin | Update position margin
[**UpdatePositionRiskLimit**](FuturesApi.md#UpdatePositionRiskLimit) | **Post** /futures/{settle}/positions/{contract}/risk_limit | Update position risk limit


# **CancelFuturesOrder**
> FuturesOrder CancelFuturesOrder(ctx, settle, orderId)
Cancel a single order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | [default to btc]
  **orderId** | **string**| ID returned on order successfully being created | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.FuturesApi
settle := "btc"; // string - Settle currency
orderId := "12345"; // string - ID returned on order successfully being created

result, _, err := api.CancelFuturesOrder(nil, settle, orderId)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**FuturesOrder**](FuturesOrder.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelFuturesOrders**
> []FuturesOrder CancelFuturesOrders(ctx, settle, contract, optional)
Cancel all `open` orders matched

Zero-fill order cannot be retrieved 60 seconds after cancellation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | [default to btc]
  **contract** | **string**| Futures contract | 
 **optional** | ***CancelFuturesOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CancelFuturesOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **side** | **optional.String**| All bids or asks. Both included in not specified | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.FuturesApi
settle := "btc"; // string - Settle currency
contract := "BTC_USD"; // string - Futures contract

result, _, err := api.CancelFuturesOrders(nil, settle, contract, nil)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**[]FuturesOrder**](FuturesOrder.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelPriceTriggeredOrder**
> FuturesPriceTriggeredOrder CancelPriceTriggeredOrder(ctx, settle, orderId)
Cancel a single order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | [default to btc]
  **orderId** | **string**| ID returned on order successfully being created | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.FuturesApi
settle := "btc"; // string - Settle currency
orderId := "orderId_example"; // string - ID returned on order successfully being created

result, _, err := api.CancelPriceTriggeredOrder(nil, settle, orderId)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**FuturesPriceTriggeredOrder**](FuturesPriceTriggeredOrder.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelPriceTriggeredOrderList**
> []FuturesPriceTriggeredOrder CancelPriceTriggeredOrderList(ctx, settle, contract)
Cancel all open orders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | [default to btc]
  **contract** | **string**| Futures contract | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.FuturesApi
settle := "btc"; // string - Settle currency
contract := "BTC_USD"; // string - Futures contract

result, _, err := api.CancelPriceTriggeredOrderList(nil, settle, contract)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**[]FuturesPriceTriggeredOrder**](FuturesPriceTriggeredOrder.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFuturesOrder**
> FuturesOrder CreateFuturesOrder(ctx, settle, futuresOrder)
Create a futures order

Zero-fill order cannot be retrieved 60 seconds after cancellation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | [default to btc]
  **futuresOrder** | [**FuturesOrder**](FuturesOrder.md)|  | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.FuturesApi
settle := "btc"; // string - Settle currency
futuresOrder := new (gateapi.FuturesOrder); // FuturesOrder - 

result, _, err := api.CreateFuturesOrder(nil, settle, futuresOrder)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**FuturesOrder**](FuturesOrder.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePriceTriggeredOrder**
> TriggerOrderResponse CreatePriceTriggeredOrder(ctx, settle, futuresPriceTriggeredOrder)
Create a price-triggered order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | [default to btc]
  **futuresPriceTriggeredOrder** | [**FuturesPriceTriggeredOrder**](FuturesPriceTriggeredOrder.md)|  | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.FuturesApi
settle := "btc"; // string - Settle currency
futuresPriceTriggeredOrder := new (gateapi.FuturesPriceTriggeredOrder); // FuturesPriceTriggeredOrder - 

result, _, err := api.CreatePriceTriggeredOrder(nil, settle, futuresPriceTriggeredOrder)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**TriggerOrderResponse**](TriggerOrderResponse.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFuturesContract**
> Contract GetFuturesContract(ctx, settle, contract)
Get a single contract

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | [default to btc]
  **contract** | **string**| Futures contract | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
api := client.FuturesApi
settle := "btc"; // string - Settle currency
contract := "BTC_USD"; // string - Futures contract

result, _, err := api.GetFuturesContract(nil, settle, contract)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**Contract**](Contract.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFuturesOrder**
> FuturesOrder GetFuturesOrder(ctx, settle, orderId)
Get a single order

Zero-fill order cannot be retrieved 60 seconds after cancellation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | [default to btc]
  **orderId** | **string**| ID returned on order successfully being created | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.FuturesApi
settle := "btc"; // string - Settle currency
orderId := "12345"; // string - ID returned on order successfully being created

result, _, err := api.GetFuturesOrder(nil, settle, orderId)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**FuturesOrder**](FuturesOrder.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMyTrades**
> []MyFuturesTrade GetMyTrades(ctx, settle, optional)
List personal trading history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | [default to btc]
 **optional** | ***GetMyTradesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetMyTradesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contract** | **optional.String**| Futures contract, return related data only if specified | 
 **order** | **optional.Int32**| Futures order ID, return related data only if specified | 
 **limit** | **optional.Int32**| Maximum number of record returned in one list | [default to 100]
 **offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]
 **lastId** | **optional.String**| Specify list staring point using the &#x60;id&#x60; of last record in previous list-query results | 
 **countTotal** | **optional.Int32**| Whether to return total number matched. Default to 0(no return) | [default to 0]

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.FuturesApi
settle := "btc"; // string - Settle currency

result, _, err := api.GetMyTrades(nil, settle, nil)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**[]MyFuturesTrade**](MyFuturesTrade.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPosition**
> Position GetPosition(ctx, settle, contract)
Get single position

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | [default to btc]
  **contract** | **string**| Futures contract | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.FuturesApi
settle := "btc"; // string - Settle currency
contract := "BTC_USD"; // string - Futures contract

result, _, err := api.GetPosition(nil, settle, contract)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**Position**](Position.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPriceTriggeredOrder**
> FuturesPriceTriggeredOrder GetPriceTriggeredOrder(ctx, settle, orderId)
Get a single order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | [default to btc]
  **orderId** | **string**| ID returned on order successfully being created | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.FuturesApi
settle := "btc"; // string - Settle currency
orderId := "orderId_example"; // string - ID returned on order successfully being created

result, _, err := api.GetPriceTriggeredOrder(nil, settle, orderId)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**FuturesPriceTriggeredOrder**](FuturesPriceTriggeredOrder.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFuturesAccountBook**
> []FuturesAccountBook ListFuturesAccountBook(ctx, settle, optional)
Query account book

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | [default to btc]
 **optional** | ***ListFuturesAccountBookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListFuturesAccountBookOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Maximum number of record returned in one list | [default to 100]
 **from** | **optional.Int32**| Start timestamp | 
 **to** | **optional.Int32**| End timestamp | 
 **type_** | **optional.String**| Changing Type: - dnw: Deposit &amp; Withdraw - pnl: Profit &amp; Loss by reducing position - fee: Trading fee - refr: Referrer rebate - fund: Funding - point_dnw: POINT Deposit &amp; Withdraw - point_fee: POINT Trading fee - point_refr: POINT Referrer rebate | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.FuturesApi
settle := "btc"; // string - Settle currency

result, _, err := api.ListFuturesAccountBook(nil, settle, nil)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**[]FuturesAccountBook**](FuturesAccountBook.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFuturesAccounts**
> FuturesAccount ListFuturesAccounts(ctx, settle)
Query futures account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | [default to btc]

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.FuturesApi
settle := "btc"; // string - Settle currency

result, _, err := api.ListFuturesAccounts(nil, settle)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**FuturesAccount**](FuturesAccount.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFuturesCandlesticks**
> []FuturesCandlestick ListFuturesCandlesticks(ctx, settle, contract, optional)
Get futures candlesticks

Return specified contract candlesticks. If prefix `contract` with `mark_`, the contract's mark price candlesticks are returned; if prefix with `index_`, index price candlesticks will be returned.  Maximum of 2000 points are returned in one query. Be sure not to exceed the limit when specifying `from`, `to` and `interval`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | [default to btc]
  **contract** | **string**| Futures contract | 
 **optional** | ***ListFuturesCandlesticksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListFuturesCandlesticksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **from** | **optional.Float32**| Start time of candlesticks, formatted in Unix timestamp in seconds. Default to&#x60;to - 100 * interval&#x60; if not specified | 
 **to** | **optional.Float32**| End time of candlesticks, formatted in Unix timestamp in seconds. Default to current time | 
 **limit** | **optional.Int32**| Maximum recent data points returned. &#x60;limit&#x60; is conflicted with &#x60;from&#x60; and &#x60;to&#x60;. If either &#x60;from&#x60; or &#x60;to&#x60; is specified, request will be rejected. | [default to 100]
 **interval** | **optional.String**| Interval time between data points | [default to 5m]

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
api := client.FuturesApi
settle := "btc"; // string - Settle currency
contract := "BTC_USD"; // string - Futures contract

result, _, err := api.ListFuturesCandlesticks(nil, settle, contract, nil)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**[]FuturesCandlestick**](FuturesCandlestick.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFuturesContracts**
> []Contract ListFuturesContracts(ctx, settle)
List all futures contracts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | [default to btc]

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
api := client.FuturesApi
settle := "btc"; // string - Settle currency

result, _, err := api.ListFuturesContracts(nil, settle)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**[]Contract**](Contract.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFuturesFundingRateHistory**
> []FundingRateRecord ListFuturesFundingRateHistory(ctx, settle, contract, optional)
Funding rate history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | [default to btc]
  **contract** | **string**| Futures contract | 
 **optional** | ***ListFuturesFundingRateHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListFuturesFundingRateHistoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int32**| Maximum number of record returned in one list | [default to 100]

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
api := client.FuturesApi
settle := "btc"; // string - Settle currency
contract := "BTC_USD"; // string - Futures contract

result, _, err := api.ListFuturesFundingRateHistory(nil, settle, contract, nil)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**[]FundingRateRecord**](FundingRateRecord.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFuturesInsuranceLedger**
> []InsuranceRecord ListFuturesInsuranceLedger(ctx, settle, optional)
Futures insurance balance history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | [default to btc]
 **optional** | ***ListFuturesInsuranceLedgerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListFuturesInsuranceLedgerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Maximum number of record returned in one list | [default to 100]

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
api := client.FuturesApi
settle := "btc"; // string - Settle currency

result, _, err := api.ListFuturesInsuranceLedger(nil, settle, nil)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**[]InsuranceRecord**](InsuranceRecord.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFuturesOrderBook**
> FuturesOrderBook ListFuturesOrderBook(ctx, settle, contract, optional)
Futures order book

Bids will be sorted by price from high to low, while asks sorted reversely

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | [default to btc]
  **contract** | **string**| Futures contract | 
 **optional** | ***ListFuturesOrderBookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListFuturesOrderBookOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **interval** | **optional.String**| Order depth. 0 means no aggregation is applied. default to 0 | [default to 0]
 **limit** | **optional.Int32**| Maximum number of order depth data in asks or bids | [default to 10]

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
api := client.FuturesApi
settle := "btc"; // string - Settle currency
contract := "BTC_USD"; // string - Futures contract

result, _, err := api.ListFuturesOrderBook(nil, settle, contract, nil)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**FuturesOrderBook**](FuturesOrderBook.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFuturesOrders**
> []FuturesOrder ListFuturesOrders(ctx, settle, contract, status, optional)
List futures orders

Zero-fill order cannot be retrieved 60 seconds after cancellation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | [default to btc]
  **contract** | **string**| Futures contract | 
  **status** | **string**| List orders based on status | 
 **optional** | ***ListFuturesOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListFuturesOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **optional.Int32**| Maximum number of record returned in one list | [default to 100]
 **offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]
 **lastId** | **optional.String**| Specify list staring point using the &#x60;id&#x60; of last record in previous list-query results | 
 **countTotal** | **optional.Int32**| Whether to return total number matched. Default to 0(no return) | [default to 0]

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.FuturesApi
settle := "btc"; // string - Settle currency
contract := "BTC_USD"; // string - Futures contract
status := "open"; // string - List orders based on status

result, _, err := api.ListFuturesOrders(nil, settle, contract, status, nil)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**[]FuturesOrder**](FuturesOrder.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFuturesTickers**
> []FuturesTicker ListFuturesTickers(ctx, settle, optional)
List futures tickers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | [default to btc]
 **optional** | ***ListFuturesTickersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListFuturesTickersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contract** | **optional.String**| Futures contract, return related data only if specified | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
api := client.FuturesApi
settle := "btc"; // string - Settle currency

result, _, err := api.ListFuturesTickers(nil, settle, nil)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**[]FuturesTicker**](FuturesTicker.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFuturesTrades**
> []FuturesTrade ListFuturesTrades(ctx, settle, contract, optional)
Futures trading history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | [default to btc]
  **contract** | **string**| Futures contract | 
 **optional** | ***ListFuturesTradesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListFuturesTradesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int32**| Maximum number of record returned in one list | [default to 100]
 **lastId** | **optional.String**| Specify list staring point using the id of last record in previous list-query results  This parameter is deprecated. Use &#x60;from&#x60; and &#x60;to&#x60; instead to limit time range | 
 **from** | **optional.Float32**| Specify starting time in Unix seconds. If not specified, &#x60;to&#x60; and &#x60;limit&#x60; will be used to limit response items. If items between &#x60;from&#x60; and &#x60;to&#x60; are more than &#x60;limit&#x60;, only &#x60;limit&#x60; number will be returned.  | 
 **to** | **optional.Float32**| Specify end time in Unix seconds, default to current time | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
api := client.FuturesApi
settle := "btc"; // string - Settle currency
contract := "BTC_USD"; // string - Futures contract

result, _, err := api.ListFuturesTrades(nil, settle, contract, nil)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**[]FuturesTrade**](FuturesTrade.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLiquidates**
> []FuturesLiquidate ListLiquidates(ctx, settle, optional)
List liquidation history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | [default to btc]
 **optional** | ***ListLiquidatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListLiquidatesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contract** | **optional.String**| Futures contract, return related data only if specified | 
 **limit** | **optional.Int32**| Maximum number of record returned in one list | [default to 100]
 **at** | **optional.Int32**| Specify a liquidation timestamp | [default to 0]

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.FuturesApi
settle := "btc"; // string - Settle currency

result, _, err := api.ListLiquidates(nil, settle, nil)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**[]FuturesLiquidate**](FuturesLiquidate.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPositionClose**
> []PositionClose ListPositionClose(ctx, settle, optional)
List position close history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | [default to btc]
 **optional** | ***ListPositionCloseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListPositionCloseOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contract** | **optional.String**| Futures contract, return related data only if specified | 
 **limit** | **optional.Int32**| Maximum number of record returned in one list | [default to 100]

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.FuturesApi
settle := "btc"; // string - Settle currency

result, _, err := api.ListPositionClose(nil, settle, nil)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**[]PositionClose**](PositionClose.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPositions**
> []Position ListPositions(ctx, settle)
List all positions of a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | [default to btc]

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.FuturesApi
settle := "btc"; // string - Settle currency

result, _, err := api.ListPositions(nil, settle)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**[]Position**](Position.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPriceTriggeredOrders**
> []FuturesPriceTriggeredOrder ListPriceTriggeredOrders(ctx, settle, status, optional)
List all auto orders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | [default to btc]
  **status** | **string**| List orders based on status | 
 **optional** | ***ListPriceTriggeredOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListPriceTriggeredOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contract** | **optional.String**| Futures contract, return related data only if specified | 
 **limit** | **optional.Int32**| Maximum number of record returned in one list | [default to 100]
 **offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.FuturesApi
settle := "btc"; // string - Settle currency
status := "status_example"; // string - List orders based on status

result, _, err := api.ListPriceTriggeredOrders(nil, settle, status, nil)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**[]FuturesPriceTriggeredOrder**](FuturesPriceTriggeredOrder.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePositionLeverage**
> Position UpdatePositionLeverage(ctx, settle, contract, leverage)
Update position leverage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | [default to btc]
  **contract** | **string**| Futures contract | 
  **leverage** | **string**| New position leverage | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.FuturesApi
settle := "btc"; // string - Settle currency
contract := "BTC_USD"; // string - Futures contract
leverage := "10"; // string - New position leverage

result, _, err := api.UpdatePositionLeverage(nil, settle, contract, leverage)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**Position**](Position.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePositionMargin**
> Position UpdatePositionMargin(ctx, settle, contract, change)
Update position margin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | [default to btc]
  **contract** | **string**| Futures contract | 
  **change** | **string**| Margin change. Use positive number to increase margin, negative number otherwise. | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.FuturesApi
settle := "btc"; // string - Settle currency
contract := "BTC_USD"; // string - Futures contract
change := "0.01"; // string - Margin change. Use positive number to increase margin, negative number otherwise.

result, _, err := api.UpdatePositionMargin(nil, settle, contract, change)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**Position**](Position.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePositionRiskLimit**
> Position UpdatePositionRiskLimit(ctx, settle, contract, riskLimit)
Update position risk limit

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | [default to btc]
  **contract** | **string**| Futures contract | 
  **riskLimit** | **string**| New position risk limit | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.FuturesApi
settle := "btc"; // string - Settle currency
contract := "BTC_USD"; // string - Futures contract
riskLimit := "10"; // string - New position risk limit

result, _, err := api.UpdatePositionRiskLimit(nil, settle, contract, riskLimit)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**Position**](Position.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

