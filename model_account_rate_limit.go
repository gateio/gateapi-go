/*
 * Gate API
 *
 * Welcome to Gate API  APIv4 provides operations related to spot, margin, and contract trading, including public interfaces for querying market data and authenticated private interfaces for implementing API-based automated trading.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type AccountRateLimit struct {
	// Frequency limit level (For detailed frequency limit rules, see [Transaction ratio frequency limit](#rate-limit-based-on-fill-ratio))
	Tier string `json:"tier,omitempty"`
	// Transaction rate.
	Ratio string `json:"ratio,omitempty"`
	// Total transaction ratio of main account.
	MainRatio string `json:"main_ratio,omitempty"`
	// Update time.
	UpdatedAt string `json:"updated_at,omitempty"`
}
