/*
 * Gate API
 *
 * Welcome to Gate API  APIv4 provides operations related to spot, margin, and contract trading, including public interfaces for querying market data and authenticated private interfaces for implementing API-based automated trading.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type MarginTransferable struct {
	// Currency detail.
	Currency string `json:"currency,omitempty"`
	// Currency pair.
	CurrencyPair string `json:"currency_pair,omitempty"`
	// Max transferable amount.
	Amount string `json:"amount,omitempty"`
}
