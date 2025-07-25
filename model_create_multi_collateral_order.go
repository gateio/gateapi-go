/*
 * Gate API
 *
 * Welcome to Gate API  APIv4 provides operations related to spot, margin, and contract trading, including public interfaces for querying market data and authenticated private interfaces for implementing API-based automated trading.
 *
 * Contact: support@mail.gate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type CreateMultiCollateralOrder struct {
	// Order ID.
	OrderId string `json:"order_id,omitempty"`
	// current - current, fixed - fixed, if not specified, default to current.
	OrderType string `json:"order_type,omitempty"`
	// Fixed interest rate loan period: 7d - 7 days, 30d - 30 days. Must be provided for fixed
	FixedType string `json:"fixed_type,omitempty"`
	// Fixed interest rate, must be specified for fixed.
	FixedRate string `json:"fixed_rate,omitempty"`
	// Fixed interest rate, automatic renewal.
	AutoRenew bool `json:"auto_renew,omitempty"`
	// Fixed interest rate, automatic repayment.
	AutoRepay bool `json:"auto_repay,omitempty"`
	// Borrowed currency.
	BorrowCurrency string `json:"borrow_currency"`
	// Borrowing amount.
	BorrowAmount string `json:"borrow_amount"`
	// Collateral currency and amount.
	CollateralCurrencies []CollateralCurrency `json:"collateral_currencies,omitempty"`
}
