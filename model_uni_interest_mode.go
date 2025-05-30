/*
 * Gate API v4
 *
 * Welcome to Gate API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Lend & Earn interest reinvestment toggle
type UniInterestMode struct {
	// Currency
	Currency string `json:"currency"`
	// Interest toggle settings, true - interest reinvestment, false - regular dividend
	Status bool `json:"status"`
}
