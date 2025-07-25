/*
 * Gate API
 *
 * Welcome to Gate API  APIv4 provides operations related to spot, margin, and contract trading, including public interfaces for querying market data and authenticated private interfaces for implementing API-based automated trading.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Flash swap order.
type FlashSwapOrder struct {
	// Flash swap order ID.
	Id int64 `json:"id,omitempty"`
	// Creation time of order (in milliseconds).
	CreateTime int64 `json:"create_time,omitempty"`
	// User ID.
	UserId int64 `json:"user_id,omitempty"`
	// Currency to sell.
	SellCurrency string `json:"sell_currency,omitempty"`
	// Amount to sell.
	SellAmount string `json:"sell_amount,omitempty"`
	// Currency to buy.
	BuyCurrency string `json:"buy_currency,omitempty"`
	// Amount to buy.
	BuyAmount string `json:"buy_amount,omitempty"`
	// Price.
	Price string `json:"price,omitempty"`
	// Flash swap order status  `1` - success `2` - failure
	Status int32 `json:"status,omitempty"`
}
