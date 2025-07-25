/*
 * Gate API
 *
 * Welcome to Gate API  APIv4 provides operations related to spot, margin, and contract trading, including public interfaces for querying market data and authenticated private interfaces for implementing API-based automated trading.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type SubCrossMarginAccount struct {
	// User ID of the cross margin account. 0 means that the subaccount has not yet opened a cross margin account
	UserId int64 `json:"user_id,omitempty"`
	// Whether account is locked.
	Locked   bool                          `json:"locked,omitempty"`
	Balances map[string]CrossMarginBalance `json:"balances,omitempty"`
	// Total account value in USDT, i.e., the sum of all currencies' `(available+freeze)*price*discount`
	Total string `json:"total,omitempty"`
	// Total borrowed value in USDT, i.e., the sum of all currencies' `borrowed*price*discount`
	Borrowed string `json:"borrowed,omitempty"`
	// Total borrowed value in USDT * borrowed factor.
	BorrowedNet string `json:"borrowed_net,omitempty"`
	// Total net assets in USDT.
	Net string `json:"net,omitempty"`
	// Position leverage.
	Leverage string `json:"leverage,omitempty"`
	// Total unpaid interests in USDT, i.e., the sum of all currencies' `interest*price*discount`
	Interest string `json:"interest,omitempty"`
	// Risk rate. When it belows 110%, liquidation will be triggered. Calculation formula: `total / (borrowed+interest)`
	Risk string `json:"risk,omitempty"`
	// Total initial margin.
	TotalInitialMargin string `json:"total_initial_margin,omitempty"`
	// Total margin balance.
	TotalMarginBalance string `json:"total_margin_balance,omitempty"`
	// Total maintenance margin.
	TotalMaintenanceMargin string `json:"total_maintenance_margin,omitempty"`
	// Total initial margin rate.
	TotalInitialMarginRate string `json:"total_initial_margin_rate,omitempty"`
	// Total maintenance margin rate.
	TotalMaintenanceMarginRate string `json:"total_maintenance_margin_rate,omitempty"`
	// Total available margin.
	TotalAvailableMargin string `json:"total_available_margin,omitempty"`
}
