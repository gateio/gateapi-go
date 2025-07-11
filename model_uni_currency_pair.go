/*
 * Gate API v4
 *
 * Welcome to Gate API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Currency pair of the loan
type UniCurrencyPair struct {
	// Currency pair
	CurrencyPair string `json:"currency_pair,omitempty"`
	// Minimum borrow amount of base currency
	BaseMinBorrowAmount string `json:"base_min_borrow_amount,omitempty"`
	// Minimum borrow amount of quote currency
	QuoteMinBorrowAmount string `json:"quote_min_borrow_amount,omitempty"`
	// Position leverage
	Leverage string `json:"leverage,omitempty"`
}
