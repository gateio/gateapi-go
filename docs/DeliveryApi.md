# \DeliveryApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelDeliveryOrder**](DeliveryApi.md#CancelDeliveryOrder) | **Delete** /delivery/{settle}/orders/{order_id} | Cancel a single order
[**CancelDeliveryOrders**](DeliveryApi.md#CancelDeliveryOrders) | **Delete** /delivery/{settle}/orders | Cancel all &#x60;open&#x60; orders matched
[**CancelPriceTriggeredDeliveryOrder**](DeliveryApi.md#CancelPriceTriggeredDeliveryOrder) | **Delete** /delivery/{settle}/price_orders/{order_id} | Cancel a single order
[**CancelPriceTriggeredDeliveryOrderList**](DeliveryApi.md#CancelPriceTriggeredDeliveryOrderList) | **Delete** /delivery/{settle}/price_orders | Cancel all open orders
[**CreateDeliveryOrder**](DeliveryApi.md#CreateDeliveryOrder) | **Post** /delivery/{settle}/orders | Create a futures order
[**CreatePriceTriggeredDeliveryOrder**](DeliveryApi.md#CreatePriceTriggeredDeliveryOrder) | **Post** /delivery/{settle}/price_orders | Create a price-triggered order
[**GetDeliveryContract**](DeliveryApi.md#GetDeliveryContract) | **Get** /delivery/{settle}/contracts/{contract} | Get a single contract
[**GetDeliveryOrder**](DeliveryApi.md#GetDeliveryOrder) | **Get** /delivery/{settle}/orders/{order_id} | Get a single order
[**GetDeliveryPosition**](DeliveryApi.md#GetDeliveryPosition) | **Get** /delivery/{settle}/positions/{contract} | Get single position
[**GetMyDeliveryTrades**](DeliveryApi.md#GetMyDeliveryTrades) | **Get** /delivery/{settle}/my_trades | List personal trading history
[**GetPriceTriggeredDeliveryOrder**](DeliveryApi.md#GetPriceTriggeredDeliveryOrder) | **Get** /delivery/{settle}/price_orders/{order_id} | Get a single order
[**ListDeliveryAccountBook**](DeliveryApi.md#ListDeliveryAccountBook) | **Get** /delivery/{settle}/account_book | Query account book
[**ListDeliveryAccounts**](DeliveryApi.md#ListDeliveryAccounts) | **Get** /delivery/{settle}/accounts | Query futures account
[**ListDeliveryCandlesticks**](DeliveryApi.md#ListDeliveryCandlesticks) | **Get** /delivery/{settle}/candlesticks | Get futures candlesticks
[**ListDeliveryContracts**](DeliveryApi.md#ListDeliveryContracts) | **Get** /delivery/{settle}/contracts | List all futures contracts
[**ListDeliveryInsuranceLedger**](DeliveryApi.md#ListDeliveryInsuranceLedger) | **Get** /delivery/{settle}/insurance | Futures insurance balance history
[**ListDeliveryLiquidates**](DeliveryApi.md#ListDeliveryLiquidates) | **Get** /delivery/{settle}/liquidates | List liquidation history
[**ListDeliveryOrderBook**](DeliveryApi.md#ListDeliveryOrderBook) | **Get** /delivery/{settle}/order_book | Futures order book
[**ListDeliveryOrders**](DeliveryApi.md#ListDeliveryOrders) | **Get** /delivery/{settle}/orders | List futures orders
[**ListDeliveryPositionClose**](DeliveryApi.md#ListDeliveryPositionClose) | **Get** /delivery/{settle}/position_close | List position close history
[**ListDeliveryPositions**](DeliveryApi.md#ListDeliveryPositions) | **Get** /delivery/{settle}/positions | List all positions of a user
[**ListDeliverySettlements**](DeliveryApi.md#ListDeliverySettlements) | **Get** /delivery/{settle}/settlements | List settlement history
[**ListDeliveryTickers**](DeliveryApi.md#ListDeliveryTickers) | **Get** /delivery/{settle}/tickers | List futures tickers
[**ListDeliveryTrades**](DeliveryApi.md#ListDeliveryTrades) | **Get** /delivery/{settle}/trades | Futures trading history
[**ListPriceTriggeredDeliveryOrders**](DeliveryApi.md#ListPriceTriggeredDeliveryOrders) | **Get** /delivery/{settle}/price_orders | List all auto orders
[**UpdateDeliveryPositionLeverage**](DeliveryApi.md#UpdateDeliveryPositionLeverage) | **Post** /delivery/{settle}/positions/{contract}/leverage | Update position leverage
[**UpdateDeliveryPositionMargin**](DeliveryApi.md#UpdateDeliveryPositionMargin) | **Post** /delivery/{settle}/positions/{contract}/margin | Update position margin
[**UpdateDeliveryPositionRiskLimit**](DeliveryApi.md#UpdateDeliveryPositionRiskLimit) | **Post** /delivery/{settle}/positions/{contract}/risk_limit | Update position risk limit


# **CancelDeliveryOrder**
> FuturesOrder CancelDeliveryOrder(ctx, settle, orderId)
Cancel a single order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | 
  **orderId** | **string**| ID returned on order successfully being created | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.DeliveryApi
settle := "usdt"; // string - Settle currency
orderId := "12345"; // string - ID returned on order successfully being created

result, _, err := api.CancelDeliveryOrder(nil, settle, orderId)
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

# **CancelDeliveryOrders**
> []FuturesOrder CancelDeliveryOrders(ctx, settle, contract, optional)
Cancel all `open` orders matched

Zero-fill order cannot be retrieved 60 seconds after cancellation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | 
  **contract** | **string**| Futures contract | 
 **optional** | ***CancelDeliveryOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CancelDeliveryOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **side** | **optional.String**| All bids or asks. Both included in not specified | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.DeliveryApi
settle := "usdt"; // string - Settle currency
contract := "BTC_USDT_WEEKLY_20200703"; // string - Futures contract

result, _, err := api.CancelDeliveryOrders(nil, settle, contract, nil)
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

# **CancelPriceTriggeredDeliveryOrder**
> FuturesPriceTriggeredOrder CancelPriceTriggeredDeliveryOrder(ctx, settle, orderId)
Cancel a single order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | 
  **orderId** | **string**| ID returned on order successfully being created | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.DeliveryApi
settle := "usdt"; // string - Settle currency
orderId := "orderId_example"; // string - ID returned on order successfully being created

result, _, err := api.CancelPriceTriggeredDeliveryOrder(nil, settle, orderId)
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

# **CancelPriceTriggeredDeliveryOrderList**
> []FuturesPriceTriggeredOrder CancelPriceTriggeredDeliveryOrderList(ctx, settle, contract)
Cancel all open orders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | 
  **contract** | **string**| Futures contract | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.DeliveryApi
settle := "usdt"; // string - Settle currency
contract := "BTC_USD"; // string - Futures contract

result, _, err := api.CancelPriceTriggeredDeliveryOrderList(nil, settle, contract)
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

# **CreateDeliveryOrder**
> FuturesOrder CreateDeliveryOrder(ctx, settle, futuresOrder)
Create a futures order

Zero-fill order cannot be retrieved 60 seconds after cancellation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | 
  **futuresOrder** | [**FuturesOrder**](FuturesOrder.md)|  | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.DeliveryApi
settle := "usdt"; // string - Settle currency
futuresOrder := new (gateapi.FuturesOrder); // FuturesOrder - 

result, _, err := api.CreateDeliveryOrder(nil, settle, futuresOrder)
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

# **CreatePriceTriggeredDeliveryOrder**
> TriggerOrderResponse CreatePriceTriggeredDeliveryOrder(ctx, settle, futuresPriceTriggeredOrder)
Create a price-triggered order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | 
  **futuresPriceTriggeredOrder** | [**FuturesPriceTriggeredOrder**](FuturesPriceTriggeredOrder.md)|  | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.DeliveryApi
settle := "usdt"; // string - Settle currency
futuresPriceTriggeredOrder := new (gateapi.FuturesPriceTriggeredOrder); // FuturesPriceTriggeredOrder - 

result, _, err := api.CreatePriceTriggeredDeliveryOrder(nil, settle, futuresPriceTriggeredOrder)
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

# **GetDeliveryContract**
> DeliveryContract GetDeliveryContract(ctx, settle, contract)
Get a single contract

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | 
  **contract** | **string**| Futures contract | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
api := client.DeliveryApi
settle := "usdt"; // string - Settle currency
contract := "BTC_USDT_WEEKLY_20200703"; // string - Futures contract

result, _, err := api.GetDeliveryContract(nil, settle, contract)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**DeliveryContract**](DeliveryContract.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeliveryOrder**
> FuturesOrder GetDeliveryOrder(ctx, settle, orderId)
Get a single order

Zero-fill order cannot be retrieved 60 seconds after cancellation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | 
  **orderId** | **string**| ID returned on order successfully being created | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.DeliveryApi
settle := "usdt"; // string - Settle currency
orderId := "12345"; // string - ID returned on order successfully being created

result, _, err := api.GetDeliveryOrder(nil, settle, orderId)
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

# **GetDeliveryPosition**
> Position GetDeliveryPosition(ctx, settle, contract)
Get single position

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | 
  **contract** | **string**| Futures contract | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.DeliveryApi
settle := "usdt"; // string - Settle currency
contract := "BTC_USDT_WEEKLY_20200703"; // string - Futures contract

result, _, err := api.GetDeliveryPosition(nil, settle, contract)
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

# **GetMyDeliveryTrades**
> []MyFuturesTrade GetMyDeliveryTrades(ctx, settle, optional)
List personal trading history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | 
 **optional** | ***GetMyDeliveryTradesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetMyDeliveryTradesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contract** | **optional.String**| Futures contract | 
 **order** | **optional.Int32**| Futures order ID, return related data only if specified | 
 **limit** | **optional.Int32**| Maximum number of records returned in one list | [default to 100]
 **offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]
 **lastId** | **optional.String**| Specify list staring point using the &#x60;id&#x60; of last record in previous list-query results | 
 **countTotal** | **optional.Int32**| Whether to return total number matched. Default to 0(no return) | [default to 0]

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.DeliveryApi
settle := "usdt"; // string - Settle currency

result, _, err := api.GetMyDeliveryTrades(nil, settle, nil)
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

# **GetPriceTriggeredDeliveryOrder**
> FuturesPriceTriggeredOrder GetPriceTriggeredDeliveryOrder(ctx, settle, orderId)
Get a single order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | 
  **orderId** | **string**| ID returned on order successfully being created | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.DeliveryApi
settle := "usdt"; // string - Settle currency
orderId := "orderId_example"; // string - ID returned on order successfully being created

result, _, err := api.GetPriceTriggeredDeliveryOrder(nil, settle, orderId)
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

# **ListDeliveryAccountBook**
> []FuturesAccountBook ListDeliveryAccountBook(ctx, settle, optional)
Query account book

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | 
 **optional** | ***ListDeliveryAccountBookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListDeliveryAccountBookOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Maximum number of records returned in one list | [default to 100]
 **from** | **optional.Int32**| Start timestamp | 
 **to** | **optional.Int32**| End timestamp | 
 **type_** | **optional.String**| Changing Type: - dnw: Deposit &amp; Withdraw - pnl: Profit &amp; Loss by reducing position - fee: Trading fee - refr: Referrer rebate - fund: Funding - point_dnw: POINT Deposit &amp; Withdraw - point_fee: POINT Trading fee - point_refr: POINT Referrer rebate | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.DeliveryApi
settle := "usdt"; // string - Settle currency

result, _, err := api.ListDeliveryAccountBook(nil, settle, nil)
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

# **ListDeliveryAccounts**
> FuturesAccount ListDeliveryAccounts(ctx, settle)
Query futures account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.DeliveryApi
settle := "usdt"; // string - Settle currency

result, _, err := api.ListDeliveryAccounts(nil, settle)
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

# **ListDeliveryCandlesticks**
> []FuturesCandlestick ListDeliveryCandlesticks(ctx, settle, contract, optional)
Get futures candlesticks

Return specified contract candlesticks. If prefix `contract` with `mark_`, the contract's mark price candlesticks are returned; if prefix with `index_`, index price candlesticks will be returned.  Maximum of 2000 points are returned in one query. Be sure not to exceed the limit when specifying `from`, `to` and `interval`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | 
  **contract** | **string**| Futures contract | 
 **optional** | ***ListDeliveryCandlesticksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListDeliveryCandlesticksOpts struct

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
api := client.DeliveryApi
settle := "usdt"; // string - Settle currency
contract := "BTC_USDT_WEEKLY_20200703"; // string - Futures contract

result, _, err := api.ListDeliveryCandlesticks(nil, settle, contract, nil)
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

# **ListDeliveryContracts**
> []DeliveryContract ListDeliveryContracts(ctx, settle)
List all futures contracts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
api := client.DeliveryApi
settle := "usdt"; // string - Settle currency

result, _, err := api.ListDeliveryContracts(nil, settle)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**[]DeliveryContract**](DeliveryContract.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDeliveryInsuranceLedger**
> []InsuranceRecord ListDeliveryInsuranceLedger(ctx, settle, optional)
Futures insurance balance history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | 
 **optional** | ***ListDeliveryInsuranceLedgerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListDeliveryInsuranceLedgerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Maximum number of records returned in one list | [default to 100]

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
api := client.DeliveryApi
settle := "usdt"; // string - Settle currency

result, _, err := api.ListDeliveryInsuranceLedger(nil, settle, nil)
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

# **ListDeliveryLiquidates**
> []FuturesLiquidate ListDeliveryLiquidates(ctx, settle, optional)
List liquidation history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | 
 **optional** | ***ListDeliveryLiquidatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListDeliveryLiquidatesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contract** | **optional.String**| Futures contract | 
 **limit** | **optional.Int32**| Maximum number of records returned in one list | [default to 100]
 **at** | **optional.Int32**| Specify a liquidation timestamp | [default to 0]

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.DeliveryApi
settle := "usdt"; // string - Settle currency

result, _, err := api.ListDeliveryLiquidates(nil, settle, nil)
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

# **ListDeliveryOrderBook**
> FuturesOrderBook ListDeliveryOrderBook(ctx, settle, contract, optional)
Futures order book

Bids will be sorted by price from high to low, while asks sorted reversely

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | 
  **contract** | **string**| Futures contract | 
 **optional** | ***ListDeliveryOrderBookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListDeliveryOrderBookOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **interval** | **optional.String**| Order depth. 0 means no aggregation is applied. default to 0 | [default to 0]
 **limit** | **optional.Int32**| Maximum number of order depth data in asks or bids | [default to 10]

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
api := client.DeliveryApi
settle := "usdt"; // string - Settle currency
contract := "BTC_USDT_WEEKLY_20200703"; // string - Futures contract

result, _, err := api.ListDeliveryOrderBook(nil, settle, contract, nil)
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

# **ListDeliveryOrders**
> []FuturesOrder ListDeliveryOrders(ctx, settle, status, optional)
List futures orders

Zero-fill order cannot be retrieved 60 seconds after cancellation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | 
  **status** | **string**| List orders based on status | 
 **optional** | ***ListDeliveryOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListDeliveryOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contract** | **optional.String**| Futures contract | 
 **limit** | **optional.Int32**| Maximum number of records returned in one list | [default to 100]
 **offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]
 **lastId** | **optional.String**| Specify list staring point using the &#x60;id&#x60; of last record in previous list-query results | 
 **countTotal** | **optional.Int32**| Whether to return total number matched. Default to 0(no return) | [default to 0]

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.DeliveryApi
settle := "usdt"; // string - Settle currency
status := "open"; // string - List orders based on status

result, _, err := api.ListDeliveryOrders(nil, settle, status, nil)
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

# **ListDeliveryPositionClose**
> []PositionClose ListDeliveryPositionClose(ctx, settle, optional)
List position close history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | 
 **optional** | ***ListDeliveryPositionCloseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListDeliveryPositionCloseOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contract** | **optional.String**| Futures contract | 
 **limit** | **optional.Int32**| Maximum number of records returned in one list | [default to 100]

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.DeliveryApi
settle := "usdt"; // string - Settle currency

result, _, err := api.ListDeliveryPositionClose(nil, settle, nil)
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

# **ListDeliveryPositions**
> []Position ListDeliveryPositions(ctx, settle)
List all positions of a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.DeliveryApi
settle := "usdt"; // string - Settle currency

result, _, err := api.ListDeliveryPositions(nil, settle)
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

# **ListDeliverySettlements**
> []DeliverySettlement ListDeliverySettlements(ctx, settle, optional)
List settlement history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | 
 **optional** | ***ListDeliverySettlementsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListDeliverySettlementsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contract** | **optional.String**| Futures contract | 
 **limit** | **optional.Int32**| Maximum number of records returned in one list | [default to 100]
 **at** | **optional.Int32**| Specify a settlement timestamp | [default to 0]

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.DeliveryApi
settle := "usdt"; // string - Settle currency

result, _, err := api.ListDeliverySettlements(nil, settle, nil)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

### Return type

[**[]DeliverySettlement**](DeliverySettlement.md)

### Authorization

Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDeliveryTickers**
> []FuturesTicker ListDeliveryTickers(ctx, settle, optional)
List futures tickers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | 
 **optional** | ***ListDeliveryTickersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListDeliveryTickersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contract** | **optional.String**| Futures contract | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
api := client.DeliveryApi
settle := "usdt"; // string - Settle currency

result, _, err := api.ListDeliveryTickers(nil, settle, nil)
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

# **ListDeliveryTrades**
> []FuturesTrade ListDeliveryTrades(ctx, settle, contract, optional)
Futures trading history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | 
  **contract** | **string**| Futures contract | 
 **optional** | ***ListDeliveryTradesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListDeliveryTradesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int32**| Maximum number of records returned in one list | [default to 100]
 **lastId** | **optional.String**| Specify list staring point using the id of last record in previous list-query results  This parameter is deprecated. Use &#x60;from&#x60; and &#x60;to&#x60; instead to limit time range | 
 **from** | **optional.Float32**| Specify starting time in Unix seconds. If not specified, &#x60;to&#x60; and &#x60;limit&#x60; will be used to limit response items. If items between &#x60;from&#x60; and &#x60;to&#x60; are more than &#x60;limit&#x60;, only &#x60;limit&#x60; number will be returned.  | 
 **to** | **optional.Float32**| Specify end time in Unix seconds, default to current time | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
api := client.DeliveryApi
settle := "usdt"; // string - Settle currency
contract := "BTC_USDT_WEEKLY_20200703"; // string - Futures contract

result, _, err := api.ListDeliveryTrades(nil, settle, contract, nil)
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

# **ListPriceTriggeredDeliveryOrders**
> []FuturesPriceTriggeredOrder ListPriceTriggeredDeliveryOrders(ctx, settle, status, optional)
List all auto orders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | 
  **status** | **string**| List orders based on status | 
 **optional** | ***ListPriceTriggeredDeliveryOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListPriceTriggeredDeliveryOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contract** | **optional.String**| Futures contract, return related data only if specified | 
 **limit** | **optional.Int32**| Maximum number of records returned in one list | [default to 100]
 **offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.DeliveryApi
settle := "usdt"; // string - Settle currency
status := "status_example"; // string - List orders based on status

result, _, err := api.ListPriceTriggeredDeliveryOrders(nil, settle, status, nil)
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

# **UpdateDeliveryPositionLeverage**
> Position UpdateDeliveryPositionLeverage(ctx, settle, contract, leverage)
Update position leverage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | 
  **contract** | **string**| Futures contract | 
  **leverage** | **string**| New position leverage | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.DeliveryApi
settle := "usdt"; // string - Settle currency
contract := "BTC_USDT_WEEKLY_20200703"; // string - Futures contract
leverage := "10"; // string - New position leverage

result, _, err := api.UpdateDeliveryPositionLeverage(nil, settle, contract, leverage)
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

# **UpdateDeliveryPositionMargin**
> Position UpdateDeliveryPositionMargin(ctx, settle, contract, change)
Update position margin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | 
  **contract** | **string**| Futures contract | 
  **change** | **string**| Margin change. Use positive number to increase margin, negative number otherwise. | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.DeliveryApi
settle := "usdt"; // string - Settle currency
contract := "BTC_USDT_WEEKLY_20200703"; // string - Futures contract
change := "0.01"; // string - Margin change. Use positive number to increase margin, negative number otherwise.

result, _, err := api.UpdateDeliveryPositionMargin(nil, settle, contract, change)
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

# **UpdateDeliveryPositionRiskLimit**
> Position UpdateDeliveryPositionRiskLimit(ctx, settle, contract, riskLimit)
Update position risk limit

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settle** | **string**| Settle currency | 
  **contract** | **string**| Futures contract | 
  **riskLimit** | **string**| New position risk limit | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.DeliveryApi
settle := "usdt"; // string - Settle currency
contract := "BTC_USDT_WEEKLY_20200703"; // string - Futures contract
riskLimit := "10"; // string - New position risk limit

result, _, err := api.UpdateDeliveryPositionRiskLimit(nil, settle, contract, riskLimit)
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

