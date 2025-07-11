/*
 * Gate API v4
 *
 * Welcome to Gate API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type MultiLoanRepayItem struct {
	// Repayment currency
	Currency string `json:"currency,omitempty"`
	// Size
	Amount string `json:"amount,omitempty"`
	// Repayment method, set to true for full repayment, false for partial repayment.
	RepaidAll bool `json:"repaid_all,omitempty"`
}
