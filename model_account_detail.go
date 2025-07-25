/*
 * Gate API
 *
 * Welcome to Gate API  APIv4 provides operations related to spot, margin, and contract trading, including public interfaces for querying market data and authenticated private interfaces for implementing API-based automated trading.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Account detail.
type AccountDetail struct {
	// IP whitelist.
	IpWhitelist []string `json:"ip_whitelist,omitempty"`
	// CurrencyPair whitelisting.
	CurrencyPairs []string `json:"currency_pairs,omitempty"`
	// User ID.
	UserId int64 `json:"user_id,omitempty"`
	// User VIP level.
	Tier int64            `json:"tier,omitempty"`
	Key  AccountDetailKey `json:"key,omitempty"`
	// User role: 0 - Normal user, 1 - Copy trading leader, follower, 3 - Both leader and follower
	CopyTradingRole int32 `json:"copy_trading_role,omitempty"`
}
