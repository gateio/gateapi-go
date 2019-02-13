# \FuturesApi

All URIs are relative to *https://fx-api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelOrder**](FuturesApi.md#CancelOrder) | **Delete** /futures/orders/{order_id} | Cancel a single order
[**CancelOrders**](FuturesApi.md#CancelOrders) | **Delete** /futures/orders | Cancel all &#x60;open&#x60; orders matched
[**CreateOrder**](FuturesApi.md#CreateOrder) | **Post** /futures/orders | Create a futures order
[**GetFuturesContract**](FuturesApi.md#GetFuturesContract) | **Get** /futures/contracts/{contract} | Get a single contract
[**GetMyTrades**](FuturesApi.md#GetMyTrades) | **Get** /futures/my_trades | List personal trading history
[**GetOrder**](FuturesApi.md#GetOrder) | **Get** /futures/orders/{order_id} | Get a single order
[**GetPosition**](FuturesApi.md#GetPosition) | **Get** /futures/positions/{contract} | Get single position
[**ListFuturesAccountBook**](FuturesApi.md#ListFuturesAccountBook) | **Get** /futures/account_book | Query account book
[**ListFuturesAccounts**](FuturesApi.md#ListFuturesAccounts) | **Get** /futures/accounts | Query futures account
[**ListFuturesCandlesticks**](FuturesApi.md#ListFuturesCandlesticks) | **Get** /futures/candlesticks | Get futures candlesticks
[**ListFuturesContracts**](FuturesApi.md#ListFuturesContracts) | **Get** /futures/contracts | List all futures contracts
[**ListFuturesFundingRateHistory**](FuturesApi.md#ListFuturesFundingRateHistory) | **Get** /futures/funding_rate | Funding rate history
[**ListFuturesInsuranceLedger**](FuturesApi.md#ListFuturesInsuranceLedger) | **Get** /futures/insurance | Futures insurance balance history
[**ListFuturesOrderBook**](FuturesApi.md#ListFuturesOrderBook) | **Get** /futures/order_book | Futures order book
[**ListFuturesTickers**](FuturesApi.md#ListFuturesTickers) | **Get** /futures/tickers | List futures tickers
[**ListFuturesTrades**](FuturesApi.md#ListFuturesTrades) | **Get** /futures/trades | Futures trading history
[**ListOrders**](FuturesApi.md#ListOrders) | **Get** /futures/orders | List futures orders
[**ListPositionClose**](FuturesApi.md#ListPositionClose) | **Get** /futures/position_close | List position close history
[**ListPositions**](FuturesApi.md#ListPositions) | **Get** /futures/positions | List all positions of a user
[**UpdatePositionLeverage**](FuturesApi.md#UpdatePositionLeverage) | **Post** /futures/positions/{contract}/leverage | Update position leverage
[**UpdatePositionMargin**](FuturesApi.md#UpdatePositionMargin) | **Post** /futures/positions/{contract}/margin | Update position margin
[**UpdatePositionRiskLimit**](FuturesApi.md#UpdatePositionRiskLimit) | **Post** /futures/positions/{contract}/risk_limit | Update position risk limit


# **CancelOrder**
> FuturesOrder CancelOrder(ctx, orderId)
Cancel a single order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **string**| ID returned on order successfully being created | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.FuturesApi
orderId := "12345"; // string - ID returned on order successfully being created

result, _, err = api.CancelOrder(nil, orderId)
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

# **CancelOrders**
> []FuturesOrder CancelOrders(ctx, contract, optional)
Cancel all `open` orders matched

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contract** | **string**| Futures contract | 
 **optional** | ***CancelOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CancelOrdersOpts struct

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
contract := "BTC_USD"; // string - Futures contract

result, _, err = api.CancelOrders(nil, contract, nil)
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

# **CreateOrder**
> FuturesOrder CreateOrder(ctx, futuresOrder)
Create a futures order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **futuresOrder** | [**FuturesOrder**](FuturesOrder.md)|  | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.FuturesApi
futuresOrder := new (gateapi.FuturesOrder); // FuturesOrder - 

result, _, err = api.CreateOrder(nil, futuresOrder)
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

# **GetFuturesContract**
> Contract GetFuturesContract(ctx, contract)
Get a single contract

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contract** | **string**| Futures contract | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
api := client.FuturesApi
contract := "BTC_USD"; // string - Futures contract

result, _, err = api.GetFuturesContract(nil, contract)
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

# **GetMyTrades**
> []MyFuturesTrade GetMyTrades(ctx, optional)
List personal trading history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetMyTradesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetMyTradesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contract** | **optional.String**| Futures contract, return related data only if specified | 
 **order** | **optional.Int32**| Futures order ID, return related data only if specified | 
 **limit** | **optional.Int32**| Maximum number of record returned in one list | [default to 100]
 **lastId** | **optional.String**| Specify list staring point using the last record of &#x60;id&#x60; in previous list-query results | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.FuturesApi

result, _, err = api.GetMyTrades(nil, nil)
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

# **GetOrder**
> FuturesOrder GetOrder(ctx, orderId)
Get a single order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **string**| ID returned on order successfully being created | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.FuturesApi
orderId := "12345"; // string - ID returned on order successfully being created

result, _, err = api.GetOrder(nil, orderId)
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

# **GetPosition**
> Position GetPosition(ctx, contract)
Get single position

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contract** | **string**| Futures contract | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.FuturesApi
contract := "BTC_USD"; // string - Futures contract

result, _, err = api.GetPosition(nil, contract)
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

# **ListFuturesAccountBook**
> []FuturesAccountBook ListFuturesAccountBook(ctx, optional)
Query account book

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListFuturesAccountBookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListFuturesAccountBookOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Maximum number of record returned in one list | [default to 100]
 **from** | **optional.Int32**| Start timestamp | 
 **to** | **optional.Int32**| End timestamp | 
 **type_** | **optional.String**| Changing Type  - dnw: Deposit &amp; Withdraw - pnl: Profit &amp; Loss by reducing position - fee: Trading fee - refr: Referrer rebate - fund: Funding | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.FuturesApi

result, _, err = api.ListFuturesAccountBook(nil, nil)
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
> FuturesAccount ListFuturesAccounts(ctx, )
Query futures account

### Required Parameters
This endpoint does not need any parameter.

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.FuturesApi
result, _, err = api.ListFuturesAccounts(nil)
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
> []FuturesCandlestick ListFuturesCandlesticks(ctx, contract, optional)
Get futures candlesticks

Return specified contract candlesticks. If prefix `contract` with `mark_`, the contract's mark price candlesticks are returned; if prefix with `index_`, index price candlesticks will be returned.  Maximum of 2000 points are returned in one query. Be sure not to exceed the limit when specifying `from`, `to` and `interval`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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
contract := "BTC_USD"; // string - Futures contract

result, _, err = api.ListFuturesCandlesticks(nil, contract, nil)
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
> []Contract ListFuturesContracts(ctx, )
List all futures contracts

### Required Parameters
This endpoint does not need any parameter.

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
api := client.FuturesApi
result, _, err = api.ListFuturesContracts(nil)
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
> []FundingRateRecord ListFuturesFundingRateHistory(ctx, contract, optional)
Funding rate history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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
contract := "BTC_USD"; // string - Futures contract

result, _, err = api.ListFuturesFundingRateHistory(nil, contract, nil)
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
> []InsuranceRecord ListFuturesInsuranceLedger(ctx, optional)
Futures insurance balance history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

result, _, err = api.ListFuturesInsuranceLedger(nil, nil)
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
> FuturesOrderBook ListFuturesOrderBook(ctx, contract, optional)
Futures order book

Bids will be sorted by price from high to low, while asks sorted reversely

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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
contract := "BTC_USD"; // string - Futures contract

result, _, err = api.ListFuturesOrderBook(nil, contract, nil)
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

# **ListFuturesTickers**
> []FuturesTicker ListFuturesTickers(ctx, optional)
List futures tickers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

result, _, err = api.ListFuturesTickers(nil, nil)
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
> []FuturesTrade ListFuturesTrades(ctx, contract, optional)
Futures trading history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contract** | **string**| Futures contract | 
 **optional** | ***ListFuturesTradesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListFuturesTradesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Maximum number of record returned in one list | [default to 100]
 **lastId** | **optional.String**| Specify list staring point using the last record of &#x60;id&#x60; in previous list-query results | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
api := client.FuturesApi
contract := "BTC_USD"; // string - Futures contract

result, _, err = api.ListFuturesTrades(nil, contract, nil)
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

# **ListOrders**
> []FuturesOrder ListOrders(ctx, contract, status, optional)
List futures orders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contract** | **string**| Futures contract | 
  **status** | **string**| List orders based on status | 
 **optional** | ***ListOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int32**| Maximum number of record returned in one list | [default to 100]
 **lastId** | **optional.String**| Specify list staring point using the last record of &#x60;id&#x60; in previous list-query results | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.FuturesApi
contract := "BTC_USD"; // string - Futures contract
status := "open"; // string - List orders based on status

result, _, err = api.ListOrders(nil, contract, status, nil)
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

# **ListPositionClose**
> []PositionClose ListPositionClose(ctx, optional)
List position close history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

result, _, err = api.ListPositionClose(nil, nil)
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
> []Position ListPositions(ctx, )
List all positions of a user

### Required Parameters
This endpoint does not need any parameter.

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.FuturesApi
result, _, err = api.ListPositions(nil)
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

# **UpdatePositionLeverage**
> Position UpdatePositionLeverage(ctx, contract, leverage)
Update position leverage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contract** | **string**| Futures contract | 
  **leverage** | **string**| New position leverage | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.FuturesApi
contract := "BTC_USD"; // string - Futures contract
leverage := "10"; // string - New position leverage

result, _, err = api.UpdatePositionLeverage(nil, contract, leverage)
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
> Position UpdatePositionMargin(ctx, contract, change)
Update position margin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contract** | **string**| Futures contract | 
  **change** | **string**| Margin change. Use positive number to increase margin, negative number otherwise. | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.FuturesApi
contract := "BTC_USD"; // string - Futures contract
change := "0.01"; // string - Margin change. Use positive number to increase margin, negative number otherwise.

result, _, err = api.UpdatePositionMargin(nil, contract, change)
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
> Position UpdatePositionRiskLimit(ctx, contract, riskLimit)
Update position risk limit

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contract** | **string**| Futures contract | 
  **riskLimit** | **string**| New position risk limit | 

### Example

```golang
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.FuturesApi
contract := "BTC_USD"; // string - Futures contract
riskLimit := "10"; // string - New position risk limit

result, _, err = api.UpdatePositionRiskLimit(nil, contract, riskLimit)
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

