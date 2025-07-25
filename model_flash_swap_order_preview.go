/*
 * Gate API
 *
 * Welcome to Gate API  APIv4 provides operations related to spot, margin, and contract trading, including public interfaces for querying market data and authenticated private interfaces for implementing API-based automated trading.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Initiate a flash swap order preview.
type FlashSwapOrderPreview struct {
	// Preview result ID.
	PreviewId string `json:"preview_id,omitempty"`
	// Name of the sold asset,  Refer to the interface Query the list of currencies supported for flash swap GET /flash_swap/currenciesto obtain
	SellCurrency string `json:"sell_currency,omitempty"`
	// Amount to sell.
	SellAmount string `json:"sell_amount,omitempty"`
	// Name of the purchased asset,  Refer to the interface Query the list of currencies supported for flash swap GET /flash_swap/currenciesto obtain
	BuyCurrency string `json:"buy_currency,omitempty"`
	// Amount to buy.
	BuyAmount string `json:"buy_amount,omitempty"`
	// Price.
	Price string `json:"price,omitempty"`
}
