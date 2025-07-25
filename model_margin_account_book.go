/*
 * Gate API
 *
 * Welcome to Gate API  APIv4 provides operations related to spot, margin, and contract trading, including public interfaces for querying market data and authenticated private interfaces for implementing API-based automated trading.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type MarginAccountBook struct {
	// Balance change record ID.
	Id string `json:"id,omitempty"`
	// Balance changed timestamp.
	Time string `json:"time,omitempty"`
	// The timestamp of the change (in milliseconds).
	TimeMs int64 `json:"time_ms,omitempty"`
	// Currency changed.
	Currency string `json:"currency,omitempty"`
	// Account currency pair.
	CurrencyPair string `json:"currency_pair,omitempty"`
	// Amount changed. Positive value means transferring in, while negative out.
	Change string `json:"change,omitempty"`
	// Balance after change.
	Balance string `json:"balance,omitempty"`
	// Account book type. Please refer to [account book type](#accountbook-type) for more detail
	Type string `json:"type,omitempty"`
}
