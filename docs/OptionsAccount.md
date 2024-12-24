# OptionsAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | **int64** | User ID | [optional] 
**Total** | **string** | 账户余额 | [optional] 
**PositionValue** | **string** | 仓位价值，做多仓位价值为正，做空仓位价值为负 | [optional] 
**Equity** | **string** | 账户权益，账户余额与仓位价值的和 | [optional] 
**ShortEnabled** | **bool** | If the account is allowed to short | [optional] 
**MmpEnabled** | **bool** | 是否启用MMP | [optional] 
**LiqTriggered** | **bool** | 是否触发仓位强平 | [optional] 
**MarginMode** | **int32** | ｜ 保证金模式： - 0：经典现货保证金模式 - 1：跨币种保证金模式 - 2：组合保证金模式 | [optional] 
**UnrealisedPnl** | **string** | Unrealized PNL | [optional] 
**InitMargin** | **string** | Initial position margin | [optional] 
**MaintMargin** | **string** | Position maintenance margin | [optional] 
**OrderMargin** | **string** | Order margin of unfinished orders | [optional] 
**AskOrderMargin** | **string** | 未完成卖单的保证金 | [optional] 
**BidOrderMargin** | **string** | 未完成买单的保证金 | [optional] 
**Available** | **string** | Available balance to transfer out or trade | [optional] 
**Point** | **string** | POINT amount | [optional] 
**Currency** | **string** | Settle currency | [optional] 
**OrdersLimit** | **int32** | 未完成订单数量上限 | [optional] 
**PositionNotionalLimit** | **int64** | 名义价值上限，包含仓位以及未完成订单的名义价值 | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


