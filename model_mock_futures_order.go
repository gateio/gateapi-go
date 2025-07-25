/*
 * Gate API
 *
 * Welcome to Gate API  APIv4 provides operations related to spot, margin, and contract trading, including public interfaces for querying market data and authenticated private interfaces for implementing API-based automated trading.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Futures order.
type MockFuturesOrder struct {
	// Futures name, currently only supports perpetual futures for BTC and ETH with USDT.
	Contract string `json:"contract"`
	// Futures quantity, representing the initial order quantity, not involved in actual settlement.
	Size string `json:"size"`
	// Unfilled contract quantity, involved in actual calculation.
	Left string `json:"left"`
}
