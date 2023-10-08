# CrossMarginBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | **string** | Available amount | [optional] 
**Freeze** | **string** | Locked amount | [optional] 
**Borrowed** | **string** | Borrowed amount | [optional] 
**Interest** | **string** | Unpaid interests | [optional] 
**NegativeLiab** | **string** | Negative Liabilities. Formula:Min[available+total+unrealized_pnl,0] | [optional] 
**FuturesPosLiab** | **string** | Borrowing to Open Positions in Futures | [optional] 
**Equity** | **string** | Equity. Formula: available + freeze - borrowed + futures account&#39;s total + unrealized_pnl | [optional] 
**TotalFreeze** | **string** | Total freeze. Formula: freeze + position_initial_margin + order_margin | [optional] 
**TotalLiab** | **string** | Total liabilities. Formula: Max[Abs[Min[quity - total_freeze,0], borrowed]] - futures_pos_liab | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


