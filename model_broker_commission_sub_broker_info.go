/*
 * Gate API v4
 *
 * Welcome to Gate API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// The sub broker info
type BrokerCommissionSubBrokerInfo struct {
	// The sub broker user ID
	UserId int64 `json:"user_id,omitempty"`
	// The sub broker original commission rate
	OriginalCommissionRate string `json:"original_commission_rate,omitempty"`
	// The sub broker relative commission rate
	RelativeCommissionRate string `json:"relative_commission_rate,omitempty"`
	// The sub broker actual commission rate
	CommissionRate string `json:"commission_rate,omitempty"`
}
