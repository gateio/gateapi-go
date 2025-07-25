/*
 * Gate API
 *
 * Welcome to Gate API  APIv4 provides operations related to spot, margin, and contract trading, including public interfaces for querying market data and authenticated private interfaces for implementing API-based automated trading.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type OptionsAccountBook struct {
	// Change time.
	Time float64 `json:"time,omitempty"`
	// Amount changed (USDT).
	Change string `json:"change,omitempty"`
	// Account total balance after change (USDT).
	Balance string `json:"balance,omitempty"`
	// Changing Type: - dnw: Deposit & Withdraw - prem: Trading premium - fee: Trading fee - refr: Referrer rebate - point_dnw: point_fee: POINT Trading fee - point_refr: POINT Referrer rebate
	Type string `json:"type,omitempty"`
	// custom text.
	Text string `json:"text,omitempty"`
}
