/*
 * Gate API
 *
 * Welcome to Gate API  APIv4 provides operations related to spot, margin, and contract trading, including public interfaces for querying market data and authenticated private interfaces for implementing API-based automated trading.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Parameters of flash swap order creation.
type FlashSwapPreviewRequest struct {
	// The name of the asset being sold, as obtained from the \"GET /flash_swap/currency_pairs\" API, which retrieves a list of supported flash swap currency pairs.
	SellCurrency string `json:"sell_currency"`
	// Amount to sell. It is required to choose one parameter between `sell_amount` and `buy_amount`
	SellAmount string `json:"sell_amount,omitempty"`
	// The name of the asset being purchased, as obtained from the \"GET /flash_swap/currency_pairs\" API, which provides a list of supported flash swap currency pairs.
	BuyCurrency string `json:"buy_currency"`
	// Amount to buy. It is required to choose one parameter between `sell_amount` and `buy_amount`
	BuyAmount string `json:"buy_amount,omitempty"`
}
