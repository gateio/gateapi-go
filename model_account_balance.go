/*
 * Gate API
 *
 * Welcome to Gate API  APIv4 provides operations related to spot, margin, and contract trading, including public interfaces for querying market data and authenticated private interfaces for implementing API-based automated trading.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Total balances calculated with specified currency unit.
type AccountBalance struct {
	// Account total balance amount.
	Amount string `json:"amount,omitempty"`
	// Currency.
	Currency string `json:"currency,omitempty"`
	// Unrealised_pnl, this field will only appear in futures, options, delivery, and total accounts
	UnrealisedPnl string `json:"unrealised_pnl,omitempty"`
	// Borrowed，this field will only appear in margin and cross_margin accounts.
	Borrowed string `json:"borrowed,omitempty"`
}
