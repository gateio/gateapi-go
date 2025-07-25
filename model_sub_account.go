/*
 * Gate API
 *
 * Welcome to Gate API  APIv4 provides operations related to spot, margin, and contract trading, including public interfaces for querying market data and authenticated private interfaces for implementing API-based automated trading.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type SubAccount struct {
	// custom text.
	Remark string `json:"remark,omitempty"`
	// Sub-account login name: Only letters, numbers and underscores are supported, and cannot contain other illegal characters
	LoginName string `json:"login_name"`
	// The sub-account's password. (Default: the same as main account's password).
	Password string `json:"password,omitempty"`
	// The sub-account's email address. (Default: the same as main account's email address)
	Email string `json:"email,omitempty"`
	// State: 1-normal, 2-locked\".
	State int32 `json:"state,omitempty"`
	// \"Sub-account type: 1 - sub-account, 3 - cross margin account.
	Type int32 `json:"type,omitempty"`
	// The user id of the sub-account.
	UserId int64 `json:"user_id,omitempty"`
	// Created time.
	CreateTime int64 `json:"create_time,omitempty"`
}
