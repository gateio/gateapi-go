/*
 * Gate API
 *
 * Welcome to Gate API  APIv4 provides operations related to spot, margin, and contract trading, including public interfaces for querying market data and authenticated private interfaces for implementing API-based automated trading.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Borrow or repay.
type CreateUniLoan struct {
	// Currency.
	Currency string `json:"currency"`
	// type: borrow - borrow, repay - repay.
	Type string `json:"type"`
	// The amount of lending or repaying.
	Amount string `json:"amount"`
	// Full repayment. Repay operation only. If the value is `true`, the amount will be ignored and repaid in full.
	RepaidAll bool `json:"repaid_all,omitempty"`
	// Currency pair.
	CurrencyPair string `json:"currency_pair"`
}
