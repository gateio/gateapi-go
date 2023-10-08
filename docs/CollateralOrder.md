# CollateralOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **int64** | 订单id | [optional] 
**CollateralCurrency** | **string** | 质押币种 | [optional] 
**CollateralAmount** | **string** | 质押数量 | [optional] 
**BorrowCurrency** | **string** | 借款币种 | [optional] 
**BorrowAmount** | **string** | 借款数量 | [optional] 
**RepaidAmount** | **string** | 已还款数量 | [optional] 
**RepaidPrincipal** | **string** | 已还本金 | [optional] 
**RepaidInterest** | **string** | 已还利息 | [optional] 
**InitLtv** | **string** | 初始质押率 | [optional] 
**CurrentLtv** | **string** | 当前质押率 | [optional] 
**LiquidateLtv** | **string** | 平仓质押率 | [optional] 
**Status** | **string** | 订单状态: - initial: 下单初始状态 - collateral_deducted: 扣除质押物成功 - collateral_returning: 放款失败-待退回质押物 - lent: 放款成功 - repaying: 还款中 - liquidating: 平仓中 - finished: 已完成 - closed_liquidated: 已结束-平仓还款结束 | [optional] 
**BorrowTime** | **int64** | 借款时间，时间戳，单位秒 | [optional] 
**LeftRepayTotal** | **string** | 待还本息（待还本金+待还利息） | [optional] 
**LeftRepayPrincipal** | **string** | 待还本金 | [optional] 
**LeftRepayInterest** | **string** | 待还利息 | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


