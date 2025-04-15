/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Spot order details
type Order struct {
	// Order ID
	Id string `json:"id,omitempty"`
	// User defined information. If not empty, must follow the rules below:  1. prefixed with `t-` 2. no longer than 28 bytes without `t-` prefix 3. can only include 0-9, A-Z, a-z, underscore(_), hyphen(-) or dot(.)  Besides user defined information, reserved contents are listed below, denoting how the order is created:  - 101: from android - 102: from IOS - 103: from IPAD - 104: from webapp - 3: from web - 2: from apiv2 - apiv4: from apiv4
	Text string `json:"text,omitempty"`
	// The custom data that the user remarked when amending the order
	AmendText string `json:"amend_text,omitempty"`
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
	// FinishAs final status  - `open`: processing - `filled`: filled totally - `cancelled`: manually cancelled - `ioc`: time in force is IOC, finish immediately - `stp`: cancelled because self trade prevention
	FinishAs string `json:"finish_as,omitempty"`
	// Currency pair
	CurrencyPair string `json:"currency_pair"`
	// Order Type    - limit : Limit Order - market : Market Order
	Type string `json:"type,omitempty"`
	// Account type, spot - spot account, margin - leveraged account, unified - unified account
	Account string `json:"account,omitempty"`
	// Order side
	Side string `json:"side"`
	// When `type` is limit, it refers to base currency.  For instance, `BTC_USDT` means `BTC`  When `type` is `market`, it refers to different currency according to `side`  - `side` : `buy` means quote currency, `BTC_USDT` means `USDT` - `side` : `sell` means base currency，`BTC_USDT` means `BTC`
	Amount string `json:"amount"`
	// Price can't be empty when `type`= `limit`
	Price string `json:"price,omitempty"`
	// Time in force  - gtc: GoodTillCancelled - ioc: ImmediateOrCancelled, taker only - poc: PendingOrCancelled, makes a post-only order that always enjoys a maker fee - fok: FillOrKill, fill either completely or none Only `ioc` and `fok` are supported when `type`=`market`
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
	// GT used to deduct maker fee
	GtMakerFee string `json:"gt_maker_fee,omitempty"`
	// GT used to deduct taker fee
	GtTakerFee string `json:"gt_taker_fee,omitempty"`
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
	// Order completion statuses include:  - open: Awaiting processing - filled: Fully filled - cancelled: Cancelled by user - liquidate_cancelled: Cancelled due to liquidation - small: Order quantity too small - depth_not_enough: Cancelled due to insufficient market depth - trader_not_enough: Cancelled due to insufficient counterparty - ioc: Not immediately filled because tif is set to ioc - poc: Not met the order strategy because tif is set to poc - fok: Not fully filled immediately because tif is set to fok - stp: Cancelled due to self-trade prevention - unknown: Unknown
	FinishAs string `json:"finish_as,omitempty"`
	// Processing Mode: When placing an order, different fields are returned based on action_mode. This field is only valid during the request and is not included in the response result ACK: Asynchronous mode, only returns key order fields RESULT: No clearing information FULL: Full mode (default)
	ActionMode string `json:"action_mode,omitempty"`
}
