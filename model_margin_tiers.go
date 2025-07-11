/*
 * Gate API v4
 *
 * Welcome to Gate API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type MarginTiers struct {
	// Tier
	Tier string `json:"tier,omitempty"`
	// Discount
	MarginRate string `json:"margin_rate,omitempty"`
	// Lower limit
	LowerLimit string `json:"lower_limit,omitempty"`
	// Upper limit, \"\" indicates greater than (the last tier)
	UpperLimit string `json:"upper_limit,omitempty"`
	// Position leverage
	Leverage string `json:"leverage,omitempty"`
}
