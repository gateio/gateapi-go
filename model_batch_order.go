/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Batch order details
type BatchOrder struct {
	// Order ID
	OrderId string `json:"order_id,omitempty"`
	// The custom data that the user remarked when amending the order
	AmendText string `json:"amend_text,omitempty"`
	// User defined information. If not empty, must follow the rules below:  1. prefixed with `t-` 2. no longer than 28 bytes without `t-` prefix 3. can only include 0-9, A-Z, a-z, underscore(_), hyphen(-) or dot(.)
	Text string `json:"text,omitempty"`
	// Whether the batch of orders succeeded
	Succeeded bool `json:"succeeded,omitempty"`
	// Error label, if any, otherwise an empty string
	Label string `json:"label,omitempty"`
	// Detailed error message, if any, otherwise an empty string
	Message string `json:"message,omitempty"`
	// Order ID
	Id string `json:"id,omitempty"`
	// Creation time of order
	CreateTime string `json:"create_time,omitempty"`
	// Last modification time of order
	UpdateTime string `json:"update_time,omitempty"`
	// Creation time of order (in milliseconds)
	CreateTimeMs int64 `json:"create_time_ms,omitempty"`
	// Last modification time of order (in milliseconds)
	UpdateTimeMs int64 `json:"update_time_ms,omitempty"`
	// Order status  - `open`: to be filled - `closed`: filled - `cancelled`: cancelled
	Status string `json:"status,omitempty"`
	// Currency pair
	CurrencyPair string `json:"currency_pair,omitempty"`
	// Order Type    - limit : Limit Order - market : Market Order
	Type string `json:"type,omitempty"`
	// Account type, spot - spot account, margin - leveraged account, unified - unified account
	Account string `json:"account,omitempty"`
	// Order side
	Side string `json:"side,omitempty"`
	// Trade amount
	Amount string `json:"amount,omitempty"`
	// Order price
	Price string `json:"price,omitempty"`
	// Time in force  - gtc: GoodTillCancelled - ioc: ImmediateOrCancelled, taker only - poc: PendingOrCancelled, makes a post-only order that always enjoys a maker fee - fok: FillOrKill, fill either completely or none
	TimeInForce string `json:"time_in_force,omitempty"`
	// Amount to display for the iceberg order. Null or 0 for normal orders.  Hiding all amount is not supported.
	Iceberg string `json:"iceberg,omitempty"`
	// Used in margin or cross margin trading to allow automatic loan of insufficient amount if balance is not enough.
	AutoBorrow bool `json:"auto_borrow,omitempty"`
	// Enable or disable automatic repayment for automatic borrow loan generated by cross margin order. Default is disabled. Note that:  1. This field is only effective for cross margin orders. Margin account does not support setting auto repayment for orders. 2. `auto_borrow` and `auto_repay` can be both set to true in one order.
	AutoRepay bool `json:"auto_repay,omitempty"`
	// Amount left to fill
	Left string `json:"left,omitempty"`
	// Amount traded to fill
	FilledAmount string `json:"filled_amount,omitempty"`
	// Total filled in quote currency. Deprecated in favor of `filled_total`
	FillPrice string `json:"fill_price,omitempty"`
	// Total filled in quote currency
	FilledTotal string `json:"filled_total,omitempty"`
	// Average fill price
	AvgDealPrice string `json:"avg_deal_price,omitempty"`
	// Fee deducted
	Fee string `json:"fee,omitempty"`
	// Fee currency unit
	FeeCurrency string `json:"fee_currency,omitempty"`
	// Points used to deduct fee
	PointFee string `json:"point_fee,omitempty"`
	// GT used to deduct fee
	GtFee string `json:"gt_fee,omitempty"`
	// Whether GT fee discount is used
	GtDiscount bool `json:"gt_discount,omitempty"`
	// Rebated fee
	RebatedFee string `json:"rebated_fee,omitempty"`
	// Rebated fee currency unit
	RebatedFeeCurrency string `json:"rebated_fee_currency,omitempty"`
	// Orders between users in the same `stp_id` group are not allowed to be self-traded  1. If the `stp_id` of two orders being matched is non-zero and equal, they will not be executed. Instead, the corresponding strategy will be executed based on the `stp_act` of the taker. 2. `stp_id` returns `0` by default for orders that have not been set for `STP group`
	StpId int32 `json:"stp_id,omitempty"`
	// Self-Trading Prevention Action. Users can use this field to set self-trade prevetion strategies  1. After users join the `STP Group`, he can pass `stp_act` to limit the user's self-trade prevetion strategy. If `stp_act` is not passed, the default is `cn` strategy。 2. When the user does not join the `STP group`, an error will be returned when passing the `stp_act` parameter。 3. If the user did not use 'stp_act' when placing the order, 'stp_act' will return '-'  - cn: Cancel newest, Cancel new orders and keep old ones - co: Cancel oldest, Cancel old orders and keep new ones - cb: Cancel both, Both old and new orders will be cancelled
	StpAct string `json:"stp_act,omitempty"`
	// How the order was finished.  - open: processing - filled: filled totally - cancelled: manually cancelled - ioc: time in force is `IOC`, finish immediately - stp: cancelled because self trade prevention
	FinishAs string `json:"finish_as,omitempty"`
}
