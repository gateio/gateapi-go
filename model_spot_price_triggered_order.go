/*
 * Gate API
 *
 * Welcome to Gate API  APIv4 provides operations related to spot, margin, and contract trading, including public interfaces for querying market data and authenticated private interfaces for implementing API-based automated trading.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Spot order detail.
type SpotPriceTriggeredOrder struct {
	Trigger SpotPriceTrigger  `json:"trigger"`
	Put     SpotPricePutOrder `json:"put"`
	// Auto order ID.
	Id int64 `json:"id,omitempty"`
	// User ID.
	User int32 `json:"user,omitempty"`
	// Currency pair.
	Market string `json:"market"`
	// Creation time.
	Ctime int64 `json:"ctime,omitempty"`
	// Finished time.
	Ftime int64 `json:"ftime,omitempty"`
	// ID of the newly created order on condition triggered.
	FiredOrderId int64 `json:"fired_order_id,omitempty"`
	// Status  - open: open - cancelled: being manually cancelled - finish: successfully executed - failed: failed to execute - expired - expired
	Status string `json:"status,omitempty"`
	// Additional remarks on how the order was finished.
	Reason string `json:"reason,omitempty"`
}
