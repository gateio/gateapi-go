/*
 * Gate API v4
 *
 * APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type FuturesLiquidate struct {
	// Liquidation time
	Time int64 `json:"time,omitempty"`
	// Futures contract
	Contract string `json:"contract,omitempty"`
	// Position leverage. Not returned in public endpoints.
	Leverage string `json:"leverage,omitempty"`
	// Position size
	Size int64 `json:"size,omitempty"`
	// Position margin. Not returned in public endpoints.
	Margin string `json:"margin,omitempty"`
	// Average entry price. Not returned in public endpoints.
	EntryPrice string `json:"entry_price,omitempty"`
	// Liquidation price. Not returned in public endpoints.
	LiqPrice string `json:"liq_price,omitempty"`
	// Mark price. Not returned in public endpoints.
	MarkPrice string `json:"mark_price,omitempty"`
	// Liquidation order ID. Not returned in public endpoints.
	OrderId int64 `json:"order_id,omitempty"`
	// Liquidation order price
	OrderPrice string `json:"order_price,omitempty"`
	// Liquidation order average taker price
	FillPrice string `json:"fill_price,omitempty"`
	// Liquidation order maker size
	Left int64 `json:"left,omitempty"`
}
