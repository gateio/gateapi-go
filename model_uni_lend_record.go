/*
 * Gate API
 *
 * Welcome to Gate API  APIv4 provides operations related to spot, margin, and contract trading, including public interfaces for querying market data and authenticated private interfaces for implementing API-based automated trading.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Interest Record.
type UniLendRecord struct {
	// Currency name.
	Currency string `json:"currency,omitempty"`
	// current amount.
	Amount string `json:"amount,omitempty"`
	// Last wallet amount.
	LastWalletAmount string `json:"last_wallet_amount,omitempty"`
	// Last lent amount.
	LastLentAmount string `json:"last_lent_amount,omitempty"`
	// Last frozen amount.
	LastFrozenAmount string `json:"last_frozen_amount,omitempty"`
	// Record type: lend - lend, redeem - redeem.
	Type string `json:"type,omitempty"`
	// Created time.
	CreateTime int64 `json:"create_time,omitempty"`
}
