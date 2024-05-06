# FuturesAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **string** | total is the balance after the user&#39;s accumulated deposit, withdraw, profit and loss (including realized profit and loss, fund, fee and referral rebate), excluding unrealized profit and loss.  total &#x3D; SUM(history_dnw, history_pnl, history_fee, history_refr, history_fund) | [optional] 
**UnrealisedPnl** | **string** | Unrealized PNL | [optional] 
**PositionMargin** | **string** | Position margin | [optional] 
**OrderMargin** | **string** | Order margin of unfinished orders | [optional] 
**Available** | **string** | The available balance for transferring or trading(including bonus.  Bonus can&#39;t be be withdrawn. The transfer amount needs to deduct the bonus) | [optional] 
**Point** | **string** | POINT amount | [optional] 
**Currency** | **string** | Settle currency | [optional] 
**InDualMode** | **bool** | Whether dual mode is enabled | [optional] 
**EnableCredit** | **bool** | Whether portfolio margin account mode is enabled | [optional] 
**PositionInitialMargin** | **string** | Initial margin position, applicable to the portfolio margin account model | [optional] 
**MaintenanceMargin** | **string** | Maintenance margin position, applicable to the portfolio margin account model | [optional] 
**Bonus** | **string** | Perpetual Contract Bonus | [optional] 
**EnableEvolvedClassic** | **bool** | Classic account margin mode, true - enable new mode, false - revert to old mode. | [optional] 
**History** | [**FuturesAccountHistory**](FuturesAccount_history.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


