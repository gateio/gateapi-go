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
**MaintenanceMargin** | **string** | The maintenance deposit occupied by the position is suitable for the new classic account margin model and unified account model | [optional] 
**Bonus** | **string** | Perpetual Contract Bonus | [optional] 
**EnableEvolvedClassic** | **bool** | Classic account margin mode, true-new mode, false-old mode | [optional] 
**CrossOrderMargin** | **string** | Full -warehouse hanging order deposit, suitable for the new classic account margin model | [optional] 
**CrossInitialMargin** | **string** | The initial security deposit of the full warehouse is suitable for the new classic account margin model | [optional] 
**CrossMaintenanceMargin** | **string** | Maintain deposit in full warehouse, suitable for new classic account margin models | [optional] 
**CrossUnrealisedPnl** | **string** | The full warehouse does not achieve profit and loss, suitable for the new classic account margin model | [optional] 
**CrossAvailable** | **string** | Full warehouse available amount, suitable for the new classic account margin model | [optional] 
**CrossMarginBalance** | **string** | Full margin balance, suitable for the new classic account margin model | [optional] 
**CrossMmr** | **string** | Maintain margin ratio for the full position, suitable for the new classic account margin model | [optional] 
**CrossImr** | **string** | The initial margin rate of the full position is suitable for the new classic account margin model | [optional] 
**IsolatedPositionMargin** | **string** | Ware -position margin, suitable for the new classic account margin model | [optional] 
**EnableNewDualMode** | **bool** | Whether to open a new two-way position mode | [optional] 
**MarginMode** | **int32** | Margin mode, 0-classic margin mode, 1-cross-currency margin mode, 2-combined margin mode | [optional] 
**EnableTieredMm** | **bool** | Whether to enable tiered maintenance margin calculation | [optional] 
**History** | [**FuturesAccountHistory**](FuturesAccount_history.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


