/*
 * Gate API v4
 *
 * APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * API version: 4.5.1
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type RepayRequest struct {
	// Currency pair
	CurrencyPair string `json:"currency_pair"`
	// Loan currency
	Currency string `json:"currency"`
	// Repay mode. all - repay all; partial - repay only some portion
	Mode string `json:"mode"`
	// Repay amount. Required in `partial` mode
	Amount string `json:"amount,omitempty"`
}
