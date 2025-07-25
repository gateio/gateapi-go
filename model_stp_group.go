/*
 * Gate API
 *
 * Welcome to Gate API  APIv4 provides operations related to spot, margin, and contract trading, including public interfaces for querying market data and authenticated private interfaces for implementing API-based automated trading.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type StpGroup struct {
	// STP Group ID.
	Id int64 `json:"id,omitempty"`
	// STP Group name.
	Name string `json:"name"`
	// Creator ID.
	CreatorId int64 `json:"creator_id,omitempty"`
	// Creation time.
	CreateTime int64 `json:"create_time,omitempty"`
}
