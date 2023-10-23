/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Total borrowed amount and pledged collateral amount by the user
type UserTotalAmount struct {
	// Total borrowing amount, calculated in USDT
	BorrowAmount string `json:"borrow_amount,omitempty"`
	// Total collateral amount, calculated in USDT
	CollateralAmount string `json:"collateral_amount,omitempty"`
}