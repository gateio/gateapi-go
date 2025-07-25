/*
 * Gate API
 *
 * Welcome to Gate API  APIv4 provides operations related to spot, margin, and contract trading, including public interfaces for querying market data and authenticated private interfaces for implementing API-based automated trading.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type SubUserMode struct {
	// User ID.
	UserId int64 `json:"user_id,omitempty"`
	// Is it a unified account?.
	IsUnified bool `json:"is_unified,omitempty"`
	// Unified account mode： - `classic`: Classic account mode - `multi_currency`: Multi-currency margin mode - `portfolio`: Portfolio margin mode
	Mode string `json:"mode,omitempty"`
}
