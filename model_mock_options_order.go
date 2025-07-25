/*
 * Gate API
 *
 * Welcome to Gate API  APIv4 provides operations related to spot, margin, and contract trading, including public interfaces for querying market data and authenticated private interfaces for implementing API-based automated trading.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Option orders.
type MockOptionsOrder struct {
	// Option name, currently only supports options for BTC and ETH with USDT.
	OptionsName string `json:"options_name"`
	// Initial order quantity, not involved in actual calculation.
	Size string `json:"size"`
	// Unfilled contract quantity, involved in actual calculation.
	Left string `json:"left"`
}
