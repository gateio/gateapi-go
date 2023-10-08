# FlashSwapPreviewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SellCurrency** | **string** | 卖出的资产名称， 根据接口&#x60;查询支持闪兑的所有交易对列表 GET /flash_swap/currency_pairs&#x60;获取 | 
**SellAmount** | **string** | Amount to sell. It is required to choose one parameter between &#x60;sell_amount&#x60; and &#x60;buy_amount&#x60; | [optional] 
**BuyCurrency** | **string** | 买入的资产名称， 根据接口&#x60;查询支持闪兑的所有交易对列表 GET /flash_swap/currency_pairs&#x60;获取 | 
**BuyAmount** | **string** | Amount to buy. It is required to choose one parameter between &#x60;sell_amount&#x60; and &#x60;buy_amount&#x60; | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


