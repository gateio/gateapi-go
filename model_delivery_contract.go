/*
 * Gate API
 *
 * Welcome to Gate API  APIv4 provides operations related to spot, margin, and contract trading, including public interfaces for querying market data and authenticated private interfaces for implementing API-based automated trading.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Futures contract details.
type DeliveryContract struct {
	// Futures contract.
	Name string `json:"name,omitempty"`
	// Underlying.
	Underlying string `json:"underlying,omitempty"`
	// Cycle type, e.g. WEEKLY, QUARTERLY.
	Cycle string `json:"cycle,omitempty"`
	// Futures contract type.
	Type string `json:"type,omitempty"`
	// Multiplier used in converting from invoicing to settlement currency.
	QuantoMultiplier string `json:"quanto_multiplier,omitempty"`
	// Minimum leverage.
	LeverageMin string `json:"leverage_min,omitempty"`
	// Maximum leverage.
	LeverageMax string `json:"leverage_max,omitempty"`
	// Maintenance rate of margin.
	MaintenanceRate string `json:"maintenance_rate,omitempty"`
	// Mark price type, internal - based on internal trading, external index price
	MarkType string `json:"mark_type,omitempty"`
	// Current mark price.
	MarkPrice string `json:"mark_price,omitempty"`
	// Current index price.
	IndexPrice string `json:"index_price,omitempty"`
	// Last trading price.
	LastPrice string `json:"last_price,omitempty"`
	// Maker fee rate, where negative means rebate.
	MakerFeeRate string `json:"maker_fee_rate,omitempty"`
	// Taker fee rate.
	TakerFeeRate string `json:"taker_fee_rate,omitempty"`
	// Minimum order price increment.
	OrderPriceRound string `json:"order_price_round,omitempty"`
	// Minimum mark price increment.
	MarkPriceRound string `json:"mark_price_round,omitempty"`
	// Fair basis rate.
	BasisRate string `json:"basis_rate,omitempty"`
	// Fair basis value.
	BasisValue string `json:"basis_value,omitempty"`
	// Funding used for calculating impact bid, ask price.
	BasisImpactValue string `json:"basis_impact_value,omitempty"`
	// Settle price.
	SettlePrice string `json:"settle_price,omitempty"`
	// Settle price update interval.
	SettlePriceInterval int32 `json:"settle_price_interval,omitempty"`
	// Settle price update duration in seconds.
	SettlePriceDuration int32 `json:"settle_price_duration,omitempty"`
	// Contract expiry timestamp.
	ExpireTime int64 `json:"expire_time,omitempty"`
	// Risk limit base.
	RiskLimitBase string `json:"risk_limit_base,omitempty"`
	// Step of adjusting risk limit.
	RiskLimitStep string `json:"risk_limit_step,omitempty"`
	// Maximum risk limit the contract allowed.
	RiskLimitMax string `json:"risk_limit_max,omitempty"`
	// Minimum order size the contract allowed.
	OrderSizeMin int64 `json:"order_size_min,omitempty"`
	// Maximum order size the contract allowed.
	OrderSizeMax int64 `json:"order_size_max,omitempty"`
	// deviation between order price and current index price. If price of an order is denoted as order_price, it must meet the following condition:   abs(order_price - mark_price) <= mark_price * order_price_deviate
	OrderPriceDeviate string `json:"order_price_deviate,omitempty"`
	// Referral fee rate discount.
	RefDiscountRate string `json:"ref_discount_rate,omitempty"`
	// Referrer commission rate.
	RefRebateRate string `json:"ref_rebate_rate,omitempty"`
	// Current orderbook ID.
	OrderbookId int64 `json:"orderbook_id,omitempty"`
	// Current trade ID.
	TradeId int64 `json:"trade_id,omitempty"`
	// Historical accumulated trade size.
	TradeSize int64 `json:"trade_size,omitempty"`
	// Current total long position size.
	PositionSize int64 `json:"position_size,omitempty"`
	// Last changed time of configuration.
	ConfigChangeTime float64 `json:"config_change_time,omitempty"`
	// Contract is delisting.
	InDelisting bool `json:"in_delisting,omitempty"`
	// Maximum number of open orders.
	OrdersLimit int32 `json:"orders_limit,omitempty"`
}
