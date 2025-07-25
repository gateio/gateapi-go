/*
 * Gate API
 *
 * Welcome to Gate API  APIv4 provides operations related to spot, margin, and contract trading, including public interfaces for querying market data and authenticated private interfaces for implementing API-based automated trading.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Retrieve user rebate information.
type RebateUserInfo struct {
	// My inviter's UID.
	InviteUid int64 `json:"invite_uid,omitempty"`
}
