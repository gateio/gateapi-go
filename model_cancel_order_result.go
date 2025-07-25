/*
 * Gate API
 *
 * Welcome to Gate API  APIv4 provides operations related to spot, margin, and contract trading, including public interfaces for querying market data and authenticated private interfaces for implementing API-based automated trading.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Order cancellation result.
type CancelOrderResult struct {
	// Order currency pair.
	CurrencyPair string `json:"currency_pair,omitempty"`
	// Order ID.
	Id string `json:"id,omitempty"`
	// Custom order information.
	Text string `json:"text,omitempty"`
	// Whether cancellation succeeded.
	Succeeded bool `json:"succeeded,omitempty"`
	// Error label when failed to cancel the order; emtpy if succeeded.
	Label string `json:"label,omitempty"`
	// Error message when failed to cancel the order; empty if succeeded.
	Message string `json:"message,omitempty"`
	// Default is empty (deprecated).
	Account string `json:"account,omitempty"`
}
