/*
 * Gate API v4
 *
 * Welcome to Gate API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Borrow or repay
type UnifiedLoan struct {
	// Currency
	Currency string `json:"currency"`
	// type: borrow - borrow, repay - repay
	Type string `json:"type"`
	// The amount of lending or repaying
	Amount string `json:"amount"`
	// Full repayment is solely for repayment operations. When set to 'true,' it overrides the 'amount,' allowing for direct full repayment.
	RepaidAll bool `json:"repaid_all,omitempty"`
	// User defined custom ID
	Text string `json:"text,omitempty"`
}
