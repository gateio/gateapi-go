/*
 * Gate API
 *
 * Welcome to Gate API  APIv4 provides operations related to spot, margin, and contract trading, including public interfaces for querying market data and authenticated private interfaces for implementing API-based automated trading.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Currency detail.
type UniCurrency struct {
	// Currency name.
	Currency string `json:"currency,omitempty"`
	// The minimum lending amount, in the unit of the currency.
	MinLendAmount string `json:"min_lend_amount,omitempty"`
	// The total maximum lending amount, in USDT.
	MaxLendAmount string `json:"max_lend_amount,omitempty"`
	// Maximum rate (Hourly).
	MaxRate string `json:"max_rate,omitempty"`
	// Minimum rate (Hourly).
	MinRate string `json:"min_rate,omitempty"`
}
