/*
 * Gate API
 *
 * Welcome to Gate API  APIv4 provides operations related to spot, margin, and contract trading, including public interfaces for querying market data and authenticated private interfaces for implementing API-based automated trading.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Currency Quota.
type CurrencyQuota struct {
	// Currency.
	Currency string `json:"currency,omitempty"`
	// Currency Index Price.
	IndexPrice string `json:"index_price,omitempty"`
	// Minimum borrowing/collateral quota for the currency.
	MinQuota string `json:"min_quota,omitempty"`
	// Remaining borrowing/collateral limit for the currency.
	LeftQuota string `json:"left_quota,omitempty"`
	// Remaining currency limit converted to USDT.
	LeftQuoteUsdt string `json:"left_quote_usdt,omitempty"`
}
