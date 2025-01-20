/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type UnifiedLeverageConfig struct {
	// Current leverage ratio
	CurrentLeverage string `json:"current_leverage,omitempty"`
	// Minimum adjustable leverage ratio
	MinLeverage string `json:"min_leverage,omitempty"`
	// Maximum adjustable leverage ratio
	MaxLeverage string `json:"max_leverage,omitempty"`
	// Current liabilities
	Debit string `json:"debit,omitempty"`
	// Available Margin
	AvailableMargin string `json:"available_margin,omitempty"`
	// The current leverage you can choose is
	Borrowable string `json:"borrowable,omitempty"`
	// The maximum amount of margin that can be borrowed and the maximum amount of Uniloan that can be borrowed, whichever is smaller
	ExceptLeverageBorrowable string `json:"except_leverage_borrowable,omitempty"`
}
