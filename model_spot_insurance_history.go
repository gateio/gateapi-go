/*
 * Gate API
 *
 * Welcome to Gate API  APIv4 provides operations related to spot, margin, and contract trading, including public interfaces for querying market data and authenticated private interfaces for implementing API-based automated trading.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type SpotInsuranceHistory struct {
	// Currency.
	Currency string `json:"currency,omitempty"`
	// balance.
	Balance string `json:"balance,omitempty"`
	// Creation time, timestamp, milliseconds.
	Time int64 `json:"time,omitempty"`
}
