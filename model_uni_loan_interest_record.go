/*
 * Gate API
 *
 * Welcome to Gate API  APIv4 provides operations related to spot, margin, and contract trading, including public interfaces for querying market data and authenticated private interfaces for implementing API-based automated trading.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Interest record.
type UniLoanInterestRecord struct {
	// Currency name.
	Currency string `json:"currency,omitempty"`
	// Currency pair.
	CurrencyPair string `json:"currency_pair,omitempty"`
	// Actual rate.
	ActualRate string `json:"actual_rate,omitempty"`
	// Interest.
	Interest string `json:"interest,omitempty"`
	// Status: 0 - fail, 1 - success.
	Status int32 `json:"status,omitempty"`
	// Type, platform - platform，margin - margin.
	Type string `json:"type,omitempty"`
	// Created time.
	CreateTime int64 `json:"create_time,omitempty"`
}
