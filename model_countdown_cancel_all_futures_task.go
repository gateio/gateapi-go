/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Countdown cancel task detail
type CountdownCancelAllFuturesTask struct {
	// Countdown time, in seconds  At least 5 seconds, 0 means cancel the countdown
	Timeout int32 `json:"timeout"`
	// Futures contract
	Contract string `json:"contract,omitempty"`
}