/*
 * Gate API
 *
 * Welcome to Gate API  APIv4 provides operations related to spot, margin, and contract trading, including public interfaces for querying market data and authenticated private interfaces for implementing API-based automated trading.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type BrokerCommission struct {
	// Total.
	Total int64 `json:"total,omitempty"`
	// List of comission history.
	List []BrokerCommission1 `json:"list,omitempty"`
}
