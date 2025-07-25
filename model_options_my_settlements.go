/*
 * Gate API
 *
 * Welcome to Gate API  APIv4 provides operations related to spot, margin, and contract trading, including public interfaces for querying market data and authenticated private interfaces for implementing API-based automated trading.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type OptionsMySettlements struct {
	// Settlement time.
	Time float64 `json:"time,omitempty"`
	// Underlying.
	Underlying string `json:"underlying,omitempty"`
	// Options contract name.
	Contract string `json:"contract,omitempty"`
	// Strike price (quote currency).
	StrikePrice string `json:"strike_price,omitempty"`
	// Settlement price (quote currency).
	SettlePrice string `json:"settle_price,omitempty"`
	// Size.
	Size int64 `json:"size,omitempty"`
	// Settlement profit (quote currency).
	SettleProfit string `json:"settle_profit,omitempty"`
	// Fee (quote currency).
	Fee string `json:"fee,omitempty"`
	// The accumulated profit and loss of opening a position, including premium, fee, settlement profit, etc. (quote currency)
	RealisedPnl string `json:"realised_pnl,omitempty"`
}
