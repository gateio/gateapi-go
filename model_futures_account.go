/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type FuturesAccount struct {
	// total is the balance after the user's accumulated deposit, withdraw, profit and loss (including realized profit and loss, fund, fee and referral rebate), excluding unrealized profit and loss.  total = SUM(history_dnw, history_pnl, history_fee, history_refr, history_fund)
	Total string `json:"total,omitempty"`
	// Unrealized PNL
	UnrealisedPnl string `json:"unrealised_pnl,omitempty"`
	// Position margin
	PositionMargin string `json:"position_margin,omitempty"`
	// Order margin of unfinished orders
	OrderMargin string `json:"order_margin,omitempty"`
	// The available balance for transferring or trading(including bonus.  Bonus can't be withdrawn. The transfer amount needs to deduct the bonus)
	Available string `json:"available,omitempty"`
	// POINT amount
	Point string `json:"point,omitempty"`
	// Settle currency
	Currency string `json:"currency,omitempty"`
	// Whether dual mode is enabled
	InDualMode bool `json:"in_dual_mode,omitempty"`
	// Whether portfolio margin account mode is enabled
	EnableCredit bool `json:"enable_credit,omitempty"`
	// Initial margin position, applicable to the portfolio margin account model
	PositionInitialMargin string `json:"position_initial_margin,omitempty"`
	// Maintenance margin position, applicable to the portfolio margin account model
	MaintenanceMargin string `json:"maintenance_margin,omitempty"`
	// Perpetual Contract Bonus
	Bonus   string                `json:"bonus,omitempty"`
	History FuturesAccountHistory `json:"history,omitempty"`
}
