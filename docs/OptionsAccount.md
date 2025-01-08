# OptionsAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | **int64** | User ID | [optional] 
**Total** | **string** | Account balance | [optional] 
**PositionValue** | **string** | Position value, long position value is positive, short position value is negative | [optional] 
**Equity** | **string** | Account equity, the sum of account balance and position value | [optional] 
**ShortEnabled** | **bool** | If the account is allowed to short | [optional] 
**MmpEnabled** | **bool** | Whether to enable MMP | [optional] 
**LiqTriggered** | **bool** | Whether to trigger position liquidation | [optional] 
**MarginMode** | **int32** | ｜ 保证金模式： - 0：经典现货保证金模式 - 1：跨币种保证金模式 - 2：组合保证金模式 | [optional] 
**UnrealisedPnl** | **string** | Unrealized PNL | [optional] 
**InitMargin** | **string** | Initial position margin | [optional] 
**MaintMargin** | **string** | Position maintenance margin | [optional] 
**OrderMargin** | **string** | Order margin of unfinished orders | [optional] 
**AskOrderMargin** | **string** | Margin for outstanding sell orders | [optional] 
**BidOrderMargin** | **string** | Margin for outstanding buy orders | [optional] 
**Available** | **string** | Available balance to transfer out or trade | [optional] 
**Point** | **string** | POINT amount | [optional] 
**Currency** | **string** | Settle currency | [optional] 
**OrdersLimit** | **int32** | Maximum number of outstanding orders | [optional] 
**PositionNotionalLimit** | **int64** | Notional value upper limit, including the nominal value of positions and outstanding orders | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


